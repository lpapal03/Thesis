package config

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
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

func parseHostsFile(fileName string, bdso string) ([]Node, error) {
	// var nodes []Node
	var tag string
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "["+bdso+"]") {
			tag = line[1 : len(line)-1]
			fmt.Println(tag)
		}
	}
	return []Node{}, nil
}

func SetServerNodes() []Node {

	working_dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	parent_dir := filepath.Dir(working_dir)
	bdso_name := filepath.Base(working_dir)

	SERVERS, err = parseHostsFile(parent_dir+"/hosts", bdso_name)
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
