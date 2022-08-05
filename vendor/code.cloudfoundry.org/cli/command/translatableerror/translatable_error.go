package translatableerror

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . TranslatableError

// TranslatableError it wraps the error interface adding a way to set the
// translation function on the error
type TranslatableError interface {
	// Returns the untranslated error string
	Error() string
	Translate(func(string, ...interface{}) string) string
}
