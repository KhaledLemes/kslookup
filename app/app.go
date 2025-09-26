package app

import (
	"fmt"
	"log"
	"net"
	"github.com/urfave/cli"
)

// Generate will return the cli application ready to be used
func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "kslookup"
	app.Usage = "Searches for websites' IPs and servers they are running on"


	// Since we will be using the same flag twice for two different commands, I inserted it into this variable
	flags := []cli.Flag{
				cli.StringFlag{					
					Name: "host",
					Value: "google.com",
				},
			}

	app.Commands = []cli.Command{
		{
			Name: "ip",
			Usage: "Looks for the public ip",

			// This part is where we add the flags
			// May be a little difficult to read
			Flags: flags,
			Action: searchIPS,
		},
		{
			Name: "server",
			Usage: "Looks for the server a website is running on.",
			Flags: flags,
			Action: searchServers,
		},
	}
	return app
}

// Searches for the public IP of a website
func searchIPS(c *cli.Context) {
	host := c.String("host")

	ips, err := net.LookupIP(host)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Public ips for %s\n", host)
	for i, ip := range ips {
		fmt.Println(i+1, "-", ip)
	}
}

// Searches for the server the website is running on
func searchServers(c *cli.Context) {
	host := c.String("host")

	servers, err := net.LookupNS(host)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s is running on:\n", host)

	for i, server := range servers {
		fmt.Println(i+1, "-", server.Host)
	}
}