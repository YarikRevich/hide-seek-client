package sources

import (
	"embed"
	"io/fs"
	"path/filepath"
	"sync"

	"github.com/sirupsen/logrus"
)

type Parser struct {
	sync.WaitGroup

	fs         embed.FS
	path, base string
	callback   func(embed.FS, string)
}

func (p *Parser) parseFiles(base string, files []fs.DirEntry) {
	for _, v := range files {
		path := filepath.Join(base, v.Name())
		if v.IsDir() {
			p.parseDirs(path)
		} else {
			p.Add(1)
			go func() {
				p.callback(p.fs, path)
				p.Done()
			}()
		}
	}
}

func (p *Parser) parseDirs(base string) {
	files, err := p.fs.ReadDir(base)
	if err != nil {
		logrus.Fatal("error happened reading dir from embedded fs", err)
	}

	p.parseFiles(base, files)
}

func (p *Parser) Parse() {
	p.parseDirs(p.base)
	p.Wait()
}

func NewParser(fs embed.FS, path string, callback func(embed.FS, string)) *Parser {
	return &Parser{fs: fs, base: path, path: path, callback: callback}
}
