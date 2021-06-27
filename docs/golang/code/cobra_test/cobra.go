package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewRootCmd() *cobra.Command {
	var in string

	cmd := &cobra.Command{
		Use:   "hugo",
		Short: "Hugo is a very fast static site generator",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Fprintf(cmd.OutOrStdout(), in)
			return nil
		},
	}
	cmd.Flags().StringVar(&in, "in", "", "This is a very important input.")
	return cmd
}
