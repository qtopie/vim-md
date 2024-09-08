package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os/exec"
	"runtime"
	"strings"

	"github.com/govim/govim"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
)

//go:embed _template/preview.html
var tpl string

type MarkdownContent struct {
	Content string
}

func previewMarkdown(g govim.Govim, flags govim.CommandFlags, args ...string) error {
	rawMessage, err := g.ChannelExpr("join(getline(1, '$'), '\n')")
	if err != nil {
		return err
	}

	content := strings.Trim(string(rawMessage), `"`)
	content = strings.Replace(content, `\n`, "\n", -1)
	g.Logf(content)
	go serveHttp(content)
	return nil
}

func serveHttp(content string) {
	t, err := template.New("preview").Parse(tpl)
	if err != nil {
		log.Print(err)
	}

	buf := renderMarkdown(content)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t.Execute(w, MarkdownContent{Content: buf.String()})
	})

	openBrowser("http://127.0.0.1:7070")
	http.ListenAndServe("127.0.0.1:7070", nil)
}

func renderMarkdown(content string) bytes.Buffer {
	md := goldmark.New(
		goldmark.WithExtensions(extension.GFM),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
			parser.WithParagraphTransformers(),
		),
		goldmark.WithRendererOptions(
			html.WithUnsafe(),
			html.WithHardWraps(),
		),
	)

	var buf bytes.Buffer

	if err := md.Convert([]byte(content), &buf); err != nil {
		panic(err)
	}
	return buf
}

func openBrowser(url string) (err error) {
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}

	return err
}
