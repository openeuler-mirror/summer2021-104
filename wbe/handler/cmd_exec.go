package handler

import (
	"fmt"
	"io/ioutil"
	"log"
	"reflect"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mitchellh/go-homedir"
	"golang.org/x/crypto/ssh"
)

func PublicKeyAuthFunc(kPath string) ssh.AuthMethod {
	keyPath, err := homedir.Expand(kPath)
	if err != nil {
		log.Fatal("find key's home dir failed", err)
	}
	key, err := ioutil.ReadFile(keyPath)
	if err != nil {
		log.Fatal("ssh key file read failed", err)
	}
	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		log.Fatal("ssh key signer failed", err)
	}
	return ssh.PublicKeys(signer)
}

func ExecCmdFunc(host string, cmd string) string {
	fmt.Println(reflect.TypeOf(cmd))
	fmt.Println(cmd)
	sshHost := host
	sshHost = "127.0.0.1"
	sshUser := "root"
	sshPassword := "123456"
	sshType := "password"
	sshKeyPath := ""
	sshPort := 22
	config := &ssh.ClientConfig{
		Timeout:         time.Second,
		User:            sshUser,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	if sshType == "password" {
		config.Auth = []ssh.AuthMethod{ssh.Password(sshPassword)}
	} else {
		config.Auth = []ssh.AuthMethod{PublicKeyAuthFunc(sshKeyPath)}
	}
	addr := fmt.Sprintf("%s:%d", sshHost, sshPort)
	sshClient, err := ssh.Dial("tcp", addr, config)
	if err != nil {
		log.Fatal("创建ssh client 失败", err)
	}
	defer sshClient.Close()
	session, err := sshClient.NewSession()
	if err != nil {
		log.Fatal("创建ssh session 失败", err)
	}
	defer session.Close()
	combo, err := session.CombinedOutput(cmd)
	if err != nil {
		log.Fatal("远程执行cmd 失败", err)
	}
	log.Println("命令输出:", string(combo))
	return string(combo)
}

func ExecCmd(c *gin.Context) {
	host := c.PostForm("host")
	cmd := c.PostForm("cmd")
	fmt.Println(cmd)
	fmt.Println(reflect.TypeOf(cmd))
	stdout := ExecCmdFunc(host, cmd)
	c.JSON(200, gin.H{
		"message": "OK",
		"stdout":  stdout,
	})
}
