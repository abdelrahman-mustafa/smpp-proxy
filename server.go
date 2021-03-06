package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"

	"./connection"
	"./proxy"
)

// var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to `file`")

func main() {
	proxyServer := os.Getenv("PROXY_ADDRESS")
	if proxyServer == "" {
		log.Fatal("There is no defined proxy server from ENV")
	}
	listner, err := net.Listen("tcp", proxyServer)

	if err != nil {
		fmt.Println(err)
		return
	}
	defer listner.Close()

	// flag.Parse()
	// if *cpuprofile != "" {
	// 	f, err := os.Create(*cpuprofile)
	// 	if err != nil {
	// 		log.Fatal("could not create CPU profile: ", err)
	// 	}
	// 	defer f.Close() // error handling omitted for example
	// 	if err := pprof.StartCPUProfile(f); err != nil {
	// 		log.Fatal("could not start CPU profile: ", err)
	// 	}
	// 	defer pprof.StopCPUProfile()
	// }
	// defer profile.Start().Stop()
	for {
		c, err := listner.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		smppCon := connection.SmppConn{c, bufio.NewReader(c), bufio.NewWriter(c)}

		proxy := proxy.NewProxy(smppCon)
		go proxy.RunProxy()
	}

}
