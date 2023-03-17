package main

import (
	"2-Atomic-Adds/config"
	"2-Atomic-Adds/modules"
	"2-Atomic-Adds/tools"
	"os"
	"strconv"
)

func main() {
	tools.ResetLogFile()
	wd := "/users/loukis/Thesis/2-Atomic-Adds-Remote"

	// hosts are just the machine names
	sbdso := config.GetHosts(wd+"/hosts", "sbdso")
	bdso_1 := config.GetHosts(wd+"/hosts", "bdso-1")
	bdso_2 := config.GetHosts(wd+"/hosts", "bdso-2")
	default_port, num_threads := config.ParseConfigFile(wd + "/config")

	sbdso_nodes := make([]config.Node, 0)
	all_bdso := make(map[string][]config.Node)

	for _, h := range sbdso {
		for p := default_port; p < default_port+num_threads; p++ {
			p_num := strconv.Itoa(p)
			sbdso_nodes = append(sbdso_nodes, config.Node{Host: h, Port: p_num})
		}
	}
	for _, h := range bdso_1 {
		for p := default_port; p < default_port+num_threads; p++ {
			p_num := strconv.Itoa(p)
			all_bdso["bdso-1"] = append(all_bdso["bdso-1"], config.Node{Host: h, Port: p_num})
		}
	}
	for _, h := range bdso_2 {
		for p := default_port; p < default_port+num_threads; p++ {
			p_num := strconv.Itoa(p)
			all_bdso["bdso-2"] = append(all_bdso["bdso-2"], config.Node{Host: h, Port: p_num})
		}
	}

	config.N = len(sbdso)
	config.F = (config.N - 1) / 3

	if len(os.Args) < 2 {
		modules.StartNormal(sbdso_nodes, default_port, num_threads, all_bdso)
	} else {
		behaviour := os.Args[1]
		switch behaviour {
		case "normal":
			modules.StartNormal(sbdso_nodes, default_port, num_threads, all_bdso)
		case "mute":
			modules.StartMute(sbdso_nodes, default_port, num_threads, all_bdso)
		case "malicious":
			modules.StartMute(sbdso_nodes, default_port, num_threads, all_bdso)
		default:
			panic("Invalid argument")
		}
	}

	select {}

}
