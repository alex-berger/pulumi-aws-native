// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package logs

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Logs::SubscriptionFilter
func LookupSubscriptionFilter(ctx *pulumi.Context, args *LookupSubscriptionFilterArgs, opts ...pulumi.InvokeOption) (*LookupSubscriptionFilterResult, error) {
	var rv LookupSubscriptionFilterResult
	err := ctx.Invoke("aws-native:logs:getSubscriptionFilter", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSubscriptionFilterArgs struct {
	Id string `pulumi:"id"`
}

type LookupSubscriptionFilterResult struct {
	Id *string `pulumi:"id"`
}

func LookupSubscriptionFilterOutput(ctx *pulumi.Context, args LookupSubscriptionFilterOutputArgs, opts ...pulumi.InvokeOption) LookupSubscriptionFilterResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSubscriptionFilterResult, error) {
			args := v.(LookupSubscriptionFilterArgs)
			r, err := LookupSubscriptionFilter(ctx, &args, opts...)
			var s LookupSubscriptionFilterResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSubscriptionFilterResultOutput)
}

type LookupSubscriptionFilterOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupSubscriptionFilterOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSubscriptionFilterArgs)(nil)).Elem()
}

type LookupSubscriptionFilterResultOutput struct{ *pulumi.OutputState }

func (LookupSubscriptionFilterResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSubscriptionFilterResult)(nil)).Elem()
}

func (o LookupSubscriptionFilterResultOutput) ToLookupSubscriptionFilterResultOutput() LookupSubscriptionFilterResultOutput {
	return o
}

func (o LookupSubscriptionFilterResultOutput) ToLookupSubscriptionFilterResultOutputWithContext(ctx context.Context) LookupSubscriptionFilterResultOutput {
	return o
}

func (o LookupSubscriptionFilterResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSubscriptionFilterResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSubscriptionFilterResultOutput{})
}
