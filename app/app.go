package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli" // v1
)

func Generate() *cli.App {
	app := cli.NewApp()

	app.Name = "PeekDomain"
	app.Usage = "Consulta IPs e nomes de servidores da internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "google.com.br",
			Usage: "Endereço que deseja consultar",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Busca os IPs de um endereço da internet",
			Flags:  flags,
			Action: getIPs,
		},
		{
			Name:   "server",
			Usage:  "Busca os nomes dos servidores DNS de um endereço",
			Flags:  flags,
			Action: getServers,
		},
	}

	return app
}

func getIPs(c *cli.Context) error {
	host := c.String("host")
	ips, err := net.LookupIP(host)
	if err != nil {
		log.Fatalf("Erro ao buscar IPs: %v", err)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
	return nil
}

func getServers(c *cli.Context) error {
	host := c.String("host")
	servers, err := net.LookupNS(host)
	if err != nil {
		log.Fatalf("Erro ao buscar servidores: %v", err)
	}

	for _, server := range servers {
		fmt.Println(server.Host)
	}
	return nil
}
