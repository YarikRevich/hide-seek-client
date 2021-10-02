package metadataloader

import (
	"fmt"
	"log"
	"regexp"
	"sync"

	"github.com/BurntSushi/toml"
	"github.com/sirupsen/logrus"
)

var (
	MetadataCollection = make(map[string]*Metadata)
	mu                 = sync.Mutex{}
)

func GetMetadata(path string) *Metadata {
	i, ok := MetadataCollection[path]
	if !ok {
		logrus.Fatal(fmt.Sprintf("image with path '%s' not found", path))
	}
	return i
}

func Load(motherDir, extension, path string, wg *sync.WaitGroup) {
	if extension != "toml" {
		return
	}

	wg.Add(1)
	go func() {
		var ds Metadata
		if _, err := toml.DecodeFile(path, &ds); err != nil {
			log.Fatalln(err)
		}

		path = path[len(motherDir):]
		reg := regexp.MustCompile(`\.[a-z]*$`)
		if reg.MatchString(path) {
			mu.Lock()
			defer mu.Unlock()
			MetadataCollection[reg.Split(path, -1)[0]] = &ds
		}

		wg.Done()
	}()
}
