package downloader

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/rogpeppe/go-internal/lockedfile"
)

type GrpcJava struct{}

func (pd GrpcJava) Download(version string) (string, error) {

	// Get the OS-specific cache directory.
	cachePath, err := os.UserCacheDir()
	if err != nil {
		return "", err
	}
	if dir := os.Getenv("GUNK_CACHE_DIR"); dir != "" {
		// Allow overriding the cache dir entirely. Mainly for
		// the tests.
		cachePath = dir
	}
	cacheDir := filepath.Join(cachePath, "gunk")
	if err := os.MkdirAll(cacheDir, 0755); err != nil {
		return "", err
	}

	// TODO remove shared functionality with protoc.go

	// The proto command path to use or download to.
	dstPath := filepath.Join(cacheDir, fmt.Sprintf("protoc-grpc-java-%s", version))

	// First, grab a lock separate from the destination file. The
	// destination file is a binary we'll want to execute, so using it
	// directly as the lock can lead to "text file busy" errors.
	unlock, err := lockedfile.MutexAt(dstPath + ".lock").Lock()
	if err != nil {
		panic(err)
	}
	defer unlock()

	// We are the only goroutine with access to dstPath. Check if it already
	// exists. Using lockedfile.OpenFile allows us to do an atomic write
	// when it doesn't exist yet.
	// TODO: isn't this lock entirely superfluous? The first lock already blocks the whole time
	dstFile, err := lockedfile.OpenFile(dstPath, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0775)
	if os.IsExist(err) {
		// It exists. Because of O_EXCL, we haven't actually opened the
		// file. Just return.
		return dstPath, nil
	}
	if err != nil {
		return "", err
	}
	defer dstFile.Close()

	// The file does not exist. Download it, using dstFile.
	url, err := pd.downloadURL(runtime.GOOS, runtime.GOARCH, version)
	if err != nil {
		return "", fmt.Errorf("downloading grpc-java: %w", err)
	}

	cl := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	res, err := cl.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return "", fmt.Errorf("could not retrieve %q (%d)", url, res.StatusCode)
	}

	// Write command to cache.
	if _, err := io.Copy(dstFile, res.Body); err != nil {
		return "", err
	}

	if err := dstFile.Close(); err != nil {
		return "", err
	}

	return dstPath, nil
}

func (GrpcJava) downloadURL(os, arch, version string) (string, error) {
	if !strings.HasPrefix(version, "v") {
		return "", fmt.Errorf("invalid version: %s", version)
	}
	version = version[1:]

	// determine the platform
	var platform string
	switch {
	case os == "darwin" && arch == "386":
		return "", fmt.Errorf("macOS 386 not supported")
	case os == "darwin" && arch == "amd64":
		platform = "osx-x86_64"

	case os == "linux" && arch == "386":
		platform = "linux-x86_32"
	case os == "linux" && arch == "amd64":
		platform = "linux-x86_64"
	case os == "linux" && arch == "arm64":
		platform = "linux-aarch_64"

	case os == "windows" && arch == "386":
		platform = "windows-x86_32"
	case os == "windows" && arch == "amd64":
		platform = "windows-x86_64"

	default:
		return "", fmt.Errorf("unknown os %q and arch %q", os, arch)
	}
	const mavenRepo = `https://repo1.maven.org/maven2/io/grpc/protoc-gen-grpc-java`
	return fmt.Sprintf("%s/%s/protoc-gen-grpc-java-%s-%s.exe",
		mavenRepo,
		version,
		version,
		platform), nil
}
