package collection

import (
	"fmt"

	"github.com/sirupsen/logrus"

	"github.com/YarikRevich/HideSeek-Client/internal/resource_manager/metadata_loader/models"
)

var MetadataCollection = make(map[string]*models.Metadata)

func GetMetadata(path string) *models.Metadata {
	i, ok := MetadataCollection[path]
	if !ok {
		logrus.Fatal(fmt.Sprintf("image with path '%s' not found", path))
	}
	return i
}
