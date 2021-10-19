package main

import (
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
