// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package lightsail

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Lightsail::LoadBalancer
func LookupLoadBalancer(ctx *pulumi.Context, args *LookupLoadBalancerArgs, opts ...pulumi.InvokeOption) (*LookupLoadBalancerResult, error) {
	var rv LookupLoadBalancerResult
	err := ctx.Invoke("aws-native:lightsail:getLoadBalancer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupLoadBalancerArgs struct {
	// The name of your load balancer.
	LoadBalancerName string `pulumi:"loadBalancerName"`
}

type LookupLoadBalancerResult struct {
	// The names of the instances attached to the load balancer.
	AttachedInstances []string `pulumi:"attachedInstances"`
	// The path you provided to perform the load balancer health check. If you didn't specify a health check path, Lightsail uses the root path of your website (e.g., "/").
	HealthCheckPath *string `pulumi:"healthCheckPath"`
	LoadBalancerArn *string `pulumi:"loadBalancerArn"`
	// Configuration option to enable session stickiness.
	SessionStickinessEnabled *bool `pulumi:"sessionStickinessEnabled"`
	// Configuration option to adjust session stickiness cookie duration parameter.
	SessionStickinessLBCookieDurationSeconds *string `pulumi:"sessionStickinessLBCookieDurationSeconds"`
	// An array of key-value pairs to apply to this resource.
	Tags []LoadBalancerTag `pulumi:"tags"`
}

func LookupLoadBalancerOutput(ctx *pulumi.Context, args LookupLoadBalancerOutputArgs, opts ...pulumi.InvokeOption) LookupLoadBalancerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupLoadBalancerResult, error) {
			args := v.(LookupLoadBalancerArgs)
			r, err := LookupLoadBalancer(ctx, &args, opts...)
			var s LookupLoadBalancerResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupLoadBalancerResultOutput)
}

type LookupLoadBalancerOutputArgs struct {
	// The name of your load balancer.
	LoadBalancerName pulumi.StringInput `pulumi:"loadBalancerName"`
}

func (LookupLoadBalancerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLoadBalancerArgs)(nil)).Elem()
}

type LookupLoadBalancerResultOutput struct{ *pulumi.OutputState }

func (LookupLoadBalancerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLoadBalancerResult)(nil)).Elem()
}

func (o LookupLoadBalancerResultOutput) ToLookupLoadBalancerResultOutput() LookupLoadBalancerResultOutput {
	return o
}

func (o LookupLoadBalancerResultOutput) ToLookupLoadBalancerResultOutputWithContext(ctx context.Context) LookupLoadBalancerResultOutput {
	return o
}

// The names of the instances attached to the load balancer.
func (o LookupLoadBalancerResultOutput) AttachedInstances() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupLoadBalancerResult) []string { return v.AttachedInstances }).(pulumi.StringArrayOutput)
}

// The path you provided to perform the load balancer health check. If you didn't specify a health check path, Lightsail uses the root path of your website (e.g., "/").
func (o LookupLoadBalancerResultOutput) HealthCheckPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLoadBalancerResult) *string { return v.HealthCheckPath }).(pulumi.StringPtrOutput)
}

func (o LookupLoadBalancerResultOutput) LoadBalancerArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLoadBalancerResult) *string { return v.LoadBalancerArn }).(pulumi.StringPtrOutput)
}

// Configuration option to enable session stickiness.
func (o LookupLoadBalancerResultOutput) SessionStickinessEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupLoadBalancerResult) *bool { return v.SessionStickinessEnabled }).(pulumi.BoolPtrOutput)
}

// Configuration option to adjust session stickiness cookie duration parameter.
func (o LookupLoadBalancerResultOutput) SessionStickinessLBCookieDurationSeconds() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLoadBalancerResult) *string { return v.SessionStickinessLBCookieDurationSeconds }).(pulumi.StringPtrOutput)
}

// An array of key-value pairs to apply to this resource.
func (o LookupLoadBalancerResultOutput) Tags() LoadBalancerTagArrayOutput {
	return o.ApplyT(func(v LookupLoadBalancerResult) []LoadBalancerTag { return v.Tags }).(LoadBalancerTagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupLoadBalancerResultOutput{})
}
