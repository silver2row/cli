/**
	* 1. Setup the server so cf can call it under main.
				e.g. `cf my-plugin` creates the callable server. now we can call the Run command
	* 2. Implement Run that is the actual code of the plugin!
	* 3. Return an error
**/

package main

import (
	"fmt"
	"os"

	"github.com/cloudfoundry/cli/plugin"
)

type CliPlugin struct{}

var commands = []plugin.Command{
	{
		Name:     "test_2_cmd1",
		HelpText: "help text for test_2_cmd1",
	},
	{
		Name:     "test_2_cmd2",
		HelpText: "help text for test_2_cmd2",
	},
}

func (c *CliPlugin) Run(args string, reply *bool) error {
	if args == "test_2_cmd1" {
		theFirstCmd()
	} else if args == "test_2_cmd2" {
		theSecondCmd()
	}
	return nil
}

func (c *CliPlugin) ListCmds(args string, cmdlist *[]plugin.Command) error {
	*cmdlist = commands
	return nil
}

func (c *CliPlugin) CmdExists(args string, exists *bool) error {
	*exists = plugin.CmdExists(args, commands)
	return nil
}

func theFirstCmd() {
	fmt.Println("You called cmd1 in test_2")
}

func theSecondCmd() {
	fmt.Println("You called cmd2 in test_2")
}

func main() {
	port := os.Args[1]
	plugin.ServeCommand(new(CliPlugin), port)
}