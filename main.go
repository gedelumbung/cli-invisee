package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gedelumbung/cli-invisee/invisee"
	"github.com/urfave/cli"
)

var app = cli.NewApp()
var invApp = invisee.Init("production")

func info() {
	app.Name = "Invisee CLI tool"
	app.Usage = "Simple CLI tool to help you connected with Invisee"
	app.Author = "Gede Lumbung"
	app.Version = "0.1"
}

func commands() {
	app.Commands = []cli.Command{
		{
			Name:  "signature",
			Usage: "Generate signature",
			Action: func(c *cli.Context) error {
				signature := invisee.Signature(invApp, c.Args().Get(0))
				fmt.Println("Signature : " + signature)
				return nil
			},
		},
		{
			Name:  "login",
			Usage: "Login to Invisee",
			Action: func(c *cli.Context) error {
				response := invisee.Login(invApp, c.Args().Get(0), c.Args().Get(1))
				fmt.Println(response)
				return nil
			},
		},
		{
			Name:  "investments",
			Usage: "Get Investments List",
			Action: func(c *cli.Context) error {
				response := invisee.Investments(invApp, c.Args().Get(0), c.Args().Get(1))
				fmt.Println(response)
				return nil
			},
		},
		{
			Name:  "transactions",
			Usage: "Get Transactions List",
			Action: func(c *cli.Context) error {
				signature := invisee.Signature(invApp, c.Args().Get(0))
				fmt.Println("Signature : " + signature)
				return nil
			},
		},
		{
			Name:  "order-status",
			Usage: "Get Order Status",
			Action: func(c *cli.Context) error {
				signature := invisee.Signature(invApp, c.Args().Get(0))
				fmt.Println("Signature : " + signature)
				return nil
			},
		},
		{
			Name:  "rop",
			Usage: "Get Range of Partial",
			Action: func(c *cli.Context) error {
				signature := invisee.Signature(invApp, c.Args().Get(0))
				fmt.Println("Signature : " + signature)
				return nil
			},
		},
	}
}

func main() {
	info()
	commands()

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
