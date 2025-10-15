package parallaxsdk

// GenerateCookies generates PX cookies using the provided task parameters.
// Sends a POST request to the /gen endpoint with the provided task parameters.
// Returns the generated PX cookies response or an error if the API responds with an error.
//
// task: The task object containing parameters for PX cookie generation.
// Returns: Pointer to PxCookieResponse and error.
func (sdk *PerimeterxSDK) GenerateCookies(task TaskGeneratePXCookies) (*PxCookieResponse, error) {
	reqBody := PayloadGenPXCookie{
		Auth:                  sdk.AuthKey,
		TaskGeneratePXCookies: task,
	}

	var resp PxCookieResponse

	if err := sdk.request("/gen", reqBody, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// GenerateHoldCaptcha generates a hold captcha response using the provided task parameters.
// Sends a POST request to the /gen endpoint with the provided task parameters.
// Returns the generated hold captcha response or an error if the API responds with an error.
//
// task: The task object containing parameters for hold captcha generation.
// Returns: Pointer to GenHoldCaptchaResponse and error.
func (sdk *PerimeterxSDK) GenerateHoldCaptcha(task TaskGenerateHoldCaptcha) (*GenHoldCaptchaResponse, error) {
	reqBody := PayloadGenHoldCaptcha{
		Auth:                    sdk.AuthKey,
		TaskGenerateHoldCaptcha: task,
	}

	var resp GenHoldCaptchaResponse

	if err := sdk.request("/gen", reqBody, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
