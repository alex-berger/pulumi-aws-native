// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cloudformation

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::CloudFormation::CustomResource
func LookupCustomResource(ctx *pulumi.Context, args *LookupCustomResourceArgs, opts ...pulumi.InvokeOption) (*LookupCustomResourceResult, error) {
	var rv LookupCustomResourceResult
	err := ctx.Invoke("aws-native:cloudformation:getCustomResource", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCustomResourceArgs struct {
	Id string `pulumi:"id"`
}

type LookupCustomResourceResult struct {
	Id *string `pulumi:"id"`
}

func LookupCustomResourceOutput(ctx *pulumi.Context, args LookupCustomResourceOutputArgs, opts ...pulumi.InvokeOption) LookupCustomResourceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCustomResourceResult, error) {
			args := v.(LookupCustomResourceArgs)
			r, err := LookupCustomResource(ctx, &args, opts...)
			var s LookupCustomResourceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupCustomResourceResultOutput)
}

type LookupCustomResourceOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupCustomResourceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCustomResourceArgs)(nil)).Elem()
}

type LookupCustomResourceResultOutput struct{ *pulumi.OutputState }

func (LookupCustomResourceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCustomResourceResult)(nil)).Elem()
}

func (o LookupCustomResourceResultOutput) ToLookupCustomResourceResultOutput() LookupCustomResourceResultOutput {
	return o
}

func (o LookupCustomResourceResultOutput) ToLookupCustomResourceResultOutputWithContext(ctx context.Context) LookupCustomResourceResultOutput {
	return o
}

func (o LookupCustomResourceResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCustomResourceResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCustomResourceResultOutput{})
}
