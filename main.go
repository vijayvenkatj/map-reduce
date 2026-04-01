package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"

	"github.com/vijayvenkatj/map-reduce/internal"
)

var (
	idArg         = flag.Int("id", 1, "node id")
	nodeTypeArg   = flag.String("type", "master", "node type: master or worker")
	nMapArg       = flag.Int("nMap", 1, "number of map tasks")
	nReduceArg    = flag.Int("nReduce", 1, "number of reduce tasks")
	portArg       = flag.Int("port", 8080, "port number")
	masterAddrArg = flag.String("master_addr", "localhost:8080", "address of the master node")
)

var mapf internal.MapFunc
var reducef internal.ReduceFunc

func main() {

	flag.Parse()

	id := *idArg
	nodeType := *nodeTypeArg
	nMap := *nMapArg
	nReduce := *nReduceArg
	port := *portArg
	addr := fmt.Sprintf(":%d", port)
	masterAddr := *masterAddrArg

	if nodeType == "master" {

		params := internal.MasterParams{
			ID:      id,
			NMap:    nMap,
			NReduce: nReduce,
		}
		master := internal.CreateMaster(params)

		err := rpc.Register(master)
		if err != nil {
			log.Fatal("rpc register error:", err)
			return
		}
		rpc.HandleHTTP()

		listener, err := net.Listen("tcp", addr)
		if err != nil {
			log.Fatal("listen error:", err)
		}

		log.Println("listening on ", addr)
		if err = http.Serve(listener, nil); err != nil {
			log.Fatal("serve error:", err)
		}
	}

	if nodeType == "worker" {
		worker := internal.CreateWorker(id, masterAddr)
		worker.Run(mapf, reducef)
	}
}
