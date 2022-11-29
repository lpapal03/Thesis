// WILL BECOME .config FILE
// with functions elsewhere that initializes stuff

package config

const MODE = "LOCAL"

type Node struct {
	Host string
	Port string
}

var Servers_LOCAL = []Node{
	// good
	{Host: "localhost:", Port: "10001"},
	{Host: "localhost:", Port: "10002"},
	{Host: "localhost:", Port: "10003"},
	{Host: "localhost:", Port: "10004"},
	{Host: "localhost:", Port: "10005"},
	{Host: "localhost:", Port: "10006"},
	// turning point
	{Host: "localhost:", Port: "99999"},
	// byzantine
	{Host: "localhost:", Port: "20002"},
	{Host: "localhost:", Port: "20003"},
	{Host: "localhost:", Port: "20004"},
}

const DEFAULT_SERVER_PORT = "10000"

var Servers = []Node{
	{Host: "node1:", Port: DEFAULT_SERVER_PORT},
	{Host: "node2:", Port: DEFAULT_SERVER_PORT},
	{Host: "node3:", Port: DEFAULT_SERVER_PORT},
	{Host: "node4:", Port: DEFAULT_SERVER_PORT},
}

var N int = len(Servers_LOCAL)
var F int = (N - 1) / 3

// 3f+1
var HIGH_THRESHOLD = 3*F + 1

// 2f+1
var MEDIUM_THRESHOLD = 2*F + 1

// f+1
var LOW_THRESHOLD = F + 1
