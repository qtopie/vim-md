package main

import (
	"bytes"
	"fmt"
	"image/png"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/chai2010/webp"
	"github.com/govim/govim"
	"golang.design/x/clipboard"
)

func pasteImage(g govim.Govim, flags govim.CommandFlags, args ...string) error {
	rawMessage, err := g.ChannelExpr("expand('%:p')")
	if err != nil {
		return err
	}

	if rawMessage == nil || len(rawMessage) == 0 {
		showErrMsg(g, "create file first, then paste image again.")
		return nil
	}

	mdFilePath := strings.Trim(string(rawMessage), `"`)
	g.Logf("Messages %s", mdFilePath)

	b := clipboard.Read(clipboard.FmtImage)
	if b == nil {
		showErrMsg(g, "no image in system clipboard.")
		return nil
	}

	img, err := png.Decode(bytes.NewReader(b))
	if err != nil {
		return err
	}

	var buf bytes.Buffer
	err = webp.Encode(&buf, img, &webp.Options{Lossless: true})
	if err != nil {
		return err
	}

	assertsDir := strings.ToLower(strings.TrimSuffix(mdFilePath, filepath.Ext(mdFilePath))) + ".assets"
	err = os.MkdirAll(assertsDir, os.ModePerm)
	if err != nil {
		return err
	}

	pictureName := generatePictureFileName()
	file := filepath.Join(assertsDir, pictureName)
	err = os.WriteFile(file, buf.Bytes(), os.ModePerm)
	if err != nil {
		showErrMsg(g, "failed to save image %s.", file)
		return err
	}

	appendLine(g, "![](%s)", filepath.Join(filepath.Base(assertsDir), pictureName))
	showMsg(g, "saved image %s to assert dir.", pictureName)
	return nil
}

func generatePictureFileName() string {
	t := time.Now()

	return fmt.Sprintf("img-%d%02d%02d%02d%02d%02d.webp", t.Year(), t.Month(),
		t.Day(), t.Hour(), t.Minute(), t.Second())
}

func cleanImage(g govim.Govim, flags govim.CommandFlags, args ...string) error {
	return nil
}
