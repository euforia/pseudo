package main

import (
	"fmt"
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
			&cli.BoolFlag{
				Name:    "value",
				Aliases: []string{"v"},
				Usage:   "Show values",
			},
		},
		Action: varsListExec,
	}
}

func getScriptVars(scriptfile string) ([][]string, error) {
	script, err := pseudo.NewScript(scriptfile)
	if err != nil {
		return nil, err
	}

	vars := script.Vars()
	data := make([][]string, len(vars))
	for i := range vars {
		data[i] = []string{vars[i]}
	}
	return data, nil
}

func varsListExec(ctx *cli.Context) error {
	tw := tablewriter.NewWriter(os.Stdout)
	tw.SetAlignment(tablewriter.ALIGN_LEFT)

	var (
		header   = make([]string, 1, 3)
		data     [][]string
		showType = ctx.Bool("type")
		showVal  = ctx.Bool("value")

		args = ctx.Args()
		err  error
	)

	header[0] = "key"

	if args.Len() > 0 {
		// Script variables
		data, err = getScriptVars(args.Get(0))
		if err != nil {
			return err
		}

	} else {
		// Context variables
		vars, err := loadContextVars(ctx)
		if err != nil {
			return err
		}

		names := vars.Names()
		data = make([][]string, len(names))

		for j, name := range names {
			data[j] = make([]string, 1, 3)
			data[j][0] = name
		}

		if showType && showVal {
			header = append(header, "type", "value")
			for j, name := range names {
				typ := strings.ToLower(vars[name].Type.String()[4:])
				val := fmt.Sprintf("%v", vars[name].Value)
				data[j] = append(data[j], typ, val)
			}
		} else if showType {
			header = append(header, "type")
			for j, name := range names {
				typ := strings.ToLower(vars[name].Type.String()[4:])
				data[j] = append(data[j], typ)
			}
		} else if showVal {
			header = append(header, "value")
			for j, name := range names {
				val := fmt.Sprintf("%v", vars[name].Value)
				data[j] = append(data[j], val)
			}
		}

	}

	tw.SetHeader(header)
	tw.AppendBulk(data)
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
