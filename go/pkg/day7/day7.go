package day7

import (
	"aoc2022/pkg/day7/model"
	"fmt"
	"strconv"
	"strings"
)

func Part1(input string) (string, error) {
	lines := strings.Split(input, "\n")
	directories := make([]*model.Directory, 0)
	rootDirectory := model.NewDirectory("/", nil)
	currentDirectory := rootDirectory
	i := 1
	for i < len(lines) {
		line := lines[i]
		if line == "$ ls" {
			for {
				i++
				if i == len(lines) {
					break
				}
				line = lines[i]
				if strings.HasPrefix(line, "$") {
					break
				}

				split := strings.Split(line, " ")
				if split[0] == "dir" {
					directory := model.NewDirectory(split[1], currentDirectory)
					directories = append(directories, directory)
					currentDirectory.AddChild(directory)
				} else {
					size, err := strconv.Atoi(split[0])
					if err != nil {
						return "", err
					}
					file := model.NewFile(split[1], size)
					currentDirectory.AddFile(file)
				}
			}
		} else {
			split := strings.Split(line, " ")
			argument := split[2]
			if argument == ".." {
				currentDirectory = currentDirectory.Parent
			} else {
				name := split[2]
				child, ok := currentDirectory.FindChild(name)
				if !ok {
					return "", fmt.Errorf("could not find directory %s", argument)
				}
				currentDirectory = child
			}
			i++
		}
	}

	var totalSize int
	for _, directory := range directories {
		directorySize := directory.Size()
		if directorySize <= 100000 {
			totalSize += directorySize
		}
	}

	result := strconv.Itoa(totalSize)
	return result, nil
}

func Part2(input string) (string, error) {
	return "", nil
}
