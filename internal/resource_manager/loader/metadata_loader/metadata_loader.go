package metadataloader

import (
	"embed"
	"fmt"
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

func Load(e embed.FS, extension, path string, wg *sync.WaitGroup) {
	if extension != "toml" {
		return
	}

	wg.Add(1)
	go func() {
		var ds Metadata

		if _, err := toml.DecodeFS(e, path, &ds); err != nil {
			logrus.Fatal("error happened decoding toml metatdata file from embedded FS", err)
		}

		reg := regexp.MustCompile(`\.[a-z]*$`)
		if reg.MatchString(path) {
			mu.Lock()
			defer mu.Unlock()
			MetadataCollection[reg.Split(path, -1)[0]] = &ds
		}

		wg.Done()
	}()
}
