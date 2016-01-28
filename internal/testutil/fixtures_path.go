package testutil

import (
	"go/build"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

var FixtureBasePath = filepath.Join(
	build.Default.GOPATH,
	"src",
	"github.com/omise/omise-go",
	"testdata/fixtures",
)

func FixturePath(req *http.Request) (int, string, error) {
	fixpath := req.URL.Host + "/" +
		req.URL.Path[1:] + "-" + strings.ToLower(req.Method) + ".json"

	// resolve exact fixture filename
	filename := filepath.Join(FixtureBasePath, fixpath)
	if _, e := os.Lstat(filename); e != nil {
		if !os.IsNotExist(e) {
			return 500, "", e
		}

		return 404, filepath.Join(
			filepath.Dir(fixpath),
			"404-"+strings.ToLower(req.Method)+".json",
		), nil
	}

	return 200, fixpath, nil
}
