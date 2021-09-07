package main

import (
	"encoding/json"
	"fmt"
	"github.com/flosch/pongo2/v4"
	"github.com/urfave/cli/v2"
	"io/ioutil"
)

// ToHtml converts the VOTD to HTML with a Hero layout
func ToHtml(votd *VOTD) string {

	tpl, err := pongo2.FromString(template)
	if err != nil {
		panic(err)
	}

	out, err := tpl.Execute(pongo2.Context{"verse": votd.Verse.Details.Text, "verse_reference": votd.Verse.Details.Reference})
	if err != nil {
		panic(err)
	}

	return out
}

// ToJson marshalls the VOTD metadata back to JSON
func ToJson(votd *VOTD) []byte {
	data, err := json.Marshal(votd)
	if err != nil { panic(err) }
	return data
}


// Parse creates a bulma derived HTML and JSON file and writes to the provided destination
// path
func Parse(c *cli.Context) error {
	votd, err := getVerse()
	if err != nil {
		return err
	}

	out := ToHtml(votd)
	htmlOutputFile := c.String("html-output-file")
	if htmlOutputFile == "" {
		fmt.Println(out)
	} else {
		err := ioutil.WriteFile(htmlOutputFile, []byte(out), 0o644)
		if err != nil { return err }
	}

	jsonOutputFile := c.String("json-output-file")
	if jsonOutputFile != "" {
		err := ioutil.WriteFile(jsonOutputFile, ToJson(votd), 0o644)
		if err != nil { return err }
	}
	return nil
}