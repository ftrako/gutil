package goutils

import (
	"github.com/tmc/scp"
	"golang.org/x/crypto/ssh"
	"net"
	"time"
)

func PrivateKeyConfig(user, privateKey string) (*ssh.ClientConfig, error) {
	key, err := ssh.ParsePrivateKey([]byte(privateKey))
	if err != nil {
		return nil, err
	}
	sshConfig := &ssh.ClientConfig{
		Timeout: time.Second * 3,
		User:    user,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(key),
		},
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}
	return sshConfig, nil
}

func SshConnect(cfg *ssh.ClientConfig, ip string) (*ssh.Client, error) {
	sshConn, err := ssh.Dial("tcp", ip+":22", cfg)
	if err != nil {
		return nil, err
	}
	return sshConn, nil
}

func SshRunCmd(sshConn *ssh.Client, cmd string) (string, error) {
	session, err := sshConn.NewSession()
	if err != nil {
		return "", err
	}
	defer session.Close()
	outPut, err := session.CombinedOutput(cmd)
	return string(outPut), err
}

func ScpFile(sshConn *ssh.Client, src, dest string) error {
	session, err := sshConn.NewSession()
	if err != nil {
		return err
	}
	defer session.Close()
	err = scp.CopyPath(src, dest, session)
	if err != nil {
		return err
	}
	return nil
}
