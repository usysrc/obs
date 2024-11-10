package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func NewConfigCmd() *cobra.Command {
	var vault string
	var targetFolder string

	cmd := &cobra.Command{
		Use:   "config",
		Short: "Configure the CLI",
		Long:  `Configure the CLI with your Obsidian vault name and other settings.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			viper.Set("vault", vault)
			viper.Set("targetFolder", targetFolder)
			configFile := filepath.Join(
				os.Getenv("HOME"),
				".config",
				"obsidian-cli",
				"config.yaml",
			)
			if err := viper.WriteConfigAs(configFile); err != nil {
				return fmt.Errorf("failed to write config: %w", err)
			}
			fmt.Printf("Configuration saved: vault = %s, targetFolder = %s\n", vault, targetFolder)
			return nil
		},
	}

	cmd.Flags().StringVarP(&vault, "vault", "v", "", "Name of the Obsidian vault")
	cmd.MarkFlagRequired("vault")

	cmd.Flags().StringVarP(&targetFolder, "targetFolder", "t", "", "Folder where new notes are created")
	cmd.MarkFlagRequired("targetFolder")

	return cmd
}
