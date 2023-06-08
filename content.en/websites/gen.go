package main

import (
	"fmt"
	"os"
	"strings"
	"text/template"
)

var websites = []string{
	"filecoin.io",
	"libp2p.io",
	"ipld.io",
	"protocol.ai",
	"docs.ipfs.tech",
	"tldr.filecoin.io",
	"green.filecoin.io",
	"drand.love",
	"web3.storage",
	"consensuslab.world",
	"strn.network",
	"research.protocol.ai",
	"ipfs.tech",
	"blog.ipfs.tech",
	"docs.libp2p.io",
	"blog.libp2p.io",
	//"en.wikipedia-on-ipfs.org/wiki",
	"specs.ipfs.tech",
}

type TemplateData struct {
	Website string
	Anchor  string
}

func main() {
	templateStr, err := os.ReadFile("./gen.template")
	if err != nil {
		panic(err)
	}

	tpl, err := template.New("website").Delims("||", "||").Parse(string(templateStr))
	if err != nil {
		panic(err)
	}

	for _, website := range websites {
		fmt.Println("Writing", website)
		out, err := os.Create(website + ".md")
		if err != nil {
			panic(err)
		}

		data := TemplateData{
			Website: website,
			Anchor:  strings.ReplaceAll(website, ".", ""),
		}

		err = tpl.Execute(out, data)
		if err != nil {
			panic(err)
		}
		if err = out.Close(); err != nil {
			panic(err)
		}
	}
}
