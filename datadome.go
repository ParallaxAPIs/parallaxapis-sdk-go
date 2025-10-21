package parallaxsdk

import (
	"encoding/json"
	"fmt"
	"net/url"
	"regexp"
	"strings"
)

var (
	// Datadome block url regexp
	datadomeBlockUrlRe = regexp.MustCompile(`geo\.captcha\-delivery\.com\/(?:interstitial|captcha)`)
	// Datadome script object regexp
	datadomeHtmlScriptRe = regexp.MustCompile(`dd=\{[^}]+\}`)
	// Regexp for single quoted keys
	singleQuotedKeyRe = regexp.MustCompile(`'((?:[^'\\]|\\.)*)'\s*:`)
	// Regexp for single quoted values
	singleQuotedValueRe = regexp.MustCompile(`:\s*'([^'\\]*(?:\\.[^'\\]*)*)'`)
)

const (
	// Perma block
	T_BV = "bv"
	// Captcha
	T_FE = "fe"
	// Intersitial
	T_IT = "it"
)

// GenerateUserAgent generates a user agent data.
func (s *SDK) GenerateUserAgent(task TaskGenUserAgent) (*UserAgentResponse, error) {
	reqBody := PayloadGenUserAgent{
		Auth:   s.AuthKey,
		Site:   task.Site,
		Region: task.Region,
	}
	var resp UserAgentResponse
	if err := s.request("/useragent", reqBody, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GenerateDatadomeCookie generates a DataDome cookie using the provided task parameters.
func (s *SDK) GenerateDatadomeCookie(task TaskDatadomeCookie) (*DatadomeCookieResponse, error) {
	reqBody := PayloadGenDatadomeCookie{
		Auth:        s.AuthKey,
		Site:        task.Site,
		Region:      task.Region,
		Proxyregion: task.Proxyregion,
		Proxy:       task.Proxy,
		Pd:          task.Pd,
		Data:        task.Data,
	}
	var resp DatadomeCookieResponse
	if err := s.request("/gen", reqBody, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GenerateDatadomeTagsCookie generates a DataDome tags cookie using the provided task parameters.
func (s *SDK) GenerateDatadomeTagsCookie(task TaskDatadomeTagsCookie) (*DatadomeCookieResponse, error) {
	reqBody := PayloadGenDatadomeCookie{
		Auth:        s.AuthKey,
		Site:        task.Site,
		Region:      task.Region,
		Proxyregion: task.Proxyregion,
		Proxy:       task.Proxy,
		Pd:          PD_Init,
		Data: TaskDatadomeCookieData{
			Cid: task.Cid,
		},
	}
	var resp DatadomeCookieResponse
	if err := s.request("/gen", reqBody, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ParseChallengeURL parses a DataDome challenge URL and extracts the challenge data and product type.
func ParseChallengeURL(challengeURL, prevCookie string) (*TaskDatadomeCookieData, string, error) {
	u, err := url.Parse(challengeURL)
	if err != nil {
		return nil, "", ErrInvalidChallengeURL
	}

	var pd string
	switch {
	case strings.HasPrefix(u.Path, "/captcha"):
		pd = PD_Captcha
	case strings.HasPrefix(u.Path, "/interstitial"):
		pd = PD_Interstitial
	case strings.HasPrefix(u.Path, "/init"):
		pd = PD_Init
	default:
		return nil, "", ErrUnknownChallengeType
	}

	q := u.Query()

	if t := q.Get("t"); t == T_BV {
		return nil, "", ErrPermanentlyBlocked
	}

	b := q.Get("b")

	if b == "" {
		b = "0"
	}

	return &TaskDatadomeCookieData{
		B:          b,
		S:          q.Get("s"),
		E:          q.Get("e"),
		Cid:        prevCookie,
		InitialCid: q.Get("initialCid"),
	}, pd, nil
}

// ParseChallengeJSON parses a JSON body containing a DataDome challenge URL.
func ParseChallengeJSON(jsonBody, prevCookie string) (*TaskDatadomeCookieData, string, error) {
	var body JsonDatadomeBlockBody

	if err := json.Unmarshal([]byte(jsonBody), &body); err != nil || len(body.URL) == 0 {
		return nil, "", ErrUnparsableDatadomeJSONBody
	}

	return ParseChallengeURL(body.URL, prevCookie)
}

// ParseChallengeHTML parses HTML and extracts DataDome challenge data from a JS object.
func ParseChallengeHTML(htmlBody, prevCookie string) (*TaskDatadomeCookieData, string, error) {
	match := datadomeHtmlScriptRe.FindString(htmlBody)

	if match == "" {
		return nil, "", ErrNoDatadomeValuesInHtml
	}

	objStr := match[3:] // skip 'dd='

	objStr = singleQuotedKeyRe.ReplaceAllString(objStr, `"$1":`)
	objStr = singleQuotedValueRe.ReplaceAllString(objStr, `:"$1"`)

	var dd HtmlScriptObject

	if err := json.Unmarshal([]byte(objStr), &dd); err != nil {
		return nil, "", ErrNoDatadomeValuesInHtml
	}

	var pd string
	switch dd.T {
	case T_IT:
		pd = PD_Interstitial
	case T_FE:
		pd = PD_Captcha
	case T_BV:
		return nil, "", ErrPermanentlyBlocked
	default:
		return nil, "", ErrUnknownChallengeType
	}

	return &TaskDatadomeCookieData{
		B:          fmt.Sprint(dd.B),
		S:          fmt.Sprint(dd.S),
		E:          dd.E,
		Cid:        prevCookie,
		InitialCid: dd.Cid,
	}, pd, nil
}

// DetectChallengeAndParse detects the challenge type and parses accordingly.
func DetectChallengeAndParse(body, prevCookie string) (bool, *TaskDatadomeCookieData, string, error) {

	if datadomeHtmlScriptRe.MatchString(body) {
		task, pd, err := ParseChallengeHTML(body, prevCookie)

		if err != nil {
			return true, nil, "", err
		}

		return true, task, pd, nil
	} else if datadomeBlockUrlRe.MatchString(body) {
		task, pd, err := ParseChallengeJSON(body, prevCookie)

		if err != nil {
			return true, nil, "", err
		}

		return true, task, pd, nil
	}

	return false, nil, "", nil
}
