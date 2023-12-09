// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: mgmt/v1alpha1/connection.proto

package mgmtv1alpha1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1alpha1 "github.com/nucleuscloud/neosync/backend/gen/go/protos/mgmt/v1alpha1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion0_1_0

const (
	// ConnectionServiceName is the fully-qualified name of the ConnectionService service.
	ConnectionServiceName = "mgmt.v1alpha1.ConnectionService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// ConnectionServiceGetConnectionsProcedure is the fully-qualified name of the ConnectionService's
	// GetConnections RPC.
	ConnectionServiceGetConnectionsProcedure = "/mgmt.v1alpha1.ConnectionService/GetConnections"
	// ConnectionServiceGetConnectionProcedure is the fully-qualified name of the ConnectionService's
	// GetConnection RPC.
	ConnectionServiceGetConnectionProcedure = "/mgmt.v1alpha1.ConnectionService/GetConnection"
	// ConnectionServiceCreateConnectionProcedure is the fully-qualified name of the ConnectionService's
	// CreateConnection RPC.
	ConnectionServiceCreateConnectionProcedure = "/mgmt.v1alpha1.ConnectionService/CreateConnection"
	// ConnectionServiceUpdateConnectionProcedure is the fully-qualified name of the ConnectionService's
	// UpdateConnection RPC.
	ConnectionServiceUpdateConnectionProcedure = "/mgmt.v1alpha1.ConnectionService/UpdateConnection"
	// ConnectionServiceDeleteConnectionProcedure is the fully-qualified name of the ConnectionService's
	// DeleteConnection RPC.
	ConnectionServiceDeleteConnectionProcedure = "/mgmt.v1alpha1.ConnectionService/DeleteConnection"
	// ConnectionServiceIsConnectionNameAvailableProcedure is the fully-qualified name of the
	// ConnectionService's IsConnectionNameAvailable RPC.
	ConnectionServiceIsConnectionNameAvailableProcedure = "/mgmt.v1alpha1.ConnectionService/IsConnectionNameAvailable"
	// ConnectionServiceCheckConnectionConfigProcedure is the fully-qualified name of the
	// ConnectionService's CheckConnectionConfig RPC.
	ConnectionServiceCheckConnectionConfigProcedure = "/mgmt.v1alpha1.ConnectionService/CheckConnectionConfig"
	// ConnectionServiceGetConnectionSchemaProcedure is the fully-qualified name of the
	// ConnectionService's GetConnectionSchema RPC.
	ConnectionServiceGetConnectionSchemaProcedure = "/mgmt.v1alpha1.ConnectionService/GetConnectionSchema"
	// ConnectionServiceCheckSqlQueryProcedure is the fully-qualified name of the ConnectionService's
	// CheckSqlQuery RPC.
	ConnectionServiceCheckSqlQueryProcedure = "/mgmt.v1alpha1.ConnectionService/CheckSqlQuery"
	// ConnectionServiceGetConnectionDataStreamProcedure is the fully-qualified name of the
	// ConnectionService's GetConnectionDataStream RPC.
	ConnectionServiceGetConnectionDataStreamProcedure = "/mgmt.v1alpha1.ConnectionService/GetConnectionDataStream"
	// ConnectionServiceGetConnectionForeignConstraintsProcedure is the fully-qualified name of the
	// ConnectionService's GetConnectionForeignConstraints RPC.
	ConnectionServiceGetConnectionForeignConstraintsProcedure = "/mgmt.v1alpha1.ConnectionService/GetConnectionForeignConstraints"
)

// ConnectionServiceClient is a client for the mgmt.v1alpha1.ConnectionService service.
type ConnectionServiceClient interface {
	// Returns a list of connections associated with the account
	GetConnections(context.Context, *connect.Request[v1alpha1.GetConnectionsRequest]) (*connect.Response[v1alpha1.GetConnectionsResponse], error)
	// Returns a single connection
	GetConnection(context.Context, *connect.Request[v1alpha1.GetConnectionRequest]) (*connect.Response[v1alpha1.GetConnectionResponse], error)
	// Creates a new connection
	CreateConnection(context.Context, *connect.Request[v1alpha1.CreateConnectionRequest]) (*connect.Response[v1alpha1.CreateConnectionResponse], error)
	// Updates an existing connection
	UpdateConnection(context.Context, *connect.Request[v1alpha1.UpdateConnectionRequest]) (*connect.Response[v1alpha1.UpdateConnectionResponse], error)
	// Removes a connection from the system.
	DeleteConnection(context.Context, *connect.Request[v1alpha1.DeleteConnectionRequest]) (*connect.Response[v1alpha1.DeleteConnectionResponse], error)
	// Connections have friendly names, this method checks if the requested name is available in the system based on the account
	IsConnectionNameAvailable(context.Context, *connect.Request[v1alpha1.IsConnectionNameAvailableRequest]) (*connect.Response[v1alpha1.IsConnectionNameAvailableResponse], error)
	// Checks if the connection config is connectable by the backend.
	// Used mostly to verify that a connection is valid prior to creating a Connection object.
	CheckConnectionConfig(context.Context, *connect.Request[v1alpha1.CheckConnectionConfigRequest]) (*connect.Response[v1alpha1.CheckConnectionConfigResponse], error)
	// Returns the schema for a specific connection. Used mostly for SQL-based connections
	GetConnectionSchema(context.Context, *connect.Request[v1alpha1.GetConnectionSchemaRequest]) (*connect.Response[v1alpha1.GetConnectionSchemaResponse], error)
	// Checks a constructed SQL query against a sql-based connection to see if it's valid based on that connection's data schema
	// This is useful when constructing subsets to see if the WHERE clause is correct
	CheckSqlQuery(context.Context, *connect.Request[v1alpha1.CheckSqlQueryRequest]) (*connect.Response[v1alpha1.CheckSqlQueryResponse], error)
	// Streaming endpoint that will stream the data available from the Connection to the client.
	// Used primarily by the CLI sync command.
	GetConnectionDataStream(context.Context, *connect.Request[v1alpha1.GetConnectionDataStreamRequest]) (*connect.ServerStreamForClient[v1alpha1.GetConnectionDataStreamResponse], error)
	// For a specific connection, returns the foreign key constraints. Mostly useful for SQL-based Connections.
	// Used primarily by the CLI sync command to determine stream order.
	GetConnectionForeignConstraints(context.Context, *connect.Request[v1alpha1.GetConnectionForeignConstraintsRequest]) (*connect.Response[v1alpha1.GetConnectionForeignConstraintsResponse], error)
}

// NewConnectionServiceClient constructs a client for the mgmt.v1alpha1.ConnectionService service.
// By default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped
// responses, and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewConnectionServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) ConnectionServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &connectionServiceClient{
		getConnections: connect.NewClient[v1alpha1.GetConnectionsRequest, v1alpha1.GetConnectionsResponse](
			httpClient,
			baseURL+ConnectionServiceGetConnectionsProcedure,
			opts...,
		),
		getConnection: connect.NewClient[v1alpha1.GetConnectionRequest, v1alpha1.GetConnectionResponse](
			httpClient,
			baseURL+ConnectionServiceGetConnectionProcedure,
			opts...,
		),
		createConnection: connect.NewClient[v1alpha1.CreateConnectionRequest, v1alpha1.CreateConnectionResponse](
			httpClient,
			baseURL+ConnectionServiceCreateConnectionProcedure,
			opts...,
		),
		updateConnection: connect.NewClient[v1alpha1.UpdateConnectionRequest, v1alpha1.UpdateConnectionResponse](
			httpClient,
			baseURL+ConnectionServiceUpdateConnectionProcedure,
			opts...,
		),
		deleteConnection: connect.NewClient[v1alpha1.DeleteConnectionRequest, v1alpha1.DeleteConnectionResponse](
			httpClient,
			baseURL+ConnectionServiceDeleteConnectionProcedure,
			opts...,
		),
		isConnectionNameAvailable: connect.NewClient[v1alpha1.IsConnectionNameAvailableRequest, v1alpha1.IsConnectionNameAvailableResponse](
			httpClient,
			baseURL+ConnectionServiceIsConnectionNameAvailableProcedure,
			opts...,
		),
		checkConnectionConfig: connect.NewClient[v1alpha1.CheckConnectionConfigRequest, v1alpha1.CheckConnectionConfigResponse](
			httpClient,
			baseURL+ConnectionServiceCheckConnectionConfigProcedure,
			opts...,
		),
		getConnectionSchema: connect.NewClient[v1alpha1.GetConnectionSchemaRequest, v1alpha1.GetConnectionSchemaResponse](
			httpClient,
			baseURL+ConnectionServiceGetConnectionSchemaProcedure,
			opts...,
		),
		checkSqlQuery: connect.NewClient[v1alpha1.CheckSqlQueryRequest, v1alpha1.CheckSqlQueryResponse](
			httpClient,
			baseURL+ConnectionServiceCheckSqlQueryProcedure,
			opts...,
		),
		getConnectionDataStream: connect.NewClient[v1alpha1.GetConnectionDataStreamRequest, v1alpha1.GetConnectionDataStreamResponse](
			httpClient,
			baseURL+ConnectionServiceGetConnectionDataStreamProcedure,
			opts...,
		),
		getConnectionForeignConstraints: connect.NewClient[v1alpha1.GetConnectionForeignConstraintsRequest, v1alpha1.GetConnectionForeignConstraintsResponse](
			httpClient,
			baseURL+ConnectionServiceGetConnectionForeignConstraintsProcedure,
			opts...,
		),
	}
}

// connectionServiceClient implements ConnectionServiceClient.
type connectionServiceClient struct {
	getConnections                  *connect.Client[v1alpha1.GetConnectionsRequest, v1alpha1.GetConnectionsResponse]
	getConnection                   *connect.Client[v1alpha1.GetConnectionRequest, v1alpha1.GetConnectionResponse]
	createConnection                *connect.Client[v1alpha1.CreateConnectionRequest, v1alpha1.CreateConnectionResponse]
	updateConnection                *connect.Client[v1alpha1.UpdateConnectionRequest, v1alpha1.UpdateConnectionResponse]
	deleteConnection                *connect.Client[v1alpha1.DeleteConnectionRequest, v1alpha1.DeleteConnectionResponse]
	isConnectionNameAvailable       *connect.Client[v1alpha1.IsConnectionNameAvailableRequest, v1alpha1.IsConnectionNameAvailableResponse]
	checkConnectionConfig           *connect.Client[v1alpha1.CheckConnectionConfigRequest, v1alpha1.CheckConnectionConfigResponse]
	getConnectionSchema             *connect.Client[v1alpha1.GetConnectionSchemaRequest, v1alpha1.GetConnectionSchemaResponse]
	checkSqlQuery                   *connect.Client[v1alpha1.CheckSqlQueryRequest, v1alpha1.CheckSqlQueryResponse]
	getConnectionDataStream         *connect.Client[v1alpha1.GetConnectionDataStreamRequest, v1alpha1.GetConnectionDataStreamResponse]
	getConnectionForeignConstraints *connect.Client[v1alpha1.GetConnectionForeignConstraintsRequest, v1alpha1.GetConnectionForeignConstraintsResponse]
}

// GetConnections calls mgmt.v1alpha1.ConnectionService.GetConnections.
func (c *connectionServiceClient) GetConnections(ctx context.Context, req *connect.Request[v1alpha1.GetConnectionsRequest]) (*connect.Response[v1alpha1.GetConnectionsResponse], error) {
	return c.getConnections.CallUnary(ctx, req)
}

// GetConnection calls mgmt.v1alpha1.ConnectionService.GetConnection.
func (c *connectionServiceClient) GetConnection(ctx context.Context, req *connect.Request[v1alpha1.GetConnectionRequest]) (*connect.Response[v1alpha1.GetConnectionResponse], error) {
	return c.getConnection.CallUnary(ctx, req)
}

// CreateConnection calls mgmt.v1alpha1.ConnectionService.CreateConnection.
func (c *connectionServiceClient) CreateConnection(ctx context.Context, req *connect.Request[v1alpha1.CreateConnectionRequest]) (*connect.Response[v1alpha1.CreateConnectionResponse], error) {
	return c.createConnection.CallUnary(ctx, req)
}

// UpdateConnection calls mgmt.v1alpha1.ConnectionService.UpdateConnection.
func (c *connectionServiceClient) UpdateConnection(ctx context.Context, req *connect.Request[v1alpha1.UpdateConnectionRequest]) (*connect.Response[v1alpha1.UpdateConnectionResponse], error) {
	return c.updateConnection.CallUnary(ctx, req)
}

// DeleteConnection calls mgmt.v1alpha1.ConnectionService.DeleteConnection.
func (c *connectionServiceClient) DeleteConnection(ctx context.Context, req *connect.Request[v1alpha1.DeleteConnectionRequest]) (*connect.Response[v1alpha1.DeleteConnectionResponse], error) {
	return c.deleteConnection.CallUnary(ctx, req)
}

// IsConnectionNameAvailable calls mgmt.v1alpha1.ConnectionService.IsConnectionNameAvailable.
func (c *connectionServiceClient) IsConnectionNameAvailable(ctx context.Context, req *connect.Request[v1alpha1.IsConnectionNameAvailableRequest]) (*connect.Response[v1alpha1.IsConnectionNameAvailableResponse], error) {
	return c.isConnectionNameAvailable.CallUnary(ctx, req)
}

// CheckConnectionConfig calls mgmt.v1alpha1.ConnectionService.CheckConnectionConfig.
func (c *connectionServiceClient) CheckConnectionConfig(ctx context.Context, req *connect.Request[v1alpha1.CheckConnectionConfigRequest]) (*connect.Response[v1alpha1.CheckConnectionConfigResponse], error) {
	return c.checkConnectionConfig.CallUnary(ctx, req)
}

// GetConnectionSchema calls mgmt.v1alpha1.ConnectionService.GetConnectionSchema.
func (c *connectionServiceClient) GetConnectionSchema(ctx context.Context, req *connect.Request[v1alpha1.GetConnectionSchemaRequest]) (*connect.Response[v1alpha1.GetConnectionSchemaResponse], error) {
	return c.getConnectionSchema.CallUnary(ctx, req)
}

// CheckSqlQuery calls mgmt.v1alpha1.ConnectionService.CheckSqlQuery.
func (c *connectionServiceClient) CheckSqlQuery(ctx context.Context, req *connect.Request[v1alpha1.CheckSqlQueryRequest]) (*connect.Response[v1alpha1.CheckSqlQueryResponse], error) {
	return c.checkSqlQuery.CallUnary(ctx, req)
}

// GetConnectionDataStream calls mgmt.v1alpha1.ConnectionService.GetConnectionDataStream.
func (c *connectionServiceClient) GetConnectionDataStream(ctx context.Context, req *connect.Request[v1alpha1.GetConnectionDataStreamRequest]) (*connect.ServerStreamForClient[v1alpha1.GetConnectionDataStreamResponse], error) {
	return c.getConnectionDataStream.CallServerStream(ctx, req)
}

// GetConnectionForeignConstraints calls
// mgmt.v1alpha1.ConnectionService.GetConnectionForeignConstraints.
func (c *connectionServiceClient) GetConnectionForeignConstraints(ctx context.Context, req *connect.Request[v1alpha1.GetConnectionForeignConstraintsRequest]) (*connect.Response[v1alpha1.GetConnectionForeignConstraintsResponse], error) {
	return c.getConnectionForeignConstraints.CallUnary(ctx, req)
}

// ConnectionServiceHandler is an implementation of the mgmt.v1alpha1.ConnectionService service.
type ConnectionServiceHandler interface {
	// Returns a list of connections associated with the account
	GetConnections(context.Context, *connect.Request[v1alpha1.GetConnectionsRequest]) (*connect.Response[v1alpha1.GetConnectionsResponse], error)
	// Returns a single connection
	GetConnection(context.Context, *connect.Request[v1alpha1.GetConnectionRequest]) (*connect.Response[v1alpha1.GetConnectionResponse], error)
	// Creates a new connection
	CreateConnection(context.Context, *connect.Request[v1alpha1.CreateConnectionRequest]) (*connect.Response[v1alpha1.CreateConnectionResponse], error)
	// Updates an existing connection
	UpdateConnection(context.Context, *connect.Request[v1alpha1.UpdateConnectionRequest]) (*connect.Response[v1alpha1.UpdateConnectionResponse], error)
	// Removes a connection from the system.
	DeleteConnection(context.Context, *connect.Request[v1alpha1.DeleteConnectionRequest]) (*connect.Response[v1alpha1.DeleteConnectionResponse], error)
	// Connections have friendly names, this method checks if the requested name is available in the system based on the account
	IsConnectionNameAvailable(context.Context, *connect.Request[v1alpha1.IsConnectionNameAvailableRequest]) (*connect.Response[v1alpha1.IsConnectionNameAvailableResponse], error)
	// Checks if the connection config is connectable by the backend.
	// Used mostly to verify that a connection is valid prior to creating a Connection object.
	CheckConnectionConfig(context.Context, *connect.Request[v1alpha1.CheckConnectionConfigRequest]) (*connect.Response[v1alpha1.CheckConnectionConfigResponse], error)
	// Returns the schema for a specific connection. Used mostly for SQL-based connections
	GetConnectionSchema(context.Context, *connect.Request[v1alpha1.GetConnectionSchemaRequest]) (*connect.Response[v1alpha1.GetConnectionSchemaResponse], error)
	// Checks a constructed SQL query against a sql-based connection to see if it's valid based on that connection's data schema
	// This is useful when constructing subsets to see if the WHERE clause is correct
	CheckSqlQuery(context.Context, *connect.Request[v1alpha1.CheckSqlQueryRequest]) (*connect.Response[v1alpha1.CheckSqlQueryResponse], error)
	// Streaming endpoint that will stream the data available from the Connection to the client.
	// Used primarily by the CLI sync command.
	GetConnectionDataStream(context.Context, *connect.Request[v1alpha1.GetConnectionDataStreamRequest], *connect.ServerStream[v1alpha1.GetConnectionDataStreamResponse]) error
	// For a specific connection, returns the foreign key constraints. Mostly useful for SQL-based Connections.
	// Used primarily by the CLI sync command to determine stream order.
	GetConnectionForeignConstraints(context.Context, *connect.Request[v1alpha1.GetConnectionForeignConstraintsRequest]) (*connect.Response[v1alpha1.GetConnectionForeignConstraintsResponse], error)
}

// NewConnectionServiceHandler builds an HTTP handler from the service implementation. It returns
// the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewConnectionServiceHandler(svc ConnectionServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	connectionServiceGetConnectionsHandler := connect.NewUnaryHandler(
		ConnectionServiceGetConnectionsProcedure,
		svc.GetConnections,
		opts...,
	)
	connectionServiceGetConnectionHandler := connect.NewUnaryHandler(
		ConnectionServiceGetConnectionProcedure,
		svc.GetConnection,
		opts...,
	)
	connectionServiceCreateConnectionHandler := connect.NewUnaryHandler(
		ConnectionServiceCreateConnectionProcedure,
		svc.CreateConnection,
		opts...,
	)
	connectionServiceUpdateConnectionHandler := connect.NewUnaryHandler(
		ConnectionServiceUpdateConnectionProcedure,
		svc.UpdateConnection,
		opts...,
	)
	connectionServiceDeleteConnectionHandler := connect.NewUnaryHandler(
		ConnectionServiceDeleteConnectionProcedure,
		svc.DeleteConnection,
		opts...,
	)
	connectionServiceIsConnectionNameAvailableHandler := connect.NewUnaryHandler(
		ConnectionServiceIsConnectionNameAvailableProcedure,
		svc.IsConnectionNameAvailable,
		opts...,
	)
	connectionServiceCheckConnectionConfigHandler := connect.NewUnaryHandler(
		ConnectionServiceCheckConnectionConfigProcedure,
		svc.CheckConnectionConfig,
		opts...,
	)
	connectionServiceGetConnectionSchemaHandler := connect.NewUnaryHandler(
		ConnectionServiceGetConnectionSchemaProcedure,
		svc.GetConnectionSchema,
		opts...,
	)
	connectionServiceCheckSqlQueryHandler := connect.NewUnaryHandler(
		ConnectionServiceCheckSqlQueryProcedure,
		svc.CheckSqlQuery,
		opts...,
	)
	connectionServiceGetConnectionDataStreamHandler := connect.NewServerStreamHandler(
		ConnectionServiceGetConnectionDataStreamProcedure,
		svc.GetConnectionDataStream,
		opts...,
	)
	connectionServiceGetConnectionForeignConstraintsHandler := connect.NewUnaryHandler(
		ConnectionServiceGetConnectionForeignConstraintsProcedure,
		svc.GetConnectionForeignConstraints,
		opts...,
	)
	return "/mgmt.v1alpha1.ConnectionService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case ConnectionServiceGetConnectionsProcedure:
			connectionServiceGetConnectionsHandler.ServeHTTP(w, r)
		case ConnectionServiceGetConnectionProcedure:
			connectionServiceGetConnectionHandler.ServeHTTP(w, r)
		case ConnectionServiceCreateConnectionProcedure:
			connectionServiceCreateConnectionHandler.ServeHTTP(w, r)
		case ConnectionServiceUpdateConnectionProcedure:
			connectionServiceUpdateConnectionHandler.ServeHTTP(w, r)
		case ConnectionServiceDeleteConnectionProcedure:
			connectionServiceDeleteConnectionHandler.ServeHTTP(w, r)
		case ConnectionServiceIsConnectionNameAvailableProcedure:
			connectionServiceIsConnectionNameAvailableHandler.ServeHTTP(w, r)
		case ConnectionServiceCheckConnectionConfigProcedure:
			connectionServiceCheckConnectionConfigHandler.ServeHTTP(w, r)
		case ConnectionServiceGetConnectionSchemaProcedure:
			connectionServiceGetConnectionSchemaHandler.ServeHTTP(w, r)
		case ConnectionServiceCheckSqlQueryProcedure:
			connectionServiceCheckSqlQueryHandler.ServeHTTP(w, r)
		case ConnectionServiceGetConnectionDataStreamProcedure:
			connectionServiceGetConnectionDataStreamHandler.ServeHTTP(w, r)
		case ConnectionServiceGetConnectionForeignConstraintsProcedure:
			connectionServiceGetConnectionForeignConstraintsHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedConnectionServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedConnectionServiceHandler struct{}

func (UnimplementedConnectionServiceHandler) GetConnections(context.Context, *connect.Request[v1alpha1.GetConnectionsRequest]) (*connect.Response[v1alpha1.GetConnectionsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("mgmt.v1alpha1.ConnectionService.GetConnections is not implemented"))
}

func (UnimplementedConnectionServiceHandler) GetConnection(context.Context, *connect.Request[v1alpha1.GetConnectionRequest]) (*connect.Response[v1alpha1.GetConnectionResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("mgmt.v1alpha1.ConnectionService.GetConnection is not implemented"))
}

func (UnimplementedConnectionServiceHandler) CreateConnection(context.Context, *connect.Request[v1alpha1.CreateConnectionRequest]) (*connect.Response[v1alpha1.CreateConnectionResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("mgmt.v1alpha1.ConnectionService.CreateConnection is not implemented"))
}

func (UnimplementedConnectionServiceHandler) UpdateConnection(context.Context, *connect.Request[v1alpha1.UpdateConnectionRequest]) (*connect.Response[v1alpha1.UpdateConnectionResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("mgmt.v1alpha1.ConnectionService.UpdateConnection is not implemented"))
}

func (UnimplementedConnectionServiceHandler) DeleteConnection(context.Context, *connect.Request[v1alpha1.DeleteConnectionRequest]) (*connect.Response[v1alpha1.DeleteConnectionResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("mgmt.v1alpha1.ConnectionService.DeleteConnection is not implemented"))
}

func (UnimplementedConnectionServiceHandler) IsConnectionNameAvailable(context.Context, *connect.Request[v1alpha1.IsConnectionNameAvailableRequest]) (*connect.Response[v1alpha1.IsConnectionNameAvailableResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("mgmt.v1alpha1.ConnectionService.IsConnectionNameAvailable is not implemented"))
}

func (UnimplementedConnectionServiceHandler) CheckConnectionConfig(context.Context, *connect.Request[v1alpha1.CheckConnectionConfigRequest]) (*connect.Response[v1alpha1.CheckConnectionConfigResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("mgmt.v1alpha1.ConnectionService.CheckConnectionConfig is not implemented"))
}

func (UnimplementedConnectionServiceHandler) GetConnectionSchema(context.Context, *connect.Request[v1alpha1.GetConnectionSchemaRequest]) (*connect.Response[v1alpha1.GetConnectionSchemaResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("mgmt.v1alpha1.ConnectionService.GetConnectionSchema is not implemented"))
}

func (UnimplementedConnectionServiceHandler) CheckSqlQuery(context.Context, *connect.Request[v1alpha1.CheckSqlQueryRequest]) (*connect.Response[v1alpha1.CheckSqlQueryResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("mgmt.v1alpha1.ConnectionService.CheckSqlQuery is not implemented"))
}

func (UnimplementedConnectionServiceHandler) GetConnectionDataStream(context.Context, *connect.Request[v1alpha1.GetConnectionDataStreamRequest], *connect.ServerStream[v1alpha1.GetConnectionDataStreamResponse]) error {
	return connect.NewError(connect.CodeUnimplemented, errors.New("mgmt.v1alpha1.ConnectionService.GetConnectionDataStream is not implemented"))
}

func (UnimplementedConnectionServiceHandler) GetConnectionForeignConstraints(context.Context, *connect.Request[v1alpha1.GetConnectionForeignConstraintsRequest]) (*connect.Response[v1alpha1.GetConnectionForeignConstraintsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("mgmt.v1alpha1.ConnectionService.GetConnectionForeignConstraints is not implemented"))
}
