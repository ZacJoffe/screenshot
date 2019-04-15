package cmd

import (
	"fmt"
	"os"

	"github.com/ZacJoffe/clipboard/xclip"
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

		var image *os.File

		if sel {
			image, err = screenshot.Select()
			if err != nil {
				return err
			}

			defer image.Close()
			/*
				err = xclip.WriteImage(image)
				if err != nil {
					return err
				}
			*/
		} else {
			image, err = screenshot.Screen()
			if err != nil {
				return err
			}

			defer image.Close()

			/*
				err = xclip.WriteImage(image)
				if err != nil {
					return err
				}
			*/
		}

		upload, err := cmd.Flags().GetBool("upload")

		if upload {
			photo, err := os.Open("/tmp/screenshot.png")
			if err != nil {
				return err
			}

			link, err := uploadImage(photo)
			if err != nil {
				return err
			}

			fmt.Println(link)
			err = xclip.WriteText(link)
			if err != nil {
				return err
			}
		} else {
			err = xclip.WriteImage(image)
			if err != nil {
				return err
			}
		}

		return nil
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
	rootCmd.Flags().BoolP("upload", "u", false, "Upload screenshot to imgur")
}
