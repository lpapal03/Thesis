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

func parseHostsFile(filename string, bdso string) ([]Node, error) {
	var nodes []Node
	var tagFound bool
	var min_port int
	var max_port int
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "["+bdso+"]") {
			tagFound = true
			continue
		}
		if tagFound {
			port_range := strings.Split(line, "-")
			min_port, err = strconv.Atoi(port_range[0])
			if err != nil {
				panic(err)
			}
			max_port, err = strconv.Atoi(port_range[1])
			if err != nil {
				panic(err)
			}
			break
		}
	}
	for p := min_port; p <= max_port; p++ {
		nodes = append(nodes, Node{Host: "localhost:", Port: strconv.Itoa(p)})
	}
	return nodes, nil
}

func getAllBdso(filename string) map[string][]Node {
	all_bdso := make(map[string][]Node)
	all_lines := make([]string, 0)
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		all_lines = append(all_lines, scanner.Text())
	}
	file.Close()

	for _, line := range all_lines {
		if strings.Contains(line, "[bdso") {
			bdso_name := line[1 : len(line)-1]
			all_bdso[bdso_name], err = parseHostsFile(filename, bdso_name)
			if err != nil {
				panic(err)
			}

		}
	}
	return all_bdso
}

func SetServerNodes() ([]Node, map[string][]Node) {

	working_dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	parent_dir := filepath.Dir(working_dir)
	bdso_name := filepath.Base(working_dir)

	SERVERS, err = parseHostsFile(parent_dir+"/hosts", bdso_name)

	BDSO := getAllBdso(parent_dir + "/hosts")
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

	return SERVERS, BDSO

}
