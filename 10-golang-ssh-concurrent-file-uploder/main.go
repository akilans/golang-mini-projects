package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

type Server struct {
	Hostname           string
	Username           string
	Password           string
	PrivateKeyPath     string
	FiletoUpload       string
	FileUploadLocation string
}

var wg sync.WaitGroup

// check for any error
func check(err error) {
	if err != nil {
		fmt.Printf("Error Happened %s \n", err)
		return
	}
}

// Main function
func main() {

	servers := []Server{
		{"127.0.0.1:2222", "vagrant", "", "./.vagrant/machines/ubuntu_1/virtualbox/private_key", "hello.txt", "/home/vagrant/"},
		{"127.0.0.1:2200", "vagrant", "", "./.vagrant/machines/ubuntu_2/virtualbox/private_key", "hello.txt", "/home/vagrant/"},
		{"127.0.0.1:2201", "vagrant", "", "./.vagrant/machines/ubuntu_3/virtualbox/private_key", "hello.txt", "/home/vagrant/"},
		{"127.0.0.1:2202", "vagrant", "", "./.vagrant/machines/ubuntu_4/virtualbox/private_key", "hello.txt", "/home/vagrant/"},
	}

	// with out go routine
	startTime := time.Now()
	for _, server := range servers {
		server.UploadFileToServer(false)
	}
	endTime := time.Now()

	fmt.Printf("Without Go Routine - Upload task took %.2f seconds \n", endTime.Sub(startTime).Seconds())

	startTime = time.Now()

	// with go routine
	wg.Add(len(servers))
	for _, server := range servers {
		go server.UploadFileToServer(true)
	}
	wg.Wait()
	endTime = time.Now()
	fmt.Printf("With Go Routine - Upload task took %.2f seconds \n", endTime.Sub(startTime).Seconds())

}

// Upload file to remote SSH server

func (s Server) UploadFileToServer(usingGoRoutine bool) {

	if usingGoRoutine {
		defer wg.Done()
		s.FiletoUpload = "hi.txt"
	}

	if s.Password == "" && s.PrivateKeyPath == "" {
		fmt.Println("Please provide ssh password or ssh password")
	}

	var conf *ssh.ClientConfig

	// get ssh config

	if s.Password != "" {
		//fmt.Println("Login with password")
		conf = s.sshDemoWithPassword()
	}

	if s.PrivateKeyPath != "" {
		//fmt.Println("Login with private key")
		conf = s.sshDemoWithPrivateKey()
	}

	// open ssh connection
	sshClient, err := ssh.Dial("tcp", s.Hostname, conf)
	check(err)

	// open sftp connection
	//fmt.Println("open sftp connection")
	sftpClient, err := sftp.NewClient(sshClient)
	check(err)
	defer sftpClient.Close()

	// Upload a file
	//fmt.Println("Upload a file")
	srcFile, err := os.Open(s.FiletoUpload)
	check(err)
	defer srcFile.Close()

	//fmt.Println("Create a file in remote machine")
	dstFile, err := sftpClient.Create(filepath.Join(s.FileUploadLocation, s.FiletoUpload))
	check(err)
	defer dstFile.Close()

	//fmt.Println("Uploading a file")
	_, err = io.Copy(dstFile, srcFile)
	check(err)
	fmt.Printf("File %v uploaded successfully to %v \n", s.FiletoUpload, s.Hostname)

}

// SSH auth with Password
func (s Server) sshDemoWithPassword() *ssh.ClientConfig {

	// ssh config with password authentication
	conf := &ssh.ClientConfig{
		User: s.Username,
		Auth: []ssh.AuthMethod{
			ssh.Password(s.Password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	return conf

}

// SSH auth with Private Key
func (s Server) sshDemoWithPrivateKey() *ssh.ClientConfig {
	keyByte, err := os.ReadFile(s.PrivateKeyPath)
	check(err)
	signer, err := ssh.ParsePrivateKey(keyByte)
	check(err)

	// ssh config
	conf := &ssh.ClientConfig{
		User: s.Username,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	return conf
}
