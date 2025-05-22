package main

import (
	"fmt"
	"log"
	"natrix/pkg/app"
	"os"
	"strconv"

	"golang.org/x/crypto/ssh"
)

func main() {

	app := app.App{}
	app.SetArgs(getArgs())

	clientConfig := &ssh.ClientConfig{
		User: app.Client.Username,
		Auth: []ssh.AuthMethod{
			ssh.Password(app.Client.Password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	connection, err := ssh.Dial("tcp", app.Host.Address+":"+strconv.Itoa(app.Host.Port), clientConfig)
	if err != nil {
		log.Printf("Failed to dial: %s", err)
		return
	}
	defer connection.Close()

	session, err := connection.NewSession()
	if err != nil {
		log.Printf("Failed to create session: %s", err)
		return
	}
	defer session.Close()

	session.Stdout = os.Stdout
	session.Stderr = os.Stderr

	err = session.Run(app.Command)
	if err != nil {
		log.Printf("Failed to run command: %s", err)
		return
	}
}

func getArgs() (string, string, string, int, string, string) {
	if len(os.Args) != 7 {
		fmt.Println("Usage: natrix run \"command\" username password host port")
		os.Exit(1)
	}
	action := os.Args[1]
	command := os.Args[2]
	host := os.Args[3]
	portStr := os.Args[4]
	username := os.Args[5]
	password := os.Args[6]

	if action != "run" {
		log.Printf("Unknown action: %s", action)
		return "", "", "", 0, "", ""
	}
	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Printf("Invalid port number: %s", portStr)
		return "", "", "", 0, "", ""
	}

	return action, command, host, port, username, password
}
