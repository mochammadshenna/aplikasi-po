package main

import (
	"fmt"
	"net/url"
	"os"
	"regexp"

	"github.com/amacneil/dbmate/pkg/dbmate"
	_ "github.com/amacneil/dbmate/pkg/driver/postgres"
	"github.com/mochammadshenna/aplikasi-po/util/helper"
	"github.com/urfave/cli/v2"
	// configs "github.com/pintarnya/pintarnya-kerja-backend/configs"
	// "github.com/pintarnya/pintarnya-kerja-backend/internal/state"
	// "github.com/pintarnya/pintarnya-kerja-backend/internal/util/helper"
	// "github.com/pintarnya/pintarnya-kerja-backend/internal/util/logger"
)

func main() {
	app := cli.NewApp()
	app.Name = "dbmate"
	app.Usage = "A lightweight, framework-independent database migration tool."
	app.Version = dbmate.Version

	app.Commands = []*cli.Command{
		{
			Name:    "new",
			Aliases: []string{"n"},
			Usage:   "Generate a new migration file",
			Action: action(func(db *dbmate.DB, c *cli.Context) error {
				name := c.Args().First()
				return db.NewMigration(name)
			}),
		},
		{
			Name:  "up",
			Usage: "Create database (if necessary) and migrate to the latest version",
			Flags: []cli.Flag{
				&cli.BoolFlag{
					Name:    "verbose",
					Aliases: []string{"v"},
					EnvVars: []string{"DBMATE_VERBOSE"},
					Usage:   "print the result of each statement execution",
				},
			},
			Action: action(func(db *dbmate.DB, c *cli.Context) error {
				db.Verbose = c.Bool("verbose")
				return db.CreateAndMigrate()
			}),
		},
		{
			Name:    "rollback",
			Aliases: []string{"down"},
			Usage:   "Rollback the most recent migration",
			Flags: []cli.Flag{
				&cli.BoolFlag{
					Name:    "verbose",
					Aliases: []string{"v"},
					EnvVars: []string{"DBMATE_VERBOSE"},
					Usage:   "print the result of each statement execution",
				},
			},
			Action: action(func(db *dbmate.DB, c *cli.Context) error {
				db.Verbose = c.Bool("verbose")
				return db.Rollback()
			}),
		},
		{
			Name:  "status",
			Usage: "List applied and pending migrations",
			Flags: []cli.Flag{
				&cli.BoolFlag{
					Name:  "exit-code",
					Usage: "return 1 if there are pending migrations",
				},
				&cli.BoolFlag{
					Name:  "quiet",
					Usage: "don't output any text (implies --exit-code)",
				},
			},
			Action: action(func(db *dbmate.DB, c *cli.Context) error {
				setExitCode := c.Bool("exit-code")
				quiet := c.Bool("quiet")
				if quiet {
					setExitCode = true
				}

				pending, err := db.Status(quiet)
				if err != nil {
					return err
				}

				if pending > 0 && setExitCode {
					return cli.Exit("", 1)
				}

				return nil
			}),
		},
	}

	err := app.Run(os.Args)

	if err != nil {
		errText := redactLogString(fmt.Sprintf("Error: %s\n", err))
		_, _ = fmt.Fprint(os.Stderr, errText)
		os.Exit(2)
	}
}

func action(f func(*dbmate.DB, *cli.Context) error) cli.ActionFunc {
	return func(c *cli.Context) error {
		dbConfig := "root"
		dbPassword := "root"
		dbHost := "localhost"
		dbPort := 5432
		dbName := "arbrion"

		// var host string
		// if host = dbConfig; dbConfig == "" {
		// 	host = dbHost
		// }

		link := fmt.Sprintf(
			"%s://%s:%s@%s:%d/%s?sslmode=disable",
			"postgres",
			dbConfig,
			dbPassword,
			dbHost,
			dbPort,
			dbName,
		)

		u, err := url.Parse(link)
		helper.PanicError(err)

		db := dbmate.New(u)
		db.AutoDumpSchema = !c.Bool("no-dump-schema")
		db.SchemaFile = c.String("schema-file")
		db.MigrationsDir = "./scripts/migrations"

		return f(db, c)
	}
}

func redactLogString(in string) string {
	re := regexp.MustCompile("([a-zA-Z]+://[^:]+:)[^@]+@")

	return re.ReplaceAllString(in, "${1}********@")
}
