package parser

import (
	"bufio"
	"encoding/xml"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
)

func xmlToken(xmlFile *os.File, resource string) resourceStruct {
	decoder := xml.NewDecoder(xmlFile)

	data := resourceStruct{}
	data.artists = make([]Artist, 0)
	data.labels = make([]Label, 0)
	data.masters = make([]Master, 0)

	counter := 0

	var resourceFunc func(decoder *xml.Decoder, se xml.StartElement)

	switch resource {
	case "artist":
		resourceFunc = func(decoder *xml.Decoder, se xml.StartElement) {
			var a Artist
			// decode a whole chunk of following XML into the
			// variable l which is a Label (se above)
			decoder.DecodeElement(&a, &se)

			counter++

			// data.artists = append(data.artists, a)
		}
	case "label":
		resourceFunc = func(decoder *xml.Decoder, se xml.StartElement) {
			var l Label

			// decode a whole chunk of following XML into the
			// variable l which is a Label (se above)
			decoder.DecodeElement(&l, &se)

			counter++

			// data.labels = append(data.labels, l)
		}
	case "master":
		resourceFunc = func(decoder *xml.Decoder, se xml.StartElement) {
			var m Master

			// decode a whole chunk of following XML into the
			// variable l which is a Label (se above)
			decoder.DecodeElement(&m, &se)

			counter++

			// data.masters = append(data.masters, m)
		}
	case "release":
		resourceFunc = func(decoder *xml.Decoder, se xml.StartElement) {
			var r Release

			// decode a whole chunk of following XML into the
			// variable l which is a Label (se above)
			decoder.DecodeElement(&r, &se)

			counter++

			// data.releases = append(data.releases, r)
		}
	}

	for {
		// Read tokens from the XML document in a stream.
		t, _ := decoder.Token()
		if t == nil {
			break
		}
		// Inspect the type of the token just read.
		switch se := t.(type) {
		case xml.StartElement:
			// If we just read a StartElement token
			if se.Name.Local == resource {
				resourceFunc(decoder, se)

				if counter%100000 == 0 {
					fmt.Printf("Decoded %d items.\n", counter)
				}
			}
		}
	}

	fmt.Printf("Decoded amount: %d\n", counter)

	return data
}

// func fanIn(
// 	done <-chan interface{},
// 	channels ...chan interface{},
// ) chan Label {
// 	var wg sync.WaitGroup
// 	multiplexedStream := make(chan Label)

// 	multiplex := func(c <-chan interface{}) {
// 		defer wg.Done()
// 		for i := range c {
// 			select {
// 			case <-done:
// 				return
// 			case multiplexedStream <- i.(Label):
// 				fmt.Println("bye")
// 			}
// 		}
// 	}

// 	wg.Add(len(channels))
// 	for _, c := range channels {
// 		go multiplex(c)
// 	}

// 	go func() {
// 		wg.Wait()
// 		close(multiplexedStream)
// 	}()

// 	return multiplexedStream
// }

func SpawnWorkerA(dataChan chan string, wg *sync.WaitGroup) chan interface{} {
	workerChan := make(chan interface{})

	go func() {
		for d := range dataChan {
			var a Artist
			err := xml.Unmarshal([]byte(d), &a)
			if err != nil {
				log.Fatal(err)
			}
			workerChan <- a
		}
		wg.Done()
		close(workerChan)
	}()

	return workerChan
}

func SpawnWorkerL(dataChan chan string, wg *sync.WaitGroup) chan interface{} {
	workerChan := make(chan interface{})

	go func() {
		for d := range dataChan {
			var l Label
			err := xml.Unmarshal([]byte(d), &l)
			if err != nil {
				log.Fatal(err)
			}
			workerChan <- l
		}
		wg.Done()
		close(workerChan)
	}()

	return workerChan
}

func SpawnWorkerM(dataChan chan string, wg *sync.WaitGroup) chan interface{} {
	workerChan := make(chan interface{})

	go func() {
		for d := range dataChan {
			var m Master
			err := xml.Unmarshal([]byte(d), &m)
			if err != nil {
				log.Fatal(err)
			}
			workerChan <- m
		}
		wg.Done()
		close(workerChan)
	}()

	return workerChan
}

func SpawnWorkerR(dataChan chan string, wg *sync.WaitGroup) chan interface{} {
	workerChan := make(chan interface{})

	go func() {
		for d := range dataChan {
			var r Release
			err := xml.Unmarshal([]byte(d), &r)
			if err != nil {
				log.Fatal(err)
			}
			workerChan <- r
		}
		wg.Done()
		close(workerChan)
	}()

	return workerChan
}

func scanLineForXMLTag(
	tagContent string,
	openTag string,
	closeTag string,
	lineContent string,
	dataChan chan string,
) string {
	startIndex := 0
	isOpen := len(tagContent) > 0

	/**
	 * If we don't have tagContent from a previous line yet we should
	 * check for the opening tag, and if found set isOpen to true.
	 */
	if !isOpen && strings.Contains(lineContent, openTag) {
		startIndex = strings.Index(lineContent, openTag)
		isOpen = true
	}

	/**
	 * If the tag is open, look for a closing tag. If it is found, find
	 * the last occurence of it on the provided line. Combined the tag
	 * content into a full tag and send this into the data channel for
	 * further XML processing.
	 */
	if isOpen && strings.HasSuffix(lineContent, closeTag) {
		endIndex := strings.LastIndex(lineContent, closeTag)
		fullTag := tagContent + lineContent[startIndex:endIndex] + closeTag
		dataChan <- fullTag
		return ""
	}

	/**
	 * If the tag is still open because of a multiline, add the line from
	 * from the startindex onwards into the tagContent string and return it.
	 * Else return an empty string so the next line knows the tag is closed.
	 */
	if isOpen {
		return tagContent + lineContent[startIndex:] + "\n"
	}

	return ""
}

type resourceStruct struct {
	artists  []Artist
	labels   []Label
	masters  []Master
	releases []Release
}

func fetchInnerXML(xmlFile *os.File, resource string, numWorkers int) resourceStruct {

	// artists := make([]Artist, 0)

	// parse string until you find the start and end tag

	// labelChans := make([]chan interface{}, numWorkers)
	itemsChan := make(chan interface{})
	// artistChan := make(chan Artist)
	data := resourceStruct{}
	data.labels = make([]Label, 0)
	data.artists = make([]Artist, 0)
	data.masters = make([]Master, 0)
	counter := 0

	chanToArray := func(itemsChan chan interface{}) chan interface{} {
		done := make(chan interface{})

		go (func() {
			// multiPlexedLabelChans := fanIn(done, labelChans...)

			if resource == "label" {
				for i := range itemsChan {
					data.labels = append(data.labels, i.(Label))
				}
			}

			if resource == "artist" {
				for i := range itemsChan {
					data.artists = append(data.artists, i.(Artist))
				}
			}

			if resource == "master" {
				for i := range itemsChan {
					data.masters = append(data.masters, i.(Master))
				}
			}

			if resource == "release" {
				for range itemsChan {
					counter++
				}
			}

			close(done)
		})()

		return done
	}

	dataChan := make(chan string)

	var wg sync.WaitGroup

	done := chanToArray(itemsChan)

	for i := 0; i < numWorkers; i++ {
		wg.Add(2)
		var workerChan chan interface{}
		if resource == "label" {
			workerChan = SpawnWorkerL(dataChan, &wg)
		}
		if resource == "artist" {
			workerChan = SpawnWorkerA(dataChan, &wg)
		}
		if resource == "master" {
			workerChan = SpawnWorkerM(dataChan, &wg)
		}
		if resource == "release" {
			workerChan = SpawnWorkerR(dataChan, &wg)
		}
		go func() {
			for w := range workerChan {
				itemsChan <- w
			}
			wg.Done()
		}()
	}

	/** Increase the buffer for our scanner to 1MB, for some humongous lines! */
	scanner := bufio.NewScanner(xmlFile)
	buf := make([]byte, 0, 64*1024)
	scanner.Buffer(buf, 1024*1024)

	tagContent := ""
	openTag := fmt.Sprintf("<%s>", resource)
	closeTag := fmt.Sprintf("</%s>", resource)
	if resource == "master" {
		openTag = "<master id"
	}

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		tagContent = scanLineForXMLTag(tagContent, openTag, closeTag, line, dataChan)
	}

	err := scanner.Err()
	if err != nil {
		log.Fatalf("Error: %s", err)
	}

	close(dataChan)

	wg.Wait()

	close(itemsChan)

	<-done

	fmt.Printf("Decoded amount: %d\n", counter)

	return data
}
