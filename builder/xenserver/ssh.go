package xenserver

import (
	"bytes"
	gossh "code.google.com/p/go.crypto/ssh"
	"fmt"
	"github.com/mitchellh/multistep"
	commonssh "github.com/mitchellh/packer/common/ssh"
	"github.com/mitchellh/packer/communicator/ssh"
	"io"
	"log"
	"net"
	"strings"
)

func sshAddress(state multistep.StateBag) (string, error) {
	sshIP := state.Get("ssh_address").(string)
	sshHostPort := 22
	return fmt.Sprintf("%s:%d", sshIP, sshHostPort), nil
}

func sshLocalAddress(state multistep.StateBag) (string, error) {
	sshLocalPort := state.Get("local_ssh_port").(uint)
	conn_str := fmt.Sprintf("%s:%d", "127.0.0.1", sshLocalPort)
	log.Printf("sshLocalAddress: %s", conn_str)
	return conn_str, nil
}

func sshConfig(state multistep.StateBag) (*gossh.ClientConfig, error) {
	config := state.Get("config").(config)

	auth := []gossh.AuthMethod{
		gossh.Password(config.SSHPassword),
		gossh.KeyboardInteractive(
			ssh.PasswordKeyboardInteractive(config.SSHPassword)),
	}

	if config.SSHKeyPath != "" {
		signer, err := commonssh.FileSigner(config.SSHKeyPath)
		if err != nil {
			return nil, err
		}

		auth = append(auth, gossh.PublicKeys(signer))
	}

	return &gossh.ClientConfig{
		User: config.SSHUser,
		Auth: auth,
	}, nil
}

func execute_ssh_cmd(cmd, host, port, username, password string) (stdout string, err error) {
	// Setup connection config
	config := &gossh.ClientConfig{
		User: username,
		Auth: []gossh.AuthMethod{
			gossh.Password(password),
		},
	}

	client, err := gossh.Dial("tcp", host+":"+port, config)

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

	session.Close()
	return strings.Trim(b.String(), "\n"), nil
}

func forward(local_conn net.Conn, config *gossh.ClientConfig, server, remote_dest string, remote_port uint) {
	ssh_client_conn, err := gossh.Dial("tcp", server+":22", config)
	if err != nil {
		log.Fatalf("local ssh.Dial error: %s", err)
	}

	remote_loc := fmt.Sprintf("%s:%d", remote_dest, remote_port)
	ssh_conn, err := ssh_client_conn.Dial("tcp", remote_loc)
	if err != nil {
		log.Fatalf("ssh.Dial error: %s", err)
	}

	go func() {
		_, err = io.Copy(ssh_conn, local_conn)
		if err != nil {
			log.Fatalf("io.copy failed: %v", err)
		}
	}()

	go func() {
		_, err = io.Copy(local_conn, ssh_conn)
		if err != nil {
			log.Fatalf("io.copy failed: %v", err)
		}
	}()
}

func ssh_port_forward(local_port uint, remote_port uint, remote_dest, host, username, password string) (err error) {

	config := &gossh.ClientConfig{
		User: username,
		Auth: []gossh.AuthMethod{
			gossh.Password(password),
		},
	}

	// Listen on a local port
	local_listener, err := net.Listen("tcp",
		fmt.Sprintf("%s:%d",
			"127.0.0.1",
			local_port))
	if err != nil {
		log.Fatalf("Local listen failed: %s", err)
		return err
	}

	for {
		local_connection, err := local_listener.Accept()

		if err != nil {
			log.Fatalf("Local accept failed: %s", err)
			return err
		}

		// Forward to a remote port
		go forward(local_connection, config, host, remote_dest, remote_port)
	}

	return nil
}
