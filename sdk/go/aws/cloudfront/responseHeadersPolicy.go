// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cloudfront

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::CloudFront::ResponseHeadersPolicy
type ResponseHeadersPolicy struct {
	pulumi.CustomResourceState

	LastModifiedTime            pulumi.StringOutput               `pulumi:"lastModifiedTime"`
	ResponseHeadersPolicyConfig ResponseHeadersPolicyConfigOutput `pulumi:"responseHeadersPolicyConfig"`
}

// NewResponseHeadersPolicy registers a new resource with the given unique name, arguments, and options.
func NewResponseHeadersPolicy(ctx *pulumi.Context,
	name string, args *ResponseHeadersPolicyArgs, opts ...pulumi.ResourceOption) (*ResponseHeadersPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResponseHeadersPolicyConfig == nil {
		return nil, errors.New("invalid value for required argument 'ResponseHeadersPolicyConfig'")
	}
	var resource ResponseHeadersPolicy
	err := ctx.RegisterResource("aws-native:cloudfront:ResponseHeadersPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetResponseHeadersPolicy gets an existing ResponseHeadersPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetResponseHeadersPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ResponseHeadersPolicyState, opts ...pulumi.ResourceOption) (*ResponseHeadersPolicy, error) {
	var resource ResponseHeadersPolicy
	err := ctx.ReadResource("aws-native:cloudfront:ResponseHeadersPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ResponseHeadersPolicy resources.
type responseHeadersPolicyState struct {
}

type ResponseHeadersPolicyState struct {
}

func (ResponseHeadersPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*responseHeadersPolicyState)(nil)).Elem()
}

type responseHeadersPolicyArgs struct {
	ResponseHeadersPolicyConfig ResponseHeadersPolicyConfig `pulumi:"responseHeadersPolicyConfig"`
}

// The set of arguments for constructing a ResponseHeadersPolicy resource.
type ResponseHeadersPolicyArgs struct {
	ResponseHeadersPolicyConfig ResponseHeadersPolicyConfigInput
}

func (ResponseHeadersPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*responseHeadersPolicyArgs)(nil)).Elem()
}

type ResponseHeadersPolicyInput interface {
	pulumi.Input

	ToResponseHeadersPolicyOutput() ResponseHeadersPolicyOutput
	ToResponseHeadersPolicyOutputWithContext(ctx context.Context) ResponseHeadersPolicyOutput
}

func (*ResponseHeadersPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**ResponseHeadersPolicy)(nil)).Elem()
}

func (i *ResponseHeadersPolicy) ToResponseHeadersPolicyOutput() ResponseHeadersPolicyOutput {
	return i.ToResponseHeadersPolicyOutputWithContext(context.Background())
}

func (i *ResponseHeadersPolicy) ToResponseHeadersPolicyOutputWithContext(ctx context.Context) ResponseHeadersPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResponseHeadersPolicyOutput)
}

type ResponseHeadersPolicyOutput struct{ *pulumi.OutputState }

func (ResponseHeadersPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResponseHeadersPolicy)(nil)).Elem()
}

func (o ResponseHeadersPolicyOutput) ToResponseHeadersPolicyOutput() ResponseHeadersPolicyOutput {
	return o
}

func (o ResponseHeadersPolicyOutput) ToResponseHeadersPolicyOutputWithContext(ctx context.Context) ResponseHeadersPolicyOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ResponseHeadersPolicyInput)(nil)).Elem(), &ResponseHeadersPolicy{})
	pulumi.RegisterOutputType(ResponseHeadersPolicyOutput{})
}
