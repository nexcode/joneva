package joneva

import (
	"encoding/json"
	"os"
	"path"
	"sync"
)

func New(filePath string, defaults map[string]interface{}) (*Joneva, error) {
	if path := path.Dir(filePath); path != "" {
		if err := os.MkdirAll(path, 0755); err != nil {
			return nil, err
		}
	}

	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		if os.IsPermission(err) {
			if file, err = os.OpenFile(filePath, os.O_RDONLY, 0); err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	fi, err := file.Stat()
	if err != nil {
		return nil, err
	}

	if fi.Size() == 0 {
		if _, err = file.WriteString("{}"); err != nil {
			return nil, err
		}
	}

	joneva := Joneva{
		file:    file,
		decoder: json.NewDecoder(file),
		encoder: json.NewEncoder(file),
	}

	joneva.encoder.SetIndent("", "    ")

	if err = joneva.Load(); err != nil {
		return nil, err
	}

	if err = joneva.loadDefaults(defaults); err != nil {
		return nil, err
	}

	return &joneva, nil
}

type Joneva struct {
	mutex   sync.Mutex
	file    *os.File
	decoder *json.Decoder
	encoder *json.Encoder
	data    map[string]interface{}
}
