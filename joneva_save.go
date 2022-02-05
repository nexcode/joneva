package joneva

func (j *Joneva) Save() error {
	j.mutex.Lock()
	defer j.mutex.Unlock()

	if _, err := j.file.Seek(0, 0); err != nil {
		return err
	}

	if err := j.file.Truncate(0); err != nil {
		return err
	}

	return j.encoder.Encode(j.data)
}
