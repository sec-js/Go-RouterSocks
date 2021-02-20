package prompt

import (
	"strings"

	prompt "github.com/c-bata/go-prompt"
)

var commands = []prompt.Suggest{
	{Text: "route", Description: "Manage routes. (default print them)"},
	{Text: "chisel", Description: "List chisel port on this machine"},
	{Text: "help", Description: "Help menu"},

	{Text: "exit", Description: "Exit this program"},
}

var actionsRoute = []prompt.Suggest{
	{Text: "add", Description: "Add new route"},
	{Text: "flush", Description: "Remove all the routes"},
	{Text: "print", Description: "Print the routes"},
	{Text: "delete", Description: "Delete one route"},
}

var helpSubCommand = []prompt.Suggest{
	{Text: "route", Description: "Add new route"},
	{Text: "chisel", Description: "List chisel port on this machine"},
}

func complete(d prompt.Document) []prompt.Suggest {
	if d.TextBeforeCursor() == "" {
		return []prompt.Suggest{}
	}
	args := strings.Split(d.TextBeforeCursor(), " ")

	if len(args) <= 1 {
		return prompt.FilterHasPrefix(commands, args[0], true)
	}
	first := args[0]
	switch first {
	case "route":
		second := args[1]
		if len(args) == 2 {
			return prompt.FilterHasPrefix(actionsRoute, second, true)
		}

	case "help":
		second := args[1]
		if len(args) == 2 {
			return prompt.FilterHasPrefix(helpSubCommand, second, true)
		}
	}
	return []prompt.Suggest{}
}
