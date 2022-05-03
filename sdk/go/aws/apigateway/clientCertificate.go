// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package apigateway

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::ApiGateway::ClientCertificate
type ClientCertificate struct {
	pulumi.CustomResourceState

	// The Primary Identifier of the Client Certficate, generated by a Create API Call
	ClientCertificateId pulumi.StringOutput `pulumi:"clientCertificateId"`
	// A description of the client certificate.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// An array of arbitrary tags (key-value pairs) to associate with the client certificate.
	Tags ClientCertificateTagArrayOutput `pulumi:"tags"`
}

// NewClientCertificate registers a new resource with the given unique name, arguments, and options.
func NewClientCertificate(ctx *pulumi.Context,
	name string, args *ClientCertificateArgs, opts ...pulumi.ResourceOption) (*ClientCertificate, error) {
	if args == nil {
		args = &ClientCertificateArgs{}
	}

	var resource ClientCertificate
	err := ctx.RegisterResource("aws-native:apigateway:ClientCertificate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetClientCertificate gets an existing ClientCertificate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetClientCertificate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClientCertificateState, opts ...pulumi.ResourceOption) (*ClientCertificate, error) {
	var resource ClientCertificate
	err := ctx.ReadResource("aws-native:apigateway:ClientCertificate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ClientCertificate resources.
type clientCertificateState struct {
}

type ClientCertificateState struct {
}

func (ClientCertificateState) ElementType() reflect.Type {
	return reflect.TypeOf((*clientCertificateState)(nil)).Elem()
}

type clientCertificateArgs struct {
	// A description of the client certificate.
	Description *string `pulumi:"description"`
	// An array of arbitrary tags (key-value pairs) to associate with the client certificate.
	Tags []ClientCertificateTag `pulumi:"tags"`
}

// The set of arguments for constructing a ClientCertificate resource.
type ClientCertificateArgs struct {
	// A description of the client certificate.
	Description pulumi.StringPtrInput
	// An array of arbitrary tags (key-value pairs) to associate with the client certificate.
	Tags ClientCertificateTagArrayInput
}

func (ClientCertificateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clientCertificateArgs)(nil)).Elem()
}

type ClientCertificateInput interface {
	pulumi.Input

	ToClientCertificateOutput() ClientCertificateOutput
	ToClientCertificateOutputWithContext(ctx context.Context) ClientCertificateOutput
}

func (*ClientCertificate) ElementType() reflect.Type {
	return reflect.TypeOf((**ClientCertificate)(nil)).Elem()
}

func (i *ClientCertificate) ToClientCertificateOutput() ClientCertificateOutput {
	return i.ToClientCertificateOutputWithContext(context.Background())
}

func (i *ClientCertificate) ToClientCertificateOutputWithContext(ctx context.Context) ClientCertificateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientCertificateOutput)
}

type ClientCertificateOutput struct{ *pulumi.OutputState }

func (ClientCertificateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClientCertificate)(nil)).Elem()
}

func (o ClientCertificateOutput) ToClientCertificateOutput() ClientCertificateOutput {
	return o
}

func (o ClientCertificateOutput) ToClientCertificateOutputWithContext(ctx context.Context) ClientCertificateOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ClientCertificateInput)(nil)).Elem(), &ClientCertificate{})
	pulumi.RegisterOutputType(ClientCertificateOutput{})
}
