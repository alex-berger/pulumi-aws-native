// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudformation

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Specifies the configuration data for a registered hook in CloudFormation Registry.
type HookTypeConfig struct {
	pulumi.CustomResourceState

	// The configuration data for the extension, in this account and region.
	Configuration pulumi.StringPtrOutput `pulumi:"configuration"`
	// An alias by which to refer to this extension configuration data.
	ConfigurationAlias HookTypeConfigConfigurationAliasPtrOutput `pulumi:"configurationAlias"`
	// The Amazon Resource Name (ARN) for the configuration data, in this account and region.
	ConfigurationArn pulumi.StringOutput `pulumi:"configurationArn"`
	// The Amazon Resource Name (ARN) of the type without version number.
	TypeArn pulumi.StringPtrOutput `pulumi:"typeArn"`
	// The name of the type being registered.
	//
	// We recommend that type names adhere to the following pattern: company_or_organization::service::type.
	TypeName pulumi.StringPtrOutput `pulumi:"typeName"`
}

// NewHookTypeConfig registers a new resource with the given unique name, arguments, and options.
func NewHookTypeConfig(ctx *pulumi.Context,
	name string, args *HookTypeConfigArgs, opts ...pulumi.ResourceOption) (*HookTypeConfig, error) {
	if args == nil {
		args = &HookTypeConfigArgs{}
	}

	var resource HookTypeConfig
	err := ctx.RegisterResource("aws-native:cloudformation:HookTypeConfig", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetHookTypeConfig gets an existing HookTypeConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHookTypeConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HookTypeConfigState, opts ...pulumi.ResourceOption) (*HookTypeConfig, error) {
	var resource HookTypeConfig
	err := ctx.ReadResource("aws-native:cloudformation:HookTypeConfig", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering HookTypeConfig resources.
type hookTypeConfigState struct {
}

type HookTypeConfigState struct {
}

func (HookTypeConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*hookTypeConfigState)(nil)).Elem()
}

type hookTypeConfigArgs struct {
	// The configuration data for the extension, in this account and region.
	Configuration *string `pulumi:"configuration"`
	// An alias by which to refer to this extension configuration data.
	ConfigurationAlias *HookTypeConfigConfigurationAlias `pulumi:"configurationAlias"`
	// The Amazon Resource Name (ARN) of the type without version number.
	TypeArn *string `pulumi:"typeArn"`
	// The name of the type being registered.
	//
	// We recommend that type names adhere to the following pattern: company_or_organization::service::type.
	TypeName *string `pulumi:"typeName"`
}

// The set of arguments for constructing a HookTypeConfig resource.
type HookTypeConfigArgs struct {
	// The configuration data for the extension, in this account and region.
	Configuration pulumi.StringPtrInput
	// An alias by which to refer to this extension configuration data.
	ConfigurationAlias HookTypeConfigConfigurationAliasPtrInput
	// The Amazon Resource Name (ARN) of the type without version number.
	TypeArn pulumi.StringPtrInput
	// The name of the type being registered.
	//
	// We recommend that type names adhere to the following pattern: company_or_organization::service::type.
	TypeName pulumi.StringPtrInput
}

func (HookTypeConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*hookTypeConfigArgs)(nil)).Elem()
}

type HookTypeConfigInput interface {
	pulumi.Input

	ToHookTypeConfigOutput() HookTypeConfigOutput
	ToHookTypeConfigOutputWithContext(ctx context.Context) HookTypeConfigOutput
}

func (*HookTypeConfig) ElementType() reflect.Type {
	return reflect.TypeOf((**HookTypeConfig)(nil)).Elem()
}

func (i *HookTypeConfig) ToHookTypeConfigOutput() HookTypeConfigOutput {
	return i.ToHookTypeConfigOutputWithContext(context.Background())
}

func (i *HookTypeConfig) ToHookTypeConfigOutputWithContext(ctx context.Context) HookTypeConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HookTypeConfigOutput)
}

type HookTypeConfigOutput struct{ *pulumi.OutputState }

func (HookTypeConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HookTypeConfig)(nil)).Elem()
}

func (o HookTypeConfigOutput) ToHookTypeConfigOutput() HookTypeConfigOutput {
	return o
}

func (o HookTypeConfigOutput) ToHookTypeConfigOutputWithContext(ctx context.Context) HookTypeConfigOutput {
	return o
}

// The configuration data for the extension, in this account and region.
func (o HookTypeConfigOutput) Configuration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HookTypeConfig) pulumi.StringPtrOutput { return v.Configuration }).(pulumi.StringPtrOutput)
}

// An alias by which to refer to this extension configuration data.
func (o HookTypeConfigOutput) ConfigurationAlias() HookTypeConfigConfigurationAliasPtrOutput {
	return o.ApplyT(func(v *HookTypeConfig) HookTypeConfigConfigurationAliasPtrOutput { return v.ConfigurationAlias }).(HookTypeConfigConfigurationAliasPtrOutput)
}

// The Amazon Resource Name (ARN) for the configuration data, in this account and region.
func (o HookTypeConfigOutput) ConfigurationArn() pulumi.StringOutput {
	return o.ApplyT(func(v *HookTypeConfig) pulumi.StringOutput { return v.ConfigurationArn }).(pulumi.StringOutput)
}

// The Amazon Resource Name (ARN) of the type without version number.
func (o HookTypeConfigOutput) TypeArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HookTypeConfig) pulumi.StringPtrOutput { return v.TypeArn }).(pulumi.StringPtrOutput)
}

// The name of the type being registered.
//
// We recommend that type names adhere to the following pattern: company_or_organization::service::type.
func (o HookTypeConfigOutput) TypeName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HookTypeConfig) pulumi.StringPtrOutput { return v.TypeName }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*HookTypeConfigInput)(nil)).Elem(), &HookTypeConfig{})
	pulumi.RegisterOutputType(HookTypeConfigOutput{})
}
