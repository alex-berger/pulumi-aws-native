// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package greengrass

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Greengrass::DeviceDefinitionVersion
func LookupDeviceDefinitionVersion(ctx *pulumi.Context, args *LookupDeviceDefinitionVersionArgs, opts ...pulumi.InvokeOption) (*LookupDeviceDefinitionVersionResult, error) {
	var rv LookupDeviceDefinitionVersionResult
	err := ctx.Invoke("aws-native:greengrass:getDeviceDefinitionVersion", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDeviceDefinitionVersionArgs struct {
	Id string `pulumi:"id"`
}

type LookupDeviceDefinitionVersionResult struct {
	Id *string `pulumi:"id"`
}

func LookupDeviceDefinitionVersionOutput(ctx *pulumi.Context, args LookupDeviceDefinitionVersionOutputArgs, opts ...pulumi.InvokeOption) LookupDeviceDefinitionVersionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDeviceDefinitionVersionResult, error) {
			args := v.(LookupDeviceDefinitionVersionArgs)
			r, err := LookupDeviceDefinitionVersion(ctx, &args, opts...)
			var s LookupDeviceDefinitionVersionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDeviceDefinitionVersionResultOutput)
}

type LookupDeviceDefinitionVersionOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupDeviceDefinitionVersionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDeviceDefinitionVersionArgs)(nil)).Elem()
}

type LookupDeviceDefinitionVersionResultOutput struct{ *pulumi.OutputState }

func (LookupDeviceDefinitionVersionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDeviceDefinitionVersionResult)(nil)).Elem()
}

func (o LookupDeviceDefinitionVersionResultOutput) ToLookupDeviceDefinitionVersionResultOutput() LookupDeviceDefinitionVersionResultOutput {
	return o
}

func (o LookupDeviceDefinitionVersionResultOutput) ToLookupDeviceDefinitionVersionResultOutputWithContext(ctx context.Context) LookupDeviceDefinitionVersionResultOutput {
	return o
}

func (o LookupDeviceDefinitionVersionResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDeviceDefinitionVersionResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDeviceDefinitionVersionResultOutput{})
}
