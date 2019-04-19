package sorting

import (
	"fmt"
	"io/ioutil"
	"os"
	"sandbox/bom-xml2excel/read"
	"testing"
)

func TestSorting(*testing.T) {
	xmlFile, err := os.Open("../assets/oribe2.xml")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err.Error())
		return
	}
	defer xmlFile.Close()
	xmlData, err := ioutil.ReadAll(xmlFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err.Error())
		return
	}
	comps, err := read.GetXML(&xmlData)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err.Error())
		return
	}
	list := CreateList(comps)
	fmt.Println(list)
}
