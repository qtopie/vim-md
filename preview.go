package main

import "github.com/govim/govim"

func previewMarkdown(g govim.Govim, flags govim.CommandFlags, args ...string) error {
	g.Logf("command preview is called")
	return nil
}
