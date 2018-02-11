package util

import (
	"path/filepath"
	"log"
	"strconv"
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

func HostAndPortToAddress(host string, port uint16) string {
	return host + ":" + strconv.Itoa(int(port))
}


