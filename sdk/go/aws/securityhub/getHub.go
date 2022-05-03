// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package securityhub

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::SecurityHub::Hub
func LookupHub(ctx *pulumi.Context, args *LookupHubArgs, opts ...pulumi.InvokeOption) (*LookupHubResult, error) {
	var rv LookupHubResult
	err := ctx.Invoke("aws-native:securityhub:getHub", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupHubArgs struct {
	Id string `pulumi:"id"`
}

type LookupHubResult struct {
	Id   *string     `pulumi:"id"`
	Tags interface{} `pulumi:"tags"`
}

func LookupHubOutput(ctx *pulumi.Context, args LookupHubOutputArgs, opts ...pulumi.InvokeOption) LookupHubResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupHubResult, error) {
			args := v.(LookupHubArgs)
			r, err := LookupHub(ctx, &args, opts...)
			var s LookupHubResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupHubResultOutput)
}

type LookupHubOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupHubOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupHubArgs)(nil)).Elem()
}

type LookupHubResultOutput struct{ *pulumi.OutputState }

func (LookupHubResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupHubResult)(nil)).Elem()
}

func (o LookupHubResultOutput) ToLookupHubResultOutput() LookupHubResultOutput {
	return o
}

func (o LookupHubResultOutput) ToLookupHubResultOutputWithContext(ctx context.Context) LookupHubResultOutput {
	return o
}

func (o LookupHubResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupHubResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupHubResultOutput) Tags() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupHubResult) interface{} { return v.Tags }).(pulumi.AnyOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupHubResultOutput{})
}
