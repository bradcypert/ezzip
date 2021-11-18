package pkg

import (
	"fmt"
	"os"
	"path/filepath"
)

func loadFiles(d string) ([]string, error) {

	var files []string

	walker := func(path string, dir os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if dir.IsDir() {
			return nil
		}

		// If its a file, lets add its path
		files = append(files, path)

		return nil
	}

	err := filepath.WalkDir(d, walker)
	if err != nil {
		panic(err)
	}

	return files, nil

}
