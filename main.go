package main

import (
	"fmt"
	"log"
	"os"

	"github.com/AYehia0/lezer/server"
	"github.com/AYehia0/lezer/utils"
)

var Addr = "0.0.0.0"    // localhost
var Port = "8069"       // the port the server is running on
var IgnoreHidden = true // ignore the hidden files/dir

func main() {
	// TODO: this is a basic arg parsing, use something better
	args := os.Args
	if len(args) != 2 {
		log.Fatalf("Serve Path is required!")
	}

	serveDir := args[1]
	if !utils.IsValidPath(serveDir) {
		log.Fatalf("Serve Path is not a vaild path!")
	}
	log.Printf("Server is serving files in : %s\n", serveDir)

	server := server.NewServer(Addr, Port)

	// configs
	server.ServerDir = serveDir
	server.Ip = fmt.Sprintf("%s:%s/", utils.GetMachineIp(), Port)
	server.IgnoreHidden = IgnoreHidden

	// show the QR to scan
	utils.GenerateQRCode(server.Ip)

	server.StartServer()
}
