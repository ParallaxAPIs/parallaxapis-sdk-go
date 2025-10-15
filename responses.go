package parallaxsdk

// UserAgentResponse is the response type for user agent generation.
type UserAgentResponse struct {
	Message            string `json:"message"`            // Error or status message.
	UserAgent          string `json:"UserAgent"`          // The generated user agent string.
	SecHeader          string `json:"secHeader"`          // SecHeader header value.
	SecFullVersionList string `json:"secFullVersionList"` // SecFullVersionList header value.
	SecPlatform        string `json:"secPlatform"`        // SecPlatform header value.
	SecArch            string `json:"secArch"`            // SecArch header value.
}

// DatadomeCookieResponse is the response type for DataDome cookie generation.
type DatadomeCookieResponse struct {
	Message   string `json:"message"`   // Error or status message.
	UserAgent string `json:"UserAgent"` // The user agent string used for cookie generation.
}

// PxCookieResponse is the response type for PX cookies generation.
type PxCookieResponse struct {
	Message        string `json:"message"`        // Error or status message.
	Cookie         string `json:"cookie"`         // Includes the generated cookie or an error message if error is true.
	Vid            string `json:"vid"`            // Used to set the _pxvid cookie.
	Cts            string `json:"cts"`            // Used to set the pxcts cookie.
	IsFlagged      bool   `json:"isFlagged"`      // Indicate if the generation might have been flagged during generation.
	IsMaybeFlagged bool   `json:"isMaybeFlagged"` // Indicate if the generation might have been flagged during generation.
	UserAgent      string `json:"UserAgent"`      // The device used for generation.
	Data           string `json:"data"`           // A string for backend purposes, used to generate the next step.
}

// GenHoldCaptchaResponse is the response type for holding a captcha challenge.
type GenHoldCaptchaResponse struct {
	PxCookieResponse
	FlaggedPow bool `json:"flaggedPOW"` // Indicates if pow is flagged.
}
