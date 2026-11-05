package requestwrapper

import "net/http"

// Response holds the HTTP status and optional error returned by a HandlerFunc.
type Response struct {
	StatusCode int
	Err        error
}

// Continue returns a Response.
//
// Status code: 100
//
// The server has received the headers and allows the client to send the request body.
func Continue() Response {
	return Response{StatusCode: http.StatusContinue}
}

// SwitchingProtocols returns a Response.
//
// Status code: 101
//
// The server is switching to the protocol requested by the client.
func SwitchingProtocols() Response {
	return Response{StatusCode: http.StatusSwitchingProtocols}
}

// Processing returns a Response.
//
// Status code: 102
//
// The server accepted the request but has not finished processing it.
func Processing() Response {
	return Response{StatusCode: http.StatusProcessing}
}

// EarlyHints returns a Response.
//
// Status code: 103
//
// The server sends preliminary headers before its final response.
func EarlyHints() Response {
	return Response{StatusCode: http.StatusEarlyHints}
}

// OK returns a Response.
//
// Status code: 200
//
// The request completed successfully.
func OK() Response {
	return Response{StatusCode: http.StatusOK}
}

// Created returns a Response.
//
// Status code: 201
//
// The request created a new resource.
func Created() Response {
	return Response{StatusCode: http.StatusCreated}
}

// Accepted returns a Response.
//
// Status code: 202
//
// The request was accepted, but processing has not finished.
func Accepted() Response {
	return Response{StatusCode: http.StatusAccepted}
}

// NonAuthoritativeInfo returns a Response.
//
// Status code: 203
//
// A proxy changed the representation supplied by the origin server.
func NonAuthoritativeInfo() Response {
	return Response{StatusCode: http.StatusNonAuthoritativeInfo}
}

// NoContent returns a Response.
//
// Status code: 204
//
// The request succeeded and there is no response content.
func NoContent() Response {
	return Response{StatusCode: http.StatusNoContent}
}

// ResetContent returns a Response.
//
// Status code: 205
//
// The request succeeded and the client should reset its current view.
func ResetContent() Response {
	return Response{StatusCode: http.StatusResetContent}
}

// PartialContent returns a Response.
//
// Status code: 206
//
// The response contains only the requested part of a resource.
func PartialContent() Response {
	return Response{StatusCode: http.StatusPartialContent}
}

// MultiStatus returns a Response.
//
// Status code: 207
//
// A WebDAV response reports separate results for multiple resources.
func MultiStatus() Response {
	return Response{StatusCode: http.StatusMultiStatus}
}

// AlreadyReported returns a Response.
//
// Status code: 208
//
// A WebDAV response omits a member already listed in an earlier part.
func AlreadyReported() Response {
	return Response{StatusCode: http.StatusAlreadyReported}
}

// IMUsed returns a Response.
//
// Status code: 226
//
// The response applies one or more instance manipulations to the resource.
func IMUsed() Response {
	return Response{StatusCode: http.StatusIMUsed}
}

// MultipleChoices returns a Response.
//
// Status code: 300
//
// The resource has multiple representations from which the client can choose.
func MultipleChoices() Response {
	return Response{StatusCode: http.StatusMultipleChoices}
}

// MovedPermanently returns a Response.
//
// Status code: 301
//
// The resource has a permanent new URI.
func MovedPermanently() Response {
	return Response{StatusCode: http.StatusMovedPermanently}
}

// Found returns a Response.
//
// Status code: 302
//
// The resource is temporarily available at another URI.
func Found() Response {
	return Response{StatusCode: http.StatusFound}
}

// SeeOther returns a Response.
//
// Status code: 303
//
// The client should retrieve another URI to view the result.
func SeeOther() Response {
	return Response{StatusCode: http.StatusSeeOther}
}

// NotModified returns a Response.
//
// Status code: 304
//
// A cached representation remains valid, so no new content is sent.
func NotModified() Response {
	return Response{StatusCode: http.StatusNotModified}
}

// UseProxy returns a Response.
//
// Status code: 305
//
// The resource is to be accessed through a specified proxy.
func UseProxy() Response {
	return Response{StatusCode: http.StatusUseProxy}
}

// TemporaryRedirect returns a Response.
//
// Status code: 307
//
// The resource is temporarily at another URI; keep the request method.
func TemporaryRedirect() Response {
	return Response{StatusCode: http.StatusTemporaryRedirect}
}

// PermanentRedirect returns a Response.
//
// Status code: 308
//
// The resource has a permanent new URI; keep the request method.
func PermanentRedirect() Response {
	return Response{StatusCode: http.StatusPermanentRedirect}
}

// BadRequest returns a Response.
//
// Status code: 400
//
// The server cannot process the request because it is malformed.
func BadRequest(err error) Response {
	return Response{StatusCode: http.StatusBadRequest, Err: err}
}

// Unauthorized returns a Response.
//
// Status code: 401
//
// The request needs valid authentication credentials.
func Unauthorized(err error) Response {
	return Response{StatusCode: http.StatusUnauthorized, Err: err}
}

// PaymentRequired returns a Response.
//
// Status code: 402
//
// This code is reserved for future use.
func PaymentRequired(err error) Response {
	return Response{StatusCode: http.StatusPaymentRequired, Err: err}
}

// Forbidden returns a Response.
//
// Status code: 403
//
// The server understood the request but refuses to fulfill it.
func Forbidden(err error) Response {
	return Response{StatusCode: http.StatusForbidden, Err: err}
}

// NotFound returns a Response.
//
// Status code: 404
//
// The server cannot find the requested resource.
func NotFound(err error) Response {
	return Response{StatusCode: http.StatusNotFound, Err: err}
}

// MethodNotAllowed returns a Response.
//
// Status code: 405
//
// The resource does not support the request method.
func MethodNotAllowed(err error) Response {
	return Response{StatusCode: http.StatusMethodNotAllowed, Err: err}
}

// NotAcceptable returns a Response.
//
// Status code: 406
//
// No available representation satisfies the request's accept headers.
func NotAcceptable(err error) Response {
	return Response{StatusCode: http.StatusNotAcceptable, Err: err}
}

// ProxyAuthRequired returns a Response.
//
// Status code: 407
//
// The client must authenticate with a proxy.
func ProxyAuthRequired(err error) Response {
	return Response{StatusCode: http.StatusProxyAuthRequired, Err: err}
}

// RequestTimeout returns a Response.
//
// Status code: 408
//
// The server timed out while waiting for the complete request.
func RequestTimeout(err error) Response {
	return Response{StatusCode: http.StatusRequestTimeout, Err: err}
}

// Conflict returns a Response.
//
// Status code: 409
//
// The request conflicts with the resource's current state.
func Conflict(err error) Response {
	return Response{StatusCode: http.StatusConflict, Err: err}
}

// Gone returns a Response.
//
// Status code: 410
//
// The requested resource is permanently unavailable.
func Gone(err error) Response {
	return Response{StatusCode: http.StatusGone, Err: err}
}

// LengthRequired returns a Response.
//
// Status code: 411
//
// The server requires a Content-Length header for this request.
func LengthRequired(err error) Response {
	return Response{StatusCode: http.StatusLengthRequired, Err: err}
}

// PreconditionFailed returns a Response.
//
// Status code: 412
//
// A request precondition evaluated to false.
func PreconditionFailed(err error) Response {
	return Response{StatusCode: http.StatusPreconditionFailed, Err: err}
}

// RequestEntityTooLarge returns a Response.
//
// Status code: 413
//
// The request content exceeds a server limit.
func RequestEntityTooLarge(err error) Response {
	return Response{StatusCode: http.StatusRequestEntityTooLarge, Err: err}
}

// RequestURITooLong returns a Response.
//
// Status code: 414
//
// The request target URI is too long for the server.
func RequestURITooLong(err error) Response {
	return Response{StatusCode: http.StatusRequestURITooLong, Err: err}
}

// UnsupportedMediaType returns a Response.
//
// Status code: 415
//
// The server does not support the request content format.
func UnsupportedMediaType(err error) Response {
	return Response{StatusCode: http.StatusUnsupportedMediaType, Err: err}
}

// RequestedRangeNotSatisfiable returns a Response.
//
// Status code: 416
//
// The requested range cannot be fulfilled for this resource.
func RequestedRangeNotSatisfiable(err error) Response {
	return Response{StatusCode: http.StatusRequestedRangeNotSatisfiable, Err: err}
}

// ExpectationFailed returns a Response.
//
// Status code: 417
//
// The server cannot meet a requirement in the Expect header.
func ExpectationFailed(err error) Response {
	return Response{StatusCode: http.StatusExpectationFailed, Err: err}
}

// Teapot returns a Response.
//
// Status code: 418
//
// This reserved code comes from the HTTP teapot joke.
func Teapot(err error) Response {
	return Response{StatusCode: http.StatusTeapot, Err: err}
}

// MisdirectedRequest returns a Response.
//
// Status code: 421
//
// The server cannot produce a response for the target URI on this connection.
func MisdirectedRequest(err error) Response {
	return Response{StatusCode: http.StatusMisdirectedRequest, Err: err}
}

// UnprocessableEntity returns a Response.
//
// Status code: 422
//
// The content is validly formed but its instructions cannot be processed.
func UnprocessableEntity(err error) Response {
	return Response{StatusCode: http.StatusUnprocessableEntity, Err: err}
}

// Locked returns a Response.
//
// Status code: 423
//
// A WebDAV resource is locked.
func Locked(err error) Response {
	return Response{StatusCode: http.StatusLocked, Err: err}
}

// FailedDependency returns a Response.
//
// Status code: 424
//
// The request failed because another required action failed.
func FailedDependency(err error) Response {
	return Response{StatusCode: http.StatusFailedDependency, Err: err}
}

// TooEarly returns a Response.
//
// Status code: 425
//
// The server will not process an early-data request that might be replayed.
func TooEarly(err error) Response {
	return Response{StatusCode: http.StatusTooEarly, Err: err}
}

// UpgradeRequired returns a Response.
//
// Status code: 426
//
// The client must switch to another protocol to complete the request.
func UpgradeRequired(err error) Response {
	return Response{StatusCode: http.StatusUpgradeRequired, Err: err}
}

// PreconditionRequired returns a Response.
//
// Status code: 428
//
// The server requires the request to include a precondition.
func PreconditionRequired(err error) Response {
	return Response{StatusCode: http.StatusPreconditionRequired, Err: err}
}

// TooManyRequests returns a Response.
//
// Status code: 429
//
// The client has sent too many requests in a given period.
func TooManyRequests(err error) Response {
	return Response{StatusCode: http.StatusTooManyRequests, Err: err}
}

// RequestHeaderFieldsTooLarge returns a Response.
//
// Status code: 431
//
// The request headers are too large for the server to process.
func RequestHeaderFieldsTooLarge(err error) Response {
	return Response{StatusCode: http.StatusRequestHeaderFieldsTooLarge, Err: err}
}

// UnavailableForLegalReasons returns a Response.
//
// Status code: 451
//
// Access to the resource is denied because of a legal demand.
func UnavailableForLegalReasons(err error) Response {
	return Response{StatusCode: http.StatusUnavailableForLegalReasons, Err: err}
}

// InternalServerError returns a Response.
//
// Status code: 500
//
// The server encountered an unexpected condition.
func InternalServerError(err error) Response {
	return Response{StatusCode: http.StatusInternalServerError, Err: err}
}

// NotImplemented returns a Response.
//
// Status code: 501
//
// The server does not support the functionality required by the request.
func NotImplemented(err error) Response {
	return Response{StatusCode: http.StatusNotImplemented, Err: err}
}

// BadGateway returns a Response.
//
// Status code: 502
//
// A gateway received an invalid response from an upstream server.
func BadGateway(err error) Response {
	return Response{StatusCode: http.StatusBadGateway, Err: err}
}

// ServiceUnavailable returns a Response.
//
// Status code: 503
//
// The server is temporarily unable to handle the request.
func ServiceUnavailable(err error) Response {
	return Response{StatusCode: http.StatusServiceUnavailable, Err: err}
}

// GatewayTimeout returns a Response.
//
// Status code: 504
//
// A gateway timed out waiting for an upstream server.
func GatewayTimeout(err error) Response {
	return Response{StatusCode: http.StatusGatewayTimeout, Err: err}
}

// HTTPVersionNotSupported returns a Response.
//
// Status code: 505
//
// The server does not support the request's HTTP version.
func HTTPVersionNotSupported(err error) Response {
	return Response{StatusCode: http.StatusHTTPVersionNotSupported, Err: err}
}

// VariantAlsoNegotiates returns a Response.
//
// Status code: 506
//
// A selected resource is itself configured for transparent negotiation.
func VariantAlsoNegotiates(err error) Response {
	return Response{StatusCode: http.StatusVariantAlsoNegotiates, Err: err}
}

// InsufficientStorage returns a Response.
//
// Status code: 507
//
// The server cannot store the representation needed to complete the request.
func InsufficientStorage(err error) Response {
	return Response{StatusCode: http.StatusInsufficientStorage, Err: err}
}

// LoopDetected returns a Response.
//
// Status code: 508
//
// The server detected a loop while processing the request.
func LoopDetected(err error) Response {
	return Response{StatusCode: http.StatusLoopDetected, Err: err}
}

// NotExtended returns a Response.
//
// Status code: 510
//
// The request lacks an extension required by the server.
func NotExtended(err error) Response {
	return Response{StatusCode: http.StatusNotExtended, Err: err}
}

// NetworkAuthenticationRequired returns a Response.
//
// Status code: 511
//
// The client must authenticate to gain network access.
func NetworkAuthenticationRequired(err error) Response {
	return Response{StatusCode: http.StatusNetworkAuthenticationRequired, Err: err}
}
