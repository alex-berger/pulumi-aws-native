// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package guardduty

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::GuardDuty::ThreatIntelSet
func LookupThreatIntelSet(ctx *pulumi.Context, args *LookupThreatIntelSetArgs, opts ...pulumi.InvokeOption) (*LookupThreatIntelSetResult, error) {
	var rv LookupThreatIntelSetResult
	err := ctx.Invoke("aws-native:guardduty:getThreatIntelSet", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupThreatIntelSetArgs struct {
	Id string `pulumi:"id"`
}

type LookupThreatIntelSetResult struct {
	Activate *bool   `pulumi:"activate"`
	Id       *string `pulumi:"id"`
	Location *string `pulumi:"location"`
	Name     *string `pulumi:"name"`
}

func LookupThreatIntelSetOutput(ctx *pulumi.Context, args LookupThreatIntelSetOutputArgs, opts ...pulumi.InvokeOption) LookupThreatIntelSetResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupThreatIntelSetResult, error) {
			args := v.(LookupThreatIntelSetArgs)
			r, err := LookupThreatIntelSet(ctx, &args, opts...)
			var s LookupThreatIntelSetResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupThreatIntelSetResultOutput)
}

type LookupThreatIntelSetOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupThreatIntelSetOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupThreatIntelSetArgs)(nil)).Elem()
}

type LookupThreatIntelSetResultOutput struct{ *pulumi.OutputState }

func (LookupThreatIntelSetResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupThreatIntelSetResult)(nil)).Elem()
}

func (o LookupThreatIntelSetResultOutput) ToLookupThreatIntelSetResultOutput() LookupThreatIntelSetResultOutput {
	return o
}

func (o LookupThreatIntelSetResultOutput) ToLookupThreatIntelSetResultOutputWithContext(ctx context.Context) LookupThreatIntelSetResultOutput {
	return o
}

func (o LookupThreatIntelSetResultOutput) Activate() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupThreatIntelSetResult) *bool { return v.Activate }).(pulumi.BoolPtrOutput)
}

func (o LookupThreatIntelSetResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupThreatIntelSetResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupThreatIntelSetResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupThreatIntelSetResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupThreatIntelSetResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupThreatIntelSetResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupThreatIntelSetResultOutput{})
}
