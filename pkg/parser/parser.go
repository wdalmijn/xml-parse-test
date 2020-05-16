package parser

import (
	"fmt"
	"log"
	"os"
)

func Parser(algorithm string, resource string, numWorkers int) {
	// Open our xmlFile
	path := fmt.Sprintf("discogs_%ss.xml", resource)
	xmlFile, err := os.Open(path)

	defer xmlFile.Close()

	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatalln(err)
	}

	// startTime := time.Now()
	var data resourceStruct

	switch algorithm {
	case "xml_token":
		data = xmlToken(xmlFile, resource)
	case "fetch_inner_xml":
		data = fetchInnerXML(xmlFile, resource, numWorkers)
	default:
		fmt.Printf("Unsupported algorithm %s provided.\n", algorithm)
	}

	if resource == "label" {
		fmt.Printf("%d decoded\n", len(data.labels))
	}
	if resource == "artist" {
		fmt.Printf("%d decoded\n", len(data.artists))
	}
	if resource == "master" {
		fmt.Printf("%d decoded\n", len(data.masters))
	}

	// elapsed := time.Since(startTime)
	// fmt.Printf("Parsing took %s\n", elapsed)

}
