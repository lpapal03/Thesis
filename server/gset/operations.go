package gset

import (
	"server/messaging"
	"server/tools"

	zmq "github.com/pebbe/zmq4"
)

func HandleGet(sender_id, my_id string, inbound_socket zmq.Socket, mygset map[string]string) {
	response := []string{sender_id, my_id, messaging.GET_RESPONSE, GsetToString(mygset, false)}
	inbound_socket.SendMessage(response)
	tools.Log(my_id, messaging.GET_RESPONSE+" to "+sender_id)
}

func HandleAdd(id string, mygset map[string]string, record string, servers []*zmq.Socket) {

	if Exists(mygset, record) {
		return
	}

	if !Exists(mygset, record) {
		Append(mygset, record)
		tools.Log(id, "Added {"+record+"} to local G-Set")
	} else {
		tools.Log(id, "Record {"+record+"} already exists in local G-Set")
	}

}
