package utils

import (
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

type Item struct {
	Type string
	Name string
}

// expand the relative path that starts with ~ (ex : ~/path/ --> /home/{username}/path/)
func expandPath(path string) (string, error) {
	if len(path) == 0 || path[0] != '~' {
		return path, nil
	}

	usr, err := user.Current()
	if err != nil {
		return "", err
	}
	return filepath.Join(usr.HomeDir, path[1:]), nil

}

// The path must be an existing path in the system and must reference a directory not a file (ex: /home/name/somefile)
func IsValidPath(path string) bool {
	path, err := expandPath(path)
	if err != nil {
		return false
	}
	if stat, err := os.Stat(path); err == nil && stat.IsDir() {
		return true
	}
	return false
}

// returns the full-path of the files and dir
func GetAllInDir(serverDir, baseDir string, ignoreHidden bool) ([]Item, error) {
	// retunr [] items, item can be dir, or file (not dir)
	var items []Item
	entries, err := os.ReadDir(baseDir)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		var itemType string
		if entry.IsDir() {
			itemType = "folder"
		} else {
			itemType = "file"
		}

		if ignoreHidden && strings.HasPrefix(entry.Name(), ".") {
			continue
		}

		item := Item{
			Type: itemType,
			Name: entry.Name(),
		}

		items = append(items, item)
	}

	return items, nil
}
