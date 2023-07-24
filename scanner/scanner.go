package scanner

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/abhijitWakchaure/git-commit-finder/hashutils"
)

// Scan ...
func Scan(dir string) map[string]string {
	fmt.Println("Scanning directory:", dir)
	m := make(map[string]string)
	filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		filteredPath := strings.TrimPrefix(strings.TrimPrefix(path, dir), "/")
		if d.IsDir() {
			return nil
		}
		if strings.HasPrefix(filteredPath, ".") {
			return nil
		}
		if strings.HasPrefix(filteredPath, ".git") {
			return filepath.SkipDir
		}
		// fmt.Println("scanning path", filteredPath)
		md5, err := hashutils.MD5(path)
		if err != nil {
			fmt.Println("failed to calculate md5 for path:", path, " due to", err.Error())
			os.Exit(1)
		}
		m[filteredPath] = md5
		return nil
	})
	return m
}

// Compare ...
func Compare(m map[string]string, gitRepo string) error {
	// fmt.Println("Comparing with git repo:", gitRepo)
	return filepath.WalkDir(gitRepo, func(path string, d fs.DirEntry, err error) error {
		filteredPath := strings.TrimPrefix(strings.TrimPrefix(path, gitRepo), "/")
		if d.IsDir() {
			return nil
		}
		if strings.HasPrefix(filteredPath, ".") {
			return nil
		}
		if strings.HasPrefix(filteredPath, ".git") {
			return filepath.SkipDir
		}
		// fmt.Println("comparing path", filteredPath)
		oldMD5, ok := m[filteredPath]
		if !ok {
			return nil
		}
		newMD5, err := hashutils.MD5(path)
		if err != nil {
			fmt.Println("failed to calculate md5 for path:", path, " due to", err.Error())
			os.Exit(1)
		}
		if oldMD5 != newMD5 {
			// fmt.Println("MD5 mismatched for path:", path)
			return fmt.Errorf("MD5 mismatched for path: %s", path)
		}
		return nil
	})
}
