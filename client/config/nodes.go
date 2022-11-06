package config

// Router sockets of servers
var Servers = []string{"node0"}

var Server_router_port = "10000"

var N int = len(Servers)
var f int = (N - 1) / 3

// 3f+1
var HIGH_THRESHOLD = 3*f + 1

// 2f+1
var MEDIUM_THRESHOLD = 2*f + 1

// f+1
var LOW_THRESHOLD = f + 1
