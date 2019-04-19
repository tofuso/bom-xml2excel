package read

import (
	"encoding/xml"
)

// Components パーツデータのリストを格納する構造体
type Components struct {
	Comp []Component `xml:"components>comp"`
}

// Component パーツの情報を含んだ構造体
type Component struct {
	Reference string        `xml:"ref,attr"`
	Val       string        `xml:"value"`
	FootPrint string        `xml:"footprint"`
	Libsource  LibrarySource `xml:"libsource"`
}

// LibrarySource ライブラリ名、パーツ区分と説明を格納する構造体
type LibrarySource struct {
	LibraryName string `xml:"lib,attr"`
	Part        string `xml:"part,attr"`
	Description string `xml:"description"`
}

// GetXML 渡されたXMLをパースして構造体へ読み込む
func GetXML(data *[]byte) (*Components, error) {
	// Component構造体の配列を持つ構造体を定義、宣言
	contents := Components{}
	err := xml.Unmarshal(*data, &contents)
	if err != nil {
		return &Components{}, err
	}
	return &contents, nil
}
