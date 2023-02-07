package config

import (
	"os"
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
