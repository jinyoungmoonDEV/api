// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: spaceone/api/identity/v2/endpoint.proto

/*
Package v2 is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package v2

import (
	"context"
	"io"
	"net/http"

	v2_0 "github.com/cloudforet-io/api/dist/go/spaceone/api/identity/v2"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/v2/utilities"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

// Suppress "imported and not used" errors
var _ codes.Code
var _ io.Reader
var _ status.Status
var _ = runtime.String
var _ = utilities.NewDoubleArray
var _ = metadata.Join

func request_Endpoint_List_0(ctx context.Context, marshaler runtime.Marshaler, client v2_0.EndpointClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq v2_0.EndpointSearchQuery
	var metadata runtime.ServerMetadata

	if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.List(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_Endpoint_List_0(ctx context.Context, marshaler runtime.Marshaler, server v2_0.EndpointServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq v2_0.EndpointSearchQuery
	var metadata runtime.ServerMetadata

	if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.List(ctx, &protoReq)
	return msg, metadata, err

}

// RegisterEndpointHandlerServer registers the http handlers for service Endpoint to "mux".
// UnaryRPC     :call EndpointServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
// Note that using this registration option will cause many gRPC library features to stop working. Consider using RegisterEndpointHandlerFromEndpoint instead.
func RegisterEndpointHandlerServer(ctx context.Context, mux *runtime.ServeMux, server v2_0.EndpointServer) error {

	mux.Handle("POST", pattern_Endpoint_List_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateIncomingContext(ctx, mux, req, "/spaceone.api.identity.v2.Endpoint/List", runtime.WithHTTPPathPattern("/identity/v2/endpoint/list"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_Endpoint_List_0(annotatedContext, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_Endpoint_List_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

// RegisterEndpointHandlerFromEndpoint is same as RegisterEndpointHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterEndpointHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.NewClient(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Errorf("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Errorf("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterEndpointHandler(ctx, mux, conn)
}

// RegisterEndpointHandler registers the http handlers for service Endpoint to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterEndpointHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterEndpointHandlerClient(ctx, mux, v2_0.NewEndpointClient(conn))
}

// RegisterEndpointHandlerClient registers the http handlers for service Endpoint
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "v2_0.EndpointClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "v2_0.EndpointClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "v2_0.EndpointClient" to call the correct interceptors.
func RegisterEndpointHandlerClient(ctx context.Context, mux *runtime.ServeMux, client v2_0.EndpointClient) error {

	mux.Handle("POST", pattern_Endpoint_List_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateContext(ctx, mux, req, "/spaceone.api.identity.v2.Endpoint/List", runtime.WithHTTPPathPattern("/identity/v2/endpoint/list"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_Endpoint_List_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_Endpoint_List_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_Endpoint_List_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3}, []string{"identity", "v2", "endpoint", "list"}, ""))
)

var (
	forward_Endpoint_List_0 = runtime.ForwardResponseMessage
)
