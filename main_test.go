package main

import(
	"github.com/urfave/cli"
	"testing"
	"os"
)

func TestRun(t *testing.T){
	app := cli.NewApp()
	app.Name = "xml2excel"
	app.Usage = "kicadで出力したxmlをexcelに変換します。"
	app.Version = "0.1"
	os.Args = []string{
		"xml2excel",
		"--out",
		"assets/out.xlsx",
		"assets/input.xml",
	}
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
		t.Errorf("error :%v", err)
		return
	}
	t.Log("ファイル指定OK")
}