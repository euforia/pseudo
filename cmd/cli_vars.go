package main

import (
	"os"
	"sort"

	"github.com/olekukonko/tablewriter"
	"gopkg.in/urfave/cli.v2"

	"github.com/euforia/pseudo"
)

func varsCommand(biname string) *cli.Command {
	return &cli.Command{
		Name:  "vars",
		Usage: "Interact with available variables",
		Subcommands: []*cli.Command{
			varsListCommand(biname),
		},
	}
}

func varsListCommand(biname string) *cli.Command {
	return &cli.Command{
		Name:      "list",
		Aliases:   []string{"ls"},
		Usage:     "List available variables",
		ArgsUsage: "[ script ]",
		UsageText: biname + ` vars ls [ script ]

   This command is used to list available source variables or source variables
   referenced in a script.`,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "type",
				Aliases: []string{"t"},
				Usage:   "Show data type",
			},
		},
		Action: varsListExec,
	}
}

func varsListExec(ctx *cli.Context) error {
	tw := tablewriter.NewWriter(os.Stdout)
	tw.SetBorder(false)
	tw.SetColumnSeparator("")

	args := ctx.Args()

	if args.Len() > 0 {
		// Script variables
		script, err := pseudo.NewScript(args.Get(0))
		if err != nil {
			return err
		}

		vars := script.Vars()
		for i := range vars {
			tw.Append([]string{vars[i]})
		}

	} else {
		// Source variables
		varmap, err := loadSourceVars(ctx)
		if err != nil {
			return err
		}

		vars := varmap.Names()
		sort.Strings(vars)

		if ctx.Bool("type") {
			for _, v := range vars {
				tw.Append([]string{v[1:], varmap[v].Type.String()[4:]})
			}

		} else {
			for _, v := range vars {
				tw.Append([]string{v[1:]})
			}
		}

	}

	tw.Render()

	return nil
}
