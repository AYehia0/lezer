package utils

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

type Item struct {
	Type string
	Path string
	Name string
}

// We want a function which take "RootPath": Current working directory path, and a name maybe directory name (DirName)
// and returns a full path to this dir
// examples :
// - func ("/home/none/Things/github/", "/Gimme") --> "/home/none/Things/github/Gimme/"
// - func ("/home/none/Things/github/Gimme", "/SomeDirNotInGithub") --> "/home/none/Things/github/Gimme/"
type FileItem struct {
	Root string // just the root "/"
	Dir  string
	Base string
	Ext  string
	Name string
}

func Scan(startPath, dir string) []FileItem {
	var items []FileItem

	currentAbsPath, _ := filepath.Abs(dir)

	toStart, _ := filepath.Rel(dir, currentAbsPath)

	log.Printf("ABS: %s \nTo: %s\n", currentAbsPath, toStart)

	return items
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

func MergePaths(root, curr string) (string, error) {
	if root == curr {
		return filepath.Abs(root)
	}
	return filepath.Abs(filepath.Join(root, curr))
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

		var newPath string
		if filepath.Base(baseDir) == filepath.Base(serverDir) {
			newPath = entry.Name()
		} else {
			newPath = fmt.Sprintf("%s/%s", filepath.Base(baseDir), entry.Name())
		}

		item := Item{
			Type: itemType,
			Path: newPath,
			Name: entry.Name(),
		}

		items = append(items, item)
	}

	return items, nil
}
