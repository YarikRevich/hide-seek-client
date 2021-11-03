package sources

import (
	"embed"
	"io/fs"
	"sync"

	"github.com/sirupsen/logrus"
)

type Parser struct {
	sync.WaitGroup

	fs embed.FS
	path string
	callback func(embed.FS, string)
}

func (p *Parser) parseFiles(files []fs.DirEntry) {
	for _, v := range files {
		p.path = p.path + "/" + v.Name()
		if v.IsDir() {
			p.parseDirs()
		} else {
			p.Add(1)
			go func(){
				p.callback(p.fs, p.path)
				p.Done()
			}()
		}
	}
}

func (p *Parser) parseDirs() {
	files, err := p.fs.ReadDir(p.path)
	if err != nil {
		logrus.Fatal("error happened reading dir from embedded fs", err)
	}

	p.parseFiles(files)
}

func (p *Parser) Parse() {
	p.parseDirs()
	p.Wait()
}

func NewParser(fs embed.FS, path string, callback func(embed.FS, string)) *Parser {
	return &Parser{fs: fs, path: path, callback: callback}
}
