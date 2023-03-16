package config

import (
	"bufio"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type Node struct {
	Host string
	Port string
}

var (
	N                int
	F                int
	HIGH_THRESHOLD   int
	MEDIUM_THRESHOLD int
	LOW_THRESHOLD    int
	SERVERS          []Node
)

func parseHostsFile(fileName string) ([]Node, error) {
	var nodes []Node
	var min_port int
	var max_port int
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "[") {
			continue
		}
		port_range := strings.Split(line, "-")
		min_port, err = strconv.Atoi(port_range[0])
		if err != nil {
			panic(err)
		}
		max_port, err = strconv.Atoi(port_range[1])
		if err != nil {
			panic(err)
		}
	}
	for p := min_port; p <= max_port; p++ {
		nodes = append(nodes, Node{Host: "localhost:", Port: strconv.Itoa(p)})
	}
	return nodes, nil
}

func SetServerNodes() []Node {

	working_dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	parent_dir := filepath.Dir(working_dir)

	SERVERS, err = parseHostsFile(parent_dir + "/hosts")
	if err != nil {
		panic(err)
	}

	N = len(SERVERS)

	F = (N - 1) / 3
	// 3f+1
	HIGH_THRESHOLD = 3*F + 1
	// 2f+1
	MEDIUM_THRESHOLD = 2*F + 1
	// f+1
	LOW_THRESHOLD = F + 1

	return SERVERS

}
