package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	ff "github.com/peterbourgon/ff/v3"
	"github.com/peterbourgon/ff/v3/ffcli"
	"moul.io/banner"
	"moul.io/godev"
	"moul.io/makerlog"
	"moul.io/motd"
)

func main() {
	if err := run(os.Args); err != nil {
		log.Fatalf("error: %v", err)
		os.Exit(1)
	}
}

// nolint:unparam
func run(args []string) error {
	var (
		token              string
		username, password string
	)
	rootFlags := flag.NewFlagSet("root", flag.ExitOnError)
	rootFlags.StringVar(&token, "token", "", "Your private API key (use the 'login' command to get one)")
	loginFlags := flag.NewFlagSet("login", flag.ExitOnError)
	loginFlags.StringVar(&username, "username", os.Getenv("USER"), "your username")
	loginFlags.StringVar(&password, "password", "", "your password")

	root := &ffcli.Command{
		Name:       "makerlog",
		ShortUsage: "makerlog <subcommand>",
		FlagSet:    rootFlags,
		Options:    []ff.Option{ff.WithEnvVarPrefix("MAKERLOG")},
		Subcommands: []*ffcli.Command{
			{
				Name:      "login",
				ShortHelp: "get a --token by giving --username and --password",
				FlagSet:   loginFlags,
				Options:   []ff.Option{ff.WithEnvVarPrefix("MAKERLOG")},
				Exec: func(ctx context.Context, _ []string) error {
					token, err := makerlog.Login(username, password)
					if err != nil {
						return err
					}
					fmt.Println(token)
					return nil
				},
			},
			{
				Name:       "raw",
				ShortHelp:  "raw API calls",
				ShortUsage: "makerlog raw <subcommand>",
				Exec: func(_ context.Context, _ []string) error {
					fmt.Fprintln(os.Stderr, banner.Inline("makerlog - raw"))
					return flag.ErrHelp
				},
				Subcommands: []*ffcli.Command{
					{
						Name:      "notifications_list",
						ShortHelp: "all read notifications in past 24hrs and all unread",
						Exec: func(ctx context.Context, _ []string) error {
							client := makerlog.New(token)
							ret, err := client.RawNotificationsList(ctx)
							if err != nil {
								return err
							}
							// FIXME: flag for pretty
							fmt.Println(godev.JSON(ret))
							return nil
						},
					},
				},
			},
		},
		Exec: func(_ context.Context, _ []string) error {
			fmt.Fprintln(os.Stderr, motd.Default())
			return flag.ErrHelp
		},
	}

	return root.ParseAndRun(context.Background(), os.Args[1:])
}
