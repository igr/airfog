package main

import (
	"fmt"
	"foursquare.com/airfog/cmd/actions"
	"github.com/alecthomas/kong"
	"github.com/fatih/color"
)

type VersionFlag string

func (v VersionFlag) Decode(_ *kong.DecodeContext) error { return nil }
func (v VersionFlag) IsBool() bool                       { return true }
func (v VersionFlag) BeforeApply(app *kong.Kong, vars kong.Vars) error {
	fmt.Println(vars["version"])
	app.Exit(0)
	return nil
}

type Globals struct {
	Debug   bool        `short:"D" help:"Enable debug mode"`
	Version VersionFlag `name:"version" help:"Print version information and quit"`
}

type CLI struct {
	Globals
	ClearTi CleartiCmd `cmd:"" help:"Clear non-successful Task Instances"`
}

func runCli(apiCtx actions.ApiCtx) {
	cli := CLI{
		Globals: Globals{
			Version: VersionFlag("0.1.1"),
		},
	}

	green := color.New(color.FgGreen).SprintFunc()

	kongCtx := kong.Parse(&cli,
		kong.Name(green("airfog")),
		kong.Description("Common FSQ tasks on AirFlow."),
		kong.UsageOnError(),
		kong.ConfigureHelp(kong.HelpOptions{
			Compact: true,
		}),
		kong.Vars{
			"version": "0.0.1",
		})
	err := kongCtx.Run(apiCtx, &cli.Globals)
	kongCtx.FatalIfErrorf(err)
}
