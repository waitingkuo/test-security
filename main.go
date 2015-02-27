package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"os"
	"os/exec"
)

func main() {

	app := cli.NewApp()
	app.Name = "belt"
	app.Usage = "tool-kit for x509"
	app.Commands = []cli.Command{
		{
			Name: "mkkey",
			Action: func(c *cli.Context) {
				cmd := exec.Command("openssl", "genrsa", "1024")
				output, err := cmd.Output()
				if err != nil {
					panic(err)
				}

				fmt.Println(string(output))

			},
		},
		{
			Name: "mkkey2048",
			Action: func(c *cli.Context) {
				cmd := exec.Command("openssl", "genrsa", "2048")
				output, err := cmd.Output()
				if err != nil {
					panic(err)
				}

				fmt.Println(string(output))

			},
		},
		{
			Name: "mkcrt",
			Action: func(c *cli.Context) {
				cmd := exec.Command("openssl", "req", "-new")
				output, err := cmd.Output()
				if err != nil {
					panic(err)
				}

				fmt.Println(string(output))

			},
		},
		{
			Name: "chkcrt",
			Action: func(c *cli.Context) {

				file := c.Args().First()
				//cmd := exec.Command("openssl", "x509", "-in", file, "-text", "-noout")
				cmd := exec.Command("openssl", "x509", "-in", file, "-text")
				output, err := cmd.Output()
				if err != nil {
					panic(err)
				}

				fmt.Println(string(output))
			},
		},
	}

	app.Run(os.Args)
}
