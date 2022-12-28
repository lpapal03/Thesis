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

const SCENARIO_LOCAL = "SCENARIO_LOCAL"
const SCENARIO_REMOTE = "SCENARIO_REMOTE"

const MODE_NORMAL = "MODE_NORMAL"
const MODE_MUTES = "MODE_MUTE"
const MODE_HALF_AND_HALF = "MODE_HALF_AND_HALF"

func CreateScenario(scenario, mode string) {

	if mode == "LOCAL" {
		N = len(LOCAL_SERVERS)
		SERVERS = LOCAL_SERVERS

	} else if mode == "REMOTE" {
		N = len(REMOTE_SERVERS)
		SERVERS = REMOTE_SERVERS
	} else {
		panic("Mode not supported!")
	}

	if scenario == "NORMAL" {

	} else {
		panic("Scenario not supported!")
	}

	F = (N - 1) / 3
	// 3f+1
	HIGH_THRESHOLD = 3*F + 1
	// 2f+1
	MEDIUM_THRESHOLD = 2*F + 1
	// f+1
	LOW_THRESHOLD = F + 1

}
