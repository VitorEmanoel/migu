package main

import (
	"fmt"
	"os"

	"github.com/naoina/migu"
	"github.com/naoina/migu/dialect"
	"github.com/spf13/cobra"
)

func init() {
	dump := &dump{}
	dumpCmd := &cobra.Command{
		Use:   "dump [OPTIONS] DATABASE [FILE]",
		Short: "dump the database schema as Go code",
		RunE: func(cmd *cobra.Command, args []string) error {
			return dump.Execute(args)
		},
	}
	dumpCmd.SetUsageTemplate(usageTemplate + "\nWith FILE, output to FILE.\n")
	rootCmd.AddCommand(dumpCmd)
}

type dump struct{}

func (d *dump) Execute(args []string) error {
	var dbname string
	var filename string
	switch len(args) {
	case 0:
		return fmt.Errorf("too few arguments")
	case 1:
		dbname = args[0]
	case 2:
		dbname, filename = args[0], args[1]
	default:
		return fmt.Errorf("too many arguments")
	}
	db, err := database(dbname)
	if err != nil {
		return err
	}
	defer db.Close()
	di := dialect.NewMySQL(db)
	return d.run(di, filename)
}

func (d *dump) run(di dialect.Dialect, filename string) error {
	out := os.Stdout
	if filename != "" {
		file, err := os.Create(filename)
		if err != nil {
			return err
		}
		defer file.Close()
		out = file
	}
	return migu.Fprint(out, di)
}
