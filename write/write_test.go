package write

import (
	"fmt"
	"io/ioutil"
	"os"
	"sandbox/bom-xml2excel/read"
	"sandbox/bom-xml2excel/sorting"
	"testing"
)

func TestWrite(t *testing.T) {
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
	list := sorting.CreateList(comps)
	err = SetDataOnExcel(list, "../assets/out.xlsx")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err.Error())
		return
	}
}
