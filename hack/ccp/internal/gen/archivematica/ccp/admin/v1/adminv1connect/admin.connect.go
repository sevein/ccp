// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: archivematica/ccp/admin/v1/admin.proto

package adminv1connect

import (
	context "context"
	errors "errors"
	v1 "github.com/artefactual/archivematica/hack/ccp/internal/gen/archivematica/ccp/admin/v1"
	connect_go "github.com/bufbuild/connect-go"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// AdminServiceName is the fully-qualified name of the AdminService service.
	AdminServiceName = "archivematica.ccp.admin.v1.AdminService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// AdminServiceListActivePackagesProcedure is the fully-qualified name of the AdminService's
	// ListActivePackages RPC.
	AdminServiceListActivePackagesProcedure = "/archivematica.ccp.admin.v1.AdminService/ListActivePackages"
	// AdminServiceListAwaitingDecisionsProcedure is the fully-qualified name of the AdminService's
	// ListAwaitingDecisions RPC.
	AdminServiceListAwaitingDecisionsProcedure = "/archivematica.ccp.admin.v1.AdminService/ListAwaitingDecisions"
)

// AdminServiceClient is a client for the archivematica.ccp.admin.v1.AdminService service.
type AdminServiceClient interface {
	ListActivePackages(context.Context, *connect_go.Request[v1.ListActivePackagesRequest]) (*connect_go.Response[v1.ListActivePackagesResponse], error)
	ListAwaitingDecisions(context.Context, *connect_go.Request[v1.ListAwaitingDecisionsRequest]) (*connect_go.Response[v1.ListAwaitingDecisionsResponse], error)
}

// NewAdminServiceClient constructs a client for the archivematica.ccp.admin.v1.AdminService
// service. By default, it uses the Connect protocol with the binary Protobuf Codec, asks for
// gzipped responses, and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply
// the connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewAdminServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) AdminServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &adminServiceClient{
		listActivePackages: connect_go.NewClient[v1.ListActivePackagesRequest, v1.ListActivePackagesResponse](
			httpClient,
			baseURL+AdminServiceListActivePackagesProcedure,
			opts...,
		),
		listAwaitingDecisions: connect_go.NewClient[v1.ListAwaitingDecisionsRequest, v1.ListAwaitingDecisionsResponse](
			httpClient,
			baseURL+AdminServiceListAwaitingDecisionsProcedure,
			opts...,
		),
	}
}

// adminServiceClient implements AdminServiceClient.
type adminServiceClient struct {
	listActivePackages    *connect_go.Client[v1.ListActivePackagesRequest, v1.ListActivePackagesResponse]
	listAwaitingDecisions *connect_go.Client[v1.ListAwaitingDecisionsRequest, v1.ListAwaitingDecisionsResponse]
}

// ListActivePackages calls archivematica.ccp.admin.v1.AdminService.ListActivePackages.
func (c *adminServiceClient) ListActivePackages(ctx context.Context, req *connect_go.Request[v1.ListActivePackagesRequest]) (*connect_go.Response[v1.ListActivePackagesResponse], error) {
	return c.listActivePackages.CallUnary(ctx, req)
}

// ListAwaitingDecisions calls archivematica.ccp.admin.v1.AdminService.ListAwaitingDecisions.
func (c *adminServiceClient) ListAwaitingDecisions(ctx context.Context, req *connect_go.Request[v1.ListAwaitingDecisionsRequest]) (*connect_go.Response[v1.ListAwaitingDecisionsResponse], error) {
	return c.listAwaitingDecisions.CallUnary(ctx, req)
}

// AdminServiceHandler is an implementation of the archivematica.ccp.admin.v1.AdminService service.
type AdminServiceHandler interface {
	ListActivePackages(context.Context, *connect_go.Request[v1.ListActivePackagesRequest]) (*connect_go.Response[v1.ListActivePackagesResponse], error)
	ListAwaitingDecisions(context.Context, *connect_go.Request[v1.ListAwaitingDecisionsRequest]) (*connect_go.Response[v1.ListAwaitingDecisionsResponse], error)
}

// NewAdminServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewAdminServiceHandler(svc AdminServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle(AdminServiceListActivePackagesProcedure, connect_go.NewUnaryHandler(
		AdminServiceListActivePackagesProcedure,
		svc.ListActivePackages,
		opts...,
	))
	mux.Handle(AdminServiceListAwaitingDecisionsProcedure, connect_go.NewUnaryHandler(
		AdminServiceListAwaitingDecisionsProcedure,
		svc.ListAwaitingDecisions,
		opts...,
	))
	return "/archivematica.ccp.admin.v1.AdminService/", mux
}

// UnimplementedAdminServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedAdminServiceHandler struct{}

func (UnimplementedAdminServiceHandler) ListActivePackages(context.Context, *connect_go.Request[v1.ListActivePackagesRequest]) (*connect_go.Response[v1.ListActivePackagesResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("archivematica.ccp.admin.v1.AdminService.ListActivePackages is not implemented"))
}

func (UnimplementedAdminServiceHandler) ListAwaitingDecisions(context.Context, *connect_go.Request[v1.ListAwaitingDecisionsRequest]) (*connect_go.Response[v1.ListAwaitingDecisionsResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("archivematica.ccp.admin.v1.AdminService.ListAwaitingDecisions is not implemented"))
}
