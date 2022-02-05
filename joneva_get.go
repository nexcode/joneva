package joneva

import "strings"

func (j *Joneva) Get(key string) interface{} {
	j.mutex.Lock()
	defer j.mutex.Unlock()

	data := j.data
	path := strings.Split(key, ".")
	var i int

	for i < len(path)-1 {
		tmp, ok := data[path[i]]
		if !ok {
			return nil
		}

		if data, ok = tmp.(map[string]interface{}); !ok {
			return nil
		}

		i++
	}

	return data[path[i]]
}
