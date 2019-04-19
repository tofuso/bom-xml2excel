package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sandbox/bom-xml2excel/read"
	"sandbox/bom-xml2excel/sorting"
	"sandbox/bom-xml2excel/write"

	"github.com/urfave/cli"
)

func main() {

	app := cli.NewApp()
	app.Name = "xml2excel"
	app.Usage = "kicadで出力したxmlをexcelに変換します。"
	app.Version = "0.1"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "out, o",
			Usage: "出力先を指定する。",
		},
		cli.BoolFlag{
			Name:  "show, s",
			Usage: "情報を表示する。",
		},
	}
	app.Action = run
	err := app.Run(os.Args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v", err)
		return
	}
}

func run(con *cli.Context) error {
	filename := con.Args().First()
	if filename == "" {
		return fmt.Errorf("読み込むファイル名を指定してください。")
	}
	if filepath.Ext(filename) != ".xml" {
		return fmt.Errorf("%sはxmlファイルではありません。", filename)
	}
	var outname string // 出力先の名前
	if outname = con.String("out"); outname == "" {
		outname = filepath.Base(filename)
		outname = outname[:len(outname)-3]
		outname += "xlsx"
	}
	if filepath.Ext(outname) != ".xlsx" {
		return fmt.Errorf("%sはxlsxファイルではありません。", outname)
	}
	xmlFile, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer xmlFile.Close()
	xmlData, err := ioutil.ReadAll(xmlFile)
	if err != nil {
		return err
	}
	if con.Bool("show") {
		fmt.Printf("%s 読み込み\n", filename)
	}
	comps, err := read.GetXML(&xmlData)
	if err != nil {
		return err
	}
	list := sorting.CreateList(comps)
	err = write.SetDataOnExcel(list, outname)
	if err != nil {
		return err
	}
	if con.Bool("show") {
		fmt.Printf("%s 書き込み\n", outname)
		fmt.Println("完了")
	}
	return nil
}
