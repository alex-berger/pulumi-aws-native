// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package events

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Events::Endpoint.
type Endpoint struct {
	pulumi.CustomResourceState

	Arn               pulumi.StringOutput                `pulumi:"arn"`
	Description       pulumi.StringPtrOutput             `pulumi:"description"`
	EndpointId        pulumi.StringOutput                `pulumi:"endpointId"`
	EndpointUrl       pulumi.StringOutput                `pulumi:"endpointUrl"`
	EventBuses        EndpointEventBusArrayOutput        `pulumi:"eventBuses"`
	Name              pulumi.StringOutput                `pulumi:"name"`
	ReplicationConfig EndpointReplicationConfigPtrOutput `pulumi:"replicationConfig"`
	RoleArn           pulumi.StringPtrOutput             `pulumi:"roleArn"`
	RoutingConfig     EndpointRoutingConfigOutput        `pulumi:"routingConfig"`
	State             EndpointStateEnumOutput            `pulumi:"state"`
	StateReason       pulumi.StringOutput                `pulumi:"stateReason"`
}

// NewEndpoint registers a new resource with the given unique name, arguments, and options.
func NewEndpoint(ctx *pulumi.Context,
	name string, args *EndpointArgs, opts ...pulumi.ResourceOption) (*Endpoint, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EventBuses == nil {
		return nil, errors.New("invalid value for required argument 'EventBuses'")
	}
	if args.RoutingConfig == nil {
		return nil, errors.New("invalid value for required argument 'RoutingConfig'")
	}
	var resource Endpoint
	err := ctx.RegisterResource("aws-native:events:Endpoint", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEndpoint gets an existing Endpoint resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEndpoint(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EndpointState, opts ...pulumi.ResourceOption) (*Endpoint, error) {
	var resource Endpoint
	err := ctx.ReadResource("aws-native:events:Endpoint", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Endpoint resources.
type endpointState struct {
}

type EndpointState struct {
}

func (EndpointState) ElementType() reflect.Type {
	return reflect.TypeOf((*endpointState)(nil)).Elem()
}

type endpointArgs struct {
	Description       *string                    `pulumi:"description"`
	EventBuses        []EndpointEventBus         `pulumi:"eventBuses"`
	Name              *string                    `pulumi:"name"`
	ReplicationConfig *EndpointReplicationConfig `pulumi:"replicationConfig"`
	RoleArn           *string                    `pulumi:"roleArn"`
	RoutingConfig     EndpointRoutingConfig      `pulumi:"routingConfig"`
}

// The set of arguments for constructing a Endpoint resource.
type EndpointArgs struct {
	Description       pulumi.StringPtrInput
	EventBuses        EndpointEventBusArrayInput
	Name              pulumi.StringPtrInput
	ReplicationConfig EndpointReplicationConfigPtrInput
	RoleArn           pulumi.StringPtrInput
	RoutingConfig     EndpointRoutingConfigInput
}

func (EndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*endpointArgs)(nil)).Elem()
}

type EndpointInput interface {
	pulumi.Input

	ToEndpointOutput() EndpointOutput
	ToEndpointOutputWithContext(ctx context.Context) EndpointOutput
}

func (*Endpoint) ElementType() reflect.Type {
	return reflect.TypeOf((**Endpoint)(nil)).Elem()
}

func (i *Endpoint) ToEndpointOutput() EndpointOutput {
	return i.ToEndpointOutputWithContext(context.Background())
}

func (i *Endpoint) ToEndpointOutputWithContext(ctx context.Context) EndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointOutput)
}

type EndpointOutput struct{ *pulumi.OutputState }

func (EndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Endpoint)(nil)).Elem()
}

func (o EndpointOutput) ToEndpointOutput() EndpointOutput {
	return o
}

func (o EndpointOutput) ToEndpointOutputWithContext(ctx context.Context) EndpointOutput {
	return o
}

func (o EndpointOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *Endpoint) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

func (o EndpointOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Endpoint) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o EndpointOutput) EndpointId() pulumi.StringOutput {
	return o.ApplyT(func(v *Endpoint) pulumi.StringOutput { return v.EndpointId }).(pulumi.StringOutput)
}

func (o EndpointOutput) EndpointUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *Endpoint) pulumi.StringOutput { return v.EndpointUrl }).(pulumi.StringOutput)
}

func (o EndpointOutput) EventBuses() EndpointEventBusArrayOutput {
	return o.ApplyT(func(v *Endpoint) EndpointEventBusArrayOutput { return v.EventBuses }).(EndpointEventBusArrayOutput)
}

func (o EndpointOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Endpoint) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o EndpointOutput) ReplicationConfig() EndpointReplicationConfigPtrOutput {
	return o.ApplyT(func(v *Endpoint) EndpointReplicationConfigPtrOutput { return v.ReplicationConfig }).(EndpointReplicationConfigPtrOutput)
}

func (o EndpointOutput) RoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Endpoint) pulumi.StringPtrOutput { return v.RoleArn }).(pulumi.StringPtrOutput)
}

func (o EndpointOutput) RoutingConfig() EndpointRoutingConfigOutput {
	return o.ApplyT(func(v *Endpoint) EndpointRoutingConfigOutput { return v.RoutingConfig }).(EndpointRoutingConfigOutput)
}

func (o EndpointOutput) State() EndpointStateEnumOutput {
	return o.ApplyT(func(v *Endpoint) EndpointStateEnumOutput { return v.State }).(EndpointStateEnumOutput)
}

func (o EndpointOutput) StateReason() pulumi.StringOutput {
	return o.ApplyT(func(v *Endpoint) pulumi.StringOutput { return v.StateReason }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EndpointInput)(nil)).Elem(), &Endpoint{})
	pulumi.RegisterOutputType(EndpointOutput{})
}
