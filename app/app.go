package app

import "github.com/urfave/cli"

func Generate() *cli.App {
	app := cli.NewApp()

	app.Name = "PeekDomain"
	app.Usage = "Consulta IPs e nome de servidores da internet"

	return app
}