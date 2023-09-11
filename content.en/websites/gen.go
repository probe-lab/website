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
	"probelab.io",
	"singularity.storage",
	"n0.computer",
	"fil.org",
}

type TemplateData struct {
	Website string
	Anchor  string
}

/*
 * Why do we need this?
 *   Hugo currently cannot generate parameterized files as part of its build
 *   process. One option we explored was to create a layout file for each
 *   website. The drawback was that we couldn't write markdown in it nor use
 *   the plotly short codes. Therefore, this 30L Go program.
 */
func main() {
	templateStr, err := os.ReadFile("./gen.template")
	if err != nil {
		panic(err)
	}

	// use || delimiter because the rest of the text is full of {{ and }}.
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
