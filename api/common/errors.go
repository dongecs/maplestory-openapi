package common

import "fmt"

// MapleStoryErrorBody mirrors the error envelope returned by the Nexon OpenAPI.
type MapleStoryErrorBody struct {
	Error MapleStoryError `json:"error"`
}

// MapleStoryError is the inner error payload from the API.
type MapleStoryError struct {
	Name    string `json:"name"`
	Message string `json:"message"`
}

// MapleStoryAPIErrorCode lists the documented error codes from the MapleStory OpenAPI guide.
type MapleStoryAPIErrorCode string

const (
	OpenAPI00001 MapleStoryAPIErrorCode = "OPENAPI00001"
	OpenAPI00002 MapleStoryAPIErrorCode = "OPENAPI00002"
	OpenAPI00003 MapleStoryAPIErrorCode = "OPENAPI00003"
	OpenAPI00004 MapleStoryAPIErrorCode = "OPENAPI00004"
	OpenAPI00005 MapleStoryAPIErrorCode = "OPENAPI00005"
	OpenAPI00006 MapleStoryAPIErrorCode = "OPENAPI00006"
	OpenAPI00007 MapleStoryAPIErrorCode = "OPENAPI00007"
	OpenAPI00009 MapleStoryAPIErrorCode = "OPENAPI00009"
	OpenAPI00010 MapleStoryAPIErrorCode = "OPENAPI00010"
	OpenAPI00011 MapleStoryAPIErrorCode = "OPENAPI00011"
)

// MapleStoryAPIError is returned when the API responds with an error envelope.
type MapleStoryAPIError struct {
	Code    MapleStoryAPIErrorCode
	Message string
	body    MapleStoryError
}

func (e *MapleStoryAPIError) Error() string {
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

// Body returns the raw error payload provided by the API.
func (e *MapleStoryAPIError) Body() MapleStoryError {
	return e.body
}

func newAPIError(body MapleStoryError) *MapleStoryAPIError {
	return &MapleStoryAPIError{
		Code:    errorCodeFromName(body.Name),
		Message: body.Message,
		body:    body,
	}
}

func errorCodeFromName(name string) MapleStoryAPIErrorCode {
	switch name {
	case string(OpenAPI00001):
		return OpenAPI00001
	case string(OpenAPI00002):
		return OpenAPI00002
	case string(OpenAPI00003):
		return OpenAPI00003
	case string(OpenAPI00004):
		return OpenAPI00004
	case string(OpenAPI00005):
		return OpenAPI00005
	case string(OpenAPI00006):
		return OpenAPI00006
	case string(OpenAPI00007):
		return OpenAPI00007
	case string(OpenAPI00009):
		return OpenAPI00009
	case string(OpenAPI00010):
		return OpenAPI00010
	case string(OpenAPI00011):
		return OpenAPI00011
	default:
		return MapleStoryAPIErrorCode(name)
	}
}
