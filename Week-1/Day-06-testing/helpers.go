package main

import (
	"os"
	"testing"
)

func FileExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}

	return !info.IsDir()
}

func saveAll(t *testing.T, path string) {
	t.Helper()
	if err := loadedAll.Save(path); err != nil {
		t.Fatalf("failed to save file: %v", err)
	}
}
