package sorting

import (
	"sandbox/bom-xml2excel/read"
)

// Data 分類するための構造体
type Data struct {
	Val           string
	FootprintName string
	LibraryName   string
	Part          string
}

// CreateList Components構造体を渡されると、
// 整理してData構造をキーとしたリファレンスの配列を返す。
func CreateList(comps *read.Components) *map[Data][]string {
	list := make(map[Data][]string)
	var d Data        // ループ毎のコンポーネントデータ
	var ok bool       // リスト内にコンポーネントデータが存在するか判定するフラグ
	var refs []string // リスト内にリファレンスの配列がある時、付け足すためのバッファ
	for _, e := range comps.Comp {
		d = createData(e)
		_, ok = list[d]
		if ok {
			refs = list[d]
			refs = append(refs, e.Reference)
			list[d] = refs
		} else {
			list[d] = []string{e.Reference}
		}
	}
	return &list
}

func createData(comp read.Component) Data {
	return Data{
		Val:           comp.Val,
		FootprintName: comp.FootPrint,
		LibraryName:   comp.Libsource.LibraryName,
		Part:          comp.Libsource.Part,
	}
}
