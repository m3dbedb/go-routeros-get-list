package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
	"github.com/go-routeros/routeros"
)

var (
	// Routeros connection parameters
	useTLS          = flag.Bool("tls", false, "use tls")
	routerIpPort    = flag.String("router-ip-port", os.Getenv("ROUTER_IP_PORT"), "mikrotik ip")
	routerUser      = flag.String("router-user", os.Getenv("ROUTER_USER"), "mikrotik username")
	routerPwd       = flag.String("router-pwd", os.Getenv("ROUTER_PWD"), "mikrotik password")
	printProperties = flag.String("print-parameters", os.Getenv("PRINT_PARAMETERS"), "Properties")

	// In app switches
	interval            = flag.Duration("interval", 1*time.Second, "Interval")

)

func dial() (*routeros.Client, error) {
	if *useTLS {
		return routeros.DialTLS(*routerIpPort, *routerUser, *routerPwd, nil)
	}

	return routeros.Dial(*routerIpPort, *routerUser, *routerPwd)
}

func main() {

	flag.Parse()
	watchOnlines()
}

func watchOnlines() {
	client, err := dial()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
		reply, runErr := client.Run("/ip/hotspot/active/print")

		if runErr != nil {
			log.Fatal(runErr)
			os.Exit(1)
		}

		for _, re := range reply.Re {
			for _, p := range strings.Split(*printProperties, ",") {
				fmt.Print(re.Map[p], "\t")
			}
		}
}
