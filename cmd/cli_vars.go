package main

import (
	"net/url"
	"os"
	"strings"

	"github.com/olekukonko/tablewriter"
	"gopkg.in/urfave/cli.v2"

	"github.com/euforia/pseudo"
	"github.com/euforia/pseudo/scope"
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
		// Context variables
		vars, err := loadContextVars(ctx)
		if err != nil {
			return err
		}

		names := vars.Names()

		if ctx.Bool("type") {

			for _, n := range names {
				typ := strings.ToLower(vars[n].Type.String()[4:])
				tw.Append([]string{n, typ})
			}

		} else {

			for i := range names {
				tw.Append([]string{names[i]})
			}

		}

	}

	tw.Render()

	return nil
}

func loadContextVars(ctx *cli.Context) (scope.Variables, error) {
	uri, err := url.Parse(ctx.String("context"))
	if err != nil {
		return nil, err
	}

	var opt pseudo.IndexOptions
	if !uri.IsAbs() {
		opt.ContentType = "pseudo"
	}

	return pseudo.LoadVariables(uri, opt)
}
