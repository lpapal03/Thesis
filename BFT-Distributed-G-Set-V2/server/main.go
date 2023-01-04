package main

import (
	"fmt"
	"os"
)

func main() {

	// the only thing i know is what i have to do
	// and the servers in the network
	// hostname, err := os.Hostname()
	// if err != nil {
	// 	panic(err)
	// }
	// hostname = strings.Split(hostname, ".")[0]

	data, err := os.ReadFile("/users/loukis/Thesis/BFT-Distributed-G-Set-V2/hosts")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))

	// files, err := ioutil.ReadDir("./")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// for _, f := range files {
	// 	fmt.Println(f.Name())
	// }

	// all_servers := []string{}
	// for i := 0; i < N; i++ {
	// 	p := "node" + strconv.Itoa(i) + ":"
	// 	all_servers = append(all_servers, p+config.DEFAULT_PORT)
	// }
	// server.Create(hostname, all_servers)
	// fmt.Println("Hello from server", hostname)

}
