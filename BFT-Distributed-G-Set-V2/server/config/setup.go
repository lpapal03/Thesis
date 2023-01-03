package config

type Node struct {
	Host string
	Port string
}

var (
	N                int
	F                int
	HIGH_THRESHOLD   int
	MEDIUM_THRESHOLD int
	LOW_THRESHOLD    int
	SERVERS          []Node
)

func SetServers(mode string) []Node {

	if mode == "LOCAL" {
		N = len(LOCAL_SERVERS)
		SERVERS = LOCAL_SERVERS

	} else if mode == "REMOTE" {
		N = len(REMOTE_SERVERS)
		SERVERS = REMOTE_SERVERS
	} else {
		panic("Mode not supported!")
	}

	F = (N - 1) / 3
	// 3f+1
	HIGH_THRESHOLD = 3*F + 1
	// 2f+1
	MEDIUM_THRESHOLD = 2*F + 1
	// f+1
	LOW_THRESHOLD = F + 1

	return SERVERS

}
