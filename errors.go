package gop5js

type errorContainer struct {
	errors []error
}

func (ec *errorContainer) add(err error) {
	if err != nil {
		ec.errors = append(ec.errors, err)
	}
}

// Errors let check you for errors of the Error container
// The first parameter returns true if there is an error
// The second parameter returns all errors over a slice
func (ec *errorContainer) Errors() (bool, []error) {
	if len(ec.errors) > 0 {
		return true, ec.errors
	}
	return false, nil
}
