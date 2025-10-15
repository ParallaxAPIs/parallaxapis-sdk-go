package parallaxsdk

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

func (sdk *PerimeterxSDK) GenerateHoldCaptcha(task TaskGenerateHoldCaptcha) (*GenHoldCaptchaResponse, error) {
	reqBody := PayloadGenHoldCaptcha{
		Auth:                    sdk.AuthKey,
		TaskGenerateHoldCaptcha: task,
	}

	var resp GenHoldCaptchaResponse

	if err := sdk.request("/holdcaptcha", reqBody, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
