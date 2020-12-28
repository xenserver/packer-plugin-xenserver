module github.com/xenserver/packer-builder-xenserver

go 1.14

require (
	github.com/amfranz/go-xmlrpc-client v0.0.0-20190612172737-76858463955d
	github.com/dylanmei/iso8601 v0.1.0 // indirect
	github.com/dylanmei/winrmtest v0.0.0-20151226195028-025617847eb2 // indirect
	github.com/hashicorp/errwrap v0.0.0-20141028054710-7554cd9344ce // indirect
	github.com/hashicorp/go-multierror v0.0.0-20150916205742-d30f09973e19 // indirect
	github.com/hashicorp/go-version v0.0.0-20160119211326-7e3c02b30806 // indirect
	github.com/hashicorp/yamux v0.0.0-20151129044643-df949784da9e // indirect
	github.com/kr/fs v0.0.0-20131111012553-2788f0dbd169 // indirect
	github.com/masterzen/simplexml v0.0.0-20140219194429-95ba30457eb1 // indirect
	github.com/masterzen/winrm v0.0.0-20151214220635-54ea5d01478c // indirect
	github.com/masterzen/xmlpath v0.0.0-20140218185901-13f4951698ad // indirect
	github.com/mitchellh/go-fs v0.0.0-20150611040906-a34c1b9334e8 // indirect
	github.com/mitchellh/go-vnc v0.0.0-20150629162542-723ed9867aed
	github.com/mitchellh/iochan v0.0.0-20150529224432-87b45ffd0e95 // indirect
	github.com/mitchellh/mapstructure v0.0.0-20150717051158-281073eb9eb0 // indirect
	github.com/mitchellh/multistep v0.0.0-20140407010115-162146fc5711
	github.com/mitchellh/packer v0.10.2-0.20160629004225-63edbd40edc5
	github.com/mitchellh/reflectwalk v0.0.0-20150527153153-eecf4c70c626 // indirect
	github.com/niemeyer/pretty v0.0.0-20200227124842-a10e7caefd8e // indirect
	github.com/nu7hatch/gouuid v0.0.0-20131221200532-179d4d0c4d8d // indirect
	github.com/packer-community/winrmcp v0.0.0-20160310040704-f1bcf36a69fa // indirect
	github.com/pkg/sftp v0.0.0-20160118190721-e84cc8c755ca // indirect
	github.com/satori/go.uuid v0.0.0-20151028231719-d41af8bb6a77 // indirect
	github.com/stretchr/testify v1.6.1 // indirect
	github.com/terra-farm/go-xen-api-client v0.0.1
	github.com/ugorji/go v0.0.0-20151218193438-646ae4a518c1 // indirect
	golang.org/x/crypto v0.0.0-20201208171446-5f87f3452ae9
	gopkg.in/check.v1 v1.0.0-20200227125254-8fa46927fb4f // indirect
	launchpad.net/gocheck v0.0.0-20140225173054-000000000087 // indirect
	launchpad.net/xmlpath v0.0.0-20130614043138-000000000004 // indirect
)

// replace github.com/glycerine/sshego => github.com/ddelnano/sshego v1.0.1
replace github.com/glycerine/sshego => /home/ddelnano/code/sshego
