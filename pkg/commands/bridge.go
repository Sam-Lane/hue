package commands

import (
	"fmt"
	"strings"

	"github.com/amimof/huego"
	"github.com/urfave/cli/v2"
)

func CreateNewUser(c *cli.Context, b *huego.Bridge) error {
	userDesc := strings.Join(c.Args().Slice(), " ")
	fmt.Print("💡 press bridge button then any key to continue...")
	fmt.Scanln()
	key, err := b.CreateUser(userDesc)
	if err != nil {
		return err
	}
	fmt.Println(fmt.Sprintf("New user api key: %s", key))
	return nil
}

func DeleteUser(c *cli.Context, b *huego.Bridge) error {
	return b.DeleteUser(c.Args().First())
}
