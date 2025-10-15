package parallaxsdk

// PayloadGenUserAgent is the payload for generating a user agent.
type PayloadGenUserAgent struct {
	Auth   string `json:"auth"` // The API key used for authenticating SDK requests.
	Site   string `json:"site"`
	Region string `json:"region"`
	Pd     string `json:"pd"`
}

// PayloadGenDatadomeCookie is the payload for generating a DataDome cookie.
type PayloadGenDatadomeCookie struct {
	Auth        string                 `json:"auth"`        // The API key used for authenticating SDK requests.
	Site        string                 `json:"site"`        // Site for which to generate the cookie.
	Region      string                 `json:"region"`      // Site region.
	Proxyregion string                 `json:"proxyregion"` // The region of your proxy (either "eu" or "us").
	Proxy       string                 `json:"proxy"`       // Proxy address.
	Pd          string                 `json:"pd"`          // Product type.
	Data        TaskDatadomeCookieData `json:"data"`        // Data required for cookie generation.
}

// PayloadGenPXCookie is the payload for generating PX cookies.
type PayloadGenPXCookie struct {
	Auth string `json:"auth"` // The API key used for authenticating SDK requests.
	TaskGeneratePXCookies
}

// PayloadGenHoldCaptcha is the payload for the holdcaptcha challenge.
type PayloadGenHoldCaptcha struct {
	Auth string `json:"auth"` // The API key used for authenticating SDK requests.
	TaskGenerateHoldCaptcha
}
