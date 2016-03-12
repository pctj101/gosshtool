package gosshtool

import (
	"testing"
)

func Test_newSSHClinet(t *testing.T) {
	config := &SSHClientConfig{
		User:     "myname",
		Password: "1231",
		Host:     "11.23.1.1",
	}
	sshclient := NewSSHClient(config)
	stdout, stderr, err := sshclient.Cmd("pwd")
	if err != nil {
		t.Error(err)
	}
	t.Log(stdout)
	t.Log(stderr)

}

func Test_mutiCmd(t *testing.T) {
	config := &SSHClientConfig{
		User:     "jack",
		Password: "assd",
		Host:     "31.11.11.11",
	}
	NewSSHClient(config)

	config2 := &SSHClientConfig{
		User:     "asd",
		Password: "daas",
		Host:     "8.8.8.8",
	}
	NewSSHClient(config2)
	stdout, _, err := ExecuteCmd("pwd", "8.8.8.8")
	if err != nil {
		t.Error(err)
	}
	t.Log(stdout)

	stdout, _, err = ExecuteCmd("pwd", "114.215.151.48")
	if err != nil {
		t.Error(err)
	}
	t.Log(stdout)
}