package day23

import (
	"advent-of-code-2024/util"
	"fmt"
	"slices"
	"strings"
)

func SolvePart1(useRealInput bool) (int, error) {
	g, err := parseInput(useRealInput)
	if err != nil {
		return 0, err
	}

	k3SubGraphs := make([][3]string, 0)
	exhaustedVertices := make(map[string]struct{})
	for u, uNeighbors := range g.neighbors {
		for _, v := range uNeighbors {
			if _, exhausted := exhaustedVertices[v]; exhausted {
				continue // skip vertex that was already checked in top level loop
			}

			vNeighbors := g.neighbors[v]
			for _, w := range vNeighbors {
				if w == u {
					continue
				}
				if _, exhausted := exhaustedVertices[w]; exhausted {
					continue // skip vertex that was already checked in top level loop
				}

				// Disabled below, instead we just divede by 2 at the end
				// if slices.ContainsFunc(k3SubGraphs, func(subgraph [3]string) bool {
				// 	return subgraph[0] == u && subgraph[1] == w && subgraph[2] == v
				// }) {
				// 	continue  // skip if we already saw this in reverse order
				// }

				wNeighbors := g.neighbors[w]
				if slices.Contains(wNeighbors, u) {
					k3SubGraphs = append(k3SubGraphs, [3]string{u, v, w})
				}
			}
		}
		exhaustedVertices[u] = struct{}{}
	}

	count := 0
	for _, h := range k3SubGraphs {
		for _, v := range h {
			if v[0] == 't' {
				count++
				break
			}
		}
	}
	count = count / 2 // Since we count loops in both directions

	return count, nil
}

func SolvePart2(useRealInput bool) (string, error) {
	g, err := parseInput(useRealInput)
	if err != nil {
		return "", err
	}

	// Me trying to be smart was basically same as Bron Kerbosch (but worse), so we use that instead
	maxClique := bronKerbosch(*g)

	slices.Sort(maxClique)
	password := maxClique[0]
	for _, v := range maxClique[1:] {
		password = fmt.Sprintf("%s,%s", password, v)
	}
	return password, nil
}

type graph struct {
	vertices  map[string]struct{}
	neighbors map[string][]string
}

func bronKerbosch(g graph) []string {
	type bkInput struct{ r, p, x []string }

	r := make([]string, 0)
	x := make([]string, 0)
	p := make([]string, 0, len(g.vertices))
	for v := range g.vertices {
		p = append(p, v)
	}

	// output of algorithm, all (?) cliques
	var cliques [][]string

	// Iteration implementation of Bron-Kerbosch, see https://en.wikipedia.org/wiki/Bron%E2%80%93Kerbosch_algorithm#Without_pivoting
	queue := make([]bkInput, 1)
	queue[0] = bkInput{r: r, p: p, x: x}
	for len(queue) > 0 {
		// Pop from back (recursion)
		current := queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		r, p, x = current.r, current.p, current.x

		if len(p) == 0 && len(x) == 0 {
			cliques = append(cliques, r)
			continue
		}

		// Backwards loop so that we can simply remove the entry from the slice
		for i := len(p) - 1; i >= 0; i-- {
			v := p[i]

			// Union of r and {v}
			newR := make([]string, len(r)+1)
			copy(newR, r)
			newR[len(r)] = v

			// Intersect p and x with the neighbors of v
			newP := make([]string, 0)
			newX := make([]string, 0)
			neighbors := g.neighbors[v]
			for _, neighbor := range neighbors {
				if slices.Contains(p, neighbor) {
					newP = append(newP, neighbor)
				}
				if slices.Contains(x, neighbor) {
					newX = append(newX, neighbor)
				}
			}

			// "Recursion" step, add to end of queue
			queue = append(queue, bkInput{r: newR, p: newP, x: newX})

			// Exclude v from consideration in future cliques
			p = p[:len(p)-1]
			x = append(x, v)
		}

	}

	max := cliques[0]
	for _, clique := range cliques {
		if len(clique) > len(max) {
			max = clique
		}
	}

	return max
}

func parseInput(useRealInput bool) (*graph, error) {
	data, err := util.ReadInputMulti(23, useRealInput)
	if err != nil {
		return nil, err
	}
	if len(data) != 1 {
		return nil, fmt.Errorf("expected single line of input")
	}

	vertices := make(map[string]struct{})
	neighbors := make(map[string][]string)

	for _, line := range data[0] {
		readNodes := strings.Split(line, "-")
		a, b := readNodes[0], readNodes[1]
		vertices[a] = struct{}{}
		vertices[b] = struct{}{}
		neighbors[a] = append(neighbors[a], b)
		neighbors[b] = append(neighbors[b], a)
	}
	return &graph{vertices: vertices, neighbors: neighbors}, nil
}
