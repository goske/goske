package main

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/alecthomas/kingpin"
)

var (
	newCommand    = kingpin.Command("new", "create project with specified skeleton")
	newOutputFlag = newCommand.Flag("output", "output name").Short('o')
	newName       string
	newOutName    string
)

func init() {
	newCommand.Action(newAction).Arg("name", "skeleton name").Required().StringVar(&newName)
	newOutputFlag.Default(newName).StringVar(&newOutName)
}

func extract(tr *tar.Reader) error {
	header, err := tr.Next()
	if err != nil {
		return err
	}
	paths := strings.SplitN(header.Name, "/", 2)
	if len(paths) < 2 || paths[1] == "" {
		return nil
	}
	fname := filepath.FromSlash(path.Join(newOutName, paths[1]))
	err = os.MkdirAll(filepath.Dir(fname), 0755)
	if err != nil {
		return err
	}
	f, err := os.Create(fname)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = io.Copy(f, tr)
	if err != nil {
		return err
	}
	return nil
}

func newAction(ctx *kingpin.ParseContext) error {
	resp, err := http.Get("https://api.github.com/repos/goske/goske-" + newName + "/tarball")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	gr, err := gzip.NewReader(resp.Body)
	if err != nil {
		return err
	}
	tr := tar.NewReader(gr)
	for {
		err = extract(tr)
		if err != nil {
			if err == io.EOF {
				err = nil
			}
			break
		}
	}
	return err
}