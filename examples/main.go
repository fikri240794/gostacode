// Package main demonstrates the usage of gostacode library for converting
// between HTTP status codes and gRPC codes in various scenarios.
package main

import (
	"fmt"
	"net/http"

	"github.com/fikri240794/gostacode"
	"google.golang.org/grpc/codes"
)

func main() {
	fmt.Println("=== GostaCode Demo: HTTP to gRPC Code Conversion ===")
	fmt.Println()

	// Simple conversion examples
	fmt.Println("1. Basic HTTP to gRPC Conversions:")
	demoHTTPToGRPC()

	fmt.Println("\n2. Basic gRPC to HTTP Conversions:")
	demoGRPCToHTTP()

	fmt.Println("\n3. Edge Cases and Error Handling:")
	demoEdgeCases()

	fmt.Println("\n4. Real-world API Gateway Scenario:")
	demoAPIGateway()

	fmt.Println("\n5. Batch Processing Example:")
	demoBatchProcessing()
}

// demoHTTPToGRPC demonstrates converting common HTTP status codes to gRPC codes
func demoHTTPToGRPC() {
	httpCodes := []int{
		http.StatusOK,
		http.StatusCreated,
		http.StatusBadRequest,
		http.StatusUnauthorized,
		http.StatusForbidden,
		http.StatusNotFound,
		http.StatusConflict,
		http.StatusTooManyRequests,
		http.StatusInternalServerError,
		http.StatusNotImplemented,
		http.StatusServiceUnavailable,
		http.StatusGatewayTimeout,
	}

	for _, httpCode := range httpCodes {
		grpcCode := gostacode.GRPCCodeFromHTTPStatusCode(httpCode)
		fmt.Printf("  HTTP %d (%s) → gRPC %s (%d)\n",
			httpCode, http.StatusText(httpCode), grpcCode.String(), grpcCode)
	}
}

// demoGRPCToHTTP demonstrates converting gRPC codes to HTTP status codes
func demoGRPCToHTTP() {
	grpcCodes := []codes.Code{
		codes.OK,
		codes.Unknown,
		codes.InvalidArgument,
		codes.DeadlineExceeded,
		codes.NotFound,
		codes.AlreadyExists,
		codes.PermissionDenied,
		codes.Unauthenticated,
		codes.ResourceExhausted,
		codes.FailedPrecondition,
		codes.Aborted,
		codes.OutOfRange,
		codes.Unimplemented,
		codes.Internal,
		codes.Unavailable,
		codes.DataLoss,
		codes.Canceled,
	}

	for _, grpcCode := range grpcCodes {
		httpCode := gostacode.HTTPStatusCodeFromGRPCCode(grpcCode)
		fmt.Printf("  gRPC %s (%d) → HTTP %d (%s)\n",
			grpcCode.String(), grpcCode, httpCode, http.StatusText(httpCode))
	}
}

// demoEdgeCases demonstrates handling of unmapped codes and edge cases
func demoEdgeCases() {
	fmt.Println("  Testing unmapped HTTP codes:")

	// Test unmapped HTTP status codes
	unmappedHTTPCodes := []int{0, -1, 418, 999}
	for _, httpCode := range unmappedHTTPCodes {
		grpcCode := gostacode.GRPCCodeFromHTTPStatusCode(httpCode)
		fmt.Printf("    HTTP %d → gRPC %s (%d) [fallback]\n",
			httpCode, grpcCode.String(), grpcCode)
	}

	fmt.Println("  Testing unmapped gRPC codes:")

	// Test unmapped gRPC codes
	unmappedGRPCCodes := []codes.Code{codes.Code(999)}
	for _, grpcCode := range unmappedGRPCCodes {
		httpCode := gostacode.HTTPStatusCodeFromGRPCCode(grpcCode)
		fmt.Printf("    gRPC %d → HTTP %d (%s) [fallback]\n",
			grpcCode, httpCode, http.StatusText(httpCode))
	}
}

// demoAPIGateway simulates an API gateway translating between protocols
func demoAPIGateway() {
	fmt.Println("  Simulating API Gateway protocol translation:")

	// Simulate incoming HTTP requests with various status codes
	incomingHTTPResponses := []struct {
		service    string
		statusCode int
	}{
		{"user-service", http.StatusOK},
		{"auth-service", http.StatusUnauthorized},
		{"inventory-service", http.StatusNotFound},
		{"payment-service", http.StatusTooManyRequests},
		{"notification-service", http.StatusInternalServerError},
	}

	fmt.Println("    HTTP → gRPC translation:")
	for _, response := range incomingHTTPResponses {
		grpcCode := gostacode.GRPCCodeFromHTTPStatusCode(response.statusCode)
		fmt.Printf("      %s: HTTP %d → gRPC %s\n",
			response.service, response.statusCode, grpcCode.String())
	}

	// Simulate gRPC responses that need to be translated to HTTP
	incomingGRPCResponses := []struct {
		service string
		code    codes.Code
	}{
		{"backend-service", codes.OK},
		{"validation-service", codes.InvalidArgument},
		{"db-service", codes.DeadlineExceeded},
		{"cache-service", codes.Unavailable},
		{"audit-service", codes.PermissionDenied},
	}

	fmt.Println("    gRPC → HTTP translation:")
	for _, response := range incomingGRPCResponses {
		httpCode := gostacode.HTTPStatusCodeFromGRPCCode(response.code)
		fmt.Printf("      %s: gRPC %s → HTTP %d (%s)\n",
			response.service, response.code.String(), httpCode, http.StatusText(httpCode))
	}
}

// demoBatchProcessing demonstrates processing multiple codes efficiently
func demoBatchProcessing() {
	fmt.Println("  Processing batch of HTTP status codes:")

	// Simulate batch processing scenario
	httpStatusCodes := []int{200, 201, 400, 401, 403, 404, 409, 429, 500, 502, 503, 504}
	results := make(map[int]codes.Code, len(httpStatusCodes))

	// Process all codes
	for _, httpCode := range httpStatusCodes {
		results[httpCode] = gostacode.GRPCCodeFromHTTPStatusCode(httpCode)
	}

	// Display results
	fmt.Printf("    Processed %d HTTP codes:\n", len(results))
	for httpCode, grpcCode := range results {
		fmt.Printf("      %d → %s\n", httpCode, grpcCode.String())
	}

	fmt.Println("  Processing batch of gRPC codes:")

	// Simulate reverse batch processing
	grpcCodes := []codes.Code{
		codes.OK, codes.InvalidArgument, codes.NotFound,
		codes.Internal, codes.Unavailable, codes.DeadlineExceeded,
	}
	reverseResults := make(map[codes.Code]int, len(grpcCodes))

	// Process all codes
	for _, grpcCode := range grpcCodes {
		reverseResults[grpcCode] = gostacode.HTTPStatusCodeFromGRPCCode(grpcCode)
	}

	// Display results
	fmt.Printf("    Processed %d gRPC codes:\n", len(reverseResults))
	for grpcCode, httpCode := range reverseResults {
		fmt.Printf("      %s → %d\n", grpcCode.String(), httpCode)
	}
}
