package admin

import (
	// "fmt"
	"bytes"
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
	"strings"
	"time"
)

type UploadController struct {
	BaseController
}

func (this *UploadController) Upload() {
	_, header, err := this.GetFile("uploadify")

	if err != nil {
		beego.Debug("upload failtrue")
	}

	suffix := strings.ToLower(header.Filename[strings.LastIndex(header.Filename, "."):])

	file_name := time.Now().Format("20060102235959") + suffix
	save_path := "./static/attachments/" + file_name
	size := "1000"

	out := fmt.Sprintf("%s,%s,%s,%s", file_name, size, save_path, save_path)

	this.SaveToFile("uploadify", save_path)

	// this.Data["out"] = out
	// o := []byte{}
	// for k, v := range out {
	// 	o[k] = v
	// }
	ibytes := bytes.NewBufferString(out)
	icontent, _ := ioutil.ReadAll(ibytes)
	this.Ctx.Output.Header("Content-Type", "text/html; charset=utf-8")
	this.Ctx.Output.Body(icontent)
}
