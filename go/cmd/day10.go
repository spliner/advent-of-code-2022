package cmd

import (
	"aoc2022/pkg/day10"
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(day10Cmd)
}

var day10Cmd = &cobra.Command{
	Use:   "day10 [input path] [part]",
	Short: "Day 10 solution",
	RunE: func(cmd *cobra.Command, args []string) error {
		input, err := readInput(args[0])
		if err != nil {
			return err
		}

		part := args[1]
		if part == "1" {
			result, err := day10.Part1(input)
			if err != nil {
				return err
			}

			fmt.Printf("Day 10 part 1: %s\n", result)
		} else if part == "2" {
			result, err := day10.Part2(input)
			if err != nil {
				return err
			}

			fmt.Printf("Day 10 part 2: %s\n", result)
		} else {
			return fmt.Errorf("invalid part: %s", part)
		}

		return nil
	},
}
