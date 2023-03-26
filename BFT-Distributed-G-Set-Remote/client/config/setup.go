package config

import (
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
	N = 0
	F = 0
)

func Initialize(network_name string) []Node {
	working_dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	parent_dir := filepath.Dir(working_dir)
	port, threads := GetPortAndThreads(parent_dir + "/config")
	servers := GetHosts(parent_dir+"/hosts", network_name)
	server_nodes := make([]Node, 0)
	for _, s := range servers {
		for i := port; i < port+threads; i++ {
			server_nodes = append(server_nodes, Node{Host: s, Port: strconv.Itoa(i)})
		}
	}

	if err != nil {
		panic(err)
	}
	N = len(server_nodes)
	F = (N - 1) / 3

	return server_nodes
}

func NetworkExists(net_name string) bool {
	working_dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	parent_dir := filepath.Dir(working_dir)
	_, err = os.Stat(parent_dir + "/" + net_name)
	return !os.IsNotExist(err)
}

func GetHosts(filename, option string) []string {
	b, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	s := string(b)
	lines := strings.Split(s, "\n")

	var master, clients, serversNormal, serversMute, serversMalicious, serversHalfAndHalf []string
	var servers []string

	currentCategory := ""
	for _, line := range lines {
		if strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]") {
			currentCategory = line
			continue
		}

		if line == "" {
			continue
		}
		switch currentCategory {
		case "[master]":
			master = append(master, line)
		case "[clients]":
			clients = append(clients, line)
		case "[servers-normal]":
			serversNormal = append(serversNormal, line)
			servers = append(servers, line)
		case "[servers-mute]":
			serversMute = append(serversMute, line)
			servers = append(servers, line)
		case "[servers-malicious]":
			serversMalicious = append(serversMalicious, line)
			servers = append(servers, line)
		case "[servers-half_and_half]":
			serversHalfAndHalf = append(serversHalfAndHalf, line)
			servers = append(servers, line)
		}
	}

	switch option {
	case "master":
		return master
	case "clients":
		return clients
	case "servers":
		return servers
	}
	return []string{}
}

func GetPortAndThreads(filename string) (int, int) {
	b, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	s := string(b)
	lines := strings.Split(s, "\n")

	default_port := strings.Split(lines[0], "=")[1]
	threads := strings.Split(lines[1], "=")[1]

	num_threads, err := strconv.Atoi(threads)
	if err != nil {
		panic(err)
	}
	num_default_port, err := strconv.Atoi(default_port)
	if err != nil {
		panic(err)
	}

	return num_default_port, num_threads
}
