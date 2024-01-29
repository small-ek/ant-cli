package cmd

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"os"
)

type Create struct {
}

func (e Create) Action(c *cli.Context) error {
	fmt.Println(c.Args().First())

	err := os.Mkdir(c.Args().First(), 0755) //create a directory and give it required permissions
	if err != nil {
		fmt.Println(err) //print the error on the console
		return err
	}
	fmt.Println("Directory created successfully!")
	return nil
}
