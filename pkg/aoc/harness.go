package aoc

import (
	"fmt"
	"os"
	"strings"
)

type RunFn func(input string)

func Run(fn RunFn, year int, identifier int) {
	data, err := os.ReadFile(fmt.Sprintf("data/%d/day%d.txt", year, identifier))
	if err != nil {
		panic(err)
	}

	fn(strings.TrimSpace(string(data)))
}
