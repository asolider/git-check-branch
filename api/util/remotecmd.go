package util

import (
	"fmt"
	"log"
	"net"
	"strings"
	"time"

	"golang.org/x/crypto/ssh"
)

func connect(user, password, host string, port int) (*ssh.Session, error) {
	var (
		auth         []ssh.AuthMethod
		addr         string
		clientConfig *ssh.ClientConfig
		client       *ssh.Client
		session      *ssh.Session
		err          error
	)
	// get auth method
	auth = make([]ssh.AuthMethod, 0)
	auth = append(auth, ssh.Password(password))

	clientConfig = &ssh.ClientConfig{
		User:    user,
		Auth:    auth,
		Timeout: 10 * time.Second,
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}

	// connet to ssh
	addr = fmt.Sprintf("%s:%d", host, port)

	if client, err = ssh.Dial("tcp", addr, clientConfig); err != nil {
		return nil, err
	}

	// create session
	if session, err = client.NewSession(); err != nil {
		return nil, err
	}

	return session, nil
}

func runCmd(server *serverItem, cmd string) ([]byte, error) {
	log.Printf("服务器 %s 开始执行命令: %s", server.Name, cmd)
	session, err := connect(server.User, server.Passwd, server.IP, server.Port)
	if err != nil {
		log.Printf("登录服务器失败， err: ", err)
		return nil, err
	}
	defer session.Close()

	res, err := session.CombinedOutput(cmd)
	log.Printf("服务器 %s 执行命令: %s 执行结果:\n%s", server.Name, cmd, res)
	return res, err
}

func formatCmdReturn(res []byte) []string {
	lines := strings.Split(string(res), "\n")
	strMap := make([]string, 0)
	for _, v := range lines {
		line := strings.TrimSpace(v)
		if line != "" {
			strMap = append(strMap, line)
		}
	}
	log.Printf("format cmd res : \n %#v", strMap)
	return strMap
}
