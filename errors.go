package parallaxsdk

import "fmt"

type APIError struct{ Message string }

type ErrorEnv struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Cookie  string `json:"cookie"`
}

func (e *APIError) Error() string {
	return fmt.Sprintf("Parallax API error: %s", e.Message)
}

// ErrUnknownChallengeType means we couldn’t identify captcha/interstitial/init.
var ErrUnknownChallengeType = fmt.Errorf("unknown challenge type in URL")

// ErrInvalidChallengeURL means the URL failed to parse.
var ErrInvalidChallengeURL = fmt.Errorf("invalid challenge URL")

// ErrPermanentlyBlocked means the challenge indicates a permanent block (t=bv)
var ErrPermanentlyBlocked = fmt.Errorf("permanently blocked by DataDome (t=bv)")

// ErrUnparsableDatadomeBody means the JSON body could not be parsed or did not contain a URL
var ErrUnparsableDatadomeJSONBody = fmt.Errorf("unparsable DataDome JSON body")

// ErrNoDatadomeValuesInHtml means no Datadome values were found in the HTML body
var ErrNoDatadomeValuesInHtml = fmt.Errorf("no DataDome values in HTML body")

// ErrUnparsableDatadomeHTMLBody means no sdk couldn't parse
var ErrUnparsableDatadomeHTMLBody = fmt.Errorf("no DataDome values in HTML body")
