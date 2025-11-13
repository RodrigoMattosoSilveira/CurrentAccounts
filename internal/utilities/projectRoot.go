package utilities

import (
	"log"
	"path/filepath"
)

// GetProjectRoot returns the root directory of the project

func GetProjectRoot() string {
	// Logic to determine the project root directory
    // Determine project root (2 levels up from cmd/web)
    projectRoot, err := filepath.Abs(filepath.Join(filepath.Dir("."), "../../"))
    if err != nil {
        log.Fatalf("failed to resolve project root: %v", err)
    }
	return projectRoot
}	