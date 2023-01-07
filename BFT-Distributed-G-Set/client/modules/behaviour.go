package modules

import (
	"BFT-Distributed-G-Set/client"
	"BFT-Distributed-G-Set/tools"
)

func StartInteractive(c client.Client) {
	tools.Log(c.Hostname, "Started interactive session")
}

func StartAutomated(c client.Client) {

}
