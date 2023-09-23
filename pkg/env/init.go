package env

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

// Just import this package to get all the env variables at the root of the project
// Import this package anonymously as shown below:
// import _ "github.com/Goboolean/common/pkg/env"


func init() {
	path, err := filepath.Abs(".")
	if err != nil {
		panic(err)
	}

	for err := os.ErrNotExist; os.IsNotExist(err); _, err = os.Stat(filepath.Join(path, "go.mod")) {
		path = filepath.Dir(path)
		if path == "/" {
			panic(errRootNotFound)
		}
	}

	if err := os.Chdir(path); err != nil {
		panic(err)
	}

	if err := godotenv.Load(); err != nil {
		path, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		panic(fmt.Errorf("%v, working directory: %s", err, path))
	}
}

var errRootNotFound = errors.New("could not find root directory, be sure to set root of the project as fetch-server")
