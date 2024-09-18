package errs

import "errors"

type ignore struct{}

func (ei ignore) unexported() {}

// SimpleClctr is a simple implementation of the Collector interface.
type SimpleClctr struct {
	errs []error
}

// HasErrors returns true if any errors have been collected.
func (a *SimpleClctr) HasErrors() bool {
	return len(a.errs) > 0
}

func (a *SimpleClctr) Collect(err error) Ignore {
	if err != nil {
		a.errs = append(a.errs, err)
	}
	return ignore{}
}

func (a *SimpleClctr) CollectF(fn func() error) Ignore {
	return a.Collect(fn())
}

func (a *SimpleClctr) Errors() error {
	if len(a.errs) == 0 {
		return nil
	}
	return errors.Join(a.errs...)
}
