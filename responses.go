package parallaxsdk

// UserAgentResponse is the response type for user agent generation.
type UserAgentResponse struct {
	Message            string `json:"message"`
	UserAgent          string `json:"UserAgent"`
	SecHeader          string `json:"secHeader"`
	SecFullVersionList string `json:"secFullVersionList"`
	SecPlatform        string `json:"secPlatform"`
	SecArch            string `json:"secArch"`
}

// DatadomeCookieResponse is the response type for DataDome cookie generation.
type DatadomeCookieResponse struct {
	Message   string `json:"message"`
	UserAgent string `json:"UserAgent"`
}

// PxCookieResponse is the response type for PX cookies generation.
type PxCookieResponse struct {
	Message        string `json:"message"`
	Cookie         string `json:"cookie"`
	Vid            string `json:"vid"`            // Used to set the _pxvid cookie.
	Cts            string `json:"cts"`            // Used to set the pxcts cookie.
	IsFlagged      bool   `json:"isFlagged"`      // Indicate if the generation might have been flagged during generation.
	IsMaybeFlagged bool   `json:"isMaybeFlagged"` // Indicate if the generation might have been flagged during generation.
	UserAgent      string `json:"UserAgent"`      // The device used for generation.
	Data           string `json:"data"`           // A string used to generate the next step.
}

// GenHoldCaptchaResponse is the response type for holdingcaptcha challenge.
type GenHoldCaptchaResponse struct {
	PxCookieResponse
	FlaggedPow bool `json:"flaggedPOW"` // Indicates if pow is flagged.
}
