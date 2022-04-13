// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package globalaccelerator

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::GlobalAccelerator::EndpointGroup
func LookupEndpointGroup(ctx *pulumi.Context, args *LookupEndpointGroupArgs, opts ...pulumi.InvokeOption) (*LookupEndpointGroupResult, error) {
	var rv LookupEndpointGroupResult
	err := ctx.Invoke("aws-native:globalaccelerator:getEndpointGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupEndpointGroupArgs struct {
	// The Amazon Resource Name (ARN) of the endpoint group
	EndpointGroupArn string `pulumi:"endpointGroupArn"`
}

type LookupEndpointGroupResult struct {
	// The list of endpoint objects.
	EndpointConfigurations []EndpointGroupEndpointConfiguration `pulumi:"endpointConfigurations"`
	// The Amazon Resource Name (ARN) of the endpoint group
	EndpointGroupArn *string `pulumi:"endpointGroupArn"`
	// The time in seconds between each health check for an endpoint. Must be a value of 10 or 30
	HealthCheckIntervalSeconds *int    `pulumi:"healthCheckIntervalSeconds"`
	HealthCheckPath            *string `pulumi:"healthCheckPath"`
	// The port that AWS Global Accelerator uses to check the health of endpoints in this endpoint group.
	HealthCheckPort *int `pulumi:"healthCheckPort"`
	// The protocol that AWS Global Accelerator uses to check the health of endpoints in this endpoint group.
	HealthCheckProtocol *EndpointGroupHealthCheckProtocol `pulumi:"healthCheckProtocol"`
	PortOverrides       []EndpointGroupPortOverride       `pulumi:"portOverrides"`
	// The number of consecutive health checks required to set the state of the endpoint to unhealthy.
	ThresholdCount *int `pulumi:"thresholdCount"`
	// The percentage of traffic to sent to an AWS Region
	TrafficDialPercentage *float64 `pulumi:"trafficDialPercentage"`
}

func LookupEndpointGroupOutput(ctx *pulumi.Context, args LookupEndpointGroupOutputArgs, opts ...pulumi.InvokeOption) LookupEndpointGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupEndpointGroupResult, error) {
			args := v.(LookupEndpointGroupArgs)
			r, err := LookupEndpointGroup(ctx, &args, opts...)
			var s LookupEndpointGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupEndpointGroupResultOutput)
}

type LookupEndpointGroupOutputArgs struct {
	// The Amazon Resource Name (ARN) of the endpoint group
	EndpointGroupArn pulumi.StringInput `pulumi:"endpointGroupArn"`
}

func (LookupEndpointGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEndpointGroupArgs)(nil)).Elem()
}

type LookupEndpointGroupResultOutput struct{ *pulumi.OutputState }

func (LookupEndpointGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEndpointGroupResult)(nil)).Elem()
}

func (o LookupEndpointGroupResultOutput) ToLookupEndpointGroupResultOutput() LookupEndpointGroupResultOutput {
	return o
}

func (o LookupEndpointGroupResultOutput) ToLookupEndpointGroupResultOutputWithContext(ctx context.Context) LookupEndpointGroupResultOutput {
	return o
}

// The list of endpoint objects.
func (o LookupEndpointGroupResultOutput) EndpointConfigurations() EndpointGroupEndpointConfigurationArrayOutput {
	return o.ApplyT(func(v LookupEndpointGroupResult) []EndpointGroupEndpointConfiguration {
		return v.EndpointConfigurations
	}).(EndpointGroupEndpointConfigurationArrayOutput)
}

// The Amazon Resource Name (ARN) of the endpoint group
func (o LookupEndpointGroupResultOutput) EndpointGroupArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEndpointGroupResult) *string { return v.EndpointGroupArn }).(pulumi.StringPtrOutput)
}

// The time in seconds between each health check for an endpoint. Must be a value of 10 or 30
func (o LookupEndpointGroupResultOutput) HealthCheckIntervalSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupEndpointGroupResult) *int { return v.HealthCheckIntervalSeconds }).(pulumi.IntPtrOutput)
}

func (o LookupEndpointGroupResultOutput) HealthCheckPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEndpointGroupResult) *string { return v.HealthCheckPath }).(pulumi.StringPtrOutput)
}

// The port that AWS Global Accelerator uses to check the health of endpoints in this endpoint group.
func (o LookupEndpointGroupResultOutput) HealthCheckPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupEndpointGroupResult) *int { return v.HealthCheckPort }).(pulumi.IntPtrOutput)
}

// The protocol that AWS Global Accelerator uses to check the health of endpoints in this endpoint group.
func (o LookupEndpointGroupResultOutput) HealthCheckProtocol() EndpointGroupHealthCheckProtocolPtrOutput {
	return o.ApplyT(func(v LookupEndpointGroupResult) *EndpointGroupHealthCheckProtocol { return v.HealthCheckProtocol }).(EndpointGroupHealthCheckProtocolPtrOutput)
}

func (o LookupEndpointGroupResultOutput) PortOverrides() EndpointGroupPortOverrideArrayOutput {
	return o.ApplyT(func(v LookupEndpointGroupResult) []EndpointGroupPortOverride { return v.PortOverrides }).(EndpointGroupPortOverrideArrayOutput)
}

// The number of consecutive health checks required to set the state of the endpoint to unhealthy.
func (o LookupEndpointGroupResultOutput) ThresholdCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupEndpointGroupResult) *int { return v.ThresholdCount }).(pulumi.IntPtrOutput)
}

// The percentage of traffic to sent to an AWS Region
func (o LookupEndpointGroupResultOutput) TrafficDialPercentage() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupEndpointGroupResult) *float64 { return v.TrafficDialPercentage }).(pulumi.Float64PtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupEndpointGroupResultOutput{})
}
