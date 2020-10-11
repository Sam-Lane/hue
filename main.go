package main

import (
	"os"

	"github.com/sam-lane/hue/pkg/commands"
	"github.com/sam-lane/hue/pkg/util"

	"github.com/amimof/huego"
	"github.com/urfave/cli/v2"
)

func main() {
	hueConf := util.ReadConfig()
	bridge, _ := setUpBridge(*hueConf)
	app := &cli.App{
		Name:    "hue",
		Usage:   "A cli app for controlling Philips Hue",
		Version: util.AppVersion,
	}
	app.EnableBashCompletion = true

	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:     "config",
			Aliases:  []string{"c"},
			Usage:    "path location for the config file",
			FilePath: "~/.hue",
		},
		&cli.BoolFlag{
			Name:  "json",
			Usage: "set output to json format",
		},
	}

	app.Commands = []*cli.Command{
		{
			Name:    "lights",
			Aliases: []string{"l"},
			Usage:   "Control philips hue lights",
			Subcommands: []*cli.Command{
				{
					Name:  "turn",
					Usage: "turn on light(s)",
					Action: func(c *cli.Context) error {
						return commands.LightsTurn(c, bridge)
					},
				},
				{
					Name:    "get",
					Aliases: []string{"g"},
					Usage:   "Return a list of light(s)",
					Action: func(c *cli.Context) error {
						return commands.LightsGet(c, bridge)
					},
				},
			},
		},
		{
			Name:    "bridge",
			Aliases: []string{"b"},
			Usage:   "Interface with the philips hue bridge",
			Subcommands: []*cli.Command{
				{
					Name:  "user",
					Usage: "create or delete users",
					Subcommands: []*cli.Command{
						{
							Name:  "add",
							Usage: "create a new user",
							Action: func(c *cli.Context) error {
								return commands.CreateNewUser(c, bridge)
							},
						},
						{
							Name:  "del",
							Usage: "delete a user",
							Action: func(c *cli.Context) error {
								return commands.DeleteUser(c, bridge)
							},
						},
					},
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}

func setUpBridge(conf util.HueConfig) (*huego.Bridge, error) {
	if conf.IPAddress == "" {
		bridge, err := huego.Discover()
		if err != nil {
			return nil, err
		}
		bridge.Login(conf.Username)
		return bridge, nil
	}
	bridge := huego.New(conf.IPAddress, conf.Username)
	return bridge, nil
}
