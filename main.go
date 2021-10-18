package main

import (
	"fmt"
	"os"

	"github.com/govim/govim"
)

func main() {
	p := &VimMdPlugin{}
	g, err := govim.NewGovim(p, os.Stdin, os.Stdout, os.Stdout, &p.tomb)
	if err != nil {
		fmt.Printf("failed to create vim-md instance: %v", err)
	}
	
	p.tomb.Go(g.Run)
	p.tomb.Wait()
}
