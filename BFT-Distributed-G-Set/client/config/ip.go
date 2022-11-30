package config

const _DEFAULT_SERVER_PORT = "10000"

var REMOTE_SERVERS = []Node{
	{Host: "node1:", Port: _DEFAULT_SERVER_PORT},
	{Host: "node2:", Port: _DEFAULT_SERVER_PORT},
	{Host: "node3:", Port: _DEFAULT_SERVER_PORT},
	{Host: "node4:", Port: _DEFAULT_SERVER_PORT},
}
