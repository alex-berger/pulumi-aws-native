// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package apigateway

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::ApiGateway::BasePathMapping
func LookupBasePathMapping(ctx *pulumi.Context, args *LookupBasePathMappingArgs, opts ...pulumi.InvokeOption) (*LookupBasePathMappingResult, error) {
	var rv LookupBasePathMappingResult
	err := ctx.Invoke("aws-native:apigateway:getBasePathMapping", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupBasePathMappingArgs struct {
	// The base path name that callers of the API must provide in the URL after the domain name.
	BasePath string `pulumi:"basePath"`
	// The DomainName of an AWS::ApiGateway::DomainName resource.
	DomainName string `pulumi:"domainName"`
}

type LookupBasePathMappingResult struct {
	Id *string `pulumi:"id"`
	// The ID of the API.
	RestApiId *string `pulumi:"restApiId"`
	// The name of the API's stage.
	Stage *string `pulumi:"stage"`
}

func LookupBasePathMappingOutput(ctx *pulumi.Context, args LookupBasePathMappingOutputArgs, opts ...pulumi.InvokeOption) LookupBasePathMappingResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupBasePathMappingResult, error) {
			args := v.(LookupBasePathMappingArgs)
			r, err := LookupBasePathMapping(ctx, &args, opts...)
			var s LookupBasePathMappingResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupBasePathMappingResultOutput)
}

type LookupBasePathMappingOutputArgs struct {
	// The base path name that callers of the API must provide in the URL after the domain name.
	BasePath pulumi.StringInput `pulumi:"basePath"`
	// The DomainName of an AWS::ApiGateway::DomainName resource.
	DomainName pulumi.StringInput `pulumi:"domainName"`
}

func (LookupBasePathMappingOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBasePathMappingArgs)(nil)).Elem()
}

type LookupBasePathMappingResultOutput struct{ *pulumi.OutputState }

func (LookupBasePathMappingResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBasePathMappingResult)(nil)).Elem()
}

func (o LookupBasePathMappingResultOutput) ToLookupBasePathMappingResultOutput() LookupBasePathMappingResultOutput {
	return o
}

func (o LookupBasePathMappingResultOutput) ToLookupBasePathMappingResultOutputWithContext(ctx context.Context) LookupBasePathMappingResultOutput {
	return o
}

func (o LookupBasePathMappingResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBasePathMappingResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// The ID of the API.
func (o LookupBasePathMappingResultOutput) RestApiId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBasePathMappingResult) *string { return v.RestApiId }).(pulumi.StringPtrOutput)
}

// The name of the API's stage.
func (o LookupBasePathMappingResultOutput) Stage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBasePathMappingResult) *string { return v.Stage }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupBasePathMappingResultOutput{})
}
