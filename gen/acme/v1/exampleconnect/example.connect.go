// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: acme/v1/example.proto

package exampleconnect

import (
	example "acme.com/example"
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_7_0

const (
	// LibraryServiceName is the fully-qualified name of the LibraryService service.
	LibraryServiceName = "acme.v1.example.LibraryService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// LibraryServiceGetBookProcedure is the fully-qualified name of the LibraryService's GetBook RPC.
	LibraryServiceGetBookProcedure = "/acme.v1.example.LibraryService/GetBook"
)

// LibraryServiceClient is a client for the acme.v1.example.LibraryService service.
type LibraryServiceClient interface {
	// Gets a book.
	GetBook(context.Context, *connect.Request[example.GetBookRequest]) (*connect.Response[example.Book], error)
}

// NewLibraryServiceClient constructs a client for the acme.v1.example.LibraryService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewLibraryServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) LibraryServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &libraryServiceClient{
		getBook: connect.NewClient[example.GetBookRequest, example.Book](
			httpClient,
			baseURL+LibraryServiceGetBookProcedure,
			connect.WithIdempotency(connect.IdempotencyNoSideEffects),
			connect.WithClientOptions(opts...),
		),
	}
}

// libraryServiceClient implements LibraryServiceClient.
type libraryServiceClient struct {
	getBook *connect.Client[example.GetBookRequest, example.Book]
}

// GetBook calls acme.v1.example.LibraryService.GetBook.
func (c *libraryServiceClient) GetBook(ctx context.Context, req *connect.Request[example.GetBookRequest]) (*connect.Response[example.Book], error) {
	return c.getBook.CallUnary(ctx, req)
}

// LibraryServiceHandler is an implementation of the acme.v1.example.LibraryService service.
type LibraryServiceHandler interface {
	// Gets a book.
	GetBook(context.Context, *connect.Request[example.GetBookRequest]) (*connect.Response[example.Book], error)
}

// NewLibraryServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewLibraryServiceHandler(svc LibraryServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	libraryServiceGetBookHandler := connect.NewUnaryHandler(
		LibraryServiceGetBookProcedure,
		svc.GetBook,
		connect.WithIdempotency(connect.IdempotencyNoSideEffects),
		connect.WithHandlerOptions(opts...),
	)
	return "/acme.v1.example.LibraryService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case LibraryServiceGetBookProcedure:
			libraryServiceGetBookHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedLibraryServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedLibraryServiceHandler struct{}

func (UnimplementedLibraryServiceHandler) GetBook(context.Context, *connect.Request[example.GetBookRequest]) (*connect.Response[example.Book], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("acme.v1.example.LibraryService.GetBook is not implemented"))
}
