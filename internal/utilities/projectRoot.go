package utilities

import (
	"log"
	"os"
	"path/filepath"
)

// FindProjectRoot finds the root directory of the project by locating the go.mod file.
// It starts from the current working directory and walks up the tree.
func FindProjectRoot() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for {
		// Check if go.mod exists in the current directory
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir, nil // Found the project root
		}

		// Move up to the parent directory
		parent := filepath.Dir(dir)
		if parent == dir {
			// Reached the filesystem root, go.mod not found
			return "", os.ErrNotExist
		}
		dir = parent
	}
}

func GetTemplateFileFN (templateRelativeFN string) string {
	projectRoot, err := FindProjectRoot()
	if (err != nil) {
        log.Fatal("getTemplateFileFN: Error calculating project's home directory")		
	}

	templateFileFN := filepath.Join(projectRoot, "internal", "templates", "root", templateRelativeFN)
	return templateFileFN
}
