package internal

import (
	"backend/shared"
	"github.com/joho/godotenv"
	"os"
)

func Prepare() error {
	if err := shared.EnsureDockerProperlyInitialized(); err != nil {
		return err
	}

	_ = godotenv.Load()
	if err := prepareFs(); err != nil {
		return err
	}
	if err := shared.EnsureEnvAreSet(); err != nil {
		return err
	}

	return nil
}

func prepareFs() error {
	if err := os.MkdirAll(shared.VarLibPath, 0755); err != nil {
		return err
	}

	type fileT struct {
		path           string
		defaultContent string
	}

	filesToPrepare := []fileT{
		{shared.InstancesPath, "[]"},
	}

	for _, file := range filesToPrepare {
		if err := ensureFileExists(file.path, file.defaultContent); err != nil {
			return err
		}
	}

	return nil
}

func ensureFileExists(path string, defaultContent string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err := os.WriteFile(path, []byte(defaultContent), 0644); err != nil {
			return err
		}
	}

	return nil
}
