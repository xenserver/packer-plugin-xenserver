package common

import (
	"fmt"
	"github.com/mitchellh/packer/packer"
	"os"
	"path/filepath"
)

// This is the common builder ID to all of these artifacts.
const BuilderId = "packer.xenserver"

type LocalArtifact struct {
	dir string
	f   []string
}

func NewArtifact(dir string) (packer.Artifact, error) {
	files := make([]string, 0, 1)
	visit := func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, path)
		}
		return err
	}
	if err := filepath.Walk(dir, visit); err != nil {
		return nil, err
	}

	return &LocalArtifact{
		dir: dir,
		f:   files,
	}, nil
}

func (*LocalArtifact) BuilderId() string {
	return BuilderId
}

func (a *LocalArtifact) Files() []string {
	return a.f
}

func (*LocalArtifact) Id() string {
	return "VM"
}

func (a *LocalArtifact) String() string {
	return fmt.Sprintf("VM files in directory: %s", a.dir)
}

func (a *LocalArtifact) State(name string) interface{} {
	return nil
}

func (a *LocalArtifact) Destroy() error {
	return os.RemoveAll(a.dir)
}
