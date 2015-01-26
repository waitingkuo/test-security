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
  app.Usage =  "tool-kit for x509"
  app.Commands = []cli.Command {
    {
      Name: "chkcrt",
      Action: func (c *cli.Context) {

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
