package xenserver

import (
    "github.com/mitchellh/multistep"
    "github.com/mitchellh/packer/packer"
    "github.com/mitchellh/packer/common"
    "fmt"
    "log"
    "errors"
)


// Set the unique ID for this builder
const BuilderId = "packer.xenserver"


type config struct {
    common.PackerConfig `mapstructure:",squash"`

    Username        string  `mapstructure:"username"`
    Password        string  `mapstructure:"password"`
    HostIp          string  `mapstructure:"host_ip"`
    IsoUrl          string  `mapstructure:"iso_url"`

    InstanceName    string  `mapstructure:"instance_name"`

    tpl *packer.ConfigTemplate
}


type Builder struct {
    config config
    runner multistep.Runner
}


func (self *Builder) Prepare (raws ...interface{}) ([]string, error) {

    md, err := common.DecodeConfig(&self.config, raws...)
    if err != nil {
        return nil, err
    }

    self.config.tpl, err = packer.NewConfigTemplate()

    if err != nil {
        return nil, err
    }

    self.config.tpl.UserVars = self.config.PackerUserVars


    errs := common.CheckUnusedConfig(md)


    // Set default vaules

    templates := map[string]*string {
        "username":         &self.config.Username,
        "password":         &self.config.Password,
        "host_ip":          &self.config.HostIp,
        "iso_url":          &self.config.IsoUrl,
        "instance_name":    &self.config.InstanceName,
    }


    for n, ptr := range templates {
        var err error
        *ptr, err = self.config.tpl.Process(*ptr, nil)
        if err != nil {
            errs = packer.MultiErrorAppend(errs, fmt.Errorf("Error processing %s: %s", n, err))
        }
    }

    if self.config.IsoUrl == "" {
        errs = packer.MultiErrorAppend(
                errs, errors.New("a iso url must be specified"))
    }
    return nil, nil

}

func (self *Builder) Run(ui packer.Ui, hook packer.Hook, cache packer.Cache) (packer.Artifact, error) {
    log.Println("Run, not yet implemented!")
    return nil, nil
}


func (self *Builder) Cancel() {
    if self.runner != nil {
        log.Println("Cancelling the step runner...")
        self.runner.Cancel()
    }
    fmt.Println("Cancelling the builder")
}

