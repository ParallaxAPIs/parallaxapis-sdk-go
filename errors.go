package parallaxsdk

import (
	"encoding/json"
	"fmt"
)

type APIError struct{ Message string }

type ErrorEnv struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
}

type PXErrorDetails struct {
	Cookie         string `json:"cookie"`
	IsFlagged      bool   `json:"isFlagged"`
	IsMaybeFlagged bool   `json:"isMaybeFlagged"`
	FlaggedPow     bool   `json:"flaggedPOW"`
}

func formatPXErrors(body []byte) string {
	var v PXErrorDetails
	if err := json.Unmarshal(body, &v); err != nil {
		return string(body)
	}
	// add message fields
	return fmt.Sprintf(
		"message: '%v', isFlagged: %t, isMaybeFlagged: %t, flaggedPOW: %t",
		v.Cookie,
		v.IsFlagged,
		v.IsMaybeFlagged,
		v.FlaggedPow,
	)
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
