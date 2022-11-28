package config

type Node struct {
	Host string
	Port string
}

var Servers_LOCAL = []Node{
	{Host: "localhost:", Port: "11111"},
	{Host: "localhost:", Port: "22222"},
	{Host: "localhost:", Port: "33333"},
	{Host: "localhost:", Port: "10003"},
	// {Host: "localhost:", Port: "10004"},
	// {Host: "localhost:", Port: "10005"},
	// {Host: "localhost:", Port: "10006"},
	// {Host: "localhost:", Port: "10007"},
}

const DEFAULT_SERVER_PORT = "10000"

var Servers = []Node{
	{Host: "node1:", Port: DEFAULT_SERVER_PORT},
	{Host: "node2:", Port: DEFAULT_SERVER_PORT},
	{Host: "node3:", Port: DEFAULT_SERVER_PORT},
	{Host: "node4:", Port: DEFAULT_SERVER_PORT},
}

var N int = len(Servers)
var F int = (N - 1) / 3

// 3f+1
var HIGH_THRESHOLD = 3*F + 1

// 2f+1
var MEDIUM_THRESHOLD = 2*F + 1

// f+1
var LOW_THRESHOLD = F + 1
