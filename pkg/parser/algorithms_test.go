package parser

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testLabel string = `{
	"profile": "Founded in 1991 by [a=Dave Clarke (7)] \u0026 [a=Slam] Soma Quality Recordings is a successful UK electronic label which has been releasing tracks in the house/tech-house/techno vein; [a=Daft Punk]'s debut releases were on Soma and the label roster includes artists like Slam, Funk D'Void, H-Foundation, Percy X and many more.\r\nThe company behind the label is [l270709].\r\nThe publishing entities are [l295038], and, since 2002, [l=SPG Publishing UK Ltd]\r\nLC 12668\r\n\r\nName variations as found on this profile:\r\n- Soma Quality Recordings\r\n- Soma Records\r\n- Soma Recordings - England\r\n",
	"releases_url": "",
	"name": "Soma Quality Recordings",
	"contact_info": "Soma Quality Recordings Ltd,\r\n2nd Floor,\r\n342 Argyle Street,\r\nGlasgow\r\nG2 8LY\r\n\r\nTel: +44 (0) 141 229 6220\r\nFax: +44 (0) 141 226 4383\r\nEmail: info@somarecords.com\r\n",
	"uri": "",
	"parent_label": {
		"resource_url": "",
		"id": 270709,
		"name": "Soma Recordings Ltd."
	},
	"sublabels": [
		{
			"resource_url": "",
			"id": 646,
			"name": "Fenetik Music"
		},
		{
			"resource_url": "",
			"id": 4133,
			"name": "Fifth Freedom"
		},
		{
			"resource_url": "",
			"id": 127392,
			"name": "Paragraph"
		},
		{
			"resource_url": "",
			"id": 57766,
			"name": "Pnuma"
		},
		{
			"resource_url": "",
			"id": 995826,
			"name": "Soma Track Series"
		}
	],
	"urls": [
		"http://www.somarecords.com/",
		"http://www.facebook.com/SomaRecords",
		"http://plus.google.com/u/0/+OfficialSomaRecords",
		"http://www.myspace.com/somarecords",
		"http://soundcloud.com/soma",
		"http://twitter.com/SomaRecords",
		"http://www.youtube.com/user/OfficialSomaRecords"
	],
	"resource_url": "",
	"id": 18,
	"data_quality": "Correct"
}`

func TestXmlToken(t *testing.T) {
	// Open our xmlFile
	path := fmt.Sprintf("micro_test_labels.xml")
	xmlFile, _ := os.Open(path)

	defer xmlFile.Close()

	data := xmlToken(xmlFile, "label")

	label := data.labels[15]

	jsonLabel, _ := json.MarshalIndent(label, "", "\t")

	assert.Equal(t, testLabel, string(jsonLabel), "Generated label should match screenshot")
}

func TestFetchInnerXML(t *testing.T) {
	path := fmt.Sprintf("micro_test_labels.xml")
	xmlFile, _ := os.Open(path)

	defer xmlFile.Close()

	data := fetchInnerXML(xmlFile, "label", 10)

	var createdLabel Label

	for _, l := range data.labels {
		if l.ID == int64(18) {
			createdLabel = l
		}
	}

	jsonLabel, _ := json.MarshalIndent(createdLabel, "", "\t")

	assert.Equal(t, testLabel, string(jsonLabel), "Generated label should match screenshot")
}
