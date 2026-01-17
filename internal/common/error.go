package common

// Error types
const (
	ErrorTypeMaxRetry           = "MAX_RETRY"
	ErrorTypeNoTokens           = "NO_TOKENS"
	ErrorTypeUserNotFound       = "USER_NOT_FOUND"
	ErrorTypeGraphQLError       = "GRAPHQL_ERROR"
	ErrorTypeGitHubRESTAPIError = "GITHUB_REST_API_ERROR"
	ErrorTypeWakaTimeError      = "WAKATIME_ERROR"
)

const TryAgainLater = "Please try again later"

// SecondaryErrorMessages maps error types to secondary error messages
var SecondaryErrorMessages = map[string]string{
	ErrorTypeMaxRetry:           "You can deploy own instance or wait until public will be no longer limited",
	ErrorTypeNoTokens:           "Please add an env variable called PAT_1 with your GitHub API token in vercel",
	ErrorTypeUserNotFound:       "Make sure the provided username is not an organization",
	ErrorTypeGraphQLError:       TryAgainLater,
	ErrorTypeGitHubRESTAPIError: TryAgainLater,
	ErrorTypeWakaTimeError:      "Make sure you have a public WakaTime profile",
}

// CustomError represents a custom error with type and secondary message
type CustomError struct {
	Message          string
	Type             string
	SecondaryMessage string
}

func (e *CustomError) Error() string {
	return e.Message
}

// NewCustomError creates a new CustomError
func NewCustomError(message, errorType string) *CustomError {
	secondaryMsg := SecondaryErrorMessages[errorType]
	if secondaryMsg == "" {
		secondaryMsg = errorType
	}
	return &CustomError{
		Message:          message,
		Type:             errorType,
		SecondaryMessage: secondaryMsg,
	}
}

// MissingParamError represents a missing parameter error
type MissingParamError struct {
	Message          string
	MissedParams     []string
	SecondaryMessage string
}

func (e *MissingParamError) Error() string {
	return e.Message
}

// NewMissingParamError creates a new MissingParamError
func NewMissingParamError(missedParams []string, secondaryMessage string) *MissingParamError {
	msg := "Missing params "
	for i, p := range missedParams {
		if i > 0 {
			msg += ", "
		}
		msg += `"` + p + `"`
	}
	msg += " make sure you pass the parameters in URL"
	return &MissingParamError{
		Message:          msg,
		MissedParams:     missedParams,
		SecondaryMessage: secondaryMessage,
	}
}

// RetrieveSecondaryMessage retrieves secondary message from an error
func RetrieveSecondaryMessage(err error) string {
	if customErr, ok := err.(*CustomError); ok {
		return customErr.SecondaryMessage
	}
	if missingErr, ok := err.(*MissingParamError); ok {
		return missingErr.SecondaryMessage
	}
	return ""
}
