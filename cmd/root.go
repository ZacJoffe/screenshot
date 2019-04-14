package cmd

import (
	"fmt"
	"os"

	"github.com/ZacJoffe/screenshot-lib/screenshot"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "screenshot",
	Short: "Screenshot CLI app",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		sel, err := cmd.Flags().GetBool("select")
		if err != nil {
			return err
		}

		if sel {
			_, err = screenshot.Select()
			return err
		}
		_, err = screenshot.Screen()
		return err
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("select", "s", false, "Select region of the screen to screenshot")
}
