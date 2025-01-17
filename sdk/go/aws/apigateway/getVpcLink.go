// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package apigateway

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::ApiGateway::VpcLink
func LookupVpcLink(ctx *pulumi.Context, args *LookupVpcLinkArgs, opts ...pulumi.InvokeOption) (*LookupVpcLinkResult, error) {
	var rv LookupVpcLinkResult
	err := ctx.Invoke("aws-native:apigateway:getVpcLink", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVpcLinkArgs struct {
	Id string `pulumi:"id"`
}

type LookupVpcLinkResult struct {
	Description *string      `pulumi:"description"`
	Id          *string      `pulumi:"id"`
	Name        *string      `pulumi:"name"`
	Tags        []VpcLinkTag `pulumi:"tags"`
}

func LookupVpcLinkOutput(ctx *pulumi.Context, args LookupVpcLinkOutputArgs, opts ...pulumi.InvokeOption) LookupVpcLinkResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVpcLinkResult, error) {
			args := v.(LookupVpcLinkArgs)
			r, err := LookupVpcLink(ctx, &args, opts...)
			var s LookupVpcLinkResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVpcLinkResultOutput)
}

type LookupVpcLinkOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupVpcLinkOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVpcLinkArgs)(nil)).Elem()
}

type LookupVpcLinkResultOutput struct{ *pulumi.OutputState }

func (LookupVpcLinkResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVpcLinkResult)(nil)).Elem()
}

func (o LookupVpcLinkResultOutput) ToLookupVpcLinkResultOutput() LookupVpcLinkResultOutput {
	return o
}

func (o LookupVpcLinkResultOutput) ToLookupVpcLinkResultOutputWithContext(ctx context.Context) LookupVpcLinkResultOutput {
	return o
}

func (o LookupVpcLinkResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVpcLinkResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupVpcLinkResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVpcLinkResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupVpcLinkResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVpcLinkResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupVpcLinkResultOutput) Tags() VpcLinkTagArrayOutput {
	return o.ApplyT(func(v LookupVpcLinkResult) []VpcLinkTag { return v.Tags }).(VpcLinkTagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVpcLinkResultOutput{})
}
