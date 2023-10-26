package fs

import (
	"example/src/lib"
	"os"
)

var logger = lib.Log

func Write(data []string, dir string, filename string) {
	createDir(dir)
	file := writeFile(filename)
	defer file.Close()
	// Write each element of the slice to a new line in the file
	for _, element := range data {
		_, err := file.WriteString(element + "\n")
		if err != nil {
			logger.WithError(err).Fatal("Failed write data in new file")
			os.Exit(1)
		}
	}
	logger.Info("Successfully added new file")
}

func Clean(dir string, file string) {
	removeFile(file)
	removeDir(dir)
}

/* Private Methods*/
func writeFile(filename string) *os.File {
	file, err := os.Create(filename)
	if err != nil {
		logger.WithError(err).Fatal("Failed to create file", filename)
		os.Exit(1)
	}
	return file
}

func removeFile(file string) {
	err := os.Remove(file)
	if err != nil {
		logger.WithError(err).Fatal("Failed to clean file")
	}
}

func createDir(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.Mkdir(dir, 0755)
		if err != nil {
			logger.WithError(err).Fatal("Failed to created folder", dir)
		}
		logger.Info("Directory has been created", dir)
	} else {
		logger.Warn("Directoty already exists", dir)
	}

}

func removeDir(dir string) {
	err := os.Remove(dir)
	if err != nil {
		logger.WithError(err).Fatal("Failed to clean directory")
	} else {
		logger.Info("Successfully cleaned tmp dir")
	}
}
