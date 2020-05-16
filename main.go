package main

import (
	"fmt"
	"time"

	"github.com/wdalmijn/golang-xml-parse-bench/pkg/parser"
)

type testStruct struct {
	numWorkers int
	test       string
	resource   string
}

func main() {
	tests := []testStruct{
		testStruct{
			numWorkers: 0,
			test:       "xml_token",
			resource:   "release",
		},
		// testStruct{
		// 	numWorkers: 2,
		// 	test:       "fetch_inner_xml",
		// 	resource:   "label",
		// },
		// testStruct{
		// 	numWorkers: 4,
		// 	test:       "fetch_inner_xml",
		// 	resource:   "label",
		// },
		// testStruct{
		// 	numWorkers: 6,
		// 	test:       "fetch_inner_xml",
		// 	resource:   "label",
		// },
		// testStruct{
		// 	numWorkers: 8,
		// 	test:       "fetch_inner_xml",
		// 	resource:   "master",
		// },
		// testStruct{
		// 	numWorkers: 10,
		// 	test:       "fetch_inner_xml",
		// 	resource:   "artist",
		// },
		testStruct{
			numWorkers: 20,
			test:       "fetch_inner_xml",
			resource:   "release",
		},
		// testStruct{
		// 	numWorkers: 30,
		// 	test:       "fetch_inner_xml",
		// 	resource:   "master",
		// },
		// testStruct{
		// 	numWorkers: 40,
		// 	test:       "fetch_inner_xml",
		// 	resource:   "master",
		// },
		// testStruct{
		// 	numWorkers: 50,
		// 	test:       "fetch_inner_xml",
		// 	resource:   "master",
		// },
		// testStruct{
		// 	numWorkers: 100,
		// 	test:       "fetch_inner_xml",
		// 	resource:   "master",
		// },
		// testStruct{
		// 	numWorkers: 200,
		// 	test:       "fetch_inner_xml",
		// 	resource:   "master",
		// },
		// testStruct{
		// 	numWorkers: 500,
		// 	test:       "fetch_inner_xml",
		// 	resource:   "label",
		// },
	}
	for _, t := range tests {
		start := time.Now()
		fmt.Printf("Running test: %s, resource: %s, with %d workers\n", t.test, t.resource, t.numWorkers)
		parser.Parser(t.test, t.resource, t.numWorkers)
		fmt.Printf("Test took %s.\n", time.Since(start))
	}
}
