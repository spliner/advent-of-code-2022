package cmd

import (
	"aoc2022/pkg/day3"
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(day3Cmd)
}

var day3Cmd = &cobra.Command{
	Use:   "day3 [input path] [part]",
	Short: "Day 3 solution",
	RunE: func(cmd *cobra.Command, args []string) error {
		input, err := readInput(args[0])
		if err != nil {
			return err
		}

		part := args[1]
		if part == "1" {
			result, err := day3.Part1(input)
			if err != nil {
				return err
			}

			fmt.Printf("Day 3 part 1: %s\n", result)
		} else if part == "2" {
			result, err := day3.Part2(input)
			if err != nil {
				return err
			}

			fmt.Printf("Day 3 part 2: %s\n", result)
		} else {
			return fmt.Errorf("invalid part: %s", part)
		}

		return nil
	},
}
