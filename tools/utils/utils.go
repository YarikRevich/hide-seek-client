package utils

import "path/filepath"

func GetBasePath(path string) string {
	r, _ := filepath.Split(path)
	return r
}

// func GetPathByHash(hash [sha256.Size]byte, m map[string][sha256.Size]byte) string {
// 	for k, v := range m {
// 		if v == hash {
// 			return k
// 		}
// 	}
// 	return ""
// }
