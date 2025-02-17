package interaction

import (
	"fmt"
	// log "github.com/sirupsen/logrus"
	"os"
)

type CliTable interface {
	Keyword() []string
	SprintAll() string
}

func ParseInputArguments(cli_tables []CliTable) {
	table_map := make(map[string]CliTable)

	for _, table := range cli_tables {
		for _, key := range table.Keyword() {
			// log.Infof("Found available keyword: %v", key)
			table_map[key] = table
		}
	}

	for _, arg := range os.Args[1:] {
		if table, exists := table_map[arg]; exists {
			fmt.Print(table.SprintAll())
		}
	}
}
