package day7_test

import (
	"aoc2022/pkg/day7"
	"testing"
)

func TestPart1(t *testing.T) {
	input := `$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k`

	got, err := day7.Part1(input)
	if err != nil {
		t.Fatal("expected no error, got", err)
	}

	want := "95437"
	if got != want {
		t.Fatalf("got %s, want %s", got, want)
	}
}
