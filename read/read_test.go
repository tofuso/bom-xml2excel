package read

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestGetXML(t *testing.T) {
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
	comps, err := GetXML(&xmlData)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err.Error())
		return
	}
	fmt.Printf("%+v\n", *comps)
}
