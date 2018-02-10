package util

import (
	"path/filepath"
	"log"
)

func HomePath() string {
	return AbsolutePath(".")
}

func AbsolutePath(relpath string) string {
	absolutePath, err := filepath.Abs(relpath)
	if err != nil {
		log.Println("current path error:", err)
	}
	return absolutePath
}


