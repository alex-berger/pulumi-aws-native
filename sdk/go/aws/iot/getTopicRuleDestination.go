// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iot

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::IoT::TopicRuleDestination
func LookupTopicRuleDestination(ctx *pulumi.Context, args *LookupTopicRuleDestinationArgs, opts ...pulumi.InvokeOption) (*LookupTopicRuleDestinationResult, error) {
	var rv LookupTopicRuleDestinationResult
	err := ctx.Invoke("aws-native:iot:getTopicRuleDestination", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTopicRuleDestinationArgs struct {
	// Amazon Resource Name (ARN).
	Arn string `pulumi:"arn"`
}

type LookupTopicRuleDestinationResult struct {
	// Amazon Resource Name (ARN).
	Arn *string `pulumi:"arn"`
	// The status of the TopicRuleDestination.
	Status *TopicRuleDestinationStatus `pulumi:"status"`
	// The reasoning for the current status of the TopicRuleDestination.
	StatusReason *string `pulumi:"statusReason"`
}

func LookupTopicRuleDestinationOutput(ctx *pulumi.Context, args LookupTopicRuleDestinationOutputArgs, opts ...pulumi.InvokeOption) LookupTopicRuleDestinationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupTopicRuleDestinationResult, error) {
			args := v.(LookupTopicRuleDestinationArgs)
			r, err := LookupTopicRuleDestination(ctx, &args, opts...)
			var s LookupTopicRuleDestinationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupTopicRuleDestinationResultOutput)
}

type LookupTopicRuleDestinationOutputArgs struct {
	// Amazon Resource Name (ARN).
	Arn pulumi.StringInput `pulumi:"arn"`
}

func (LookupTopicRuleDestinationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTopicRuleDestinationArgs)(nil)).Elem()
}

type LookupTopicRuleDestinationResultOutput struct{ *pulumi.OutputState }

func (LookupTopicRuleDestinationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTopicRuleDestinationResult)(nil)).Elem()
}

func (o LookupTopicRuleDestinationResultOutput) ToLookupTopicRuleDestinationResultOutput() LookupTopicRuleDestinationResultOutput {
	return o
}

func (o LookupTopicRuleDestinationResultOutput) ToLookupTopicRuleDestinationResultOutputWithContext(ctx context.Context) LookupTopicRuleDestinationResultOutput {
	return o
}

// Amazon Resource Name (ARN).
func (o LookupTopicRuleDestinationResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTopicRuleDestinationResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

// The status of the TopicRuleDestination.
func (o LookupTopicRuleDestinationResultOutput) Status() TopicRuleDestinationStatusPtrOutput {
	return o.ApplyT(func(v LookupTopicRuleDestinationResult) *TopicRuleDestinationStatus { return v.Status }).(TopicRuleDestinationStatusPtrOutput)
}

// The reasoning for the current status of the TopicRuleDestination.
func (o LookupTopicRuleDestinationResultOutput) StatusReason() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTopicRuleDestinationResult) *string { return v.StatusReason }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTopicRuleDestinationResultOutput{})
}
