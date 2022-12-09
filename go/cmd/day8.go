package cmd

import (
	"aoc2022/pkg/day8"
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(day8Cmd)
}

var day8Cmd = &cobra.Command{
	Use:   "day8 [input path] [part]",
	Short: "Day 8 solution",
	RunE: func(cmd *cobra.Command, args []string) error {
		input, err := readInput(args[0])
		if err != nil {
			return err
		}

		part := args[1]
		if part == "1" {
			result, err := day8.Part1(input)
			if err != nil {
				return err
			}

			fmt.Printf("Day 8 part 1: %s\n", result)
		} else if part == "2" {
			result, err := day8.Part2(input)
			if err != nil {
				return err
			}

			fmt.Printf("Day 8 part 2: %s\n", result)
		} else {
			return fmt.Errorf("invalid part: %s", part)
		}

		return nil
	},
}
