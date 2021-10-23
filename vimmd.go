package main

import (
	"fmt"
	"strconv"

	"github.com/govim/govim"
	"gopkg.in/tomb.v2"
)

type VimMdPlugin struct {
	tomb tomb.Tomb
}

func (p *VimMdPlugin) Init(g govim.Govim, ch chan error) error {
	g.DefineCommand("MarkdownPreview", previewMarkdown)
	g.DefineCommand("MarkdownImagePaste", pasteImage)
	g.DefineCommand("MarkdownImageClean", cleanImage)
	return nil
}

func (p *VimMdPlugin) Shutdown() error {
	return nil
}

func showMsg(g govim.Govim, format string, args ...interface{}) error {
	msg := fmt.Sprintf(format, args...)
	return g.ChannelEx("echomsg " + strconv.Quote(msg))
}

func showErrMsg(g govim.Govim, format string, args ...interface{}) error {
	msg := fmt.Sprintf(format, args...)
	return g.ChannelEx("echoerr " + strconv.Quote(msg))
}

func appendLine(g govim.Govim, format string, args ...interface{}) error {
	line := fmt.Sprintf(format, args...)
	return g.ChannelNormal("A" + line)
}
