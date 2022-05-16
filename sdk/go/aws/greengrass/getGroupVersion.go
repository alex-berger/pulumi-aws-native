// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package greengrass

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Greengrass::GroupVersion
func LookupGroupVersion(ctx *pulumi.Context, args *LookupGroupVersionArgs, opts ...pulumi.InvokeOption) (*LookupGroupVersionResult, error) {
	var rv LookupGroupVersionResult
	err := ctx.Invoke("aws-native:greengrass:getGroupVersion", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupGroupVersionArgs struct {
	Id string `pulumi:"id"`
}

type LookupGroupVersionResult struct {
	Id *string `pulumi:"id"`
}

func LookupGroupVersionOutput(ctx *pulumi.Context, args LookupGroupVersionOutputArgs, opts ...pulumi.InvokeOption) LookupGroupVersionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupGroupVersionResult, error) {
			args := v.(LookupGroupVersionArgs)
			r, err := LookupGroupVersion(ctx, &args, opts...)
			var s LookupGroupVersionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupGroupVersionResultOutput)
}

type LookupGroupVersionOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupGroupVersionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGroupVersionArgs)(nil)).Elem()
}

type LookupGroupVersionResultOutput struct{ *pulumi.OutputState }

func (LookupGroupVersionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGroupVersionResult)(nil)).Elem()
}

func (o LookupGroupVersionResultOutput) ToLookupGroupVersionResultOutput() LookupGroupVersionResultOutput {
	return o
}

func (o LookupGroupVersionResultOutput) ToLookupGroupVersionResultOutputWithContext(ctx context.Context) LookupGroupVersionResultOutput {
	return o
}

func (o LookupGroupVersionResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGroupVersionResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupGroupVersionResultOutput{})
}