// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package resourcemanager

import (
	"context"
	"fmt"
	"math"
	"net/url"
	"time"

	"cloud.google.com/go/longrunning"
	lroauto "cloud.google.com/go/longrunning/autogen"
	resourcemanagerpb "cloud.google.com/go/resourcemanager/apiv3/resourcemanagerpb"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
)

var newTagBindingsClientHook clientHook

// TagBindingsCallOptions contains the retry settings for each method of TagBindingsClient.
type TagBindingsCallOptions struct {
	ListTagBindings  []gax.CallOption
	CreateTagBinding []gax.CallOption
	DeleteTagBinding []gax.CallOption
}

func defaultTagBindingsGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("cloudresourcemanager.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("cloudresourcemanager.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://cloudresourcemanager.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultTagBindingsCallOptions() *TagBindingsCallOptions {
	return &TagBindingsCallOptions{
		ListTagBindings: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		CreateTagBinding: []gax.CallOption{},
		DeleteTagBinding: []gax.CallOption{},
	}
}

// internalTagBindingsClient is an interface that defines the methods available from Cloud Resource Manager API.
type internalTagBindingsClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	ListTagBindings(context.Context, *resourcemanagerpb.ListTagBindingsRequest, ...gax.CallOption) *TagBindingIterator
	CreateTagBinding(context.Context, *resourcemanagerpb.CreateTagBindingRequest, ...gax.CallOption) (*CreateTagBindingOperation, error)
	CreateTagBindingOperation(name string) *CreateTagBindingOperation
	DeleteTagBinding(context.Context, *resourcemanagerpb.DeleteTagBindingRequest, ...gax.CallOption) (*DeleteTagBindingOperation, error)
	DeleteTagBindingOperation(name string) *DeleteTagBindingOperation
}

// TagBindingsClient is a client for interacting with Cloud Resource Manager API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Allow users to create and manage TagBindings between TagValues and
// different cloud resources throughout the GCP resource hierarchy.
type TagBindingsClient struct {
	// The internal transport-dependent client.
	internalClient internalTagBindingsClient

	// The call options for this service.
	CallOptions *TagBindingsCallOptions

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient *lroauto.OperationsClient
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *TagBindingsClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *TagBindingsClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *TagBindingsClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// ListTagBindings lists the TagBindings for the given cloud resource, as specified with
// parent.
//
// NOTE: The parent field is expected to be a full resource name:
// https://cloud.google.com/apis/design/resource_names#full_resource_name (at https://cloud.google.com/apis/design/resource_names#full_resource_name)
func (c *TagBindingsClient) ListTagBindings(ctx context.Context, req *resourcemanagerpb.ListTagBindingsRequest, opts ...gax.CallOption) *TagBindingIterator {
	return c.internalClient.ListTagBindings(ctx, req, opts...)
}

// CreateTagBinding creates a TagBinding between a TagValue and a cloud resource
// (currently project, folder, or organization).
func (c *TagBindingsClient) CreateTagBinding(ctx context.Context, req *resourcemanagerpb.CreateTagBindingRequest, opts ...gax.CallOption) (*CreateTagBindingOperation, error) {
	return c.internalClient.CreateTagBinding(ctx, req, opts...)
}

// CreateTagBindingOperation returns a new CreateTagBindingOperation from a given name.
// The name must be that of a previously created CreateTagBindingOperation, possibly from a different process.
func (c *TagBindingsClient) CreateTagBindingOperation(name string) *CreateTagBindingOperation {
	return c.internalClient.CreateTagBindingOperation(name)
}

// DeleteTagBinding deletes a TagBinding.
func (c *TagBindingsClient) DeleteTagBinding(ctx context.Context, req *resourcemanagerpb.DeleteTagBindingRequest, opts ...gax.CallOption) (*DeleteTagBindingOperation, error) {
	return c.internalClient.DeleteTagBinding(ctx, req, opts...)
}

// DeleteTagBindingOperation returns a new DeleteTagBindingOperation from a given name.
// The name must be that of a previously created DeleteTagBindingOperation, possibly from a different process.
func (c *TagBindingsClient) DeleteTagBindingOperation(name string) *DeleteTagBindingOperation {
	return c.internalClient.DeleteTagBindingOperation(name)
}

// tagBindingsGRPCClient is a client for interacting with Cloud Resource Manager API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type tagBindingsGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing TagBindingsClient
	CallOptions **TagBindingsCallOptions

	// The gRPC API client.
	tagBindingsClient resourcemanagerpb.TagBindingsClient

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient **lroauto.OperationsClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewTagBindingsClient creates a new tag bindings client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Allow users to create and manage TagBindings between TagValues and
// different cloud resources throughout the GCP resource hierarchy.
func NewTagBindingsClient(ctx context.Context, opts ...option.ClientOption) (*TagBindingsClient, error) {
	clientOpts := defaultTagBindingsGRPCClientOptions()
	if newTagBindingsClientHook != nil {
		hookOpts, err := newTagBindingsClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	disableDeadlines, err := checkDisableDeadlines()
	if err != nil {
		return nil, err
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := TagBindingsClient{CallOptions: defaultTagBindingsCallOptions()}

	c := &tagBindingsGRPCClient{
		connPool:          connPool,
		disableDeadlines:  disableDeadlines,
		tagBindingsClient: resourcemanagerpb.NewTagBindingsClient(connPool),
		CallOptions:       &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	client.LROClient, err = lroauto.NewOperationsClient(ctx, gtransport.WithConnPool(connPool))
	if err != nil {
		// This error "should not happen", since we are just reusing old connection pool
		// and never actually need to dial.
		// If this does happen, we could leak connp. However, we cannot close conn:
		// If the user invoked the constructor with option.WithGRPCConn,
		// we would close a connection that's still in use.
		// TODO: investigate error conditions.
		return nil, err
	}
	c.LROClient = &client.LROClient
	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *tagBindingsGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *tagBindingsGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *tagBindingsGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *tagBindingsGRPCClient) ListTagBindings(ctx context.Context, req *resourcemanagerpb.ListTagBindingsRequest, opts ...gax.CallOption) *TagBindingIterator {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append((*c.CallOptions).ListTagBindings[0:len((*c.CallOptions).ListTagBindings):len((*c.CallOptions).ListTagBindings)], opts...)
	it := &TagBindingIterator{}
	req = proto.Clone(req).(*resourcemanagerpb.ListTagBindingsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*resourcemanagerpb.TagBinding, string, error) {
		resp := &resourcemanagerpb.ListTagBindingsResponse{}
		if pageToken != "" {
			req.PageToken = pageToken
		}
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else if pageSize != 0 {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.tagBindingsClient.ListTagBindings(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetTagBindings(), resp.GetNextPageToken(), nil
	}
	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}

	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.GetPageSize())
	it.pageInfo.Token = req.GetPageToken()

	return it
}

func (c *tagBindingsGRPCClient) CreateTagBinding(ctx context.Context, req *resourcemanagerpb.CreateTagBindingRequest, opts ...gax.CallOption) (*CreateTagBindingOperation, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append((*c.CallOptions).CreateTagBinding[0:len((*c.CallOptions).CreateTagBinding):len((*c.CallOptions).CreateTagBinding)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.tagBindingsClient.CreateTagBinding(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &CreateTagBindingOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

func (c *tagBindingsGRPCClient) DeleteTagBinding(ctx context.Context, req *resourcemanagerpb.DeleteTagBindingRequest, opts ...gax.CallOption) (*DeleteTagBindingOperation, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).DeleteTagBinding[0:len((*c.CallOptions).DeleteTagBinding):len((*c.CallOptions).DeleteTagBinding)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.tagBindingsClient.DeleteTagBinding(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &DeleteTagBindingOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

// CreateTagBindingOperation manages a long-running operation from CreateTagBinding.
type CreateTagBindingOperation struct {
	lro *longrunning.Operation
}

// CreateTagBindingOperation returns a new CreateTagBindingOperation from a given name.
// The name must be that of a previously created CreateTagBindingOperation, possibly from a different process.
func (c *tagBindingsGRPCClient) CreateTagBindingOperation(name string) *CreateTagBindingOperation {
	return &CreateTagBindingOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *CreateTagBindingOperation) Wait(ctx context.Context, opts ...gax.CallOption) (*resourcemanagerpb.TagBinding, error) {
	var resp resourcemanagerpb.TagBinding
	if err := op.lro.WaitWithInterval(ctx, &resp, time.Minute, opts...); err != nil {
		return nil, err
	}
	return &resp, nil
}

// Poll fetches the latest state of the long-running operation.
//
// Poll also fetches the latest metadata, which can be retrieved by Metadata.
//
// If Poll fails, the error is returned and op is unmodified. If Poll succeeds and
// the operation has completed with failure, the error is returned and op.Done will return true.
// If Poll succeeds and the operation has completed successfully,
// op.Done will return true, and the response of the operation is returned.
// If Poll succeeds and the operation has not completed, the returned response and error are both nil.
func (op *CreateTagBindingOperation) Poll(ctx context.Context, opts ...gax.CallOption) (*resourcemanagerpb.TagBinding, error) {
	var resp resourcemanagerpb.TagBinding
	if err := op.lro.Poll(ctx, &resp, opts...); err != nil {
		return nil, err
	}
	if !op.Done() {
		return nil, nil
	}
	return &resp, nil
}

// Metadata returns metadata associated with the long-running operation.
// Metadata itself does not contact the server, but Poll does.
// To get the latest metadata, call this method after a successful call to Poll.
// If the metadata is not available, the returned metadata and error are both nil.
func (op *CreateTagBindingOperation) Metadata() (*resourcemanagerpb.CreateTagBindingMetadata, error) {
	var meta resourcemanagerpb.CreateTagBindingMetadata
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *CreateTagBindingOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *CreateTagBindingOperation) Name() string {
	return op.lro.Name()
}

// DeleteTagBindingOperation manages a long-running operation from DeleteTagBinding.
type DeleteTagBindingOperation struct {
	lro *longrunning.Operation
}

// DeleteTagBindingOperation returns a new DeleteTagBindingOperation from a given name.
// The name must be that of a previously created DeleteTagBindingOperation, possibly from a different process.
func (c *tagBindingsGRPCClient) DeleteTagBindingOperation(name string) *DeleteTagBindingOperation {
	return &DeleteTagBindingOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *DeleteTagBindingOperation) Wait(ctx context.Context, opts ...gax.CallOption) error {
	return op.lro.WaitWithInterval(ctx, nil, time.Minute, opts...)
}

// Poll fetches the latest state of the long-running operation.
//
// Poll also fetches the latest metadata, which can be retrieved by Metadata.
//
// If Poll fails, the error is returned and op is unmodified. If Poll succeeds and
// the operation has completed with failure, the error is returned and op.Done will return true.
// If Poll succeeds and the operation has completed successfully,
// op.Done will return true, and the response of the operation is returned.
// If Poll succeeds and the operation has not completed, the returned response and error are both nil.
func (op *DeleteTagBindingOperation) Poll(ctx context.Context, opts ...gax.CallOption) error {
	return op.lro.Poll(ctx, nil, opts...)
}

// Metadata returns metadata associated with the long-running operation.
// Metadata itself does not contact the server, but Poll does.
// To get the latest metadata, call this method after a successful call to Poll.
// If the metadata is not available, the returned metadata and error are both nil.
func (op *DeleteTagBindingOperation) Metadata() (*resourcemanagerpb.DeleteTagBindingMetadata, error) {
	var meta resourcemanagerpb.DeleteTagBindingMetadata
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *DeleteTagBindingOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *DeleteTagBindingOperation) Name() string {
	return op.lro.Name()
}

// TagBindingIterator manages a stream of *resourcemanagerpb.TagBinding.
type TagBindingIterator struct {
	items    []*resourcemanagerpb.TagBinding
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// Response is the raw response for the current page.
	// It must be cast to the RPC response type.
	// Calling Next() or InternalFetch() updates this value.
	Response interface{}

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*resourcemanagerpb.TagBinding, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *TagBindingIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *TagBindingIterator) Next() (*resourcemanagerpb.TagBinding, error) {
	var item *resourcemanagerpb.TagBinding
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *TagBindingIterator) bufLen() int {
	return len(it.items)
}

func (it *TagBindingIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}
