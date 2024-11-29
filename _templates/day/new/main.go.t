---
to: cmd/aoc<%= locals.year || '2024' %>_day<%= day %>/main.go
---
package main

import (
	"github.com/hreeder/aocgo/pkg/aoc"
	"github.com/hreeder/aocgo/pkg/aoc<%= locals.year || '2024' %>"
)

func main() {
	aoc.Run(aoc<%= locals.year || '2024' %>.Day<%= day %>Part1, <%= locals.year || '2024' %>, <%= day %>)
    aoc.Run(aoc<%= locals.year || '2024' %>.Day<%= day %>Part2, <%= locals.year || '2024' %>, <%= day %>)
}
