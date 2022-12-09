package cmd

import (
	"aoc2022/pkg/day9"
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(day9Cmd)
}

var day9Cmd = &cobra.Command{
	Use:   "day9 [input path] [part]",
	Short: "Day 9 solution",
	RunE: func(cmd *cobra.Command, args []string) error {
		input, err := readInput(args[0])
		if err != nil {
			return err
		}

		part := args[1]
		if part == "1" {
			result, err := day9.Part1(input)
			if err != nil {
				return err
			}

			fmt.Printf("Day 9 part 1: %s\n", result)
		} else if part == "2" {
			result, err := day9.Part2(input)
			if err != nil {
				return err
			}

			fmt.Printf("Day 9 part 2: %s\n", result)
		} else {
			return fmt.Errorf("invalid part: %s", part)
		}

		return nil
	},
}
