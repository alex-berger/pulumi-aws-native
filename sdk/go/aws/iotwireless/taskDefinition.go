// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package iotwireless

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a gateway task definition.
type TaskDefinition struct {
	pulumi.CustomResourceState

	// TaskDefinition arn. Returned after successful create.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Whether to automatically create tasks using this task definition for all gateways with the specified current version. If false, the task must me created by calling CreateWirelessGatewayTask.
	AutoCreateTasks pulumi.BoolOutput `pulumi:"autoCreateTasks"`
	// The list of task definitions.
	LoRaWANUpdateGatewayTaskEntry TaskDefinitionLoRaWANUpdateGatewayTaskEntryPtrOutput `pulumi:"loRaWANUpdateGatewayTaskEntry"`
	// The name of the new resource.
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// A list of key-value pairs that contain metadata for the destination.
	Tags TaskDefinitionTagArrayOutput `pulumi:"tags"`
	// A filter to list only the wireless gateway task definitions that use this task definition type
	TaskDefinitionType TaskDefinitionTypePtrOutput `pulumi:"taskDefinitionType"`
	// Information about the gateways to update.
	Update TaskDefinitionUpdateWirelessGatewayTaskCreatePtrOutput `pulumi:"update"`
}

// NewTaskDefinition registers a new resource with the given unique name, arguments, and options.
func NewTaskDefinition(ctx *pulumi.Context,
	name string, args *TaskDefinitionArgs, opts ...pulumi.ResourceOption) (*TaskDefinition, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AutoCreateTasks == nil {
		return nil, errors.New("invalid value for required argument 'AutoCreateTasks'")
	}
	var resource TaskDefinition
	err := ctx.RegisterResource("aws-native:iotwireless:TaskDefinition", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTaskDefinition gets an existing TaskDefinition resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTaskDefinition(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TaskDefinitionState, opts ...pulumi.ResourceOption) (*TaskDefinition, error) {
	var resource TaskDefinition
	err := ctx.ReadResource("aws-native:iotwireless:TaskDefinition", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TaskDefinition resources.
type taskDefinitionState struct {
}

type TaskDefinitionState struct {
}

func (TaskDefinitionState) ElementType() reflect.Type {
	return reflect.TypeOf((*taskDefinitionState)(nil)).Elem()
}

type taskDefinitionArgs struct {
	// Whether to automatically create tasks using this task definition for all gateways with the specified current version. If false, the task must me created by calling CreateWirelessGatewayTask.
	AutoCreateTasks bool `pulumi:"autoCreateTasks"`
	// The list of task definitions.
	LoRaWANUpdateGatewayTaskEntry *TaskDefinitionLoRaWANUpdateGatewayTaskEntry `pulumi:"loRaWANUpdateGatewayTaskEntry"`
	// The name of the new resource.
	Name *string `pulumi:"name"`
	// A list of key-value pairs that contain metadata for the destination.
	Tags []TaskDefinitionTag `pulumi:"tags"`
	// A filter to list only the wireless gateway task definitions that use this task definition type
	TaskDefinitionType *TaskDefinitionType `pulumi:"taskDefinitionType"`
	// Information about the gateways to update.
	Update *TaskDefinitionUpdateWirelessGatewayTaskCreate `pulumi:"update"`
}

// The set of arguments for constructing a TaskDefinition resource.
type TaskDefinitionArgs struct {
	// Whether to automatically create tasks using this task definition for all gateways with the specified current version. If false, the task must me created by calling CreateWirelessGatewayTask.
	AutoCreateTasks pulumi.BoolInput
	// The list of task definitions.
	LoRaWANUpdateGatewayTaskEntry TaskDefinitionLoRaWANUpdateGatewayTaskEntryPtrInput
	// The name of the new resource.
	Name pulumi.StringPtrInput
	// A list of key-value pairs that contain metadata for the destination.
	Tags TaskDefinitionTagArrayInput
	// A filter to list only the wireless gateway task definitions that use this task definition type
	TaskDefinitionType TaskDefinitionTypePtrInput
	// Information about the gateways to update.
	Update TaskDefinitionUpdateWirelessGatewayTaskCreatePtrInput
}

func (TaskDefinitionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*taskDefinitionArgs)(nil)).Elem()
}

type TaskDefinitionInput interface {
	pulumi.Input

	ToTaskDefinitionOutput() TaskDefinitionOutput
	ToTaskDefinitionOutputWithContext(ctx context.Context) TaskDefinitionOutput
}

func (*TaskDefinition) ElementType() reflect.Type {
	return reflect.TypeOf((**TaskDefinition)(nil)).Elem()
}

func (i *TaskDefinition) ToTaskDefinitionOutput() TaskDefinitionOutput {
	return i.ToTaskDefinitionOutputWithContext(context.Background())
}

func (i *TaskDefinition) ToTaskDefinitionOutputWithContext(ctx context.Context) TaskDefinitionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TaskDefinitionOutput)
}

type TaskDefinitionOutput struct{ *pulumi.OutputState }

func (TaskDefinitionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TaskDefinition)(nil)).Elem()
}

func (o TaskDefinitionOutput) ToTaskDefinitionOutput() TaskDefinitionOutput {
	return o
}

func (o TaskDefinitionOutput) ToTaskDefinitionOutputWithContext(ctx context.Context) TaskDefinitionOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TaskDefinitionInput)(nil)).Elem(), &TaskDefinition{})
	pulumi.RegisterOutputType(TaskDefinitionOutput{})
}
