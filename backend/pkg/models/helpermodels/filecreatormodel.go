package helpermodels

import "os"

// FilePathManager is an interface that provides methods to retrieve directory and file name.
type FilePathManager interface {
	GetDirectory() string
	GetFileName() string
	GetFileExtension() string
}

// FilePathComponents is a struct that holds the directory and file name.
type FilePathComponents struct {
	Directory string
	FileName  string
	Extension string
}

// GetDirectory returns a directory.
func (f *FilePathComponents) GetDirectory() string {
	return f.Directory
}

// GetFileName returns a file name.
func (f *FilePathComponents) GetFileName() string {
	return f.FileName
}

// GetFileExtension returns a file extension.
func (f *FilePathComponents) GetFileExtension() string {
	return f.Extension
}

// FileCreator is an interface that abstracts the creation of a file.
type FileCreator interface {
	Create(name string) (*os.File, error)
}

// OSFileCreator is a concrete implementation of the FileCreator interface,
// wrapping the os.Create function.
type OSFileCreator struct{}

// Create wraps the os.Create function to create a file with the provided name.
func (f *OSFileCreator) Create(name string) (*os.File, error) {
	return os.Create(name)
}
