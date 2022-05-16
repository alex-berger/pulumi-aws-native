// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package apigateway

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::ApiGateway::UsagePlanKey
func LookupUsagePlanKey(ctx *pulumi.Context, args *LookupUsagePlanKeyArgs, opts ...pulumi.InvokeOption) (*LookupUsagePlanKeyResult, error) {
	var rv LookupUsagePlanKeyResult
	err := ctx.Invoke("aws-native:apigateway:getUsagePlanKey", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupUsagePlanKeyArgs struct {
	// An autogenerated ID which is a combination of the ID of the key and ID of the usage plan combined with a : such as 123abcdef:abc123.
	Id string `pulumi:"id"`
}

type LookupUsagePlanKeyResult struct {
	// An autogenerated ID which is a combination of the ID of the key and ID of the usage plan combined with a : such as 123abcdef:abc123.
	Id *string `pulumi:"id"`
}

func LookupUsagePlanKeyOutput(ctx *pulumi.Context, args LookupUsagePlanKeyOutputArgs, opts ...pulumi.InvokeOption) LookupUsagePlanKeyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupUsagePlanKeyResult, error) {
			args := v.(LookupUsagePlanKeyArgs)
			r, err := LookupUsagePlanKey(ctx, &args, opts...)
			var s LookupUsagePlanKeyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupUsagePlanKeyResultOutput)
}

type LookupUsagePlanKeyOutputArgs struct {
	// An autogenerated ID which is a combination of the ID of the key and ID of the usage plan combined with a : such as 123abcdef:abc123.
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupUsagePlanKeyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupUsagePlanKeyArgs)(nil)).Elem()
}

type LookupUsagePlanKeyResultOutput struct{ *pulumi.OutputState }

func (LookupUsagePlanKeyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupUsagePlanKeyResult)(nil)).Elem()
}

func (o LookupUsagePlanKeyResultOutput) ToLookupUsagePlanKeyResultOutput() LookupUsagePlanKeyResultOutput {
	return o
}

func (o LookupUsagePlanKeyResultOutput) ToLookupUsagePlanKeyResultOutputWithContext(ctx context.Context) LookupUsagePlanKeyResultOutput {
	return o
}

// An autogenerated ID which is a combination of the ID of the key and ID of the usage plan combined with a : such as 123abcdef:abc123.
func (o LookupUsagePlanKeyResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUsagePlanKeyResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupUsagePlanKeyResultOutput{})
}