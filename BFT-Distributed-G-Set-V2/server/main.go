package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {

	// the only thing i know is what i have to do
	// and the servers in the network
	// hostname, err := os.Hostname()
	// if err != nil {
	// 	panic(err)
	// }
	// hostname = strings.Split(hostname, ".")[0]

	// config_file := os.Args[1]
	// data, err := os.ReadFile(os.Args[1])
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(string(data))

	files, err := ioutil.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		fmt.Println(f.Name())
	}

	// all_servers := []string{}
	// for i := 0; i < N; i++ {
	// 	p := "node" + strconv.Itoa(i) + ":"
	// 	all_servers = append(all_servers, p+config.DEFAULT_PORT)
	// }
	// server.Create(hostname, all_servers)
	// fmt.Println("Hello from server", hostname)

}
