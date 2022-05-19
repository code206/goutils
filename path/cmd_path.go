package path

import (
	"errors"
	"os"
	"path/filepath"
)

func CmdDir() (string, error) {
	CmdDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return "", err
	} else {
		return CmdDir, nil
	}
}

func CmdName() (string, error) {
	CmdName := filepath.Base(os.Args[0])
	if CmdName == "" {
		return "", errors.New("cmd name is empty")
	} else {
		return CmdName, nil
	}
}
