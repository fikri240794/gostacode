package gostacode

import (
	"net/http"
	"testing"

	"google.golang.org/grpc/codes"
)

func TestGRPCCodeFromHTTPStatusCode(t *testing.T) {
	var testCases []struct {
		Name           string
		HTTPStatusCode int
		Expectation    codes.Code
	} = []struct {
		Name           string
		HTTPStatusCode int
		Expectation    codes.Code
	}{
		{
			Name:           http.StatusText(http.StatusHTTPVersionNotSupported),
			HTTPStatusCode: http.StatusHTTPVersionNotSupported,
			Expectation:    codes.Unknown,
		},
		{
			Name:           http.StatusText(http.StatusOK),
			HTTPStatusCode: http.StatusOK,
			Expectation:    codes.OK,
		},
		{
			Name:           http.StatusText(http.StatusCreated),
			HTTPStatusCode: http.StatusCreated,
			Expectation:    codes.OK,
		},
		{
			Name:           http.StatusText(http.StatusBadRequest),
			HTTPStatusCode: http.StatusBadRequest,
			Expectation:    codes.InvalidArgument,
		},
		{
			Name:           http.StatusText(http.StatusUnauthorized),
			HTTPStatusCode: http.StatusUnauthorized,
			Expectation:    codes.Unauthenticated,
		},
		{
			Name:           http.StatusText(http.StatusForbidden),
			HTTPStatusCode: http.StatusForbidden,
			Expectation:    codes.PermissionDenied,
		},
		{
			Name:           http.StatusText(http.StatusNotFound),
			HTTPStatusCode: http.StatusNotFound,
			Expectation:    codes.NotFound,
		},
		{
			Name:           http.StatusText(http.StatusConflict),
			HTTPStatusCode: http.StatusConflict,
			Expectation:    codes.AlreadyExists,
		},
		{
			Name:           http.StatusText(http.StatusTooManyRequests),
			HTTPStatusCode: http.StatusTooManyRequests,
			Expectation:    codes.ResourceExhausted,
		},
		{
			Name:           http.StatusText(http.StatusTooManyRequests),
			HTTPStatusCode: http.StatusTooManyRequests,
			Expectation:    codes.ResourceExhausted,
		},
		{
			Name:           http.StatusText(http.StatusInternalServerError),
			HTTPStatusCode: http.StatusInternalServerError,
			Expectation:    codes.Internal,
		},
		{
			Name:           http.StatusText(http.StatusNotImplemented),
			HTTPStatusCode: http.StatusNotImplemented,
			Expectation:    codes.Unimplemented,
		},
		{
			Name:           http.StatusText(http.StatusBadGateway),
			HTTPStatusCode: http.StatusBadGateway,
			Expectation:    codes.Unavailable,
		},
		{
			Name:           http.StatusText(http.StatusServiceUnavailable),
			HTTPStatusCode: http.StatusServiceUnavailable,
			Expectation:    codes.Unavailable,
		},
		{
			Name:           http.StatusText(http.StatusGatewayTimeout),
			HTTPStatusCode: http.StatusGatewayTimeout,
			Expectation:    codes.DeadlineExceeded,
		},
	}

	for i := range testCases {
		t.Run(testCases[i].Name, func(t *testing.T) {
			var actual codes.Code = GRPCCodeFromHTTPStatusCode(testCases[i].HTTPStatusCode)

			if testCases[i].Expectation != actual {
				t.Errorf("expectation is %d, got %d", testCases[i].Expectation, actual)
			}
		})
	}
}

func TestHTTPStatusCodeFromGRPCCode(t *testing.T) {
	var testCases []struct {
		Name        string
		GRPCCode    codes.Code
		Expectation int
	} = []struct {
		Name        string
		GRPCCode    codes.Code
		Expectation int
	}{
		{
			Name:        codes.Canceled.String(),
			GRPCCode:    codes.Canceled,
			Expectation: http.StatusInternalServerError,
		},
		{
			Name:        codes.OK.String(),
			GRPCCode:    codes.OK,
			Expectation: http.StatusOK,
		},
		{
			Name:        codes.Unknown.String(),
			GRPCCode:    codes.Unknown,
			Expectation: http.StatusInternalServerError,
		},
		{
			Name:        codes.InvalidArgument.String(),
			GRPCCode:    codes.InvalidArgument,
			Expectation: http.StatusBadRequest,
		},
		{
			Name:        codes.DeadlineExceeded.String(),
			GRPCCode:    codes.DeadlineExceeded,
			Expectation: http.StatusGatewayTimeout,
		},
		{
			Name:        codes.NotFound.String(),
			GRPCCode:    codes.NotFound,
			Expectation: http.StatusNotFound,
		},
		{
			Name:        codes.AlreadyExists.String(),
			GRPCCode:    codes.AlreadyExists,
			Expectation: http.StatusConflict,
		},
		{
			Name:        codes.PermissionDenied.String(),
			GRPCCode:    codes.PermissionDenied,
			Expectation: http.StatusForbidden,
		},
		{
			Name:        codes.Unauthenticated.String(),
			GRPCCode:    codes.Unauthenticated,
			Expectation: http.StatusUnauthorized,
		},
		{
			Name:        codes.ResourceExhausted.String(),
			GRPCCode:    codes.ResourceExhausted,
			Expectation: http.StatusTooManyRequests,
		},
		{
			Name:        codes.FailedPrecondition.String(),
			GRPCCode:    codes.FailedPrecondition,
			Expectation: http.StatusBadRequest,
		},
		{
			Name:        codes.Aborted.String(),
			GRPCCode:    codes.Aborted,
			Expectation: http.StatusConflict,
		},
		{
			Name:        codes.OutOfRange.String(),
			GRPCCode:    codes.OutOfRange,
			Expectation: http.StatusBadRequest,
		},
		{
			Name:        codes.Unimplemented.String(),
			GRPCCode:    codes.Unimplemented,
			Expectation: http.StatusNotImplemented,
		},
		{
			Name:        codes.Internal.String(),
			GRPCCode:    codes.Internal,
			Expectation: http.StatusInternalServerError,
		},
		{
			Name:        codes.Unavailable.String(),
			GRPCCode:    codes.Unavailable,
			Expectation: http.StatusServiceUnavailable,
		},
		{
			Name:        codes.DataLoss.String(),
			GRPCCode:    codes.DataLoss,
			Expectation: http.StatusInternalServerError,
		},
	}

	for i := range testCases {
		t.Run(testCases[i].Name, func(t *testing.T) {
			var actual int = HTTPStatusCodeFromGRPCCode(testCases[i].GRPCCode)

			if testCases[i].Expectation != actual {
				t.Errorf("expectation is %d, got %d", testCases[i].Expectation, actual)
			}
		})
	}
}
