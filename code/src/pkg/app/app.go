package app

import (
	"natrix/pkg/api"

	"golang.org/x/crypto/ssh"
)

type App struct {
	ClientConfig *ssh.ClientConfig
	Host         api.Host
	Client       api.Client
	Command      string
	Action       string
}

func (app *App) SetArgs(action, command, host string, port int, username, password string) {
	app.Action = action
	app.Command = command
	app.Client.Username = username
	app.Client.Password = password
	app.Host.Address = host
	app.Host.Port = port
}

func (app *App) BuildConfig() {

}
