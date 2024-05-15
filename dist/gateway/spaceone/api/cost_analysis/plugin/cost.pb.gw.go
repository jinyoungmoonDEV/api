// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: spaceone/api/cost_analysis/plugin/cost.proto

/*
Package plugin is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package plugin

import (
	"context"
	"io"
	"net/http"

	extPlugin "github.com/cloudforet-io/api/dist/go/spaceone/api/cost_analysis/plugin"
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

func request_Cost_GetLinkedAccounts_0(ctx context.Context, marshaler runtime.Marshaler, client extPlugin.CostClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq extPlugin.GetLinkedAccountsRequest
	var metadata runtime.ServerMetadata

	if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.GetLinkedAccounts(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_Cost_GetLinkedAccounts_0(ctx context.Context, marshaler runtime.Marshaler, server extPlugin.CostServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq extPlugin.GetLinkedAccountsRequest
	var metadata runtime.ServerMetadata

	if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.GetLinkedAccounts(ctx, &protoReq)
	return msg, metadata, err

}

func request_Cost_GetData_0(ctx context.Context, marshaler runtime.Marshaler, client extPlugin.CostClient, req *http.Request, pathParams map[string]string) (extPlugin.Cost_GetDataClient, runtime.ServerMetadata, error) {
	var protoReq extPlugin.GetDataRequest
	var metadata runtime.ServerMetadata

	if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	stream, err := client.GetData(ctx, &protoReq)
	if err != nil {
		return nil, metadata, err
	}
	header, err := stream.Header()
	if err != nil {
		return nil, metadata, err
	}
	metadata.HeaderMD = header
	return stream, metadata, nil

}

// RegisterCostHandlerServer registers the http handlers for service Cost to "mux".
// UnaryRPC     :call CostServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
// Note that using this registration option will cause many gRPC library features to stop working. Consider using RegisterCostHandlerFromEndpoint instead.
func RegisterCostHandlerServer(ctx context.Context, mux *runtime.ServeMux, server extPlugin.CostServer) error {

	mux.Handle("POST", pattern_Cost_GetLinkedAccounts_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateIncomingContext(ctx, mux, req, "/spaceone.api.cost_analysis.plugin.Cost/GetLinkedAccounts", runtime.WithHTTPPathPattern("/spaceone.api.cost_analysis.plugin.Cost/get_linked_accounts"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_Cost_GetLinkedAccounts_0(annotatedContext, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_Cost_GetLinkedAccounts_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_Cost_GetData_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		err := status.Error(codes.Unimplemented, "streaming calls are not yet supported in the in-process transport")
		_, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
		return
	})

	return nil
}

// RegisterCostHandlerFromEndpoint is same as RegisterCostHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterCostHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
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

	return RegisterCostHandler(ctx, mux, conn)
}

// RegisterCostHandler registers the http handlers for service Cost to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterCostHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterCostHandlerClient(ctx, mux, extPlugin.NewCostClient(conn))
}

// RegisterCostHandlerClient registers the http handlers for service Cost
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "extPlugin.CostClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "extPlugin.CostClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "extPlugin.CostClient" to call the correct interceptors.
func RegisterCostHandlerClient(ctx context.Context, mux *runtime.ServeMux, client extPlugin.CostClient) error {

	mux.Handle("POST", pattern_Cost_GetLinkedAccounts_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateContext(ctx, mux, req, "/spaceone.api.cost_analysis.plugin.Cost/GetLinkedAccounts", runtime.WithHTTPPathPattern("/spaceone.api.cost_analysis.plugin.Cost/get_linked_accounts"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_Cost_GetLinkedAccounts_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_Cost_GetLinkedAccounts_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_Cost_GetData_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateContext(ctx, mux, req, "/spaceone.api.cost_analysis.plugin.Cost/GetData", runtime.WithHTTPPathPattern("/spaceone.api.cost_analysis.plugin.Cost/get_data"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_Cost_GetData_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_Cost_GetData_0(annotatedContext, mux, outboundMarshaler, w, req, func() (proto.Message, error) { return resp.Recv() }, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_Cost_GetLinkedAccounts_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1}, []string{"spaceone.api.cost_analysis.plugin.Cost", "get_linked_accounts"}, ""))

	pattern_Cost_GetData_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1}, []string{"spaceone.api.cost_analysis.plugin.Cost", "get_data"}, ""))
)

var (
	forward_Cost_GetLinkedAccounts_0 = runtime.ForwardResponseMessage

	forward_Cost_GetData_0 = runtime.ForwardResponseStream
)
