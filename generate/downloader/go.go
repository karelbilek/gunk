package downloader

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/rogpeppe/go-internal/lockedfile"

	"github.com/gunk/gunk/log"
)

type Go struct{}

func (g Go) Download(version string) (string, error) {
	if version == "" {
		// require version. this is used only with version explicitly set.
		return "", fmt.Errorf("must provide protoc-gen-go version")
	}

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

	gitCloneDir := filepath.Join(cachePath, "gunk", fmt.Sprintf("protoc-gen-go-%s", version))


	repoPath, cmdBuildPath, err := g.repoPath(version)
	if err != nil {
		return "", err
	}

	cmdBuildPath = append([]string{gitCloneDir}, cmdBuildPath...)

	binaryDir := filepath.Join(cmdBuildPath...)
	binaryPath := filepath.Join(binaryDir, "protoc-gen-go")

	// lock cannot be in the dir where we git clone, otherwise git is unhappy
	lockPath := gitCloneDir + ".lock"

	// First, grab a lock separate from the destination file. The
	// destination file is a binary we'll want to execute, so using it
	// directly as the lock can lead to "text file busy" errors.
	unlock, err := lockedfile.MutexAt(lockPath).Lock()
	if err != nil {
		panic(err)
	}
	defer unlock()

	_, fErr := os.Stat(binaryPath)
	if !os.IsNotExist(fErr) {
		if fErr != nil {
			return "", fErr
		}

		return binaryPath, nil
	}

	// remove the whole git clone dir if something happened wrong
	_, fErr = os.Stat(gitCloneDir)
	if !os.IsNotExist(fErr) {
		if fErr != nil {
			return "", fErr
		}

		if err := os.RemoveAll(gitCloneDir); err != nil {
			return "", err
		}
	}


	cmdArgs := []string{"clone", "--depth", "1", "--branch", version, repoPath, gitCloneDir}

	gitCmd := log.ExecCommand("git", cmdArgs...)
	err = gitCmd.Run()
	if err != nil {
		all := "git " + strings.Join(cmdArgs, " ")
		return "", log.ExecError(all, err)
	}

	buildCmd := log.ExecCommand("go", "build")
	buildCmd.Dir = binaryDir

	var stderr bytes.Buffer
	buildCmd.Stderr = &stderr

	err = buildCmd.Run()
	if err != nil {
		all := "go build"
		return "", log.ExecError(all, err)
	}

	return binaryPath, nil
}


func (Go) repoPath(version string) (string, []string, error) {
	const oldPath = `https://github.com/golang/protobuf`
	const newPath = `https://github.com/protocolbuffers/protobuf-go`
	var oldSrcPath = []string{`protoc-gen-go`}
	var newSrcPath = []string{`cmd`,`protoc-gen-go`}

	version = strings.TrimPrefix(version, "v")
	split := strings.Split(version, ".")
	if len(split) != 3 {
		return "", nil, fmt.Errorf("cannot interpret %q as version number: not 3 parts", split)
	}
	major, err := strconv.Atoi(split[0])
	if err != nil {
		return "",  nil, fmt.Errorf("cannot interpret %q as version number: major not a number", split)
	}
	minor, err := strconv.Atoi(split[1])
	if err != nil {
		return "", nil, fmt.Errorf("cannot interpret %q as version number: major not a number", split)
	}
	if (major > 1) {
		return newPath, newSrcPath, nil
	}
	if minor >=20 {
		return newPath, newSrcPath, nil
	}
	return oldPath, oldSrcPath, nil
}