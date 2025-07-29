package main

import (
	"io"
	"log"
	"os"
	"path/filepath"
)

func copyFile(src, dst string) error {
	source, err := os.Open(src)
	if err != nil {
		return err
	}
	defer source.Close()

	err = os.MkdirAll(filepath.Dir(dst), 0755)
	if err != nil {
		return err
	}

	destination, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destination.Close()

	_, err = io.Copy(destination, source)
	return err
}

func movePublicToBuild() {
	publicDir := "public"
	buildDir := "build"

	if _, err := os.Stat(publicDir); os.IsNotExist(err) {
		return // no public folder, nothing to do
	}

	err := filepath.Walk(publicDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relPath, err := filepath.Rel(publicDir, path)
		if err != nil {
			return err
		}

		destPath := filepath.Join(buildDir, relPath)

		if info.IsDir() {
			return os.MkdirAll(destPath, 0755)
		}

		return copyFile(path, destPath)
	})

	if err != nil {
		log.Fatalf("failed to copy public files to build folder : %v", err)
	}
}
