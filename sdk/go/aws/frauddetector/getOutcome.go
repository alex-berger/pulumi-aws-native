// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package frauddetector

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// An outcome for rule evaluation.
func LookupOutcome(ctx *pulumi.Context, args *LookupOutcomeArgs, opts ...pulumi.InvokeOption) (*LookupOutcomeResult, error) {
	var rv LookupOutcomeResult
	err := ctx.Invoke("aws-native:frauddetector:getOutcome", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupOutcomeArgs struct {
	// The outcome ARN.
	Arn string `pulumi:"arn"`
}

type LookupOutcomeResult struct {
	// The outcome ARN.
	Arn *string `pulumi:"arn"`
	// The timestamp when the outcome was created.
	CreatedTime *string `pulumi:"createdTime"`
	// The outcome description.
	Description *string `pulumi:"description"`
	// The timestamp when the outcome was last updated.
	LastUpdatedTime *string `pulumi:"lastUpdatedTime"`
	// Tags associated with this outcome.
	Tags []OutcomeTag `pulumi:"tags"`
}

func LookupOutcomeOutput(ctx *pulumi.Context, args LookupOutcomeOutputArgs, opts ...pulumi.InvokeOption) LookupOutcomeResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupOutcomeResult, error) {
			args := v.(LookupOutcomeArgs)
			r, err := LookupOutcome(ctx, &args, opts...)
			var s LookupOutcomeResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupOutcomeResultOutput)
}

type LookupOutcomeOutputArgs struct {
	// The outcome ARN.
	Arn pulumi.StringInput `pulumi:"arn"`
}

func (LookupOutcomeOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOutcomeArgs)(nil)).Elem()
}

type LookupOutcomeResultOutput struct{ *pulumi.OutputState }

func (LookupOutcomeResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOutcomeResult)(nil)).Elem()
}

func (o LookupOutcomeResultOutput) ToLookupOutcomeResultOutput() LookupOutcomeResultOutput {
	return o
}

func (o LookupOutcomeResultOutput) ToLookupOutcomeResultOutputWithContext(ctx context.Context) LookupOutcomeResultOutput {
	return o
}

// The outcome ARN.
func (o LookupOutcomeResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupOutcomeResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

// The timestamp when the outcome was created.
func (o LookupOutcomeResultOutput) CreatedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupOutcomeResult) *string { return v.CreatedTime }).(pulumi.StringPtrOutput)
}

// The outcome description.
func (o LookupOutcomeResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupOutcomeResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// The timestamp when the outcome was last updated.
func (o LookupOutcomeResultOutput) LastUpdatedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupOutcomeResult) *string { return v.LastUpdatedTime }).(pulumi.StringPtrOutput)
}

// Tags associated with this outcome.
func (o LookupOutcomeResultOutput) Tags() OutcomeTagArrayOutput {
	return o.ApplyT(func(v LookupOutcomeResult) []OutcomeTag { return v.Tags }).(OutcomeTagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupOutcomeResultOutput{})
}
