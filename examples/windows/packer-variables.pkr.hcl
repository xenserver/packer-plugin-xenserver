#------------------------------------------------#
#virtual machine variables

variable "vm_name" {
  type      = string
  default   = ""
  sensitive = false
}

variable "vm_description" {
  type    = string
  default   = ""
  sensitive = false
}

variable "firmware" {
  type      = string
  default   = "efi"
  sensitive = false
}

variable "cpu" {
  type      = number
  default   = 4
  sensitive = false
}

variable "memory" {
  type      = number
  default   = 8192
  sensitive = false
}

variable "disksize" {
  type      = number
  default   = 51200
  sensitive = false
}

variable "secure_boot" {
    type = bool
    default = false
    sensitive = false
}

variable "vtpm" {
  type      = bool
  default   = false
  sensitive = false
}

#------------------------------------------------#
#GPU variables

variable "vgpu_profile" {
  type      = string
  default   = ""
  sensitive = false
}

#------------------------------------------------#
#operating system variables

variable "os_iso_url" {
    description = "The download url for the ISO"
    default = ""
}
variable "boot_command" {
  type      = list(string)
  default   = [""]
}

variable "os_iso_path" {
  type      = string
  default   = ""
  sensitive = false
}

variable "connection_username" {
  type      = string
  default   = ""
  sensitive = false
}

variable "connection_password" {
  type      = string
  default   = ""
  sensitive = true
}

variable "http_directory" {
  type      = string
  default   = ""
  sensitive = false
}

#------------------------------------------------#
#Export variables

variable "converttotemplate" {
    type = bool
    default = false
    sensitive = false
}

variable "create_snapshot" {
    type = bool
    default = true
    sensitive = false
}

variable "templatename" {
  type      = string
  default   = ""
  sensitive = false
}


#------------------------------------------------#
# Virtual network

variable "virtual_network_name" {
  type    = string
  default   = ""
  sensitive = false
}

#------------------------------------------------#
#winrm variables

variable "winrm_timeout" {
  type    = string
  default   = ""
  sensitive = false
}

variable "winrm_username" {
  type    = string
  default   = ""
  sensitive = false
}

variable "winrm_password" {
  type    = string
  default   = ""
  sensitive = false
}

#------------------------------------------------#
#XenServer variables

variable "iso_checksum" {
  type    = string
  default   = ""
  sensitive = false
}

variable "iso_checksum_type" {
  type    = string
  default   = ""
  sensitive = false
}

variable "iso_url" {
  type    = string
  default   = ""
  sensitive = false
}

variable "sr_iso_name" {
  type    = string
  default   = ""
  sensitive = false
}

variable "sr_name" {
  type    = string
  default   = ""
  sensitive = false
}

variable "tools_iso_name" {
  type    = string
  default   = ""
  sensitive = false
}

variable "remote_host" {
  type    = string
  default   = ""
  sensitive = false
}

variable "remote_password" {
  type    = string
  default   = ""
  sensitive = false
}

variable "remote_username" {
  type    = string
  default   = ""
  sensitive = false
}

variable "ssh_username" {
  type    = string
  default   = ""
  sensitive = false
}

variable "ssh_password" {
  type    = string
  default   = ""
  sensitive = false
}

variable "clone_template" {
  type    = string
  default   = ""
  sensitive = false
}


