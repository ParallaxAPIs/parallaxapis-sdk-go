package parallaxsdk

// TaskGenUserAgent represents a task for generating useragent data.
type TaskGenUserAgent struct {
	Site   string `json:"site"`
	Region string `json:"region"`
}

// TaskDatadomeCookieData contains data required for generating a Datadome cookie.
type TaskDatadomeCookieData struct {
	Cid        string `json:"cid"`
	E          string `json:"e"`
	S          string `json:"s"`
	B          string `json:"b"`
	InitialCid string `json:"initialCid"`
}

// TaskDatadomeTagsCookie represents a task for generating a DataDome tags cookie.
type TaskDatadomeTagsCookie struct {
	Site        string                 `json:"site"`        // Site for which to generate the cookie.
	Region      string                 `json:"region"`      // Site region.
	Proxyregion string                 `json:"proxyregion"` // The region of your proxy (either "eu" or "us").
	Proxy       string                 `json:"proxy"`       // Proxy address.
	Data        TaskDatadomeCookieData `json:"data"`        // Data required for cookie generation.
}

// TaskDatadomeCookie represents a task for generating a DataDome cookie.
type TaskDatadomeCookie struct {
	Site        string                 `json:"site"`        // Site for which to generate the cookie.
	Region      string                 `json:"region"`      // Site region.
	Proxyregion string                 `json:"proxyregion"` // The region of your proxy (either "eu" or "us").
	Proxy       string                 `json:"proxy"`       // Proxy address.
	Pd          string                 `json:"pd"`          // Product type.
	Data        TaskDatadomeCookieData `json:"data"`        // Data required for cookie generation.
}

// TaskGeneratePXCookies represents a task for generating PX cookies.
type TaskGeneratePXCookies struct {
	Site        string `json:"site"`        // Site for which to generate cookies.
	Region      string `json:"region"`      // Site region.
	Proxyregion string `json:"proxyregion"` // Proxy region.
	Proxy       string `json:"proxy"`       // Proxy address.
}

// TaskGenerateHoldCaptcha represents a task for hold captcha challenge.
type TaskGenerateHoldCaptcha struct {
	Site        string `json:"site"`
	Region      string `json:"region"`
	Proxyregion string `json:"proxyregion"` // The region of your proxy (either "eu" or "us").
	Proxy       string `json:"proxy"`
	Data        string `json:"data"`    // Data required for cookie generation.
	PowPro      string `json:"POW_PRO"` // (Optional) Insert your Cuda POW solver key here.
}
