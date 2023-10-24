package testutil

import (
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func FixtureBasePath() string {
	_, b, _, _ := runtime.Caller(0)
	parentDir := strings.Replace(filepath.Dir(b), "internal/testutil", "", -1)
	return filepath.Join(parentDir, "testdata/fixtures")
}

func FixturePath(req *http.Request) (int, string, error) {
	fixpath := req.URL.Host + "/" +
		req.URL.Path[1:] + "-" + strings.ToLower(req.Method) + ".json"

	// resolve exact fixture filename
	filename := filepath.Join(FixtureBasePath(), fixpath)

	if _, err := os.Lstat(filename); err != nil {
		if !os.IsNotExist(err) {
			return 500, "", err
		}

		return 404, filepath.Join(
			filepath.Dir(fixpath),
			"404-"+strings.ToLower(req.Method)+".json",
		), nil
	}

	return 200, fixpath, nil
}
