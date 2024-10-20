package gostacode

import (
	"net/http"

	"google.golang.org/grpc/codes"
)

var (
	httpGRPCCodeMap map[int]codes.Code = map[int]codes.Code{
		http.StatusOK:      codes.OK,
		http.StatusCreated: codes.OK,

		http.StatusBadRequest:      codes.InvalidArgument,
		http.StatusUnauthorized:    codes.Unauthenticated,
		http.StatusForbidden:       codes.PermissionDenied,
		http.StatusNotFound:        codes.NotFound,
		http.StatusConflict:        codes.AlreadyExists,
		http.StatusTooManyRequests: codes.ResourceExhausted,

		http.StatusInternalServerError: codes.Internal,
		http.StatusNotImplemented:      codes.Unimplemented,
		http.StatusBadGateway:          codes.Unavailable,
		http.StatusServiceUnavailable:  codes.Unavailable,
		http.StatusGatewayTimeout:      codes.DeadlineExceeded,
	}

	grpcHTTPCodeMap map[codes.Code]int = map[codes.Code]int{
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
	}
)

func GRPCCodeFromHTTPStatusCode(httpStatusCode int) codes.Code {
	var (
		grpcCode codes.Code
		ok       bool
	)

	grpcCode, ok = httpGRPCCodeMap[httpStatusCode]
	if !ok {
		return codes.Unknown
	}

	return grpcCode
}

func HTTPStatusCodeFromGRPCCode(grpcCode codes.Code) int {
	var (
		httpStatusCode int
		ok             bool
	)

	httpStatusCode, ok = grpcHTTPCodeMap[grpcCode]
	if !ok {
		return http.StatusInternalServerError
	}

	return httpStatusCode
}
