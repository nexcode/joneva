package joneva

import (
	"encoding/json"
	"io"
	"os"
)

func (j *Joneva) override(fileName string) error {
	file, err := os.OpenFile(fileName, os.O_RDONLY, 0)
	if err != nil {
		return err
	}

	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	return json.Unmarshal(bytes, &j.data)
}
