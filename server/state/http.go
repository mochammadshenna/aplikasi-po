package state

var _httpHeaders httpHeaders
var _httpContentTypeValues httpContentTypeValues

func init() {
	_httpHeaders = newHttpHeaders()
	_httpContentTypeValues = newHttpContentTypeValues()
}

type HttpContentWrapper struct {
	ApplicationJson      string
	ApplicationJsonPatch string
}

type httpContentTypeValues struct {
	ApplicationJson      string
	ApplicationJsonPatch string
}

func newHttpContentTypeValues() httpContentTypeValues {
	return httpContentTypeValues{
		ApplicationJson:      "application/json",
		ApplicationJsonPatch: "application/json-patch+json",
	}
}

func HttpContentTypeValues() httpContentTypeValues {
	return _httpContentTypeValues
}

type httpHeader string

func (h httpHeader) String() string {
	return string(h)
}

type httpHeaders struct {
	Authorization httpHeader
	ContentType   httpHeader
	StartTime     httpHeader
	RequestId     httpHeader
	PlatformType  httpHeader
	Platform      httpHeader
	Version       httpHeader
	UserAgent     httpHeader
	XForwardedFor httpHeader
	CacheControl  httpHeader
	Accept        httpHeader
}

func newHttpHeaders() httpHeaders {
	return httpHeaders{
		Authorization: "Authorization",
		ContentType:   "Content-Type",
		PlatformType:  "Platform-Type",
		Platform:      "Platform",
		Version:       "Version",
		UserAgent:     "User-Agent",
		StartTime:     "Start-Time",
		RequestId:     "Request-Id",
		XForwardedFor: "X-Forwarded-For",
		CacheControl:  "Cache-Control",
		Accept:        "Accept",
	}
}

func HttpHeaders() httpHeaders {
	return _httpHeaders
}