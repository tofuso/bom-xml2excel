package write

import (
	"github.com/loadoff/excl"
	"sandbox/bom-xml2excel/sorting"
	"sort"
	"strings"
)

const (
	// 列に書かれる内容を定義
	libCol  = iota + 1 // Libraryの列
	partCol            // partの列
	fpCol              // footprintの列
	valCol             // valueの列
	numCol             // numberの列
	refCol             // referenceの列
)

// SetDataOnExcel 渡されたデータを元にExcelファイルを作成する関数
func SetDataOnExcel(list *map[sorting.Data][]string, filename string) error {
	file, err := excl.Create() // ファイル作成
	defer file.Close()
	sheet, err := file.OpenSheet("Sheet1")
	if err != nil {
		return err
	}

	// リストのmapに格納されているDataの一覧を取得し、
	// ソートしてエクセルに書き込む際はソートされたものを用いる
	var dlis []sorting.Data
	for d := range *list {
		dlis = append(dlis, d)
	}
	// Data構造体のフィールドの優先順位に従ってソート

	// Valの値に従ってソート
	less := func(i, j int) bool {
		return dlis[i].Val < dlis[j].Val
	}
	sort.SliceStable(dlis, less)
	// footprintに従ってソート
	less = func(i, j int) bool {
		return dlis[i].FootprintName < dlis[j].FootprintName
	}
	sort.SliceStable(dlis, less)
	// Partに従ってソート
	less = func(i, j int) bool {
		return dlis[i].Part < dlis[j].Part
	}
	sort.SliceStable(dlis, less)
	// Libraryに従ってソート
	less = func(i, j int) bool {
		return dlis[i].LibraryName < dlis[j].LibraryName
	}
	sort.SliceStable(dlis, less)

	header := "" // 1行目のセルにヘッダーを書き込み
	border := excl.Border{
		Left: &excl.BorderSetting{
			Style: "medium",
		},
		Right: &excl.BorderSetting{
			Style: "medium",
		},
		Bottom: &excl.BorderSetting{
			Style: "medium",
		},
		Top: &excl.BorderSetting{
			Style: "medium",
		},
	}
	for i := 1; i <= 6; i++ {
		switch i {
		case libCol:
			header = "library"
		case partCol:
			header = "part"
		case fpCol:
			header = "footprint"
		case valCol:
			header = "value"
		case numCol:
			header = "number"
		case refCol:
			header = "reference"
		}
		sheet.GetRow(1).GetCell(i).SetString(header).SetBorder(border).SetBackgroundColor("AFEEEE")
	}

	// リファレンスを書き込む
	var refs []string
	var row int
	for i, d := range dlis {
		refs = (*list)[d]
		row = i + 2
		sheet.GetRow(row).GetCell(libCol).SetString(d.LibraryName).SetBorder(border)
		sheet.GetRow(row).GetCell(partCol).SetString(d.Part).SetBorder(border)
		sheet.GetRow(row).GetCell(fpCol).SetString(d.FootprintName).SetBorder(border)
		sheet.GetRow(row).GetCell(valCol).SetString(d.Val).SetBorder(border)
		sheet.GetRow(row).GetCell(numCol).SetNumber(len(refs)).SetBorder(border)
		sheet.GetRow(row).GetCell(refCol).SetString(strings.Join(refs, ",")).SetBorder(border)
	}
	err = file.Save(filename)
	if err != nil {
		return err
	}
	return nil
}
