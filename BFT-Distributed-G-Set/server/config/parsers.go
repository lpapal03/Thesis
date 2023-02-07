package config

import (
	"os"
	"strconv"
	"strings"
)

func GetHosts(filename, option string) []string {
	b, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	s := string(b)
	lines := strings.Split(s, "\n")

	var master, clients, serversNormal, serversMute, serversMalicious []string
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
