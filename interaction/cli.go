package interaction

import (
	// "fmt"
	log "github.com/sirupsen/logrus"
	"strings"
	// "os"
)

type CliSnippet map[string]func([]string)

type CliTableCommand interface {
	SpecialKeyword() string
	Arguments() CliSnippet
}

type TestCommand struct{}

func (c TestCommand) Arguments() CliSnippet {
	return CliSnippet{
		"-h": func(words []string) {
			log.Info(strings.Join(words, " "))
		},
	}
}

func (c TestCommand) SpecialKeyword() string {
	return "--test"
}

func ParseInputArguments(cli_tables []CliTableCommand, args []string) {
	for _, cli_table := range cli_tables {
		special := cli_table.SpecialKeyword()
		if args[1] != special {
			continue
		}
		cmds := cli_table.Arguments()
		for i := 0; i < len(args); i++ {
			if cmd, exists := cmds[args[i]]; exists {
				cmd(args[i+1:])
			}
		}
	}
}
