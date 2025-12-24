package task

import (
	"encoding/json"
	"os"
)

const fileName = "tasks.json"

func Load() (List, error) {
	var list List

	data, err := os.ReadFile(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			os.Create(fileName)
			return list, nil
		}
		return nil, err
	}

	if len(data) == 0 {
		return list, nil
	}

	err = json.Unmarshal(data, &list)
	return list, err
}

func Save(list List) error {
	data, err := json.MarshalIndent(list, "", "	")

	if err != nil {
		return err
	}

	return os.WriteFile(fileName, data, 0644)
}

func Delete(id int) error {
	list, err := Load()
	if err != nil {
		return err
	}
	list.Delete(id)
	return Save(list)
}
