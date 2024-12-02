---
to: pkg/aoc<%= locals.year || '2024' %>/day<%= day %>.go
---
package aoc<%= locals.year || '2024' %>

import "fmt"

func Day<%= day %>Part1(input string) {
    fmt.Printf("Day <%= day %> Part 1: \n")
}

func Day<%= day %>Part2(input string) {
    fmt.Printf("Day <%= day %> Part 2: \n")
}
