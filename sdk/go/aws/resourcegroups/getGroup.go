// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package resourcegroups

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Schema for ResourceGroups::Group
func LookupGroup(ctx *pulumi.Context, args *LookupGroupArgs, opts ...pulumi.InvokeOption) (*LookupGroupResult, error) {
	var rv LookupGroupResult
	err := ctx.Invoke("aws-native:resourcegroups:getGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupGroupArgs struct {
	// The name of the resource group
	Name string `pulumi:"name"`
}

type LookupGroupResult struct {
	// The Resource Group ARN.
	Arn           *string                  `pulumi:"arn"`
	Configuration []GroupConfigurationItem `pulumi:"configuration"`
	// The description of the resource group
	Description   *string             `pulumi:"description"`
	ResourceQuery *GroupResourceQuery `pulumi:"resourceQuery"`
	Resources     []string            `pulumi:"resources"`
	Tags          []GroupTag          `pulumi:"tags"`
}

func LookupGroupOutput(ctx *pulumi.Context, args LookupGroupOutputArgs, opts ...pulumi.InvokeOption) LookupGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupGroupResult, error) {
			args := v.(LookupGroupArgs)
			r, err := LookupGroup(ctx, &args, opts...)
			return *r, err
		}).(LookupGroupResultOutput)
}

type LookupGroupOutputArgs struct {
	// The name of the resource group
	Name pulumi.StringInput `pulumi:"name"`
}

func (LookupGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGroupArgs)(nil)).Elem()
}

type LookupGroupResultOutput struct{ *pulumi.OutputState }

func (LookupGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGroupResult)(nil)).Elem()
}

func (o LookupGroupResultOutput) ToLookupGroupResultOutput() LookupGroupResultOutput {
	return o
}

func (o LookupGroupResultOutput) ToLookupGroupResultOutputWithContext(ctx context.Context) LookupGroupResultOutput {
	return o
}

// The Resource Group ARN.
func (o LookupGroupResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGroupResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

func (o LookupGroupResultOutput) Configuration() GroupConfigurationItemArrayOutput {
	return o.ApplyT(func(v LookupGroupResult) []GroupConfigurationItem { return v.Configuration }).(GroupConfigurationItemArrayOutput)
}

// The description of the resource group
func (o LookupGroupResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGroupResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupGroupResultOutput) ResourceQuery() GroupResourceQueryPtrOutput {
	return o.ApplyT(func(v LookupGroupResult) *GroupResourceQuery { return v.ResourceQuery }).(GroupResourceQueryPtrOutput)
}

func (o LookupGroupResultOutput) Resources() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupGroupResult) []string { return v.Resources }).(pulumi.StringArrayOutput)
}

func (o LookupGroupResultOutput) Tags() GroupTagArrayOutput {
	return o.ApplyT(func(v LookupGroupResult) []GroupTag { return v.Tags }).(GroupTagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupGroupResultOutput{})
}