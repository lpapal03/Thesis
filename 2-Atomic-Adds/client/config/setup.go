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
	N = 0
	F = 0
)

func NetworkExists(filename, net_name string) bool {
	working_dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	parent_dir := filepath.Dir(working_dir)
	file, err := os.Open(parent_dir + "/" + filename)
	if err != nil {
		return false
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "["+net_name+"]") {
			return true
		}
	}
	return false
}

func GetHosts(filename, option string) []string {
	b, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	s := string(b)
	lines := strings.Split(s, "\n")

	var (
		master           []string
		clientsAutomated []string
		sbdsoNormal      []string
		sbdsoMute        []string
		sbdsoMalicious   []string
		bdso1Normal      []string
		bdso1Mute        []string
		bdso1Malicious   []string
		bdso2Normal      []string
		bdso2Mute        []string
		bdso2Malicious   []string
		sbdsoServers     []string
		bdso1Servers     []string
		bdso2Servers     []string
	)

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
		case "[clients-automated]":
			clientsAutomated = append(clientsAutomated, line)
		case "[sbdso-normal]":
			sbdsoNormal = append(sbdsoNormal, line)
			sbdsoServers = append(sbdsoServers, line)
		case "[sbdso-mute]":
			sbdsoMute = append(sbdsoMute, line)
			sbdsoServers = append(sbdsoServers, line)
		case "[sbdso-malicious]":
			sbdsoMalicious = append(sbdsoMalicious, line)
			sbdsoServers = append(sbdsoServers, line)
		case "[bdso-1-normal]":
			bdso1Normal = append(bdso1Normal, line)
			bdso1Servers = append(bdso1Servers, line)
		case "[bdso-1-mute]":
			bdso1Mute = append(bdso1Mute, line)
			bdso1Servers = append(bdso1Servers, line)
		case "[bdso-1-malicious]":
			bdso1Malicious = append(bdso1Malicious, line)
			bdso1Servers = append(bdso1Servers, line)
		case "[bdso-2-normal]":
			bdso2Normal = append(bdso2Normal, line)
			bdso2Servers = append(bdso2Servers, line)
		case "[bdso-2-mute]":
			bdso2Mute = append(bdso2Mute, line)
			bdso2Servers = append(bdso2Servers, line)
		case "[bdso-2-malicious]":
			bdso2Malicious = append(bdso2Malicious, line)
			bdso2Servers = append(bdso2Servers, line)
		}
	}

	switch option {
	case "master":
		return master
	case "clients":
		return clientsAutomated
	case "sbdso":
		return sbdsoServers
	case "bdso-1":
		return bdso1Servers
	case "bdso-2":
		return bdso2Servers
	case "sbdso-normal":
		return sbdsoNormal
	case "sbdso-mute":
		return sbdsoMute
	case "sbdso-malicious":
		return sbdsoMalicious
	case "bdso1-normal":
		return bdso1Normal
	case "bdso1-mute":
		return bdso1Mute
	case "bdso1-malicious":
		return bdso1Malicious
	case "bdso2-normal":
		return bdso2Normal
	case "bdso2-mute":
		return bdso2Mute
	case "bdso2-malicious":
		return bdso2Malicious
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
