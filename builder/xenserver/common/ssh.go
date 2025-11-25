package common

import (
	"bytes"
	"encoding/pem"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strings"

	"github.com/hashicorp/packer-plugin-sdk/multistep"
	gossh "golang.org/x/crypto/ssh"
)

func SSHAddress(state multistep.StateBag) (string, error) {
	sshIP := state.Get("ssh_address").(string)
	sshHostPort := state.Get("ssh_port").(uint)
	return fmt.Sprintf("%s:%d", sshIP, sshHostPort), nil
}

func SSHLocalAddress(state multistep.StateBag) (string, error) {
	sshLocalPort, ok := state.Get("local_ssh_port").(uint)
	if !ok {
		return "", fmt.Errorf("SSH port forwarding hasn't been set up yet")
	}
	conn_str := fmt.Sprintf("%s:%d", "127.0.0.1", sshLocalPort)
	return conn_str, nil
}

func SSHPort(state multistep.StateBag) (int, error) {
	sshHostPort := state.Get("local_ssh_port").(uint)
	return int(sshHostPort), nil
}

func CommHost(state multistep.StateBag) (string, error) {
	return "127.0.0.1", nil
}

func SSHConfigFunc(config SSHConfig) func(multistep.StateBag) (*gossh.ClientConfig, error) {
	return func(state multistep.StateBag) (*gossh.ClientConfig, error) {
		config := state.Get("commonconfig").(CommonConfig)
		auth := []gossh.AuthMethod{
			gossh.Password(config.SSHPassword),
		}

		if config.SSHKeyPath != "" {
			signer, err := FileSigner(config.SSHKeyPath)
			if err != nil {
				return nil, err
			}

			auth = append(auth, gossh.PublicKeys(signer))
		}

		return &gossh.ClientConfig{
			User:            config.SSHUser,
			Auth:            auth,
			HostKeyCallback: gossh.InsecureIgnoreHostKey(),
		}, nil
	}
}

func doExecuteSSHCmd(cmd, target string, config *gossh.ClientConfig) (stdout string, err error) {
	client, err := gossh.Dial("tcp", target, config)
	if err != nil {
		return "", err
	}

	//Create session
	session, err := client.NewSession()
	if err != nil {
		return "", err
	}

	defer session.Close()

	var b bytes.Buffer
	session.Stdout = &b
	if err := session.Run(cmd); err != nil {
		return "", err
	}

	return strings.Trim(b.String(), "\n"), nil
}

func ExecuteHostSSHCmd(state multistep.StateBag, cmd string) (stdout string, err error) {
	config := state.Get("commonconfig").(CommonConfig)
	sshAddress, _ := SSHAddress(state)
	// Setup connection config
	sshConfig := &gossh.ClientConfig{
		User: config.Username,
		Auth: []gossh.AuthMethod{
			gossh.Password(config.Password),
		},
		HostKeyCallback: gossh.InsecureIgnoreHostKey(),
	}
	return doExecuteSSHCmd(cmd, sshAddress, sshConfig)
}

func ExecuteGuestSSHCmd(state multistep.StateBag, cmd string) (stdout string, err error) {
	config := state.Get("commonconfig").(CommonConfig)
	localAddress, err := SSHLocalAddress(state)
	if err != nil {
		return
	}
	sshConfig, err := SSHConfigFunc(config.SSHConfig)(state)
	if err != nil {
		return
	}

	return doExecuteSSHCmd(cmd, localAddress, sshConfig)
}

func forward(local_conn net.Conn, config *gossh.ClientConfig, server string, server_ssh_port int, remote_dest string, remote_port uint) error {
	defer local_conn.Close()

	ssh_client_conn, err := gossh.Dial("tcp", fmt.Sprintf("%s:%d", server, server_ssh_port), config)
	if err != nil {
		log.Printf("local ssh.Dial error: %s", err)
		return err
	}
	defer ssh_client_conn.Close()

	remote_loc := fmt.Sprintf("%s:%d", remote_dest, remote_port)
	ssh_conn, err := ssh_client_conn.Dial("tcp", remote_loc)
	if err != nil {
		log.Printf("ssh.Dial error: %s", err)
		return err
	}
	defer ssh_conn.Close()

	txDone := make(chan struct{})
	rxDone := make(chan struct{})

	go func() {
		_, err = io.Copy(ssh_conn, local_conn)
		if err != nil {
			log.Printf("io.copy failed: %v", err)
		}
		close(txDone)
	}()

	go func() {
		_, err = io.Copy(local_conn, ssh_conn)
		if err != nil {
			log.Printf("io.copy failed: %v", err)
		}
		close(rxDone)
	}()

	<-txDone
	<-rxDone

	return nil
}

func ssh_port_forward(local_listener net.Listener, remote_port int, remote_dest, host string, host_ssh_port int, username, password string) error {

	config := &gossh.ClientConfig{
		User: username,
		Auth: []gossh.AuthMethod{
			gossh.Password(password),
		},
		HostKeyCallback: gossh.InsecureIgnoreHostKey(),
	}

	for {
		local_connection, err := local_listener.Accept()

		if err != nil {
			log.Printf("Local accept failed: %s", err)
			return err
		}

		// Forward to a remote port
		go forward(local_connection, config, host, host_ssh_port, remote_dest, uint(remote_port))
	}

	return nil
}

// FileSigner returns an gossh.Signer for a key file.
func FileSigner(path string) (gossh.Signer, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	keyBytes, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	// We parse the private key on our own first so that we can
	// show a nicer error if the private key has a password.
	block, _ := pem.Decode(keyBytes)
	if block == nil {
		return nil, fmt.Errorf(
			"Failed to read key '%s': no key found", path)
	}
	if block.Headers["Proc-Type"] == "4,ENCRYPTED" {
		return nil, fmt.Errorf(
			"Failed to read key '%s': password protected keys are\n"+
				"not supported. Please decrypt the key prior to use.", path)
	}

	signer, err := gossh.ParsePrivateKey(keyBytes)
	if err != nil {
		return nil, fmt.Errorf("Error setting up SSH config: %s", err)
	}

	return signer, nil
}
