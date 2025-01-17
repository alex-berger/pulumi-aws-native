// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package configuration

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Config::ConfigRule
//
// Deprecated: ConfigRule is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type ConfigRule struct {
	pulumi.CustomResourceState

	Arn                       pulumi.StringOutput      `pulumi:"arn"`
	ComplianceType            pulumi.StringPtrOutput   `pulumi:"complianceType"`
	ConfigRuleId              pulumi.StringOutput      `pulumi:"configRuleId"`
	ConfigRuleName            pulumi.StringPtrOutput   `pulumi:"configRuleName"`
	Description               pulumi.StringPtrOutput   `pulumi:"description"`
	InputParameters           pulumi.AnyOutput         `pulumi:"inputParameters"`
	MaximumExecutionFrequency pulumi.StringPtrOutput   `pulumi:"maximumExecutionFrequency"`
	Scope                     ConfigRuleScopePtrOutput `pulumi:"scope"`
	Source                    ConfigRuleSourceOutput   `pulumi:"source"`
}

// NewConfigRule registers a new resource with the given unique name, arguments, and options.
func NewConfigRule(ctx *pulumi.Context,
	name string, args *ConfigRuleArgs, opts ...pulumi.ResourceOption) (*ConfigRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Source == nil {
		return nil, errors.New("invalid value for required argument 'Source'")
	}
	var resource ConfigRule
	err := ctx.RegisterResource("aws-native:configuration:ConfigRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConfigRule gets an existing ConfigRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConfigRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConfigRuleState, opts ...pulumi.ResourceOption) (*ConfigRule, error) {
	var resource ConfigRule
	err := ctx.ReadResource("aws-native:configuration:ConfigRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ConfigRule resources.
type configRuleState struct {
}

type ConfigRuleState struct {
}

func (ConfigRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*configRuleState)(nil)).Elem()
}

type configRuleArgs struct {
	ComplianceType            *string          `pulumi:"complianceType"`
	ConfigRuleName            *string          `pulumi:"configRuleName"`
	Description               *string          `pulumi:"description"`
	InputParameters           interface{}      `pulumi:"inputParameters"`
	MaximumExecutionFrequency *string          `pulumi:"maximumExecutionFrequency"`
	Scope                     *ConfigRuleScope `pulumi:"scope"`
	Source                    ConfigRuleSource `pulumi:"source"`
}

// The set of arguments for constructing a ConfigRule resource.
type ConfigRuleArgs struct {
	ComplianceType            pulumi.StringPtrInput
	ConfigRuleName            pulumi.StringPtrInput
	Description               pulumi.StringPtrInput
	InputParameters           pulumi.Input
	MaximumExecutionFrequency pulumi.StringPtrInput
	Scope                     ConfigRuleScopePtrInput
	Source                    ConfigRuleSourceInput
}

func (ConfigRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*configRuleArgs)(nil)).Elem()
}

type ConfigRuleInput interface {
	pulumi.Input

	ToConfigRuleOutput() ConfigRuleOutput
	ToConfigRuleOutputWithContext(ctx context.Context) ConfigRuleOutput
}

func (*ConfigRule) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigRule)(nil)).Elem()
}

func (i *ConfigRule) ToConfigRuleOutput() ConfigRuleOutput {
	return i.ToConfigRuleOutputWithContext(context.Background())
}

func (i *ConfigRule) ToConfigRuleOutputWithContext(ctx context.Context) ConfigRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigRuleOutput)
}

type ConfigRuleOutput struct{ *pulumi.OutputState }

func (ConfigRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigRule)(nil)).Elem()
}

func (o ConfigRuleOutput) ToConfigRuleOutput() ConfigRuleOutput {
	return o
}

func (o ConfigRuleOutput) ToConfigRuleOutputWithContext(ctx context.Context) ConfigRuleOutput {
	return o
}

func (o ConfigRuleOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *ConfigRule) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

func (o ConfigRuleOutput) ComplianceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigRule) pulumi.StringPtrOutput { return v.ComplianceType }).(pulumi.StringPtrOutput)
}

func (o ConfigRuleOutput) ConfigRuleId() pulumi.StringOutput {
	return o.ApplyT(func(v *ConfigRule) pulumi.StringOutput { return v.ConfigRuleId }).(pulumi.StringOutput)
}

func (o ConfigRuleOutput) ConfigRuleName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigRule) pulumi.StringPtrOutput { return v.ConfigRuleName }).(pulumi.StringPtrOutput)
}

func (o ConfigRuleOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigRule) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o ConfigRuleOutput) InputParameters() pulumi.AnyOutput {
	return o.ApplyT(func(v *ConfigRule) pulumi.AnyOutput { return v.InputParameters }).(pulumi.AnyOutput)
}

func (o ConfigRuleOutput) MaximumExecutionFrequency() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigRule) pulumi.StringPtrOutput { return v.MaximumExecutionFrequency }).(pulumi.StringPtrOutput)
}

func (o ConfigRuleOutput) Scope() ConfigRuleScopePtrOutput {
	return o.ApplyT(func(v *ConfigRule) ConfigRuleScopePtrOutput { return v.Scope }).(ConfigRuleScopePtrOutput)
}

func (o ConfigRuleOutput) Source() ConfigRuleSourceOutput {
	return o.ApplyT(func(v *ConfigRule) ConfigRuleSourceOutput { return v.Source }).(ConfigRuleSourceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ConfigRuleInput)(nil)).Elem(), &ConfigRule{})
	pulumi.RegisterOutputType(ConfigRuleOutput{})
}
