// DO NOT TOUCH UNLESS YOU KNOW WHAT YOU ARE DOING. We are using this filepath
// to serve the client distribution folder that was generated upon running `npm run build`

package config

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-chi/chi"
)

var (
	// WorkDir dynamically fetches the working directory
	WorkDir, _ = os.Getwd()

	// FilesDir points to the client distribution folder
	// Use an environment variable for flexibility, with a fallback default
	FilesDir = http.Dir(func() string {
		if path := os.Getenv("CLIENT_DIST_PATH"); path != "" {
			return path
		}
		return filepath.Join(WorkDir, "../../client/dist") // Default fallback for local development
	}())
)

// FileServer serves static files from a given directory.
// The root path ("/") can end with a slash, but non-root paths will redirect to remove the trailing slash.
func FileServer(r chi.Router, path string, root http.FileSystem) {
	if path == "/" {
		// Serve files from the root directory
		fs := http.FileServer(root)
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			fs.ServeHTTP(w, r)
		})
		r.Get("/*", func(w http.ResponseWriter, r *http.Request) {
			fs.ServeHTTP(w, r)
		})
		return
	}

	// Handle non-root paths
	fs := http.StripPrefix(path, http.FileServer(root))

	// Redirect requests with trailing slash to the path without it
	r.Get(path+"/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, strings.TrimSuffix(r.URL.Path, "/"), http.StatusMovedPermanently)
	})

	// Serve static files for requests without a trailing slash
	r.Get(path+"*", func(w http.ResponseWriter, r *http.Request) {
		fs.ServeHTTP(w, r)
	})
}
