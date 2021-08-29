package configparser

import (
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

func ParseConfig(name string)(frames [][]float64) {
	f, err := os.OpenFile(name, os.O_RDONLY, 0755)
	if err != nil {
		log.Fatalln(err)
	}

	s, err := f.Stat()
	if err != nil {
		log.Fatalln(err)
	}

	buff := make([]byte, s.Size())
	if _, err = f.Read(buff); err != nil {
		log.Fatalln(err)
	}

	var d [][]float64
	if _, err = toml.Decode(string(buff), &d); err != nil{
		log.Fatalln(err)
	}

	return d
}
