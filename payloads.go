package parallaxsdk

// PayloadGenUserAgent is the payload for generating a user agent.
type PayloadGenUserAgent struct {
	Auth   string `json:"auth"`   // The API key used for authenticating SDK requests.
	Site   string `json:"site"`   // The site's top-level domain (e.g., "com").
	Region string `json:"region"` // Defines the product type, with three possible values: captcha, interstitial, init.
	Pd     string `json:"pd"`     // Currently not important, any value works.
}

// PayloadGenDatadomeCookie is the payload for generating a DataDome cookie.
type PayloadGenDatadomeCookie struct {
	Auth        string                 `json:"auth"`        // The API key used for authenticating SDK requests.
	Site        string                 `json:"site"`        // Site for which to generate the cookie.
	Region      string                 `json:"region"`      // Site region.
	Proxyregion string                 `json:"proxyregion"` // Proxy region.
	Proxy       string                 `json:"proxy"`       // Proxy address.
	Pd          string                 `json:"pd"`          // Product type.
	Data        TaskDatadomeCookieData `json:"data"`        // Data required for cookie generation.
}

// PayloadGenPXCookie is the payload for generating PX cookies.
type PayloadGenPXCookie struct {
	Auth string `json:"auth"` // The API key used for authenticating SDK requests.
	TaskGeneratePXCookies
}

// PayloadGenHoldCaptcha is the payload for holding a captcha challenge.
type PayloadGenHoldCaptcha struct {
	Auth string `json:"auth"` // The API key used for authenticating SDK requests.
	TaskGenerateHoldCaptcha
}
