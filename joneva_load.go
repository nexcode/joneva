package joneva

func (j *Joneva) Load(override ...string) error {
	j.mutex.Lock()
	defer j.mutex.Unlock()

	if len(override) != 0 && override[0] != "" {
		return j.override(override[0])
	}

	if _, err := j.file.Seek(0, 0); err != nil {
		return err
	}

	if j.data != nil {
		j.data = nil
	}

	return j.decoder.Decode(&j.data)
}
