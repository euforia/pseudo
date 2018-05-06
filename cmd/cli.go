package main

import (
	"errors"
	"fmt"
	"net/url"
	"os"
	"path/filepath"

	"github.com/euforia/pseudo"
	"gopkg.in/urfave/cli.v2"
)

// CLI is the command line interface
type CLI struct {
	*cli.App
}

// NewCLI inits a new CLI with the given version
func NewCLI(version string) *CLI {
	c := &CLI{
		App: &cli.App{
			Name:                  "pseudo",
			Version:               version,
			EnableShellCompletion: true,
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:    "context",
					Aliases: []string{"ctx", "c"},
					Usage:   "source variables `path`",
					Value:   "./etc/context.hcl", // Temporary
				},
				&cli.StringFlag{
					Name:    "out",
					Aliases: []string{"o"},
					Usage:   "output `file`",
				},
				&cli.BoolFlag{
					Name:  "debug",
					Usage: "turn on debug mode",
				},
			},
			Before: before,
			Action: execScript,
		},
	}
	c.UsageText = c.Name + ` [global options] [command] [command options] [script]`

	c.init()

	return c
}

func (c *CLI) init() {
	cli.VersionPrinter = func(ctx *cli.Context) {
		fmt.Println(ctx.App.Name + " " + ctx.App.Version)
	}
	c.Commands = []*cli.Command{
		varsCommand(c.Name),
	}
}

func before(ctx *cli.Context) error {
	if ctx.Bool("debug") {

	}

	return nil
}

func execScript(ctx *cli.Context) error {
	args := ctx.Args()

	fdpath := args.First()
	if len(fdpath) == 0 {
		cli.ShowAppHelpAndExit(ctx, 1)
	}

	fpath, err := filepath.Abs(fdpath)
	if err != nil {
		return err
	}

	stat, err := os.Stat(fpath)
	if err != nil {
		return err
	}

	if stat.IsDir() {
		err = execScriptDir(ctx, fpath)
	} else {
		err = execScriptFile(ctx, fpath)
	}

	return err
}

func execScriptDir(ctx *cli.Context, fpath string) error {
	_, err := pseudo.ReadDirFiles(fpath)
	if err != nil {
		return err
	}

	return errors.New("not yet implemented")
}

func execScriptFile(ctx *cli.Context, fpath string) error {
	vars, err := loadVarsMap(ctx)
	if err != nil {
		return err
	}

	script, err := pseudo.NewScript(fpath)
	if err != nil {
		return err
	}

	vm := pseudo.NewVM()
	vm.SetVars(vars)

	err = vm.Parse(script.Contents())
	if err != nil {
		return err
	}

	result, err := vm.Eval()
	if err != nil {
		return err
	}

	out := ctx.String("out")
	if len(out) == 0 {
		fmt.Printf("%s", result.Value)
		return nil
	}

	outURL, err := url.Parse(out)
	if err != nil {
		return err
	}

	var fh *os.File
	fh, err = os.Create(outURL.Path)
	if err == nil {
		fmt.Fprintln(fh, result.Value)
		err = fh.Close()
	}
	return err
}

func loadVarsMap(ctx *cli.Context) (pseudo.VarsMap, error) {
	uri, err := url.Parse(ctx.String("context"))
	if err != nil {
		return nil, err
	}

	var varsmap pseudo.VarsMap

	switch uri.Scheme {

	case "http", "https":
		err = fmt.Errorf("scheme='%s' not yet supported", uri.Scheme)

	default:
		varsmap, err = pseudo.LoadHCLScopeVarsFromFile(uri.Path)

	}

	return varsmap, err
}
