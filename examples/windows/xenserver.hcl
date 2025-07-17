source "xenserver-iso" "windows" {

    iso_checksum_type           = "none"
	iso_name					= var.os_iso_path
    sr_iso_name                 = var.sr_iso_name
    sr_name                     = var.sr_name
    tools_iso_name              = var.tools_iso_name

    remote_host                 = var.remote_host
    remote_password             = var.remote_password
    remote_username             = var.remote_username

    vm_name                     = var.vm_name
    vm_description              = var.vm_description
    vm_memory                   = var.memory
    disk_size                   = var.disksize
	vcpus				        = var.cpu
	cores_per_socket			= var.cpu
	network_names				= ["${var.virtual_network_name}"]
	firmware                    = var.firmware
    secure_boot 				= var.secure_boot
    vTPM                        = var.vtpm
    vgpu_profile                = var.vgpu_profile

	clone_template				= var.clone_template
	
    #Do not use Floppy files as it does not work with EFI boot
    cd_files                	= []

	ssh_username            	= var.ssh_username
	ssh_password            	= var.ssh_password
	ssh_wait_timeout        	= "60000s"
	
    boot_command = [
        "<spacebar><wait2><spacebar>"
    ]
    boot_wait                   = "1s"
	
	#Snapshot / Template Options
    create_snapshot             = var.create_snapshot
	snapshot_name               = "Created by Packer"
	convert_to_template         = var.converttotemplate

    #Export Options
	format 						= "none"
    keep_vm 					= "on_success"

    communicator                = "winrm"
    winrm_insecure              = true
    winrm_password              = var.winrm_password
    winrm_timeout               = "5m"
    winrm_use_ssl               = true
    winrm_username              = var.winrm_username
}

source "xenserver-clone" "windows" {
    # Pool information
    remote_host                 = var.remote_host
    remote_username             = var.remote_username
    remote_password             = var.remote_password

    # Storage information
    sr_name                     = var.sr_name
    clone_template				= var.clone_template

    #VM information
    vm_name                     = var.vm_name
    vm_description              = var.vm_description
    vm_memory                   = var.memory
    vcpus_atstartup				= var.cpu
	vcpus_max					= var.cpu
	network_names				= ["${var.virtual_network_name}"]
    vgpu_profile                = var.vgpu_profile
	firmware                    = var.firmware
    secure_boot 				= var.secure_boot
    vTPM                        = var.vtpm
    
    #Snapshot / Template Options
    create_snapshot             = var.create_snapshot
	snapshot_name               = "Created by Packer"
	convert_to_template         = var.converttotemplate

    #Export Options
    format 						= "none"
    keep_vm 					= "on_success"

	ssh_username            	= var.ssh_username
	ssh_password            	= var.ssh_password
	ssh_wait_timeout        	= "60000s"

    communicator                = "winrm"
    winrm_insecure              = true
    winrm_password              = var.winrm_password
    winrm_timeout               = "5m"
    winrm_use_ssl               = true
    winrm_username              = var.winrm_username
}

build {
    name = "GoldenImageFromISO"
    sources = ["xenserver-iso.windows"]
  
}

build {
    name = "GoldenImageFromTemplate"
    sources = ["xenserver-clone.windows"]
  
}