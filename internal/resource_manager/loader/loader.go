package loader

import (
	_ "image/png"
	"io/fs"
	"io/ioutil"
	"log"
	"strings"
	"sync"
)

func processResource(motherDir, sourcePath string, files []fs.FileInfo, motherWg *sync.WaitGroup, loaders ...func(string, string, string, *sync.WaitGroup)) {
	for _, v := range files {
		path := sourcePath + "/" + v.Name()
		if v.IsDir() {
			processResourceDir(motherDir, path, motherWg, loaders...)
		} else {
			nameSplit := strings.Split(v.Name(), ".")
			extension := nameSplit[len(nameSplit)-1]

			for _, l := range loaders {
				l(motherDir, extension, path, motherWg)
			}
		}
	}
}

func processResourceDir(motherDir, path string, wg *sync.WaitGroup, loaders ...func(string, string, string, *sync.WaitGroup)) {
	d, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatalln(err)
	}

	processResource(motherDir, path, d, wg, loaders...)
}

func LoadResources(loaders map[string][]func(string, string, string, *sync.WaitGroup)) {
	var wg sync.WaitGroup
	for mp, l := range loaders {
		wg.Add(1)
		go func(mp string, l []func(string, string, string, *sync.WaitGroup)){
			defer wg.Done()
			processResourceDir(mp, mp, &wg, l...)
		}(mp, l) 
	}
	wg.Wait()
}
