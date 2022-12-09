package day7

import (
	"aoc2022/pkg/day7/model"
	"fmt"
	"strconv"
	"strings"
)

func Part1(input string) (string, error) {
	directories, _, err := parseInput(input)
	if err != nil {
		return "", err
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

func parseInput(input string) ([]*model.Directory, *model.Directory, error) {
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
						return nil, nil, err
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
					return nil, nil, fmt.Errorf("could not find directory %s", argument)
				}
				currentDirectory = child
			}
			i++
		}
	}

	return directories, rootDirectory, nil
}

func Part2(input string) (string, error) {
	directories, root, err := parseInput(input)
	if err != nil {
		return "", err
	}

	diskSpace := 70000000
	usedSpace := root.Size()
	targetSpace := 30000000
	candidateSize := usedSpace
	for _, dir := range directories[1:] {
		dirSize := dir.Size()
		if diskSpace-usedSpace+dirSize < targetSpace {
			continue
		}

		if dirSize < candidateSize {
			candidateSize = dirSize
		}
	}

	result := strconv.Itoa(candidateSize)
	return result, nil
}
