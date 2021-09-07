package main

import (
	"github.com/urfave/cli/v2"
	"os"
)



func main() {
	app := &cli.App{
		Name: "verse",
		Usage: "Write a verse HTML and json file",
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "html-output-file", Usage: "Path to output the HTML file" },
			&cli.StringFlag{Name: "json-output-file", Usage: "Path to output the json file" },
		},
		Action: Parse,
	}
	
	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}

