package directory

import (
	"testing"
)

func Test_Get(t *testing.T) {
	tests := []struct {
		name       string
		rootDir    *Directory
		shouldFind bool
		dirToFind  string
	}{
		{"Test Get a single directory", NewDefaultDirectory("test", nil, nil), true, "test"},
		{"Test Get a single non-existing directory", NewDefaultDirectory("test", nil, nil), false, "test2"},
		{"Test Get a level 2 directory", NewDefaultDirectory("test", makeDirectories("test", "test2", "test3"), nil), true, "test2"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, ok := tt.rootDir.Get(tt.dirToFind)

			if tt.shouldFind != ok {
				t.Errorf("expected directory found: %v, but was %v", tt.shouldFind, ok)
			}
		})
	}
}

func Test_InsertAt(t *testing.T) {
	tests := []struct {
		name        string
		shouldPanic bool
		rootDir     *Directory
		insertAt    string
	}{
		{"Test when Insert and directory exists then should not panic", false, NewDefaultDirectory("test", nil, nil), "test"},
		{"Test when Insert and directory does not exist then should panic", true, NewDefaultDirectory("test", nil, nil), "notExistingDir"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assertPanic(t, tt.rootDir.InsertAt, tt.shouldPanic, tt.insertAt)
		})
	}
}

func assertPanic(t *testing.T, f func(string, *Directory), shouldPanic bool, insertAt string) {
	didPanic := true
	defer func() {
		if r := recover(); r == nil {
			didPanic = false
		}
		if shouldPanic != didPanic {
			t.Errorf("expected panic to be: %v, but was: %v", shouldPanic, didPanic)
		}
	}()
	f(insertAt, NewDefaultDirectory("insertedDir", nil, nil))
}

func makeDirectories(names ...string) []*Directory {
	dirs := []*Directory{}
	for _, n := range names {
		dirs = append(dirs, NewDefaultDirectory(n, nil, nil))
	}
	return dirs
}
