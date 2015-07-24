// GANS : Go Automated Nmap Scanner
// This permit to launch a scanner and add adresses to be scanned.
package main

import (
	"github.com/codegangsta/cli"
	"github.com/HugoDelval/gans/cmd"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "Gans"
	app.Usage = "Go automated Nmap scanner"
	app.Commands = []cli.Command{
		cmd.CmdScan,
		cmd.CmdRun,
		cmd.CmdParse,
	}
	app.Run(os.Args)
}
