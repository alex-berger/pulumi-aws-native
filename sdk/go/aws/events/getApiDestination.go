// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package events

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Events::ApiDestination.
func LookupApiDestination(ctx *pulumi.Context, args *LookupApiDestinationArgs, opts ...pulumi.InvokeOption) (*LookupApiDestinationResult, error) {
	var rv LookupApiDestinationResult
	err := ctx.Invoke("aws-native:events:getApiDestination", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupApiDestinationArgs struct {
	// Name of the apiDestination.
	Name string `pulumi:"name"`
}

type LookupApiDestinationResult struct {
	// The arn of the api destination.
	Arn *string `pulumi:"arn"`
	// The arn of the connection.
	ConnectionArn *string                   `pulumi:"connectionArn"`
	Description   *string                   `pulumi:"description"`
	HttpMethod    *ApiDestinationHttpMethod `pulumi:"httpMethod"`
	// Url endpoint to invoke.
	InvocationEndpoint           *string `pulumi:"invocationEndpoint"`
	InvocationRateLimitPerSecond *int    `pulumi:"invocationRateLimitPerSecond"`
}

func LookupApiDestinationOutput(ctx *pulumi.Context, args LookupApiDestinationOutputArgs, opts ...pulumi.InvokeOption) LookupApiDestinationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupApiDestinationResult, error) {
			args := v.(LookupApiDestinationArgs)
			r, err := LookupApiDestination(ctx, &args, opts...)
			var s LookupApiDestinationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupApiDestinationResultOutput)
}

type LookupApiDestinationOutputArgs struct {
	// Name of the apiDestination.
	Name pulumi.StringInput `pulumi:"name"`
}

func (LookupApiDestinationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApiDestinationArgs)(nil)).Elem()
}

type LookupApiDestinationResultOutput struct{ *pulumi.OutputState }

func (LookupApiDestinationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApiDestinationResult)(nil)).Elem()
}

func (o LookupApiDestinationResultOutput) ToLookupApiDestinationResultOutput() LookupApiDestinationResultOutput {
	return o
}

func (o LookupApiDestinationResultOutput) ToLookupApiDestinationResultOutputWithContext(ctx context.Context) LookupApiDestinationResultOutput {
	return o
}

// The arn of the api destination.
func (o LookupApiDestinationResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApiDestinationResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

// The arn of the connection.
func (o LookupApiDestinationResultOutput) ConnectionArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApiDestinationResult) *string { return v.ConnectionArn }).(pulumi.StringPtrOutput)
}

func (o LookupApiDestinationResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApiDestinationResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupApiDestinationResultOutput) HttpMethod() ApiDestinationHttpMethodPtrOutput {
	return o.ApplyT(func(v LookupApiDestinationResult) *ApiDestinationHttpMethod { return v.HttpMethod }).(ApiDestinationHttpMethodPtrOutput)
}

// Url endpoint to invoke.
func (o LookupApiDestinationResultOutput) InvocationEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApiDestinationResult) *string { return v.InvocationEndpoint }).(pulumi.StringPtrOutput)
}

func (o LookupApiDestinationResultOutput) InvocationRateLimitPerSecond() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupApiDestinationResult) *int { return v.InvocationRateLimitPerSecond }).(pulumi.IntPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupApiDestinationResultOutput{})
}