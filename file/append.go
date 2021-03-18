package file

import "os"

func Append(path, content string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		file, err := os.Create(path)
		defer file.Close()
		if err != nil {
			return err
		}
		file.Write([]byte(content))
		return nil
	}

	file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, os.ModePerm)
	defer file.Close()
	if err != nil {
		return err
	}

	file.Write([]byte(content))
	return nil
}
