package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/alecthomas/kingpin"
	"github.com/google/go-github/github"
)

var (
	listCommand = kingpin.Command("list", "Show all skeletons")
)

func init() {
	listCommand.Action(listAction)
}

func listAction(ctx *kingpin.ParseContext) error {
	repos, _, err := github.NewClient(nil).Repositories.ListByOrg(context.Background(), "goske", nil)
	if err != nil {
		return err
	}
	for _, repo := range repos {
		if !strings.HasPrefix(*repo.Name, "goske-") {
			continue
		}
		var desc string
		if repo.Description != nil {
			desc = *repo.Description
		}
		fmt.Printf("%s: %s\n", stripPrefix(*repo.Name), desc)
	}
	return nil
}
