package pkg

import (
	"archive/zip"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

func ZipAssets(directoryPath string, shouldEncrypt bool) (*string, error) {
	file, err := os.Create(directoryPath + ".zip")
	var key string

	if err != nil {
		panic(err)
	}
	defer file.Close()

	w := zip.NewWriter(file)
	defer w.Close()

	paths, _ := loadFiles(directoryPath)

	for _, path := range paths {
		f, err := w.Create(path)
		if err != nil {
			return nil, err
		}

		contents, err := os.ReadFile(path)

		if err != nil {
			return nil, err
		}

		fmt.Println(shouldEncrypt)

		if shouldEncrypt {
			key, err = encrypt(&contents)

			// need to encrypt and something went wrong
			if err != nil {
				return nil, err
			}
		}

		// write file data
		_, err = f.Write(contents)
		if err != nil {
			return nil, err
		}
	}

	return &key, nil
}

func UnzipAssets(zipPath string, key string) {
	dir := strings.ReplaceAll(zipPath, ".zip", "")
	err := os.Mkdir(dir, 0755)

	if err != nil {
		panic(err)
	}

	rc, err := zip.OpenReader(zipPath)
	if err != nil {
		panic(err)
	}
	defer rc.Close()

	for _, file := range rc.File {
		nonDir := path.Join(strings.Split(file.Name, string(os.PathSeparator))[1:]...)
		newPath := path.Join(dir, nonDir)
		f, err := os.Create(newPath)
		if err != nil {
			panic("unable to create file while unzipping")
		}

		defer f.Close()

		rc, err := file.Open()
		if err != nil {
			panic("unable to open zipped file")
		}
		defer rc.Close()

		contents, err := ioutil.ReadAll(rc)

		if err != nil {
			panic("unable to read zipped files")
		}

		if key != "" {
			err = decrypt(&contents, key)

			// need to decrypt and something went wrong
			if err != nil {
				panic(err)
			}
		}

		// write file data
		_, err = f.Write(contents)
		if err != nil {
			panic("unable to write file to disk")
		}
	}
}
