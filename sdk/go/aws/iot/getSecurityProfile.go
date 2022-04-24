// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package iot

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A security profile defines a set of expected behaviors for devices in your account.
func LookupSecurityProfile(ctx *pulumi.Context, args *LookupSecurityProfileArgs, opts ...pulumi.InvokeOption) (*LookupSecurityProfileResult, error) {
	var rv LookupSecurityProfileResult
	err := ctx.Invoke("aws-native:iot:getSecurityProfile", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSecurityProfileArgs struct {
	// A unique identifier for the security profile.
	SecurityProfileName string `pulumi:"securityProfileName"`
}

type LookupSecurityProfileResult struct {
	// A list of metrics whose data is retained (stored). By default, data is retained for any metric used in the profile's behaviors, but it is also retained for any metric specified here.
	AdditionalMetricsToRetainV2 []SecurityProfileMetricToRetain `pulumi:"additionalMetricsToRetainV2"`
	// Specifies the destinations to which alerts are sent.
	AlertTargets interface{} `pulumi:"alertTargets"`
	// Specifies the behaviors that, when violated by a device (thing), cause an alert.
	Behaviors []SecurityProfileBehavior `pulumi:"behaviors"`
	// The ARN (Amazon resource name) of the created security profile.
	SecurityProfileArn *string `pulumi:"securityProfileArn"`
	// A description of the security profile.
	SecurityProfileDescription *string `pulumi:"securityProfileDescription"`
	// Metadata that can be used to manage the security profile.
	Tags []SecurityProfileTag `pulumi:"tags"`
	// A set of target ARNs that the security profile is attached to.
	TargetArns []string `pulumi:"targetArns"`
}

func LookupSecurityProfileOutput(ctx *pulumi.Context, args LookupSecurityProfileOutputArgs, opts ...pulumi.InvokeOption) LookupSecurityProfileResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSecurityProfileResult, error) {
			args := v.(LookupSecurityProfileArgs)
			r, err := LookupSecurityProfile(ctx, &args, opts...)
			var s LookupSecurityProfileResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSecurityProfileResultOutput)
}

type LookupSecurityProfileOutputArgs struct {
	// A unique identifier for the security profile.
	SecurityProfileName pulumi.StringInput `pulumi:"securityProfileName"`
}

func (LookupSecurityProfileOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSecurityProfileArgs)(nil)).Elem()
}

type LookupSecurityProfileResultOutput struct{ *pulumi.OutputState }

func (LookupSecurityProfileResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSecurityProfileResult)(nil)).Elem()
}

func (o LookupSecurityProfileResultOutput) ToLookupSecurityProfileResultOutput() LookupSecurityProfileResultOutput {
	return o
}

func (o LookupSecurityProfileResultOutput) ToLookupSecurityProfileResultOutputWithContext(ctx context.Context) LookupSecurityProfileResultOutput {
	return o
}

// A list of metrics whose data is retained (stored). By default, data is retained for any metric used in the profile's behaviors, but it is also retained for any metric specified here.
func (o LookupSecurityProfileResultOutput) AdditionalMetricsToRetainV2() SecurityProfileMetricToRetainArrayOutput {
	return o.ApplyT(func(v LookupSecurityProfileResult) []SecurityProfileMetricToRetain {
		return v.AdditionalMetricsToRetainV2
	}).(SecurityProfileMetricToRetainArrayOutput)
}

// Specifies the destinations to which alerts are sent.
func (o LookupSecurityProfileResultOutput) AlertTargets() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupSecurityProfileResult) interface{} { return v.AlertTargets }).(pulumi.AnyOutput)
}

// Specifies the behaviors that, when violated by a device (thing), cause an alert.
func (o LookupSecurityProfileResultOutput) Behaviors() SecurityProfileBehaviorArrayOutput {
	return o.ApplyT(func(v LookupSecurityProfileResult) []SecurityProfileBehavior { return v.Behaviors }).(SecurityProfileBehaviorArrayOutput)
}

// The ARN (Amazon resource name) of the created security profile.
func (o LookupSecurityProfileResultOutput) SecurityProfileArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSecurityProfileResult) *string { return v.SecurityProfileArn }).(pulumi.StringPtrOutput)
}

// A description of the security profile.
func (o LookupSecurityProfileResultOutput) SecurityProfileDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSecurityProfileResult) *string { return v.SecurityProfileDescription }).(pulumi.StringPtrOutput)
}

// Metadata that can be used to manage the security profile.
func (o LookupSecurityProfileResultOutput) Tags() SecurityProfileTagArrayOutput {
	return o.ApplyT(func(v LookupSecurityProfileResult) []SecurityProfileTag { return v.Tags }).(SecurityProfileTagArrayOutput)
}

// A set of target ARNs that the security profile is attached to.
func (o LookupSecurityProfileResultOutput) TargetArns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupSecurityProfileResult) []string { return v.TargetArns }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSecurityProfileResultOutput{})
}
