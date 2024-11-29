---
to: cmd/aoc<%= year || '2024' %>_day<%= day %>/main.go
---
package main

import (
	"github.com/hreeder/aocgo/pkg/aoc"
	"github.com/hreeder/aocgo/pkg/aoc<%= year || '2024' %>"
)

func main() {
	aoc.Run(aoc<%= year || '2024' %>.Day<%= day %>Part1, <%= year || '2024' %>, <%= day %>)
    aoc.Run(aoc<%= year || '2024' %>.Day<%= day %>Part2, <%= year || '2024' %>, <%= day %>)
}
