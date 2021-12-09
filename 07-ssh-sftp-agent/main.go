package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

// check for any error
func check(err error) {
	if err != nil {
		fmt.Printf("Error Happened %s \n", err)
		os.Exit(1)
	}
}

// In case if you have private key authentication
// uncomment the sshKeyPath variable
// uncomment sshDemoWithPrivateKey() function
// call the sshDemoWithPrivateKey() function instead of sshDemoWithPassword()
var (
	sshUserName        = "vagrant"
	sshPassword        = "vagrant"
	sshKeyPath         = "/home/akilan/Desktop/linux-vagrant/.vagrant/machines/ubuntu_ssh_server/virtualbox/private_key"
	sshHostname        = "192.168.33.10:22"
	commandToExec      = "ls -la /home/vagrant"
	fileToUpload       = "./upload.txt"
	fileUploadLocation = "/home/vagrant/upload.txt"
	fileToDownload     = "/home/vagrant/download.txt"
)

func main() {

	fmt.Println("....Golang SSH Demo......")

	//conf := sshDemoWithPassword() // username and password authentication
	conf := sshDemoWithPrivateKey() // username and private key authentication

	// open ssh connection
	sshClient, err := ssh.Dial("tcp", sshHostname, conf)
	check(err)
	session, err := sshClient.NewSession()
	check(err)
	defer session.Close()

	// execute command on remote server
	var b bytes.Buffer
	session.Stdout = &b
	err = session.Run(commandToExec)
	check(err)
	log.Printf("%s: %s", commandToExec, b.String())

	// open sftp connection
	sftpClient, err := sftp.NewClient(sshClient)
	check(err)
	defer sftpClient.Close()

	// create a file
	createFile, err := sftpClient.Create(fileToDownload)
	check(err)
	text := "This file created by Golang SSH.\nThis will be downloaded by Golang SSH\n"
	_, err = createFile.Write([]byte(text))
	check(err)
	fmt.Println("Created file ", fileToDownload)

	// Upload a file
	srcFile, err := os.Open(fileToUpload)
	check(err)
	defer srcFile.Close()

	dstFile, err := sftpClient.Create(fileUploadLocation)
	check(err)
	defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	check(err)
	fmt.Println("File uploaded successfully ", fileUploadLocation)

	// Download a file
	remoteFile, err := sftpClient.Open(fileToDownload)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to open remote file: %v\n", err)
		return
	}
	defer remoteFile.Close()

	localFile, err := os.Create("./download.txt")
	check(err)
	defer localFile.Close()

	_, err = io.Copy(localFile, remoteFile)
	check(err)
	fmt.Println("File downloaded successfully")

}

func sshDemoWithPassword() *ssh.ClientConfig {

	// ssh config with password authentication
	conf := &ssh.ClientConfig{
		User: sshUserName,
		Auth: []ssh.AuthMethod{
			ssh.Password(sshPassword),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	return conf

}

func sshDemoWithPrivateKey() *ssh.ClientConfig {
	keyByte, err := ioutil.ReadFile(sshKeyPath)
	check(err)
	key, err := ssh.ParsePrivateKey(keyByte)
	check(err)

	// ssh config
	conf := &ssh.ClientConfig{
		User: sshUserName,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(key),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	return conf
}
