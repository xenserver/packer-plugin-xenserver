# -----------------------------------------------------------------------------
# XenServer ISO builder: Windows golden image over WinRM
#
# Installs Windows from an ISO on an ISO SR, runs a provisioner over WinRM, and
# leaves the result as a template on the pool. It needs the answer file and the
# VM Tools MSI described at cd_files below.
#
# Credentials come from the environment, for example:
#   export PKR_VAR_xenserver_host=pool-master.example.com
#   export PKR_VAR_xenserver_password='...'
#   export PKR_VAR_winrm_password='...'
#
# Run packer from the plugin's root directory (the one with go.mod); the cd_files
# paths below are relative to it:
#   packer init     examples/windows-iso/windows-iso-winrm.pkr.hcl
#   packer validate examples/windows-iso/windows-iso-winrm.pkr.hcl
#   packer build    examples/windows-iso/windows-iso-winrm.pkr.hcl
# -----------------------------------------------------------------------------

packer {
  required_version = ">= 1.7.0"

  required_plugins {
    xenserver = {
      version = ">= 1.0.0"
      source  = "github.com/xenserver/xenserver"
    }
  }
}

# --- Connection -------------------------------------------------------------

variable "xenserver_host" {
  type        = string
  description = "Hostname or IP of the pool coordinator (master)."
}

variable "xenserver_username" {
  type    = string
  default = "root"
}

variable "xenserver_password" {
  type      = string
  sensitive = true
}

variable "ca_cert_path" {
  type        = string
  default     = null
  description = "The pool's certificate (xe host-get-server-certificate). Leave unset to take it from XENSERVER_CA_CERT_PATH."
}

# --- Install media and placement --------------------------------------------

variable "iso_name" {
  type        = string
  description = "Name label of the Windows install ISO, as it appears on iso_sr."
}

variable "iso_sr" {
  type        = string
  description = "Name label of the SR holding the install ISO."
}

# Empty means the pool's default SR. A named SR must be unique on the pool, so
# "Local storage" fails on a pool of more than one host.
variable "vm_sr" {
  type    = string
  default = ""
}

variable "clone_template" {
  type        = string
  default     = "Windows 11"
  description = "XenServer template the build VM starts from; it sets the guest-OS hints."
}

variable "virtual_network_name" {
  type    = string
  default = "Pool-wide network associated with eth0"
}

# --- Hardware ---------------------------------------------------------------

variable "vm_name" {
  type    = string
  default = "packer-windows-golden"
}

variable "vm_description" {
  type    = string
  default = "Created by Packer"
}

variable "memory_gb" {
  type    = number
  default = 8
}

variable "disk_size_gb" {
  type    = number
  default = 64
}

variable "cpu" {
  type    = number
  default = 4
}

# Overrides clone_template's boot mode. vtpm and secure_boot need "uefi".
variable "firmware" {
  type    = string
  default = "uefi"
}

# true requires firmware = "uefi"; false turns Secure Boot off. Remove the line
# from the source block below to keep the template's setting.
variable "secure_boot" {
  type    = bool
  default = true
}

# Windows 11 needs UEFI and a vTPM. false adds none, but keeps a vTPM that
# clone_template is set up with.
variable "vtpm" {
  type    = bool
  default = true
}

# --- Communicator -----------------------------------------------------------

variable "winrm_username" {
  type    = string
  default = "Administrator"
}

variable "winrm_password" {
  type      = string
  sensitive = true
}

source "xenserver-vm-iso" "windows" {
  # Connection to the pool coordinator
  xenserver_host     = var.xenserver_host
  xenserver_username = var.xenserver_username
  xenserver_password = var.xenserver_password
  ca_cert_path       = var.ca_cert_path

  # Install media, already present on an ISO SR
  iso_name = var.iso_name
  iso_sr   = var.iso_sr

  # Placement
  vm_sr          = var.vm_sr
  clone_template = var.clone_template
  network_names  = [var.virtual_network_name]

  # Hardware
  vm_name          = var.vm_name
  vm_description   = var.vm_description
  vm_memory_gb     = var.memory_gb
  disk_size_gb     = var.disk_size_gb
  vcpus            = var.cpu
  cores_per_socket = var.cpu
  firmware         = var.firmware
  secure_boot      = var.secure_boot
  vtpm             = var.vtpm

  # Repeat <enter> to catch the UEFI "Press any key to boot from CD" prompt, whose
  # timing varies. <enter>, not <spacebar>: a repeat that reaches Windows Boot
  # Manager takes its default instead of stopping the countdown.
  boot_wait               = "3s"
  boot_command            = ["<enter>"]
  boot_key_groups         = 20
  boot_key_group_interval = "1s"

  # Supply Autounattend.xml and the XenServer VM Tools MSI yourself; this
  # repository ships neither. The answer file must install the VM Tools, from
  # which the build learns the VM's address, and create the WinRM 5986 HTTPS
  # listener with a certificate the machine running Packer trusts, issued for the
  # address Packer connects to. See
  # https://learn.microsoft.com/en-us/windows-hardware/customize/desktop/unattend/
  #   cd_files = [
  #     "examples/windows-iso/windows/Autounattend.xml",
  #     "examples/windows-iso/windows/managementagentx64.msi",
  #   ]

  # Communicator: WinRM, over the listener the answer file creates.
  communicator   = "winrm"
  winrm_username = var.winrm_username
  winrm_password = var.winrm_password
  winrm_use_ssl  = true
  winrm_timeout  = "30m"

  # Output: format = "none" exports no artifact and leaves the finished VM as a
  # template.
  format = "none"
}

build {
  name    = "windows-golden-from-iso"
  sources = ["source.xenserver-vm-iso.windows"]

  provisioner "powershell" {
    inline = [
      "Write-Host 'Provisioning the golden image...'",
      "Get-Service -Name WinRM | Format-List Name, Status"
    ]
  }
}
