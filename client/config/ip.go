package config

var Server_ids = []string{"192.168.1.1", "192.168.1.2", "192.168.1.3"}
var Client_ids = []string{"192.168.1.4", "192.168.1.5", "192.168.1.6"}

// Client - server communication
var Client_dealer_port = 5555
var Server_router_port = 5555

// Server = server communication
var Server_dealer_starting_port = 10000

// var Server_dealer_ports = make([]int, len(Server_ips))
