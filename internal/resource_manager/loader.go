package resource_manager

import (
	"embed"
	_ "image/png"
	"io/fs"
	"strings"
	"sync"

	"github.com/sirupsen/logrus"
)

func processResource(e embed.FS, sourcePath string, files []fs.DirEntry, motherWg *sync.WaitGroup, loaders ...Loader) {
	for _, v := range files {
		path := sourcePath + "/" + v.Name()
		if v.IsDir() {
			processResourceDir(e, path, motherWg, loaders...)
		} else {
			nameSplit := strings.Split(v.Name(), ".")
			extension := nameSplit[len(nameSplit)-1]

			for _, l := range loaders {
				l(e, extension, path, motherWg)
			}
		}
	}
}

func processResourceDir(e embed.FS, path string, wg *sync.WaitGroup, loaders ...Loader) {
	d, err := e.ReadDir(path)
	if err != nil {
		logrus.Fatal("error happened reading dir from embedded fs", err)
	}
	processResource(e, path, d, wg, loaders...)
}

func LoadResources(loaders map[Component][]Loader) {
	var wg sync.WaitGroup
	for c, l := range loaders {
		wg.Add(1)
		go func(c Component, l []Loader){
			defer wg.Done()
			processResourceDir(c.Embed, c.Path, &wg, l...)
		}(c, l) 
	}
	wg.Wait()
}
