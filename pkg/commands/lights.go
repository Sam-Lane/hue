package commands

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/amimof/huego"
	"github.com/urfave/cli/v2"
)

func LightsTurn(c *cli.Context, b *huego.Bridge) error {
	lightID, err := strconv.Atoi(c.Args().First())
	if err != nil {
		return fmt.Errorf(c.Args().First(), "is not a valid light id")
	}
	currentState, err := b.GetLight(lightID)
	if err != nil {
		return err
	}
	state := c.Args().Slice()[1]
	if state == "on" {
		fmt.Println("💡 turning on light")
		currentState.State.On = true
		b.SetLightState(lightID, *currentState.State)
		return nil
	}
	if state == "off" {
		fmt.Println("💀 turning off light")
		currentState.State.On = false
		b.SetLightState(lightID, *currentState.State)
		return nil
	}
	return fmt.Errorf(state, "is not a valid command. Use on/off")
}

func LightsGet(c *cli.Context, b *huego.Bridge) error {
	if c.Args().First() == "all" {
		lights, err := b.GetLights()
		if err != nil {
			return err
		}
		for _, l := range lights {
			printLight(c, &l)
		}
		return nil
	}
	i, err := strconv.Atoi(c.Args().First())
	if err != nil {
		return err
	}
	light, err := b.GetLight(i)
	if err != nil {
		return err
	}
	printLight(c, light)
	return nil
}

func printLight(c *cli.Context, l *huego.Light) {
	if c.Bool("json") {
		b, err := json.Marshal(l)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(b))
	} else {
		fmt.Println(l.Name, l.State.On, l.ID)
	}
}
