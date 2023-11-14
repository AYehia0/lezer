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
var DownloadLocation = "/home/none/test/"

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
	server := server.NewServer(Addr, Port)

	// configs
	server.ServerDir = serveDir
	server.Config.BrowseRoute = "browse"
	server.Ip = fmt.Sprintf("%s:%s/", utils.GetMachineIp(), Port)
	server.DownloadLocation = DownloadLocation
	server.IgnoreHidden = IgnoreHidden

	// show the QR to scan
	fmt.Println()
	utils.GenerateQRCode(server.Ip)
	fmt.Printf("\n\nShare Files : http://%s\n", server.Ip)

	server.StartServer()
}
