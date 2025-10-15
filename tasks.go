package parallaxsdk

// TaskGenUserAgent represents a task for generating a user agent data.
type TaskGenUserAgent struct {
	Site   string `json:"site"`   // The site's top-level domain (e.g., "com").
	Region string `json:"region"` // Defines the product type, with three possible values: captcha, interstitial, init.
}

// TaskDatadomeCookieData contains data required for generating a Datadome cookie.
type TaskDatadomeCookieData struct {
	Cid        string `json:"cid"`        // The DataDome cookie (mandatory for all products).
	E          string `json:"e"`          // Challenge parameter e (Required only for the captcha product).
	S          string `json:"s"`          // Challenge parameter s (Exists in both captcha and interstitial).
	B          string `json:"b"`          // Always exists for interstitial. Some sites also require it for captcha.
	InitialCid string `json:"initialCid"` // Present in captcha and interstitial responses; derived from DataDome's response block by splitting the response value at "cid:".
}

// TaskDatadomeCookie represents a task for generating a DataDome cookie.
type TaskDatadomeCookie struct {
	Site        string                 `json:"site"`        // Site for which to generate the cookie.
	Region      string                 `json:"region"`      // Site region.
	Proxyregion string                 `json:"proxyregion"` // Proxy region.
	Proxy       string                 `json:"proxy"`       // Proxy address.
	Pd          string                 `json:"pd"`          // Product type.
	Data        TaskDatadomeCookieData `json:"data"`        // Data required for cookie generation.
}

// TaskGeneratePXCookies represents a task for generating PX cookies.
type TaskGeneratePXCookies struct {
	Site        string `json:"site"`        // Site for which to generate PX cookies.
	Region      string `json:"region"`      // Site region.
	Proxyregion string `json:"proxyregion"` // Proxy region.
	Proxy       string `json:"proxy"`       // Proxy address.
}

// TaskGenerateHoldCaptcha represents a task for hold captcha challenge.
type TaskGenerateHoldCaptcha struct {
	Site        string `json:"site"`        // Site for which to solve hold captcha (e.g., "stockx").
	Region      string `json:"region"`      // The region of the site (e.g., "com" for .com sites or other TLDs like ".fr" or ".ch").
	Proxyregion string `json:"proxyregion"` // The region of your proxy (either "eu" or "us").
	Proxy       string `json:"proxy"`       // The proxy used for the request in HTTP format.
	Data        string `json:"data"`        // Hold captcha challenge data string.
	PowPro      string `json:"POW_PRO"`     // (Optional) Insert your Cuda POW solver key here.
}
