# -----------------------------------------------------------------------------
# XenServer clone builder: Windows golden image over WinRM
#
# Clones an existing Windows template, runs a provisioner over WinRM, and leaves
# the result as a template on the pool. The template must already have the
# XenServer VM Tools installed: the build learns the VM's address from them.
#
# Credentials come from the environment, for example:
#   export PKR_VAR_xenserver_host=pool-master.example.com
#   export PKR_VAR_xenserver_password='...'
#   export PKR_VAR_winrm_password='...'
#
# Run packer from the plugin's root directory (the one with go.mod):
#   packer init     examples/windows-clone-winrm/windows-clone-winrm.pkr.hcl
#   packer validate examples/windows-clone-winrm/windows-clone-winrm.pkr.hcl
#   packer build    examples/windows-clone-winrm/windows-clone-winrm.pkr.hcl
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

# --- Source and placement ---------------------------------------------------

variable "clone_template" {
  type        = string
  default     = "Windows 11 (64-bit) golden"
  description = "Name label of the source template to clone. Use clone_vm instead to clone a VM."
}

# A named SR must be unique on the pool, so "Local storage" fails on a pool of
# more than one host.
variable "vm_sr" {
  type        = string
  default     = ""
  description = "Name label of the SR for the VM's disks. Empty keeps them on the template's SR; another SR copies them in full."
}

variable "virtual_network_name" {
  type    = string
  default = "Pool-wide network associated with eth0"
}

# --- Hardware ---------------------------------------------------------------

variable "vm_name" {
  type    = string
  default = "packer-win11-golden"
}

variable "memory_gb" {
  type    = number
  default = 8
}

variable "cpu" {
  type    = number
  default = 4
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

source "xenserver-vm-clone" "windows" {
  # Connection to the pool coordinator
  xenserver_host     = var.xenserver_host
  xenserver_username = var.xenserver_username
  xenserver_password = var.xenserver_password
  ca_cert_path       = var.ca_cert_path

  # Source and placement. With vm_sr empty, the clone's disks stay on the
  # template's SR.
  clone_template = var.clone_template
  vm_sr          = var.vm_sr
  # Replaces the template's NICs, so the guest sees new MAC addresses. Remove it
  # to keep them.
  network_names = [var.virtual_network_name]

  # Hardware
  vm_name      = var.vm_name
  vm_memory_gb = var.memory_gb
  vcpus        = var.cpu

  # Communicator: WinRM. The template must already have a WinRM HTTPS listener on
  # 5986, with a certificate the machine running Packer trusts, issued for the
  # address Packer connects to.
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
  name    = "windows-golden-from-template"
  sources = ["source.xenserver-vm-clone.windows"]

  provisioner "powershell" {
    inline = [
      "Write-Host 'Provisioning the golden image...'",
      # Stand-in for your real provisioning.
      "Get-Service -Name WinRM | Format-List Name, Status"
    ]
  }
}
