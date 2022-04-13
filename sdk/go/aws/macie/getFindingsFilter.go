// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package macie

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Macie FindingsFilter resource schema.
func LookupFindingsFilter(ctx *pulumi.Context, args *LookupFindingsFilterArgs, opts ...pulumi.InvokeOption) (*LookupFindingsFilterResult, error) {
	var rv LookupFindingsFilterResult
	err := ctx.Invoke("aws-native:macie:getFindingsFilter", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupFindingsFilterArgs struct {
	// Findings filter ID.
	Id string `pulumi:"id"`
}

type LookupFindingsFilterResult struct {
	// Findings filter action.
	Action *FindingsFilterFindingFilterAction `pulumi:"action"`
	// Findings filter ARN.
	Arn *string `pulumi:"arn"`
	// Findings filter description
	Description *string `pulumi:"description"`
	// Findings filter criteria.
	FindingCriteria *FindingsFilterFindingCriteria `pulumi:"findingCriteria"`
	// Findings filters list.
	FindingsFilterListItems []FindingsFilterListItem `pulumi:"findingsFilterListItems"`
	// Findings filter ID.
	Id *string `pulumi:"id"`
	// Findings filter name
	Name *string `pulumi:"name"`
	// Findings filter position.
	Position *int `pulumi:"position"`
}

func LookupFindingsFilterOutput(ctx *pulumi.Context, args LookupFindingsFilterOutputArgs, opts ...pulumi.InvokeOption) LookupFindingsFilterResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFindingsFilterResult, error) {
			args := v.(LookupFindingsFilterArgs)
			r, err := LookupFindingsFilter(ctx, &args, opts...)
			var s LookupFindingsFilterResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupFindingsFilterResultOutput)
}

type LookupFindingsFilterOutputArgs struct {
	// Findings filter ID.
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupFindingsFilterOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFindingsFilterArgs)(nil)).Elem()
}

type LookupFindingsFilterResultOutput struct{ *pulumi.OutputState }

func (LookupFindingsFilterResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFindingsFilterResult)(nil)).Elem()
}

func (o LookupFindingsFilterResultOutput) ToLookupFindingsFilterResultOutput() LookupFindingsFilterResultOutput {
	return o
}

func (o LookupFindingsFilterResultOutput) ToLookupFindingsFilterResultOutputWithContext(ctx context.Context) LookupFindingsFilterResultOutput {
	return o
}

// Findings filter action.
func (o LookupFindingsFilterResultOutput) Action() FindingsFilterFindingFilterActionPtrOutput {
	return o.ApplyT(func(v LookupFindingsFilterResult) *FindingsFilterFindingFilterAction { return v.Action }).(FindingsFilterFindingFilterActionPtrOutput)
}

// Findings filter ARN.
func (o LookupFindingsFilterResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFindingsFilterResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

// Findings filter description
func (o LookupFindingsFilterResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFindingsFilterResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// Findings filter criteria.
func (o LookupFindingsFilterResultOutput) FindingCriteria() FindingsFilterFindingCriteriaPtrOutput {
	return o.ApplyT(func(v LookupFindingsFilterResult) *FindingsFilterFindingCriteria { return v.FindingCriteria }).(FindingsFilterFindingCriteriaPtrOutput)
}

// Findings filters list.
func (o LookupFindingsFilterResultOutput) FindingsFilterListItems() FindingsFilterListItemArrayOutput {
	return o.ApplyT(func(v LookupFindingsFilterResult) []FindingsFilterListItem { return v.FindingsFilterListItems }).(FindingsFilterListItemArrayOutput)
}

// Findings filter ID.
func (o LookupFindingsFilterResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFindingsFilterResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// Findings filter name
func (o LookupFindingsFilterResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFindingsFilterResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// Findings filter position.
func (o LookupFindingsFilterResultOutput) Position() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupFindingsFilterResult) *int { return v.Position }).(pulumi.IntPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFindingsFilterResultOutput{})
}
