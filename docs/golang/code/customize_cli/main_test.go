package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func ExampleScript() {
	NewScript := func(action func(c *Context) error) *Script {
		return &Script{
			Before: func(c *Context) error {
				c.valueStore["test"] = "test"
				return nil
			},
			Action: action,
		}
	}

	cmd := &cobra.Command{
		RunE: func(cmd *cobra.Command, args []string) error {
			NewScript(func(c *Context) error {
				fmt.Fprintf(cmd.OutOrStdout(), c.valueStore["test"].(string))
				return nil
			}).Run()
			return nil
		},
	}
	cmd.Execute()

	// Output:
	// test
}
