package constants

// DefaultScheme is the scheme that should be prepended to user-provided
// URLs that do not specify a scheme.
const DefaultScheme string = "http"

// SupportedUrlSchemes contains the schemes that the goget application
// currently supports. It is intended to be used for checking if user-provided
// URLs are suitable.
var SupportedUrlSchemes = map[string]struct{}{
	"http":  struct{}{},
	"https": struct{}{},
}
