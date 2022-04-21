// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package apigatewayv2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::ApiGatewayV2::ApiMapping
func LookupApiMapping(ctx *pulumi.Context, args *LookupApiMappingArgs, opts ...pulumi.InvokeOption) (*LookupApiMappingResult, error) {
	var rv LookupApiMappingResult
	err := ctx.Invoke("aws-native:apigatewayv2:getApiMapping", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupApiMappingArgs struct {
	Id string `pulumi:"id"`
}

type LookupApiMappingResult struct {
	ApiMappingKey *string `pulumi:"apiMappingKey"`
	Id            *string `pulumi:"id"`
	Stage         *string `pulumi:"stage"`
}

func LookupApiMappingOutput(ctx *pulumi.Context, args LookupApiMappingOutputArgs, opts ...pulumi.InvokeOption) LookupApiMappingResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupApiMappingResult, error) {
			args := v.(LookupApiMappingArgs)
			r, err := LookupApiMapping(ctx, &args, opts...)
			var s LookupApiMappingResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupApiMappingResultOutput)
}

type LookupApiMappingOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupApiMappingOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApiMappingArgs)(nil)).Elem()
}

type LookupApiMappingResultOutput struct{ *pulumi.OutputState }

func (LookupApiMappingResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApiMappingResult)(nil)).Elem()
}

func (o LookupApiMappingResultOutput) ToLookupApiMappingResultOutput() LookupApiMappingResultOutput {
	return o
}

func (o LookupApiMappingResultOutput) ToLookupApiMappingResultOutputWithContext(ctx context.Context) LookupApiMappingResultOutput {
	return o
}

func (o LookupApiMappingResultOutput) ApiMappingKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApiMappingResult) *string { return v.ApiMappingKey }).(pulumi.StringPtrOutput)
}

func (o LookupApiMappingResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApiMappingResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupApiMappingResultOutput) Stage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApiMappingResult) *string { return v.Stage }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupApiMappingResultOutput{})
}
