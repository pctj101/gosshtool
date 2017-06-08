package gosshtool

import (
	"golang.org/x/crypto/ssh"
)

type ForwardConfig struct {
	LocalBindAddress string
	RemoteAddress    string
	SshServerAddress string
	SshUserName      string
	SshUserPassword  string
	SshPrivateKey    string
	HostKeyCallback  ssh.HostKeyCallback
}
