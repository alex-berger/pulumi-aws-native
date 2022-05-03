// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package events

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Events::EventBusPolicy
func LookupEventBusPolicy(ctx *pulumi.Context, args *LookupEventBusPolicyArgs, opts ...pulumi.InvokeOption) (*LookupEventBusPolicyResult, error) {
	var rv LookupEventBusPolicyResult
	err := ctx.Invoke("aws-native:events:getEventBusPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupEventBusPolicyArgs struct {
	Id string `pulumi:"id"`
}

type LookupEventBusPolicyResult struct {
	Action    *string                  `pulumi:"action"`
	Condition *EventBusPolicyCondition `pulumi:"condition"`
	Id        *string                  `pulumi:"id"`
	Principal *string                  `pulumi:"principal"`
	Statement interface{}              `pulumi:"statement"`
}

func LookupEventBusPolicyOutput(ctx *pulumi.Context, args LookupEventBusPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupEventBusPolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupEventBusPolicyResult, error) {
			args := v.(LookupEventBusPolicyArgs)
			r, err := LookupEventBusPolicy(ctx, &args, opts...)
			var s LookupEventBusPolicyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupEventBusPolicyResultOutput)
}

type LookupEventBusPolicyOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupEventBusPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEventBusPolicyArgs)(nil)).Elem()
}

type LookupEventBusPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupEventBusPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEventBusPolicyResult)(nil)).Elem()
}

func (o LookupEventBusPolicyResultOutput) ToLookupEventBusPolicyResultOutput() LookupEventBusPolicyResultOutput {
	return o
}

func (o LookupEventBusPolicyResultOutput) ToLookupEventBusPolicyResultOutputWithContext(ctx context.Context) LookupEventBusPolicyResultOutput {
	return o
}

func (o LookupEventBusPolicyResultOutput) Action() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEventBusPolicyResult) *string { return v.Action }).(pulumi.StringPtrOutput)
}

func (o LookupEventBusPolicyResultOutput) Condition() EventBusPolicyConditionPtrOutput {
	return o.ApplyT(func(v LookupEventBusPolicyResult) *EventBusPolicyCondition { return v.Condition }).(EventBusPolicyConditionPtrOutput)
}

func (o LookupEventBusPolicyResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEventBusPolicyResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupEventBusPolicyResultOutput) Principal() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEventBusPolicyResult) *string { return v.Principal }).(pulumi.StringPtrOutput)
}

func (o LookupEventBusPolicyResultOutput) Statement() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupEventBusPolicyResult) interface{} { return v.Statement }).(pulumi.AnyOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupEventBusPolicyResultOutput{})
}
