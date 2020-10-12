package commands

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/amimof/huego"
	"github.com/urfave/cli/v2"
)

func LightsToggle(c *cli.Context, b *huego.Bridge) error {
	lightID, err := strconv.Atoi(c.Args().First())
	if err != nil {
		return fmt.Errorf(c.Args().First(), "is not a valid light id")
	}
	currentState, err := b.GetLight(lightID)
	if err != nil {
		return err
	}
	if !currentState.State.On {
		fmt.Println("💡 turning on light")
		currentState.State.On = !currentState.State.On
		b.SetLightState(lightID, *currentState.State)
		return nil
	} else {
		fmt.Println("💀 turning off light")
		currentState.State.On = !currentState.State.On
		b.SetLightState(lightID, *currentState.State)
		return nil
	}
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
		lString := fmt.Sprintf("%s\n\t%d\n\t%t", l.Name, l.ID, l.State.On)
		fmt.Println(lString)
	}
}
