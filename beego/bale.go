package main

import(
	"os"
	"strings"
	"path"

	""
)

func isExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

func init() {
	files := []string{
		"",
	}

	funcs := []func() []byte{
		bale.R,
	}

	for i, f := range funcs {
		fp := getFilePath(files[i])
		if !isExist(fp) {
			saveFile(fp, f())
		}
	}
}

func getFilePath(name string) string {
	name = strings.Replace(name, "_4_", "/", -1)
	name = strings.Replace(name, "_3_", " ", -1)
	name = strings.Replace(name, "_2_", "-", -1)
	name = strings.Replace(name, "_1_", ".", -1)
	name = strings.Replace(name, "_0_", "_", -1)
	return name
}

func saveFile(filePath string, b []byte) (int, error) {
	os.MkdirAll(path.Dir(filePath), os.ModePerm)
	fw, err := os.Create(filePath)
	if err != nil {
		return 0, err
	}
	defer fw.Close()
	return fw.Write(b)
}
