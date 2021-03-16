package main

import (
	"github.com/spf13/cobra"
)

// AddFileFlag ...
func AddFileFlag(cmd *cobra.Command) {
	cmd.Flags().StringP("input_file", "i", "./input.csv", "Raw phone number csv file.")
	cmd.Flags().StringP("output_file", "o", "./output.csv", "Encrypted phone number csv file.")
}
