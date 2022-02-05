package joneva

import "strings"

func (j *Joneva) Set(key string, value interface{}) {
	j.mutex.Lock()
	defer j.mutex.Unlock()

	data := j.data
	path := strings.Split(key, ".")
	var i int

	for i < len(path)-1 {
		tmp, ok := data[path[i]]
		if !ok {
			tmp = make(map[string]interface{})
			data[path[i]] = tmp
		}

		if data, ok = tmp.(map[string]interface{}); !ok {
			data = make(map[string]interface{})
		}

		i++
	}

	switch value.(type) {
	case int:
		value = float64(value.(int))
	case int8:
		value = float64(value.(int8))
	case int16:
		value = float64(value.(int16))
	case int32:
		value = float64(value.(int32))
	case int64:
		value = float64(value.(int64))
	case uint:
		value = float64(value.(uint))
	case uint8:
		value = float64(value.(uint8))
	case uint16:
		value = float64(value.(uint16))
	case uint32:
		value = float64(value.(uint32))
	case uint64:
		value = float64(value.(uint64))
	case float32:
		value = float64(value.(float32))
	}

	data[path[i]] = value
}
