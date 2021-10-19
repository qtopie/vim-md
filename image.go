package main

import "github.com/govim/govim"

func pasteImage(g govim.Govim, flags govim.CommandFlags, args ...string) error {
	rawMessage, err := g.ChannelExpr("expand('%:p')")
	if err != nil {
		g.Logf("%v", err)
	}
	g.Logf("Messages %s", rawMessage)
	return nil
}

func cleanImage(g govim.Govim, flags govim.CommandFlags, args ...string) error {
	return nil
}
