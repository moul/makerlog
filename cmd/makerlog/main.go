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
	"moul.io/makerlog/makerlogtypes"
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
		prettyJSON         bool
		tasksListRequest   makerlogtypes.TasksListRequest
		tasksCreateRequest makerlogtypes.TasksCreateRequest
	)
	rootFlags := flag.NewFlagSet("root", flag.ExitOnError)
	rootFlags.StringVar(&token, "token", "", "Your private API key (use the 'login' command to get one)")
	loginFlags := flag.NewFlagSet("login", flag.ExitOnError)
	loginFlags.StringVar(&username, "username", os.Getenv("USER"), "your username")
	loginFlags.StringVar(&password, "password", "", "your password")
	rawFlags := flag.NewFlagSet("raw", flag.ExitOnError)
	rawFlags.BoolVar(&prettyJSON, "pretty", false, "pretty JSON")
	tasksListFlags := flag.NewFlagSet("tasks-lists", flag.ExitOnError)
	tasksListFlags.StringVar(&tasksListRequest.User, "user", "", "filter by username")
	tasksListFlags.IntVar(&tasksListRequest.Limit, "limit", 20, "number of results to return per page")
	tasksListFlags.IntVar(&tasksListRequest.Offset, "offset", 0, "the initial index from which to return the results")
	tasksCreateFlags := flag.NewFlagSet("tasks-creates", flag.ExitOnError)
	tasksCreateFlags.StringVar(&tasksCreateRequest.Content, "content", "", "")
	tasksCreateFlags.StringVar(&tasksCreateRequest.Description, "description", "", "")
	tasksCreateFlags.BoolVar(&tasksCreateRequest.Done, "done", false, "")
	tasksCreateFlags.BoolVar(&tasksCreateRequest.InProgress, "in-progress", false, "")
	//tasksCreateFlags.TimeVar(&tasksCreateRequest.DueAt, "due-at", nil, "")

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
			}, {
				Name:       "raw",
				FlagSet:    rawFlags,
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
							if prettyJSON {
								fmt.Println(godev.PrettyJSON(ret))
							} else {
								fmt.Println(godev.JSON(ret))
							}
							return nil
						},
					}, {
						Name:      "tasks_list",
						ShortHelp: "all public tasks",
						FlagSet:   tasksListFlags,
						Exec: func(ctx context.Context, _ []string) error {
							client := makerlog.New(token)
							ret, err := client.RawTasksList(ctx, &tasksListRequest)
							if err != nil {
								return err
							}
							if prettyJSON {
								fmt.Println(godev.PrettyJSON(ret))
							} else {
								fmt.Println(godev.JSON(ret))
							}
							return nil
						},
					}, {
						Name:      "tasks_create",
						ShortHelp: "creates a new task",
						FlagSet:   tasksCreateFlags,
						Exec: func(ctx context.Context, _ []string) error {
							client := makerlog.New(token)
							ret, err := client.RawTasksCreate(ctx, &tasksCreateRequest)
							if err != nil {
								return err
							}
							if prettyJSON {
								fmt.Println(godev.PrettyJSON(ret))
							} else {
								fmt.Println(godev.JSON(ret))
							}
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
