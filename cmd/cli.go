package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"path/filepath"

	"github.com/euforia/pseudo"
	"github.com/hashicorp/hil/ast"
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
					Name:    "scope",
					Aliases: []string{"s"},
					Usage:   "source variables `path`",
					Value:   "./etc/scope.hcl", // Temporary
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
	dirfiles, err := ioutil.ReadDir(fpath)
	if err == nil {
		for _, f := range dirfiles {
			fmt.Println(f.Name())
		}
		err = errors.New("not yet implemented")
	}
	return err
}

func execScriptFile(ctx *cli.Context, fpath string) error {
	rvars, err := loadSourceVars(ctx)
	if err != nil {
		return err
	}
	vars := make(map[string]ast.Variable)
	for k, v := range rvars {
		vars[k[1:]] = v
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
		fmt.Println(result.Value)
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

func loadSourceVars(ctx *cli.Context) (pseudo.VarsMap, error) {
	varSrc, err := url.Parse(ctx.String("scope"))
	if err == nil {
		return pseudo.LoadScopeVarsFromFile(varSrc.Path)
	}
	return nil, err
}
