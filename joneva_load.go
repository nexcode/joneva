package joneva

func (j *Joneva) Load() error {
	j.mutex.Lock()
	defer j.mutex.Unlock()

	if _, err := j.file.Seek(0, 0); err != nil {
		return err
	}

	return j.decoder.Decode(&j.data)
}
