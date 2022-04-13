// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package inspectorv2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Inspector Filter resource schema
func LookupFilter(ctx *pulumi.Context, args *LookupFilterArgs, opts ...pulumi.InvokeOption) (*LookupFilterResult, error) {
	var rv LookupFilterResult
	err := ctx.Invoke("aws-native:inspectorv2:getFilter", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupFilterArgs struct {
	// Findings filter ARN.
	Arn string `pulumi:"arn"`
}

type LookupFilterResult struct {
	// Findings filter ARN.
	Arn *string `pulumi:"arn"`
	// Findings filter description.
	Description *string `pulumi:"description"`
	// Findings filter action.
	FilterAction *FilterAction `pulumi:"filterAction"`
	// Findings filter criteria.
	FilterCriteria *FilterCriteria `pulumi:"filterCriteria"`
	// Findings filter name.
	Name *string `pulumi:"name"`
}

func LookupFilterOutput(ctx *pulumi.Context, args LookupFilterOutputArgs, opts ...pulumi.InvokeOption) LookupFilterResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFilterResult, error) {
			args := v.(LookupFilterArgs)
			r, err := LookupFilter(ctx, &args, opts...)
			var s LookupFilterResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupFilterResultOutput)
}

type LookupFilterOutputArgs struct {
	// Findings filter ARN.
	Arn pulumi.StringInput `pulumi:"arn"`
}

func (LookupFilterOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFilterArgs)(nil)).Elem()
}

type LookupFilterResultOutput struct{ *pulumi.OutputState }

func (LookupFilterResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFilterResult)(nil)).Elem()
}

func (o LookupFilterResultOutput) ToLookupFilterResultOutput() LookupFilterResultOutput {
	return o
}

func (o LookupFilterResultOutput) ToLookupFilterResultOutputWithContext(ctx context.Context) LookupFilterResultOutput {
	return o
}

// Findings filter ARN.
func (o LookupFilterResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFilterResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

// Findings filter description.
func (o LookupFilterResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFilterResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// Findings filter action.
func (o LookupFilterResultOutput) FilterAction() FilterActionPtrOutput {
	return o.ApplyT(func(v LookupFilterResult) *FilterAction { return v.FilterAction }).(FilterActionPtrOutput)
}

// Findings filter criteria.
func (o LookupFilterResultOutput) FilterCriteria() FilterCriteriaPtrOutput {
	return o.ApplyT(func(v LookupFilterResult) *FilterCriteria { return v.FilterCriteria }).(FilterCriteriaPtrOutput)
}

// Findings filter name.
func (o LookupFilterResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFilterResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFilterResultOutput{})
}
