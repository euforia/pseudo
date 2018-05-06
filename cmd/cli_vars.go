package main

import (
	"net/url"
	"os"
	"reflect"

	"github.com/olekukonko/tablewriter"
	"gopkg.in/urfave/cli.v2"

	"github.com/euforia/pseudo"
	"github.com/euforia/pseudo/ewok"
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
		ew, err := loadContextVars(ctx)
		if err != nil {
			return err
		}

		if ctx.Bool("type") {
			ew.Iter(func(key string, value reflect.Value) bool {
				val := value.Kind().String()
				if val == "invalid" {
					val = "unknown"
				}
				tw.Append([]string{key, val})
				return true
			})
		} else {
			ew.Iter(func(key string, value reflect.Value) bool {
				tw.Append([]string{key})
				return true
			})
		}

	}

	tw.Render()

	return nil
}

func loadContextVars(ctx *cli.Context) (*ewok.Ewok, error) {
	uri, err := url.Parse(ctx.String("context"))
	if err != nil {
		return nil, err
	}

	var opt pseudo.IndexOptions
	if !uri.IsAbs() {
		opt.ContentType = "hcl"
	}

	return pseudo.LoadIndex(uri, opt)
}
