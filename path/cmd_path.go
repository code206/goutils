package path

import (
	"errors"
	"os"
	"path/filepath"
)

func BinDirPath() (string, error) {
	BinDirPath, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return "", err
	} else {
		return BinDirPath, nil
	}
}

func BinName() (string, error) {
	BinName := filepath.Base(os.Args[0])
	if BinName == "" {
		return "", errors.New("bin name is empty")
	} else {
		return BinName, nil
	}
}

func BinDirName() (string, error) {
	BinDirName := filepath.Base(filepath.Dir(os.Args[0]))
	if BinDirName == "" {
		return "", errors.New("bin dir name is empty")
	} else {
		return BinDirName, nil
	}
}
