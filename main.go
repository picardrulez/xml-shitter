package main

import (
	"flag"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"io/ioutil"
	"os"
)

var NAME string = "xml_shitter"
var VERSION string = "v1.0.0"

func main() {
	versionFlag := flag.Bool("v", false, "display version")
	flag.Parse()
	if *versionFlag {
		fmt.Println("xml-shitter " + VERSION)
		return
	}
	xmlFile, err := os.Open("shit.xml")
	if err != nil {
		log.Info("failed reading shit.xml file ")
		log.Info(err)
	}
	defer xmlFile.Close()
	bytes, _ := ioutil.ReadAll(xmlFile)
	fmt.Print(string(bytes))
}
