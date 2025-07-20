// Package gostacode provides efficient conversion functions between HTTP status codes
// and gRPC status codes. This package is useful for services that need to translate
// between REST API and gRPC API responses.
package gostacode

import (
	"net/http"

	"google.golang.org/grpc/codes"
)

// httpToGRPCMap maps HTTP status codes to their corresponding gRPC codes.
// This mapping follows the standard conversion rules between HTTP and gRPC protocols.
var httpToGRPCMap = map[int]codes.Code{
	// 2xx Success codes
	http.StatusOK:      codes.OK,
	http.StatusCreated: codes.OK,

	// 4xx Client error codes
	http.StatusBadRequest:      codes.InvalidArgument,
	http.StatusUnauthorized:    codes.Unauthenticated,
	http.StatusForbidden:       codes.PermissionDenied,
	http.StatusNotFound:        codes.NotFound,
	http.StatusConflict:        codes.AlreadyExists,
	http.StatusTooManyRequests: codes.ResourceExhausted,

	// 5xx Server error codes
	http.StatusInternalServerError: codes.Internal,
	http.StatusNotImplemented:      codes.Unimplemented,
	http.StatusBadGateway:          codes.Unavailable,
	http.StatusServiceUnavailable:  codes.Unavailable,
	http.StatusGatewayTimeout:      codes.DeadlineExceeded,
}

// grpcToHTTPMap maps gRPC codes to their corresponding HTTP status codes.
// This mapping follows the standard conversion rules between gRPC and HTTP protocols.
var grpcToHTTPMap = map[codes.Code]int{
	codes.OK:                 http.StatusOK,
	codes.Unknown:            http.StatusInternalServerError,
	codes.InvalidArgument:    http.StatusBadRequest,
	codes.DeadlineExceeded:   http.StatusGatewayTimeout,
	codes.NotFound:           http.StatusNotFound,
	codes.AlreadyExists:      http.StatusConflict,
	codes.PermissionDenied:   http.StatusForbidden,
	codes.Unauthenticated:    http.StatusUnauthorized,
	codes.ResourceExhausted:  http.StatusTooManyRequests,
	codes.FailedPrecondition: http.StatusBadRequest,
	codes.Aborted:            http.StatusConflict,
	codes.OutOfRange:         http.StatusBadRequest,
	codes.Unimplemented:      http.StatusNotImplemented,
	codes.Internal:           http.StatusInternalServerError,
	codes.Unavailable:        http.StatusServiceUnavailable,
	codes.DataLoss:           http.StatusInternalServerError,
	codes.Canceled:           http.StatusInternalServerError, // Added missing mapping
}

// GRPCCodeFromHTTPStatusCode converts an HTTP status code to its corresponding gRPC code.
// If the HTTP status code is not found in the mapping, it returns codes.Unknown.
//
// Parameters:
//   - httpStatusCode: The HTTP status code to convert (e.g., 200, 404, 500)
//
// Returns:
//   - codes.Code: The corresponding gRPC code or codes.Unknown if not mapped
func GRPCCodeFromHTTPStatusCode(httpStatusCode int) codes.Code {
	if grpcCode, ok := httpToGRPCMap[httpStatusCode]; ok {
		return grpcCode
	}
	return codes.Unknown
}

// HTTPStatusCodeFromGRPCCode converts a gRPC code to its corresponding HTTP status code.
// If the gRPC code is not found in the mapping, it returns http.StatusInternalServerError.
//
// Parameters:
//   - grpcCode: The gRPC code to convert (e.g., codes.OK, codes.NotFound)
//
// Returns:
//   - int: The corresponding HTTP status code or 500 if not mapped
func HTTPStatusCodeFromGRPCCode(grpcCode codes.Code) int {
	if httpStatusCode, ok := grpcToHTTPMap[grpcCode]; ok {
		return httpStatusCode
	}
	return http.StatusInternalServerError
}
