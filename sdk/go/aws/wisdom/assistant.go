// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package wisdom

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Definition of AWS::Wisdom::Assistant Resource Type
type Assistant struct {
	pulumi.CustomResourceState

	AssistantArn                      pulumi.StringOutput                                 `pulumi:"assistantArn"`
	AssistantId                       pulumi.StringOutput                                 `pulumi:"assistantId"`
	Description                       pulumi.StringPtrOutput                              `pulumi:"description"`
	Name                              pulumi.StringOutput                                 `pulumi:"name"`
	ServerSideEncryptionConfiguration AssistantServerSideEncryptionConfigurationPtrOutput `pulumi:"serverSideEncryptionConfiguration"`
	Tags                              AssistantTagArrayOutput                             `pulumi:"tags"`
	Type                              AssistantTypeOutput                                 `pulumi:"type"`
}

// NewAssistant registers a new resource with the given unique name, arguments, and options.
func NewAssistant(ctx *pulumi.Context,
	name string, args *AssistantArgs, opts ...pulumi.ResourceOption) (*Assistant, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	var resource Assistant
	err := ctx.RegisterResource("aws-native:wisdom:Assistant", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAssistant gets an existing Assistant resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAssistant(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AssistantState, opts ...pulumi.ResourceOption) (*Assistant, error) {
	var resource Assistant
	err := ctx.ReadResource("aws-native:wisdom:Assistant", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Assistant resources.
type assistantState struct {
}

type AssistantState struct {
}

func (AssistantState) ElementType() reflect.Type {
	return reflect.TypeOf((*assistantState)(nil)).Elem()
}

type assistantArgs struct {
	Description                       *string                                     `pulumi:"description"`
	Name                              *string                                     `pulumi:"name"`
	ServerSideEncryptionConfiguration *AssistantServerSideEncryptionConfiguration `pulumi:"serverSideEncryptionConfiguration"`
	Tags                              []AssistantTag                              `pulumi:"tags"`
	Type                              AssistantType                               `pulumi:"type"`
}

// The set of arguments for constructing a Assistant resource.
type AssistantArgs struct {
	Description                       pulumi.StringPtrInput
	Name                              pulumi.StringPtrInput
	ServerSideEncryptionConfiguration AssistantServerSideEncryptionConfigurationPtrInput
	Tags                              AssistantTagArrayInput
	Type                              AssistantTypeInput
}

func (AssistantArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*assistantArgs)(nil)).Elem()
}

type AssistantInput interface {
	pulumi.Input

	ToAssistantOutput() AssistantOutput
	ToAssistantOutputWithContext(ctx context.Context) AssistantOutput
}

func (*Assistant) ElementType() reflect.Type {
	return reflect.TypeOf((**Assistant)(nil)).Elem()
}

func (i *Assistant) ToAssistantOutput() AssistantOutput {
	return i.ToAssistantOutputWithContext(context.Background())
}

func (i *Assistant) ToAssistantOutputWithContext(ctx context.Context) AssistantOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssistantOutput)
}

type AssistantOutput struct{ *pulumi.OutputState }

func (AssistantOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Assistant)(nil)).Elem()
}

func (o AssistantOutput) ToAssistantOutput() AssistantOutput {
	return o
}

func (o AssistantOutput) ToAssistantOutputWithContext(ctx context.Context) AssistantOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AssistantInput)(nil)).Elem(), &Assistant{})
	pulumi.RegisterOutputType(AssistantOutput{})
}
