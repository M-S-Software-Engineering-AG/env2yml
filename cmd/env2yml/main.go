package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/M-S-Software-Engineering-AG/env2yml"
)

func main() {
	prefix := flag.String("prefix", "", "only convert variables whose name starts with this prefix; the prefix is stripped from the key")
	out := flag.String("out", "", "write YAML to this file instead of stdout")
	flag.Parse()

	doc, err := env2yml.Marshal(os.Environ(), *prefix)
	if err == nil {
		if *out == "" {
			_, err = os.Stdout.Write(doc)
		} else if err = os.MkdirAll(filepath.Dir(*out), 0o755); err == nil {
			err = os.WriteFile(*out, doc, 0o644)
		}
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, "env2yml:", err)
		os.Exit(1)
	}
}
