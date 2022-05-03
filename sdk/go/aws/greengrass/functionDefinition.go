// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package greengrass

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Greengrass::FunctionDefinition
//
// Deprecated: FunctionDefinition is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type FunctionDefinition struct {
	pulumi.CustomResourceState

	Arn              pulumi.StringOutput                    `pulumi:"arn"`
	InitialVersion   FunctionDefinitionVersionTypePtrOutput `pulumi:"initialVersion"`
	LatestVersionArn pulumi.StringOutput                    `pulumi:"latestVersionArn"`
	Name             pulumi.StringOutput                    `pulumi:"name"`
	Tags             pulumi.AnyOutput                       `pulumi:"tags"`
}

// NewFunctionDefinition registers a new resource with the given unique name, arguments, and options.
func NewFunctionDefinition(ctx *pulumi.Context,
	name string, args *FunctionDefinitionArgs, opts ...pulumi.ResourceOption) (*FunctionDefinition, error) {
	if args == nil {
		args = &FunctionDefinitionArgs{}
	}

	var resource FunctionDefinition
	err := ctx.RegisterResource("aws-native:greengrass:FunctionDefinition", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFunctionDefinition gets an existing FunctionDefinition resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFunctionDefinition(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FunctionDefinitionState, opts ...pulumi.ResourceOption) (*FunctionDefinition, error) {
	var resource FunctionDefinition
	err := ctx.ReadResource("aws-native:greengrass:FunctionDefinition", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FunctionDefinition resources.
type functionDefinitionState struct {
}

type FunctionDefinitionState struct {
}

func (FunctionDefinitionState) ElementType() reflect.Type {
	return reflect.TypeOf((*functionDefinitionState)(nil)).Elem()
}

type functionDefinitionArgs struct {
	InitialVersion *FunctionDefinitionVersionType `pulumi:"initialVersion"`
	Name           *string                        `pulumi:"name"`
	Tags           interface{}                    `pulumi:"tags"`
}

// The set of arguments for constructing a FunctionDefinition resource.
type FunctionDefinitionArgs struct {
	InitialVersion FunctionDefinitionVersionTypePtrInput
	Name           pulumi.StringPtrInput
	Tags           pulumi.Input
}

func (FunctionDefinitionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*functionDefinitionArgs)(nil)).Elem()
}

type FunctionDefinitionInput interface {
	pulumi.Input

	ToFunctionDefinitionOutput() FunctionDefinitionOutput
	ToFunctionDefinitionOutputWithContext(ctx context.Context) FunctionDefinitionOutput
}

func (*FunctionDefinition) ElementType() reflect.Type {
	return reflect.TypeOf((**FunctionDefinition)(nil)).Elem()
}

func (i *FunctionDefinition) ToFunctionDefinitionOutput() FunctionDefinitionOutput {
	return i.ToFunctionDefinitionOutputWithContext(context.Background())
}

func (i *FunctionDefinition) ToFunctionDefinitionOutputWithContext(ctx context.Context) FunctionDefinitionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FunctionDefinitionOutput)
}

type FunctionDefinitionOutput struct{ *pulumi.OutputState }

func (FunctionDefinitionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FunctionDefinition)(nil)).Elem()
}

func (o FunctionDefinitionOutput) ToFunctionDefinitionOutput() FunctionDefinitionOutput {
	return o
}

func (o FunctionDefinitionOutput) ToFunctionDefinitionOutputWithContext(ctx context.Context) FunctionDefinitionOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FunctionDefinitionInput)(nil)).Elem(), &FunctionDefinition{})
	pulumi.RegisterOutputType(FunctionDefinitionOutput{})
}
