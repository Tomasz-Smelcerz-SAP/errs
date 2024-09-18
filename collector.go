package errs

// Collector is an interface for collecting errors.
// Unlike the idiomatic Go convention of returning errors as values,
// this pattern involves passing a Collector to functions to handle errors by collecting them.
type Collector interface {
	// Collect collects a single error.
	Collect(err error) Ignore

	// CollectF collects an error returned by a parameterless function.
	// This is very useful in defer statements - you can provide a one liner
	// that ensures the function will be called and any error it returns will be collected.
	CollectF(func() error) Ignore
}

// Ignore is a return type replacement for the standard golang "error" type returned from functions.
// This value, as the name iplies, should be ignored by the caller.
// The purpose of having this return value is to force the compiler to complain in case the user forgets to collect errors.
// In this way we put some constraints on the user, but it's for greater good - to prevent sneaky, silent "error not handled" type of errors.
type Ignore interface {
	unexported() // to prevent external implementations
}
