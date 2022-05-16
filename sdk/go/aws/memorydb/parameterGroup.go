// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package memorydb

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The AWS::MemoryDB::ParameterGroup resource creates an Amazon MemoryDB ParameterGroup.
type ParameterGroup struct {
	pulumi.CustomResourceState

	// The Amazon Resource Name (ARN) of the parameter group.
	ARN pulumi.StringOutput `pulumi:"aRN"`
	// A description of the parameter group.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The name of the parameter group family that this parameter group is compatible with.
	Family pulumi.StringOutput `pulumi:"family"`
	// The name of the parameter group.
	ParameterGroupName pulumi.StringOutput `pulumi:"parameterGroupName"`
	// An map of parameter names and values for the parameter update. You must supply at least one parameter name and value; subsequent arguments are optional.
	Parameters pulumi.AnyOutput `pulumi:"parameters"`
	// An array of key-value pairs to apply to this parameter group.
	Tags ParameterGroupTagArrayOutput `pulumi:"tags"`
}

// NewParameterGroup registers a new resource with the given unique name, arguments, and options.
func NewParameterGroup(ctx *pulumi.Context,
	name string, args *ParameterGroupArgs, opts ...pulumi.ResourceOption) (*ParameterGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Family == nil {
		return nil, errors.New("invalid value for required argument 'Family'")
	}
	var resource ParameterGroup
	err := ctx.RegisterResource("aws-native:memorydb:ParameterGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetParameterGroup gets an existing ParameterGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetParameterGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ParameterGroupState, opts ...pulumi.ResourceOption) (*ParameterGroup, error) {
	var resource ParameterGroup
	err := ctx.ReadResource("aws-native:memorydb:ParameterGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ParameterGroup resources.
type parameterGroupState struct {
}

type ParameterGroupState struct {
}

func (ParameterGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*parameterGroupState)(nil)).Elem()
}

type parameterGroupArgs struct {
	// A description of the parameter group.
	Description *string `pulumi:"description"`
	// The name of the parameter group family that this parameter group is compatible with.
	Family string `pulumi:"family"`
	// The name of the parameter group.
	ParameterGroupName *string `pulumi:"parameterGroupName"`
	// An map of parameter names and values for the parameter update. You must supply at least one parameter name and value; subsequent arguments are optional.
	Parameters interface{} `pulumi:"parameters"`
	// An array of key-value pairs to apply to this parameter group.
	Tags []ParameterGroupTag `pulumi:"tags"`
}

// The set of arguments for constructing a ParameterGroup resource.
type ParameterGroupArgs struct {
	// A description of the parameter group.
	Description pulumi.StringPtrInput
	// The name of the parameter group family that this parameter group is compatible with.
	Family pulumi.StringInput
	// The name of the parameter group.
	ParameterGroupName pulumi.StringPtrInput
	// An map of parameter names and values for the parameter update. You must supply at least one parameter name and value; subsequent arguments are optional.
	Parameters pulumi.Input
	// An array of key-value pairs to apply to this parameter group.
	Tags ParameterGroupTagArrayInput
}

func (ParameterGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*parameterGroupArgs)(nil)).Elem()
}

type ParameterGroupInput interface {
	pulumi.Input

	ToParameterGroupOutput() ParameterGroupOutput
	ToParameterGroupOutputWithContext(ctx context.Context) ParameterGroupOutput
}

func (*ParameterGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**ParameterGroup)(nil)).Elem()
}

func (i *ParameterGroup) ToParameterGroupOutput() ParameterGroupOutput {
	return i.ToParameterGroupOutputWithContext(context.Background())
}

func (i *ParameterGroup) ToParameterGroupOutputWithContext(ctx context.Context) ParameterGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ParameterGroupOutput)
}

type ParameterGroupOutput struct{ *pulumi.OutputState }

func (ParameterGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ParameterGroup)(nil)).Elem()
}

func (o ParameterGroupOutput) ToParameterGroupOutput() ParameterGroupOutput {
	return o
}

func (o ParameterGroupOutput) ToParameterGroupOutputWithContext(ctx context.Context) ParameterGroupOutput {
	return o
}

// The Amazon Resource Name (ARN) of the parameter group.
func (o ParameterGroupOutput) ARN() pulumi.StringOutput {
	return o.ApplyT(func(v *ParameterGroup) pulumi.StringOutput { return v.ARN }).(pulumi.StringOutput)
}

// A description of the parameter group.
func (o ParameterGroupOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ParameterGroup) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The name of the parameter group family that this parameter group is compatible with.
func (o ParameterGroupOutput) Family() pulumi.StringOutput {
	return o.ApplyT(func(v *ParameterGroup) pulumi.StringOutput { return v.Family }).(pulumi.StringOutput)
}

// The name of the parameter group.
func (o ParameterGroupOutput) ParameterGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v *ParameterGroup) pulumi.StringOutput { return v.ParameterGroupName }).(pulumi.StringOutput)
}

// An map of parameter names and values for the parameter update. You must supply at least one parameter name and value; subsequent arguments are optional.
func (o ParameterGroupOutput) Parameters() pulumi.AnyOutput {
	return o.ApplyT(func(v *ParameterGroup) pulumi.AnyOutput { return v.Parameters }).(pulumi.AnyOutput)
}

// An array of key-value pairs to apply to this parameter group.
func (o ParameterGroupOutput) Tags() ParameterGroupTagArrayOutput {
	return o.ApplyT(func(v *ParameterGroup) ParameterGroupTagArrayOutput { return v.Tags }).(ParameterGroupTagArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ParameterGroupInput)(nil)).Elem(), &ParameterGroup{})
	pulumi.RegisterOutputType(ParameterGroupOutput{})
}
