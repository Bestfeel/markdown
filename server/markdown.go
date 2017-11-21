// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package server

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	md "github.com/russross/blackfriday"
	"fmt"
)

const tpl = `
<!DOCTYPE html>
<html>
<head>
  <meta charset='UTF-8' />
{{if .isHighlight}}
 <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/highlight.js/9.12.0/styles/monokai.min.css">
<script src="https://cdnjs.cloudflare.com/ajax/libs/highlight.js/9.12.0/highlight.min.js"></script>
<script>hljs.initHighlightingOnLoad();</script>
{{end}}

</head>
<style>
{{.css}}
</style>
</head>

<body>
<article class="{{.class}}">
{{.body}}
</article>
</body>

</html>
`
const LOGO = `
                                                  _
                                          _______| |
                                         |_________|
                                          _________
                                         |  _______|   万物互联
                                         | |   ____
                                         | |  |__  |   机智云
                                         | |_____| |
                                         |_________|   Gizwits
                                          机智云只为硬件而生的云服务


`
const (
	commonHTMLFlags = 0 |
		md.HTML_USE_XHTML |
		md.HTML_USE_SMARTYPANTS |
		md.HTML_SMARTYPANTS_FRACTIONS |
		md.HTML_SMARTYPANTS_LATEX_DASHES

	commonExtensions = 0 |
		md.EXTENSION_NO_INTRA_EMPHASIS |
		md.EXTENSION_TABLES |
		md.EXTENSION_FENCED_CODE |
		md.EXTENSION_AUTOLINK |
		md.EXTENSION_STRIKETHROUGH |
		md.EXTENSION_SPACE_HEADERS |
		md.EXTENSION_HEADER_IDS |
		md.EXTENSION_BACKSLASH_LINE_BREAK |
		md.EXTENSION_DEFINITION_LISTS
)

var (
	globalAddr = ":8080"
	globalPath = "."
	globalCss  = "github"
)

var cssMap = map[string]struct {
	name, css   string;
	isHighlight bool
}{
	"github":  {"markdown-body", GITHUB, false},
	"mou":     {"", MOU, true},
	"marxico": {"container", MARXICO, true},
}

/**
 gen markdown html
 */
func markdown(in io.Reader, out io.Writer) error {
	buf, err := ioutil.ReadAll(in)
	if err != nil {
		return err
	}
	flg := commonHTMLFlags

	render := md.HtmlRenderer(flg, "", cssMap[globalCss].css)
	body := md.MarkdownOptions(buf, render, md.Options{
		Extensions: commonExtensions,
	})
	m := map[string]interface{}{
		"css":         cssMap[globalCss].css,
		"body":        string(body),
		"class":       cssMap[globalCss].name,
		"isHighlight": cssMap[globalCss].isHighlight,
	}
	return template.Must(template.New("markdown").Parse(tpl)).Execute(out, m)
}

func handleServerMarkdown(w http.ResponseWriter, r *http.Request) {
	code := 200
	var err error
	defer func() {
		log.Printf("%s %d %s", r.Method, code, r.URL.Path)
		if err != nil {
			w.WriteHeader(code)
			io.WriteString(w, err.Error())
		}
	}()
	file := filepath.Join(globalPath, r.URL.Path)
	if !(strings.HasSuffix(file, ".md") || strings.HasSuffix(file, ".markdown")) {
		http.FileServer(http.Dir(globalPath)).ServeHTTP(w, r)
		return
	}
	f, err := os.Open(file)
	if err != nil {
		if os.IsNotExist(err) {
			code = 404
		}
		return
	}
	defer f.Close()
	err = markdown(f, w)
	if err != nil {
		code = 500
		return
	}
}

/**
  run http server
 */
func RunMarkDownServer(addr string, rootPath string, css string) {
	globalAddr = addr
	globalPath = rootPath
	globalCss = css

	fmt.Println(LOGO)
	log.Printf("Listening on %s,  path  %s", globalAddr, globalPath)
	go http.HandleFunc("/", handleServerMarkdown)
	log.Fatal(http.ListenAndServe(globalAddr, nil))
}
