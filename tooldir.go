package main

import (
	"os"
	"path"
)

var tooldir string

func init() {
	cwd, err := os.Getwd()
	if err != nil {
		fatal("cannot get current working directory: %s", err)
	}
	tooldir = path.Join(cwd, "_tools")
}

func ensureTooldir() {
	stat, err := os.Stat(tooldir)
	if os.IsNotExist(err) {
		err = os.Mkdir(tooldir, 0777)
		if err != nil {
			fatal("unable to create tooldir", err)
		}
		return
	}
	if err != nil {
		fatal("unable to stat tool directory: %s", err)
	}
	if !stat.IsDir() {
		fatal("%s already exists, but it is not a directory. This is where tools would go, so giving up.", nil)
	}
}