// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package servicediscovery

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::ServiceDiscovery::HttpNamespace
func LookupHttpNamespace(ctx *pulumi.Context, args *LookupHttpNamespaceArgs, opts ...pulumi.InvokeOption) (*LookupHttpNamespaceResult, error) {
	var rv LookupHttpNamespaceResult
	err := ctx.Invoke("aws-native:servicediscovery:getHttpNamespace", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupHttpNamespaceArgs struct {
	Id string `pulumi:"id"`
}

type LookupHttpNamespaceResult struct {
	Arn         *string            `pulumi:"arn"`
	Description *string            `pulumi:"description"`
	Id          *string            `pulumi:"id"`
	Tags        []HttpNamespaceTag `pulumi:"tags"`
}

func LookupHttpNamespaceOutput(ctx *pulumi.Context, args LookupHttpNamespaceOutputArgs, opts ...pulumi.InvokeOption) LookupHttpNamespaceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupHttpNamespaceResult, error) {
			args := v.(LookupHttpNamespaceArgs)
			r, err := LookupHttpNamespace(ctx, &args, opts...)
			var s LookupHttpNamespaceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupHttpNamespaceResultOutput)
}

type LookupHttpNamespaceOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupHttpNamespaceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupHttpNamespaceArgs)(nil)).Elem()
}

type LookupHttpNamespaceResultOutput struct{ *pulumi.OutputState }

func (LookupHttpNamespaceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupHttpNamespaceResult)(nil)).Elem()
}

func (o LookupHttpNamespaceResultOutput) ToLookupHttpNamespaceResultOutput() LookupHttpNamespaceResultOutput {
	return o
}

func (o LookupHttpNamespaceResultOutput) ToLookupHttpNamespaceResultOutputWithContext(ctx context.Context) LookupHttpNamespaceResultOutput {
	return o
}

func (o LookupHttpNamespaceResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupHttpNamespaceResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

func (o LookupHttpNamespaceResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupHttpNamespaceResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupHttpNamespaceResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupHttpNamespaceResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupHttpNamespaceResultOutput) Tags() HttpNamespaceTagArrayOutput {
	return o.ApplyT(func(v LookupHttpNamespaceResult) []HttpNamespaceTag { return v.Tags }).(HttpNamespaceTagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupHttpNamespaceResultOutput{})
}
