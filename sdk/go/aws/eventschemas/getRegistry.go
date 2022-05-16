// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package eventschemas

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::EventSchemas::Registry
func LookupRegistry(ctx *pulumi.Context, args *LookupRegistryArgs, opts ...pulumi.InvokeOption) (*LookupRegistryResult, error) {
	var rv LookupRegistryResult
	err := ctx.Invoke("aws-native:eventschemas:getRegistry", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRegistryArgs struct {
	Id string `pulumi:"id"`
}

type LookupRegistryResult struct {
	Description *string             `pulumi:"description"`
	Id          *string             `pulumi:"id"`
	RegistryArn *string             `pulumi:"registryArn"`
	Tags        []RegistryTagsEntry `pulumi:"tags"`
}

func LookupRegistryOutput(ctx *pulumi.Context, args LookupRegistryOutputArgs, opts ...pulumi.InvokeOption) LookupRegistryResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRegistryResult, error) {
			args := v.(LookupRegistryArgs)
			r, err := LookupRegistry(ctx, &args, opts...)
			var s LookupRegistryResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRegistryResultOutput)
}

type LookupRegistryOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupRegistryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRegistryArgs)(nil)).Elem()
}

type LookupRegistryResultOutput struct{ *pulumi.OutputState }

func (LookupRegistryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRegistryResult)(nil)).Elem()
}

func (o LookupRegistryResultOutput) ToLookupRegistryResultOutput() LookupRegistryResultOutput {
	return o
}

func (o LookupRegistryResultOutput) ToLookupRegistryResultOutputWithContext(ctx context.Context) LookupRegistryResultOutput {
	return o
}

func (o LookupRegistryResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRegistryResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupRegistryResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRegistryResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupRegistryResultOutput) RegistryArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRegistryResult) *string { return v.RegistryArn }).(pulumi.StringPtrOutput)
}

func (o LookupRegistryResultOutput) Tags() RegistryTagsEntryArrayOutput {
	return o.ApplyT(func(v LookupRegistryResult) []RegistryTagsEntry { return v.Tags }).(RegistryTagsEntryArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRegistryResultOutput{})
}
