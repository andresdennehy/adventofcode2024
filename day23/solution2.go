package main

import (
	"adventofcode2024/utils"
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func bronKerbosch(
	graph map[string]map[string]struct{},
	R, P, X []string,
	maxClique *[]string,
) {
	if len(P) == 0 && len(X) == 0 {
		// Found a maximal clique
		if len(R) > len(*maxClique) {
			*maxClique = append([]string(nil), R...) // Copy the clique
		}
		return
	}

	pivot := selectPivot(graph, append(P, X...))
	for _, v := range difference(P, graph[pivot]) {
		bronKerbosch(
			graph,
			append(R, v),
			intersection(P, graph[v]),
			intersection(X, graph[v]),
			maxClique,
		)
		P = difference(P, map[string]struct{}{v: {}})
		X = append(X, v)
	}
}

func selectPivot(graph map[string]map[string]struct{}, nodes []string) string {
	var pivot string
	maxDegree := -1
	for _, node := range nodes {
		degree := len(graph[node])
		if degree > maxDegree {
			maxDegree = degree
			pivot = node
		}
	}
	return pivot
}

func intersection(a []string, b map[string]struct{}) []string {
	var result []string
	for _, item := range a {
		if _, ok := b[item]; ok {
			result = append(result, item)
		}
	}
	return result
}

func difference(a []string, b map[string]struct{}) []string {
	var result []string
	for _, item := range a {
		if _, ok := b[item]; !ok {
			result = append(result, item)
		}
	}
	return result
}

func largestClique(graph map[string]map[string]struct{}) []string {
	nodes := make([]string, 0, len(graph))
	for node := range graph {
		nodes = append(nodes, node)
	}

	var maxClique []string
	bronKerbosch(graph, []string{}, nodes, []string{}, &maxClique)
	slices.Sort(maxClique)
	return maxClique
}

func main() {

	file, err := os.Open("input.txt")
	utils.Check(err)
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic("couldn't close file!")
		}
	}(file)

	scanner := bufio.NewScanner(file)
	var computers = make(map[string]map[string]struct{})
	for scanner.Scan() {
		var splitLine = strings.Split(scanner.Text(), "-")
		var computerFrom, computerTo = splitLine[0], splitLine[1]
		if computers[computerFrom] == nil {
			computers[computerFrom] = make(map[string]struct{})
		}
		computers[computerFrom][computerTo] = struct{}{}
		if computers[computerTo] == nil {
			computers[computerTo] = make(map[string]struct{})
		}
		computers[computerTo][computerFrom] = struct{}{}
	}

	var password = largestClique(computers)
	fmt.Printf("Result is %v\n", strings.Join(password, ","))

}
