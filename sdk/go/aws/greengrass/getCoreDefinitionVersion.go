// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package greengrass

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Greengrass::CoreDefinitionVersion
func LookupCoreDefinitionVersion(ctx *pulumi.Context, args *LookupCoreDefinitionVersionArgs, opts ...pulumi.InvokeOption) (*LookupCoreDefinitionVersionResult, error) {
	var rv LookupCoreDefinitionVersionResult
	err := ctx.Invoke("aws-native:greengrass:getCoreDefinitionVersion", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCoreDefinitionVersionArgs struct {
	Id string `pulumi:"id"`
}

type LookupCoreDefinitionVersionResult struct {
	Id *string `pulumi:"id"`
}

func LookupCoreDefinitionVersionOutput(ctx *pulumi.Context, args LookupCoreDefinitionVersionOutputArgs, opts ...pulumi.InvokeOption) LookupCoreDefinitionVersionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCoreDefinitionVersionResult, error) {
			args := v.(LookupCoreDefinitionVersionArgs)
			r, err := LookupCoreDefinitionVersion(ctx, &args, opts...)
			var s LookupCoreDefinitionVersionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupCoreDefinitionVersionResultOutput)
}

type LookupCoreDefinitionVersionOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupCoreDefinitionVersionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCoreDefinitionVersionArgs)(nil)).Elem()
}

type LookupCoreDefinitionVersionResultOutput struct{ *pulumi.OutputState }

func (LookupCoreDefinitionVersionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCoreDefinitionVersionResult)(nil)).Elem()
}

func (o LookupCoreDefinitionVersionResultOutput) ToLookupCoreDefinitionVersionResultOutput() LookupCoreDefinitionVersionResultOutput {
	return o
}

func (o LookupCoreDefinitionVersionResultOutput) ToLookupCoreDefinitionVersionResultOutputWithContext(ctx context.Context) LookupCoreDefinitionVersionResultOutput {
	return o
}

func (o LookupCoreDefinitionVersionResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCoreDefinitionVersionResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCoreDefinitionVersionResultOutput{})
}
