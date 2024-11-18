package main

import (
	"math/rand"
	"net/http"
	"time"
)

// Comprehensive list of all HTTP status codes
var statusCodes = []int{
	// 1xx Informational
	http.StatusContinue,           // 100
	http.StatusSwitchingProtocols, // 101
	http.StatusProcessing,         // 102 (WebDAV)

	// 2xx Success
	http.StatusOK,                   // 200
	http.StatusCreated,              // 201
	http.StatusAccepted,             // 202
	http.StatusNonAuthoritativeInfo, // 203
	http.StatusNoContent,            // 204
	http.StatusResetContent,         // 205
	http.StatusPartialContent,       // 206
	http.StatusMultiStatus,          // 207 (WebDAV)
	http.StatusAlreadyReported,      // 208 (WebDAV)
	http.StatusIMUsed,               // 226

	// 3xx Redirection
	http.StatusMultipleChoices,   // 300
	http.StatusMovedPermanently,  // 301
	http.StatusFound,             // 302
	http.StatusSeeOther,          // 303
	http.StatusNotModified,       // 304
	http.StatusUseProxy,          // 305
	http.StatusTemporaryRedirect, // 307
	http.StatusPermanentRedirect, // 308

	// 4xx Client Errors
	http.StatusBadRequest,                   // 400
	http.StatusUnauthorized,                 // 401
	http.StatusPaymentRequired,              // 402
	http.StatusForbidden,                    // 403
	http.StatusNotFound,                     // 404
	http.StatusMethodNotAllowed,             // 405
	http.StatusNotAcceptable,                // 406
	http.StatusProxyAuthRequired,            // 407
	http.StatusRequestTimeout,               // 408
	http.StatusConflict,                     // 409
	http.StatusGone,                         // 410
	http.StatusLengthRequired,               // 411
	http.StatusPreconditionFailed,           // 412
	http.StatusRequestEntityTooLarge,        // 413
	http.StatusRequestURITooLong,            // 414
	http.StatusUnsupportedMediaType,         // 415
	http.StatusRequestedRangeNotSatisfiable, // 416
	http.StatusExpectationFailed,            // 417
	http.StatusTeapot,                       // 418 (I'm a teapot, Easter egg)
	http.StatusMisdirectedRequest,           // 421
	http.StatusUnprocessableEntity,          // 422 (WebDAV)
	http.StatusLocked,                       // 423 (WebDAV)
	http.StatusFailedDependency,             // 424 (WebDAV)
	http.StatusTooEarly,                     // 425
	http.StatusUpgradeRequired,              // 426
	http.StatusPreconditionRequired,         // 428
	http.StatusTooManyRequests,              // 429
	http.StatusRequestHeaderFieldsTooLarge,  // 431
	http.StatusUnavailableForLegalReasons,   // 451

	// 5xx Server Errors
	http.StatusInternalServerError,           // 500
	http.StatusNotImplemented,                // 501
	http.StatusBadGateway,                    // 502
	http.StatusServiceUnavailable,            // 503
	http.StatusGatewayTimeout,                // 504
	http.StatusHTTPVersionNotSupported,       // 505
	http.StatusVariantAlsoNegotiates,         // 506
	http.StatusInsufficientStorage,           // 507 (WebDAV)
	http.StatusLoopDetected,                  // 508 (WebDAV)
	http.StatusNotExtended,                   // 510
	http.StatusNetworkAuthenticationRequired, // 511
}

// Function to get a random status code
func getRandomStatusCode() int {
	rand.Seed(time.Now().UnixNano())
	return statusCodes[rand.Intn(len(statusCodes))]
}
