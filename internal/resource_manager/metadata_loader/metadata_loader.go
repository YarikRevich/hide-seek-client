package metadataloader

import (
	"embed"
	"regexp"
	"sync"

	"github.com/BurntSushi/toml"
	"github.com/sirupsen/logrus"

	"github.com/YarikRevich/HideSeek-Client/internal/resource_manager/metadata_loader/collection"
	"github.com/YarikRevich/HideSeek-Client/internal/resource_manager/metadata_loader/models"
)

var mu = sync.Mutex{}

func Load(e embed.FS, extension, path string, wg *sync.WaitGroup) {
	if extension != "toml" {
		return
	}

	wg.Add(1)
	go func() {
		var ds models.Metadata

		if _, err := toml.DecodeFS(e, path, &ds); err != nil {
			logrus.Fatal("error happened decoding toml metatdata file from embedded FS", err)
		}

		reg := regexp.MustCompile(`\.[a-z0-9]*$`)
		if reg.MatchString(path) {
			mu.Lock()
			collection.MetadataCollection[reg.Split(path, -1)[0]] = &ds
			mu.Unlock()
		}

		wg.Done()
	}()
}
