package main

import "github.com/spf13/cobra"

func newRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sas",
		Short: "Systems Architecture Spec — a machine-readable semantic model for systems architecture",
	}
	cmd.AddCommand(newValidateCmd())
	return cmd
}
