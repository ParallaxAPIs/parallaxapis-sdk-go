package parallaxsdk

type HtmlScriptObject struct {
	B      int    `json:"b"`
	Rt     string `json:"rt"`
	Cid    string `json:"cid"`
	Hsh    string `json:"hsh"`
	T      string `json:"t"`
	Qp     string `json:"qp"`
	S      int    `json:"s"`
	E      string `json:"e"`
	Host   string `json:"host"`
	Cookie string `json:"cookie"`
}

type JsonDatadomeBlockBody struct {
	URL string `json:"url"`
}
