// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package globalaccelerator

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::GlobalAccelerator::EndpointGroup
type EndpointGroup struct {
	pulumi.CustomResourceState

	// The list of endpoint objects.
	EndpointConfigurations EndpointGroupEndpointConfigurationArrayOutput `pulumi:"endpointConfigurations"`
	// The Amazon Resource Name (ARN) of the endpoint group
	EndpointGroupArn pulumi.StringOutput `pulumi:"endpointGroupArn"`
	// The name of the AWS Region where the endpoint group is located
	EndpointGroupRegion pulumi.StringOutput `pulumi:"endpointGroupRegion"`
	// The time in seconds between each health check for an endpoint. Must be a value of 10 or 30
	HealthCheckIntervalSeconds pulumi.IntPtrOutput    `pulumi:"healthCheckIntervalSeconds"`
	HealthCheckPath            pulumi.StringPtrOutput `pulumi:"healthCheckPath"`
	// The port that AWS Global Accelerator uses to check the health of endpoints in this endpoint group.
	HealthCheckPort pulumi.IntPtrOutput `pulumi:"healthCheckPort"`
	// The protocol that AWS Global Accelerator uses to check the health of endpoints in this endpoint group.
	HealthCheckProtocol EndpointGroupHealthCheckProtocolPtrOutput `pulumi:"healthCheckProtocol"`
	// The Amazon Resource Name (ARN) of the listener
	ListenerArn   pulumi.StringOutput                  `pulumi:"listenerArn"`
	PortOverrides EndpointGroupPortOverrideArrayOutput `pulumi:"portOverrides"`
	// The number of consecutive health checks required to set the state of the endpoint to unhealthy.
	ThresholdCount pulumi.IntPtrOutput `pulumi:"thresholdCount"`
	// The percentage of traffic to sent to an AWS Region
	TrafficDialPercentage pulumi.Float64PtrOutput `pulumi:"trafficDialPercentage"`
}

// NewEndpointGroup registers a new resource with the given unique name, arguments, and options.
func NewEndpointGroup(ctx *pulumi.Context,
	name string, args *EndpointGroupArgs, opts ...pulumi.ResourceOption) (*EndpointGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EndpointGroupRegion == nil {
		return nil, errors.New("invalid value for required argument 'EndpointGroupRegion'")
	}
	if args.ListenerArn == nil {
		return nil, errors.New("invalid value for required argument 'ListenerArn'")
	}
	var resource EndpointGroup
	err := ctx.RegisterResource("aws-native:globalaccelerator:EndpointGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEndpointGroup gets an existing EndpointGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEndpointGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EndpointGroupState, opts ...pulumi.ResourceOption) (*EndpointGroup, error) {
	var resource EndpointGroup
	err := ctx.ReadResource("aws-native:globalaccelerator:EndpointGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EndpointGroup resources.
type endpointGroupState struct {
}

type EndpointGroupState struct {
}

func (EndpointGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*endpointGroupState)(nil)).Elem()
}

type endpointGroupArgs struct {
	// The list of endpoint objects.
	EndpointConfigurations []EndpointGroupEndpointConfiguration `pulumi:"endpointConfigurations"`
	// The name of the AWS Region where the endpoint group is located
	EndpointGroupRegion string `pulumi:"endpointGroupRegion"`
	// The time in seconds between each health check for an endpoint. Must be a value of 10 or 30
	HealthCheckIntervalSeconds *int    `pulumi:"healthCheckIntervalSeconds"`
	HealthCheckPath            *string `pulumi:"healthCheckPath"`
	// The port that AWS Global Accelerator uses to check the health of endpoints in this endpoint group.
	HealthCheckPort *int `pulumi:"healthCheckPort"`
	// The protocol that AWS Global Accelerator uses to check the health of endpoints in this endpoint group.
	HealthCheckProtocol *EndpointGroupHealthCheckProtocol `pulumi:"healthCheckProtocol"`
	// The Amazon Resource Name (ARN) of the listener
	ListenerArn   string                      `pulumi:"listenerArn"`
	PortOverrides []EndpointGroupPortOverride `pulumi:"portOverrides"`
	// The number of consecutive health checks required to set the state of the endpoint to unhealthy.
	ThresholdCount *int `pulumi:"thresholdCount"`
	// The percentage of traffic to sent to an AWS Region
	TrafficDialPercentage *float64 `pulumi:"trafficDialPercentage"`
}

// The set of arguments for constructing a EndpointGroup resource.
type EndpointGroupArgs struct {
	// The list of endpoint objects.
	EndpointConfigurations EndpointGroupEndpointConfigurationArrayInput
	// The name of the AWS Region where the endpoint group is located
	EndpointGroupRegion pulumi.StringInput
	// The time in seconds between each health check for an endpoint. Must be a value of 10 or 30
	HealthCheckIntervalSeconds pulumi.IntPtrInput
	HealthCheckPath            pulumi.StringPtrInput
	// The port that AWS Global Accelerator uses to check the health of endpoints in this endpoint group.
	HealthCheckPort pulumi.IntPtrInput
	// The protocol that AWS Global Accelerator uses to check the health of endpoints in this endpoint group.
	HealthCheckProtocol EndpointGroupHealthCheckProtocolPtrInput
	// The Amazon Resource Name (ARN) of the listener
	ListenerArn   pulumi.StringInput
	PortOverrides EndpointGroupPortOverrideArrayInput
	// The number of consecutive health checks required to set the state of the endpoint to unhealthy.
	ThresholdCount pulumi.IntPtrInput
	// The percentage of traffic to sent to an AWS Region
	TrafficDialPercentage pulumi.Float64PtrInput
}

func (EndpointGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*endpointGroupArgs)(nil)).Elem()
}

type EndpointGroupInput interface {
	pulumi.Input

	ToEndpointGroupOutput() EndpointGroupOutput
	ToEndpointGroupOutputWithContext(ctx context.Context) EndpointGroupOutput
}

func (*EndpointGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**EndpointGroup)(nil)).Elem()
}

func (i *EndpointGroup) ToEndpointGroupOutput() EndpointGroupOutput {
	return i.ToEndpointGroupOutputWithContext(context.Background())
}

func (i *EndpointGroup) ToEndpointGroupOutputWithContext(ctx context.Context) EndpointGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointGroupOutput)
}

type EndpointGroupOutput struct{ *pulumi.OutputState }

func (EndpointGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EndpointGroup)(nil)).Elem()
}

func (o EndpointGroupOutput) ToEndpointGroupOutput() EndpointGroupOutput {
	return o
}

func (o EndpointGroupOutput) ToEndpointGroupOutputWithContext(ctx context.Context) EndpointGroupOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EndpointGroupInput)(nil)).Elem(), &EndpointGroup{})
	pulumi.RegisterOutputType(EndpointGroupOutput{})
}
