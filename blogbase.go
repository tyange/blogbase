package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

type Config struct {
	ProjectURL string `json:"project_url"`
	APIKey     string `json:"api_key"`
}

func main() {
	app := &cli.App{
		Name:  "blogbase",
		Usage: "Blog post management CLI tool",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "url",
				Aliases:  []string{"u"},
				Usage:    "Supabase project URL",
				EnvVars:  []string{"SUPABASE_URL"},
				Required: false,
			},
			&cli.StringFlag{
				Name:     "key",
				Aliases:  []string{"k"},
				Usage:    "Supabase API key",
				EnvVars:  []string{"SUPABASE_KEY"},
				Required: false,
			},
		},
		Commands: []*cli.Command{
			{
				Name:  "configure",
				Usage: "Configure Supabase credentials",
				Action: func(c *cli.Context) error {
					return configureCredentials(c)
				},
			},
			{
				Name:  "test",
				Usage: "Test Supabase connection",
				Action: func(c *cli.Context) error {
					fmt.Println("Test")
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func configureCredentials(c *cli.Context) error {
	config := Config{
		ProjectURL: c.String("url"),
		APIKey:     c.String("key"),
	}

	if config.ProjectURL == "" {
		fmt.Print("Enter Supabase Project URL: ")
		fmt.Scanln(&config.ProjectURL)
	}
	if config.APIKey == "" {
		fmt.Print("Enter Supabase API Key: ")
		fmt.Scanln(&config.APIKey)
	}

	fmt.Println(config)

	return nil
}
