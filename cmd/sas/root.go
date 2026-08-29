package main

import "github.com/spf13/cobra"

func newRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sas",
		Short: "Systems Architecture Spec — a machine-readable semantic model for systems architecture",
	}
	cmd.AddCommand(newValidateCmd())
	cmd.AddCommand(newViewCmd())
	cmd.AddCommand(newBindCmd())
	cmd.AddCommand(newExportCmd())
	return cmd
}
