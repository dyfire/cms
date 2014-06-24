// Copyright 2013 wetalk authors
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package utils

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/beego/i18n"
	"html/template"
	"os"
	"reflect"
	"strings"
)

var (
	Lang  string
	Langs []string
)

// get HTML i18n string
func i18nHTML(lang, format string, args ...interface{}) template.HTML {
	return template.HTML(i18n.Tr(lang, format, args...))
}

func ToInt64(value interface{}) (d int64, err error) {
	val := reflect.ValueOf(value)
	switch value.(type) {
	case int, int8, int16, int32, int64:
		d = val.Int()
	case uint, uint8, uint16, uint32, uint64:
		d = int64(val.Uint())
	default:
		err = fmt.Errorf("ToInt64 need numeric not `%T`", value)
	}
	return
}

func settingLocales() {
	// load locales with locale_LANG.ini files
	langs := "zh-CN|en-US"
	for _, lang := range strings.Split(langs, "|") {
		lang = strings.TrimSpace(lang)
		files := []string{"conf/" + "locale_" + lang + ".ini"}
		if fh, err := os.Open(files[0]); err == nil {
			fh.Close()
		} else {
			files = nil
		}
		if err := i18n.SetMessage(lang, "conf/"+"locale_"+lang+".ini", files...); err != nil {
			beego.Error("Fail to set message file: " + err.Error())
			os.Exit(2)
		}
	}

	Lang = "zh-CN"
	Langs = i18n.ListLangs()
}

func init() {
	// Register template functions.
	settingLocales()
	beego.AddFuncMap("i18n", i18nHTML)
}
