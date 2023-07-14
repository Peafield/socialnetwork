package helpermodels

import "os"

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
