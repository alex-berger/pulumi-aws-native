// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iotwireless

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Destination's resource schema demonstrating some basic constructs and validation rules.
type Destination struct {
	pulumi.CustomResourceState

	// Destination arn. Returned after successful create.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Destination description
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Destination expression
	Expression pulumi.StringOutput `pulumi:"expression"`
	// Must be RuleName
	ExpressionType DestinationExpressionTypeOutput `pulumi:"expressionType"`
	// Unique name of destination
	Name pulumi.StringOutput `pulumi:"name"`
	// AWS role ARN that grants access
	RoleArn pulumi.StringOutput `pulumi:"roleArn"`
	// A list of key-value pairs that contain metadata for the destination.
	Tags DestinationTagArrayOutput `pulumi:"tags"`
}

// NewDestination registers a new resource with the given unique name, arguments, and options.
func NewDestination(ctx *pulumi.Context,
	name string, args *DestinationArgs, opts ...pulumi.ResourceOption) (*Destination, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Expression == nil {
		return nil, errors.New("invalid value for required argument 'Expression'")
	}
	if args.ExpressionType == nil {
		return nil, errors.New("invalid value for required argument 'ExpressionType'")
	}
	if args.RoleArn == nil {
		return nil, errors.New("invalid value for required argument 'RoleArn'")
	}
	var resource Destination
	err := ctx.RegisterResource("aws-native:iotwireless:Destination", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDestination gets an existing Destination resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDestination(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DestinationState, opts ...pulumi.ResourceOption) (*Destination, error) {
	var resource Destination
	err := ctx.ReadResource("aws-native:iotwireless:Destination", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Destination resources.
type destinationState struct {
}

type DestinationState struct {
}

func (DestinationState) ElementType() reflect.Type {
	return reflect.TypeOf((*destinationState)(nil)).Elem()
}

type destinationArgs struct {
	// Destination description
	Description *string `pulumi:"description"`
	// Destination expression
	Expression string `pulumi:"expression"`
	// Must be RuleName
	ExpressionType DestinationExpressionType `pulumi:"expressionType"`
	// Unique name of destination
	Name *string `pulumi:"name"`
	// AWS role ARN that grants access
	RoleArn string `pulumi:"roleArn"`
	// A list of key-value pairs that contain metadata for the destination.
	Tags []DestinationTag `pulumi:"tags"`
}

// The set of arguments for constructing a Destination resource.
type DestinationArgs struct {
	// Destination description
	Description pulumi.StringPtrInput
	// Destination expression
	Expression pulumi.StringInput
	// Must be RuleName
	ExpressionType DestinationExpressionTypeInput
	// Unique name of destination
	Name pulumi.StringPtrInput
	// AWS role ARN that grants access
	RoleArn pulumi.StringInput
	// A list of key-value pairs that contain metadata for the destination.
	Tags DestinationTagArrayInput
}

func (DestinationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*destinationArgs)(nil)).Elem()
}

type DestinationInput interface {
	pulumi.Input

	ToDestinationOutput() DestinationOutput
	ToDestinationOutputWithContext(ctx context.Context) DestinationOutput
}

func (*Destination) ElementType() reflect.Type {
	return reflect.TypeOf((**Destination)(nil)).Elem()
}

func (i *Destination) ToDestinationOutput() DestinationOutput {
	return i.ToDestinationOutputWithContext(context.Background())
}

func (i *Destination) ToDestinationOutputWithContext(ctx context.Context) DestinationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DestinationOutput)
}

type DestinationOutput struct{ *pulumi.OutputState }

func (DestinationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Destination)(nil)).Elem()
}

func (o DestinationOutput) ToDestinationOutput() DestinationOutput {
	return o
}

func (o DestinationOutput) ToDestinationOutputWithContext(ctx context.Context) DestinationOutput {
	return o
}

// Destination arn. Returned after successful create.
func (o DestinationOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *Destination) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// Destination description
func (o DestinationOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Destination) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Destination expression
func (o DestinationOutput) Expression() pulumi.StringOutput {
	return o.ApplyT(func(v *Destination) pulumi.StringOutput { return v.Expression }).(pulumi.StringOutput)
}

// Must be RuleName
func (o DestinationOutput) ExpressionType() DestinationExpressionTypeOutput {
	return o.ApplyT(func(v *Destination) DestinationExpressionTypeOutput { return v.ExpressionType }).(DestinationExpressionTypeOutput)
}

// Unique name of destination
func (o DestinationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Destination) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// AWS role ARN that grants access
func (o DestinationOutput) RoleArn() pulumi.StringOutput {
	return o.ApplyT(func(v *Destination) pulumi.StringOutput { return v.RoleArn }).(pulumi.StringOutput)
}

// A list of key-value pairs that contain metadata for the destination.
func (o DestinationOutput) Tags() DestinationTagArrayOutput {
	return o.ApplyT(func(v *Destination) DestinationTagArrayOutput { return v.Tags }).(DestinationTagArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DestinationInput)(nil)).Elem(), &Destination{})
	pulumi.RegisterOutputType(DestinationOutput{})
}
