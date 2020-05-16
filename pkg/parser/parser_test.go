package parser

import "testing"

func BenchmarkParser_xmlToken(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Parser("xml_token", "label", 0)
	}
}

func BenchmarkParser_fetchInnerXML_2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Parser("fetch_inner_xml", "label", 2)
	}
}

func BenchmarkParser_fetchInnerXML_5(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Parser("fetch_inner_xml", "label", 5)
	}
}

func BenchmarkParser_fetchInnerXML_10(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Parser("fetch_inner_xml", "label", 10)
	}
}

func BenchmarkParser_fetchInnerXML_20(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Parser("fetch_inner_xml", "label", 20)
	}
}

func BenchmarkParser_fetchInnerXML_50(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Parser("fetch_inner_xml", "label", 50)
	}
}

func BenchmarkParser_fetchInnerXML_100(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Parser("fetch_inner_xml", "label", 100)
	}
}

func BenchmarkParser_fetchInnerXML_200(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Parser("fetch_inner_xml", "label", 200)
	}
}
