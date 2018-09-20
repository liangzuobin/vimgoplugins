package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {
	s := []string{
		"github.com/klauspost/asmfmt/cmd/asmfmt",
		"github.com/derekparker/delve/cmd/dlv",
		"github.com/kisielk/errcheck",
		"github.com/davidrjenni/reftools/cmd/fillstruct",
		"github.com/mdempsky/gocode",
		"github.com/rogpeppe/godef",
		"github.com/zmb3/gogetdoc",
		"golang.org/x/tools/cmd/goimports",
		"github.com/golang/lint/golint",
		"github.com/alecthomas/gometalinter",
		"github.com/fatih/gomodifytags",
		"golang.org/x/tools/cmd/gorename",
		"github.com/jstemmer/gotags",
		"golang.org/x/tools/cmd/guru",
		"github.com/josharian/impl",
		"honnef.co/go/tools/cmd/keyify",
		"github.com/fatih/motion",
		"github.com/koron/iferr",
	}

	for _, p := range s {
		if err := get(p); err != nil {
			fmt.Printf("go get -u -v %s failed", p)
		}
	}
}

func get(path string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	cmd := exec.CommandContext(ctx, "meowproxy", "go", "get", "-u", path)
	cmd.Env = append(os.Environ(), "GOPATH=/Users/liangzuobin/Code/Go")
	return cmd.Run()
}
