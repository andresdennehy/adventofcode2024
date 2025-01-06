package main

import (
	"adventofcode2024/utils"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func generateSets(computerGraph map[string]map[string]struct{}) [][]string {
	var nodes []string
	for computer, _ := range computerGraph {
		nodes = append(nodes, computer)
	}

	var triples [][]string
	for i := 0; i < len(nodes)-2; i++ {
		for j := i + 1; j < len(nodes)-1; j++ {
			for k := j + 1; k < len(nodes); k++ {
				var a, b, c = nodes[i], nodes[j], nodes[k]
				_, ok1 := computerGraph[a][b]
				_, ok2 := computerGraph[b][c]
				_, ok3 := computerGraph[c][a]
				var triple = []string{a, b, c}
				if ok1 && ok2 && ok3 {
					triples = append(triples, triple)
				}
			}
		}
	}

	return triples
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

	var triples = generateSets(computers)
	var result int
	for _, triple := range triples {
		fmt.Println(triple)
		for _, comp := range triple {
			if strings.HasPrefix(comp, "t") {
				result++
				break
			}
		}
	}

	fmt.Printf("Result is %v\n", result)

}
