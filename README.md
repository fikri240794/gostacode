# GoStaCode

Go library for bidirectional conversion between HTTP status codes and gRPC status codes.

## 🚀 Features

- **Bidirectional conversion** between HTTP and gRPC status codes
- **Comprehensive mapping** covering all standard HTTP and gRPC codes
- **Graceful fallbacks** for unmapped codes

## 📦 Installation

```bash
go get github.com/fikri240794/gostacode
```

## 🔧 Quick Start

```go
package main

import (
    "fmt"
    "net/http"
    
    "github.com/fikri240794/gostacode"
    "google.golang.org/grpc/codes"
)

func main() {
    // Convert HTTP status code to gRPC code
    grpcCode := gostacode.GRPCCodeFromHTTPStatusCode(http.StatusNotFound)
    fmt.Printf("HTTP 404 → gRPC %s\n", grpcCode.String()) // gRPC NotFound
    
    // Convert gRPC code to HTTP status code
    httpCode := gostacode.HTTPStatusCodeFromGRPCCode(codes.InvalidArgument)
    fmt.Printf("gRPC InvalidArgument → HTTP %d\n", httpCode) // HTTP 400
}
```

## 📋 API Reference

### Functions

#### `GRPCCodeFromHTTPStatusCode(httpStatusCode int) codes.Code`

Converts an HTTP status code to its corresponding gRPC code.

**Parameters:**
- `httpStatusCode`: HTTP status code (e.g., 200, 404, 500)

**Returns:**
- `codes.Code`: Corresponding gRPC code or `codes.Unknown` if not mapped

**Example:**
```go
grpcCode := gostacode.GRPCCodeFromHTTPStatusCode(400)
// Returns: codes.InvalidArgument
```

#### `HTTPStatusCodeFromGRPCCode(grpcCode codes.Code) int`

Converts a gRPC code to its corresponding HTTP status code.

**Parameters:**
- `grpcCode`: gRPC status code (e.g., codes.OK, codes.NotFound)

**Returns:**
- `int`: Corresponding HTTP status code or `500` if not mapped

**Example:**
```go
httpCode := gostacode.HTTPStatusCodeFromGRPCCode(codes.NotFound)
// Returns: 404
```

## 🗺️ Conversion Mapping

### HTTP to gRPC Mapping

| HTTP Status Code | HTTP Status Text | gRPC Code |
|------------------|------------------|-----------|
| 200 | OK | `codes.OK` |
| 201 | Created | `codes.OK` |
| 400 | Bad Request | `codes.InvalidArgument` |
| 401 | Unauthorized | `codes.Unauthenticated` |
| 403 | Forbidden | `codes.PermissionDenied` |
| 404 | Not Found | `codes.NotFound` |
| 409 | Conflict | `codes.AlreadyExists` |
| 429 | Too Many Requests | `codes.ResourceExhausted` |
| 500 | Internal Server Error | `codes.Internal` |
| 501 | Not Implemented | `codes.Unimplemented` |
| 502 | Bad Gateway | `codes.Unavailable` |
| 503 | Service Unavailable | `codes.Unavailable` |
| 504 | Gateway Timeout | `codes.DeadlineExceeded` |

### gRPC to HTTP Mapping

| gRPC Code | HTTP Status Code | HTTP Status Text |
|-----------|------------------|------------------|
| `codes.OK` | 200 | OK |
| `codes.Unknown` | 500 | Internal Server Error |
| `codes.InvalidArgument` | 400 | Bad Request |
| `codes.DeadlineExceeded` | 504 | Gateway Timeout |
| `codes.NotFound` | 404 | Not Found |
| `codes.AlreadyExists` | 409 | Conflict |
| `codes.PermissionDenied` | 403 | Forbidden |
| `codes.Unauthenticated` | 401 | Unauthorized |
| `codes.ResourceExhausted` | 429 | Too Many Requests |
| `codes.FailedPrecondition` | 400 | Bad Request |
| `codes.Aborted` | 409 | Conflict |
| `codes.OutOfRange` | 400 | Bad Request |
| `codes.Unimplemented` | 501 | Not Implemented |
| `codes.Internal` | 500 | Internal Server Error |
| `codes.Unavailable` | 503 | Service Unavailable |
| `codes.DataLoss` | 500 | Internal Server Error |
| `codes.Canceled` | 500 | Internal Server Error |

## 🚀 Examples

Check out the [examples](examples/) directory for comprehensive usage demonstrations:

```bash
cd examples
go run main.go
```

The example covers:
- Basic conversions
- Edge case handling
- API gateway scenarios
- Batch processing

---