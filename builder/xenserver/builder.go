package xenserver

import (
	"errors"
	"fmt"
	"github.com/mitchellh/multistep"
	"github.com/mitchellh/packer/common"
	commonssh "github.com/mitchellh/packer/common/ssh"
	"github.com/mitchellh/packer/packer"
	"log"
	"os"
	"time"
)

// Set the unique ID for this builder
const BuilderId = "packer.xenserver"

type config struct {
	common.PackerConfig `mapstructure:",squash"`

	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	HostIp   string `mapstructure:"host_ip"`
	IsoUrl   string `mapstructure:"iso_url"`

	InstanceName   string `mapstructure:"instance_name"`
	InstanceMemory string `mapstructure:"instance_memory"`
	RootDiskSize   string `mapstructure:"root_disk_size"`
	CloneTemplate  string `mapstructure:"clone_template"`
	IsoName        string `mapstructure:"iso_name"`
	SrName         string `mapstructure:"sr_name"`
	NetworkName    string `mapstructure:"network_name"`

	HostPortMin uint `mapstructure:"host_port_min"`
	HostPortMax uint `mapstructure:"host_port_max"`

	BootCommand []string `mapstructure:"boot_command"`

	RawBootWait string        `mapstructure:"boot_wait"`
	BootWait    time.Duration ``

	ISOChecksum     string   `mapstructure:"iso_checksum"`
	ISOChecksumType string   `mapstructure:"iso_checksum_type"`
	ISOUrls         []string `mapstructure:"iso_urls"`
	ISOUrl          string   `mapstructure:"iso_url"`

	HTTPDir     string `mapstructure:"http_directory"`
	HTTPPortMin uint   `mapstructure:"http_port_min"`
	HTTPPortMax uint   `mapstructure:"http_port_max"`

	LocalIp      string            `mapstructure:"local_ip"`
	PlatformArgs map[string]string `mapstructure:"platform_args"`

	RawInstallTimeout string        `mapstructure:"install_timeout"`
	InstallTimeout    time.Duration ``

	RawSSHWaitTimeout string        `mapstructure:"ssh_wait_timeout"`
	SSHWaitTimeout    time.Duration ``

	SSHPassword string `mapstructure:"ssh_password"`
	SSHUser     string `mapstructure:"ssh_username"`
	SSHKeyPath  string `mapstructure:"ssh_key_path"`

	OutputDir string `mapstructure:"output_directory"`

	KeepInstance string `mapstructure:"keep_instance"`

	tpl *packer.ConfigTemplate
}

type Builder struct {
	config config
	runner multistep.Runner
}

func (self *Builder) Prepare(raws ...interface{}) (params []string, retErr error) {

	md, err := common.DecodeConfig(&self.config, raws...)
	if err != nil {
		return nil, err
	}

	errs := common.CheckUnusedConfig(md)
	if errs == nil {
		errs = &packer.MultiError{}
	}

	self.config.tpl, err = packer.NewConfigTemplate()

	if err != nil {
		return nil, err
	}

	// Set default vaules

	if self.config.HostPortMin == 0 {
		self.config.HostPortMin = 5900
	}

	if self.config.HostPortMax == 0 {
		self.config.HostPortMax = 6000
	}

	if self.config.RawBootWait == "" {
		self.config.RawBootWait = "5s"
	}

	if self.config.RawInstallTimeout == "" {
		self.config.RawInstallTimeout = "200m"
	}

	if self.config.HTTPPortMin == 0 {
		self.config.HTTPPortMin = 8000
	}

	if self.config.HTTPPortMax == 0 {
		self.config.HTTPPortMax = 9000
	}

	if self.config.RawSSHWaitTimeout == "" {
		self.config.RawSSHWaitTimeout = "200m"
	}

	if self.config.InstanceMemory == "" {
		self.config.InstanceMemory = "1024000000"
	}

	if self.config.CloneTemplate == "" {
		self.config.CloneTemplate = "Other install media"
	}

	if self.config.OutputDir == "" {
		self.config.OutputDir = fmt.Sprintf("output-%s", self.config.PackerBuildName)
	}

	if self.config.KeepInstance == "" {
		self.config.KeepInstance = "never"
	}

	if len(self.config.PlatformArgs) == 0 {
		pargs := make(map[string]string)
		pargs["viridian"] = "false"
		pargs["nx"] = "true"
		pargs["pae"] = "true"
		pargs["apic"] = "true"
		pargs["timeoffset"] = "0"
		pargs["acpi"] = "1"
		self.config.PlatformArgs = pargs
	}

	// Template substitution

	templates := map[string]*string{
		"username":          &self.config.Username,
		"password":          &self.config.Password,
		"host_ip":           &self.config.HostIp,
		"iso_url":           &self.config.IsoUrl,
		"instance_name":     &self.config.InstanceName,
		"instance_memory":   &self.config.InstanceMemory,
		"root_disk_size":    &self.config.RootDiskSize,
		"clone_template":    &self.config.CloneTemplate,
		"iso_name":          &self.config.IsoName,
		"sr_name":           &self.config.SrName,
		"network_name":      &self.config.NetworkName,
		"boot_wait":         &self.config.RawBootWait,
		"iso_checksum":      &self.config.ISOChecksum,
		"iso_checksum_type": &self.config.ISOChecksumType,
		"http_directory":    &self.config.HTTPDir,
		"local_ip":          &self.config.LocalIp,
		"install_timeout":   &self.config.RawInstallTimeout,
		"ssh_wait_timeout":  &self.config.RawSSHWaitTimeout,
		"ssh_username":      &self.config.SSHUser,
		"ssh_password":      &self.config.SSHPassword,
		"ssh_key_path":      &self.config.SSHKeyPath,
		"output_directory":  &self.config.OutputDir,
		"keep_instance":     &self.config.KeepInstance,
	}

	for n, ptr := range templates {
		var err error
		*ptr, err = self.config.tpl.Process(*ptr, nil)
		if err != nil {
			errs = packer.MultiErrorAppend(errs, fmt.Errorf("Error processing %s: %s", n, err))
		}
	}

	// Validation

	/*
	   if self.config.IsoUrl == "" {
	       errs = packer.MultiErrorAppend(
	               errs, errors.New("a iso url must be specified"))
	   }
	*/

	self.config.BootWait, err = time.ParseDuration(self.config.RawBootWait)
	if err != nil {
		errs = packer.MultiErrorAppend(
			errs, fmt.Errorf("Failed to parse boot_wait: %s", err))
	}

	self.config.SSHWaitTimeout, err = time.ParseDuration(self.config.RawSSHWaitTimeout)
	if err != nil {
		errs = packer.MultiErrorAppend(
			errs, fmt.Errorf("Failed to parse ssh_wait_timeout: %s", err))
	}

	self.config.InstallTimeout, err = time.ParseDuration(self.config.RawInstallTimeout)
	if err != nil {
		errs = packer.MultiErrorAppend(
			errs, fmt.Errorf("Failed to parse install_timeout: %s", err))
	}

	for i, command := range self.config.BootCommand {
		if err := self.config.tpl.Validate(command); err != nil {
			errs = packer.MultiErrorAppend(errs,
				fmt.Errorf("Error processing boot_command[%d]: %s", i, err))
		}
	}

	if self.config.SSHUser == "" {
		errs = packer.MultiErrorAppend(
			errs, errors.New("An ssh_username must be specified."))
	}

	if self.config.SSHKeyPath != "" {
		if _, err := os.Stat(self.config.SSHKeyPath); err != nil {
			errs = packer.MultiErrorAppend(
				errs, fmt.Errorf("ssh_key_path is invalid: %s", err))
		} else if _, err := commonssh.FileSigner(self.config.SSHKeyPath); err != nil {
			errs = packer.MultiErrorAppend(
				errs, fmt.Errorf("ssh_key_path is invalid: %s", err))
		}
	}

	if self.config.Username == "" {
		errs = packer.MultiErrorAppend(
			errs, errors.New("A username for the xenserver host must be specified."))
	}

	if self.config.Password == "" {
		errs = packer.MultiErrorAppend(
			errs, errors.New("A password for the xenserver host must be specified."))
	}

	if self.config.HostIp == "" {
		errs = packer.MultiErrorAppend(
			errs, errors.New("An ip for the xenserver host must be specified."))
	}

	if self.config.InstanceName == "" {
		errs = packer.MultiErrorAppend(
			errs, errors.New("An instance name must be specified."))
	}

	if self.config.RootDiskSize == "" {
		errs = packer.MultiErrorAppend(
			errs, errors.New("A root disk size must be specified."))
	}

	switch self.config.KeepInstance {
	case "always", "never", "on_success":
	default:
		errs = packer.MultiErrorAppend(
			errs, errors.New("keep_instance must be one of 'always', 'never', 'on_success'"))
	}

	/*
	   if self.config.LocalIp == "" {
	       errs = packer.MultiErrorAppend(
	               errs, errors.New("A local IP visible to XenServer's mangement interface is required to serve files."))
	   }
	*/

	if self.config.HTTPPortMin > self.config.HTTPPortMax {
		errs = packer.MultiErrorAppend(
			errs, errors.New("the HTTP min port must be less than the max"))
	}

	if self.config.HostPortMin > self.config.HostPortMax {
		errs = packer.MultiErrorAppend(
			errs, errors.New("the host min port must be less than the max"))
	}
	/*
	   if self.config.ISOChecksumType == "" {
	       errs = packer.MultiErrorAppend(
	               errs, errors.New("The iso_checksum_type must be specified."))
	   } else {
	       self.config.ISOChecksumType = strings.ToLower(self.config.ISOChecksumType)
	       if self.config.ISOChecksumType != "none" {
	           if self.config.ISOChecksum == "" {
	               errs = packer.MultiErrorAppend(
	                       errs, errors.New("Due to the file size being large, an iso_checksum is required."))
	           } else {
	               self.config.ISOChecksum = strings.ToLower(self.config.ISOChecksum)
	           }

	           if hash := common.HashForType(self.config.ISOChecksumType); hash == nil {
	               errs = packer.MultiErrorAppend(
	                       errs, fmt.Errorf("Unsupported checksum type: %s", self.config.ISOChecksumType))
	           }

	       }
	   }

	   if self.config.ISOUrl == "" {
	       errs = packer.MultiErrorAppend(
	               errs, errors.New("A ISO URL must be specfied."))
	   } else {
	       self.config.ISOUrls = []string{self.config.ISOUrl}
	   }

	   for i, url := range self.config.ISOUrls {
	       self.config.ISOUrls[i], err = common.DownloadableURL(url)
	       if err != nil {
	           errs = packer.MultiErrorAppend(
	                   errs, fmt.Errorf("Failed to parse the iso_url (%d): %s", i, err))
	       }
	   }
	*/
	if len(errs.Errors) > 0 {
		retErr = errors.New(errs.Error())
	}

	return nil, retErr

}

func (self *Builder) Run(ui packer.Ui, hook packer.Hook, cache packer.Cache) (packer.Artifact, error) {
	//Setup XAPI client
	client := NewXenAPIClient(self.config.HostIp, self.config.Username, self.config.Password)

	err := client.Login()
	if err != nil {
		return nil, err.(error)
	}
	ui.Say("XAPI client session established")

	client.GetHosts()

	//Share state between the other steps using a statebag
	state := new(multistep.BasicStateBag)
	state.Put("cache", cache)
	state.Put("client", client)
	state.Put("config", self.config)
	state.Put("hook", hook)
	state.Put("ui", ui)

	//Build the steps
	steps := []multistep.Step{
		/*
		   &common.StepDownload{
		       Checksum:       self.config.ISOChecksum,
		       ChecksumType:   self.config.ISOChecksumType,
		       Description:    "ISO",
		       ResultKey:      "iso_path",
		       Url:            self.config.ISOUrls,
		   },
		*/
		new(stepPrepareOutputDir),
		new(stepHTTPServer),
		//new(stepUploadIso),
		new(stepCreateInstance),
		new(stepStartVmPaused),
		new(stepGetVNCPort),
		&stepForwardPortOverSSH{
			RemotePort:  instanceVNCPort,
			RemoteDest:  instanceVNCIP,
			HostPortMin: self.config.HostPortMin,
			HostPortMax: self.config.HostPortMax,
			ResultKey:   "local_vnc_port",
		},
		new(stepBootWait),
		new(stepTypeBootCommand),
		new(stepWait),
		new(stepStartOnHIMN),
		&stepForwardPortOverSSH{
			RemotePort:  himnSSHPort,
			RemoteDest:  himnSSHIP,
			HostPortMin: self.config.HostPortMin,
			HostPortMax: self.config.HostPortMax,
			ResultKey:   "local_ssh_port",
		},
		&common.StepConnectSSH{
			SSHAddress:     sshLocalAddress,
			SSHConfig:      sshConfig,
			SSHWaitTimeout: self.config.SSHWaitTimeout,
		},
		new(common.StepProvision),
		new(stepShutdownAndExport),
	}

	self.runner = &multistep.BasicRunner{Steps: steps}
	self.runner.Run(state)

	if rawErr, ok := state.GetOk("error"); ok {
		return nil, rawErr.(error)
	}

	// If we were interrupted or cancelled, then just exit.
	if _, ok := state.GetOk(multistep.StateCancelled); ok {
		return nil, errors.New("Build was cancelled.")
	}
	if _, ok := state.GetOk(multistep.StateHalted); ok {
		return nil, errors.New("Build was halted.")
	}

	artifact, _ := NewArtifact(self.config.OutputDir)

	return artifact, nil
}

// all steps should check config.ShouldKeepInstance first before cleaning up
func (cfg config) ShouldKeepInstance(state multistep.StateBag) bool {
	switch cfg.KeepInstance {
	case "always":
		return true
	case "never":
		return false
	case "on_success":
		// only keep instance if build was successful
		_, cancelled := state.GetOk(multistep.StateCancelled)
		_, halted := state.GetOk(multistep.StateHalted)
		return !(cancelled || halted)
	default:
		panic(fmt.Sprintf("Unknown keep_instance value '%s'", cfg.KeepInstance))
	}
}

func (self *Builder) Cancel() {
	if self.runner != nil {
		log.Println("Cancelling the step runner...")
		self.runner.Cancel()
	}
	fmt.Println("Cancelling the builder")
}
