package util

import (
	"encoding/json"
	"os"
)

func LoadJsonConfig(config interface{}, filename string) (err error) {
	var decoder *json.Decoder
	var file *os.File
	file, err = os.Open(filename)
	if err != nil {
		return
	}
	defer file.Close()

	decoder = json.NewDecoder(file)
	if err = decoder.Decode(config); err != nil {
		return
	}

	_, err = json.Marshal(config)
	return
}
