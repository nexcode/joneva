package joneva

func (j *Joneva) loadDefaults(defaults map[string]interface{}) error {
	if defaults != nil {
		var needSave bool

		for key, value := range defaults {
			if j.Get(key) == nil {
				j.Set(key, value)
				needSave = true
			}
		}

		if needSave {
			if err := j.Save(); err != nil {
				return err
			}
		}
	}

	return nil
}
