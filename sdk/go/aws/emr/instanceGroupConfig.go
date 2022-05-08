// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package emr

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::EMR::InstanceGroupConfig
//
// Deprecated: InstanceGroupConfig is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type InstanceGroupConfig struct {
	pulumi.CustomResourceState

	AutoScalingPolicy InstanceGroupConfigAutoScalingPolicyPtrOutput `pulumi:"autoScalingPolicy"`
	BidPrice          pulumi.StringPtrOutput                        `pulumi:"bidPrice"`
	Configurations    InstanceGroupConfigConfigurationArrayOutput   `pulumi:"configurations"`
	CustomAmiId       pulumi.StringPtrOutput                        `pulumi:"customAmiId"`
	EbsConfiguration  InstanceGroupConfigEbsConfigurationPtrOutput  `pulumi:"ebsConfiguration"`
	InstanceCount     pulumi.IntOutput                              `pulumi:"instanceCount"`
	InstanceRole      pulumi.StringOutput                           `pulumi:"instanceRole"`
	InstanceType      pulumi.StringOutput                           `pulumi:"instanceType"`
	JobFlowId         pulumi.StringOutput                           `pulumi:"jobFlowId"`
	Market            pulumi.StringPtrOutput                        `pulumi:"market"`
	Name              pulumi.StringPtrOutput                        `pulumi:"name"`
}

// NewInstanceGroupConfig registers a new resource with the given unique name, arguments, and options.
func NewInstanceGroupConfig(ctx *pulumi.Context,
	name string, args *InstanceGroupConfigArgs, opts ...pulumi.ResourceOption) (*InstanceGroupConfig, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceCount == nil {
		return nil, errors.New("invalid value for required argument 'InstanceCount'")
	}
	if args.InstanceRole == nil {
		return nil, errors.New("invalid value for required argument 'InstanceRole'")
	}
	if args.InstanceType == nil {
		return nil, errors.New("invalid value for required argument 'InstanceType'")
	}
	if args.JobFlowId == nil {
		return nil, errors.New("invalid value for required argument 'JobFlowId'")
	}
	var resource InstanceGroupConfig
	err := ctx.RegisterResource("aws-native:emr:InstanceGroupConfig", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInstanceGroupConfig gets an existing InstanceGroupConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstanceGroupConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InstanceGroupConfigState, opts ...pulumi.ResourceOption) (*InstanceGroupConfig, error) {
	var resource InstanceGroupConfig
	err := ctx.ReadResource("aws-native:emr:InstanceGroupConfig", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering InstanceGroupConfig resources.
type instanceGroupConfigState struct {
}

type InstanceGroupConfigState struct {
}

func (InstanceGroupConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceGroupConfigState)(nil)).Elem()
}

type instanceGroupConfigArgs struct {
	AutoScalingPolicy *InstanceGroupConfigAutoScalingPolicy `pulumi:"autoScalingPolicy"`
	BidPrice          *string                               `pulumi:"bidPrice"`
	Configurations    []InstanceGroupConfigConfiguration    `pulumi:"configurations"`
	CustomAmiId       *string                               `pulumi:"customAmiId"`
	EbsConfiguration  *InstanceGroupConfigEbsConfiguration  `pulumi:"ebsConfiguration"`
	InstanceCount     int                                   `pulumi:"instanceCount"`
	InstanceRole      string                                `pulumi:"instanceRole"`
	InstanceType      string                                `pulumi:"instanceType"`
	JobFlowId         string                                `pulumi:"jobFlowId"`
	Market            *string                               `pulumi:"market"`
	Name              *string                               `pulumi:"name"`
}

// The set of arguments for constructing a InstanceGroupConfig resource.
type InstanceGroupConfigArgs struct {
	AutoScalingPolicy InstanceGroupConfigAutoScalingPolicyPtrInput
	BidPrice          pulumi.StringPtrInput
	Configurations    InstanceGroupConfigConfigurationArrayInput
	CustomAmiId       pulumi.StringPtrInput
	EbsConfiguration  InstanceGroupConfigEbsConfigurationPtrInput
	InstanceCount     pulumi.IntInput
	InstanceRole      pulumi.StringInput
	InstanceType      pulumi.StringInput
	JobFlowId         pulumi.StringInput
	Market            pulumi.StringPtrInput
	Name              pulumi.StringPtrInput
}

func (InstanceGroupConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceGroupConfigArgs)(nil)).Elem()
}

type InstanceGroupConfigInput interface {
	pulumi.Input

	ToInstanceGroupConfigOutput() InstanceGroupConfigOutput
	ToInstanceGroupConfigOutputWithContext(ctx context.Context) InstanceGroupConfigOutput
}

func (*InstanceGroupConfig) ElementType() reflect.Type {
	return reflect.TypeOf((**InstanceGroupConfig)(nil)).Elem()
}

func (i *InstanceGroupConfig) ToInstanceGroupConfigOutput() InstanceGroupConfigOutput {
	return i.ToInstanceGroupConfigOutputWithContext(context.Background())
}

func (i *InstanceGroupConfig) ToInstanceGroupConfigOutputWithContext(ctx context.Context) InstanceGroupConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceGroupConfigOutput)
}

type InstanceGroupConfigOutput struct{ *pulumi.OutputState }

func (InstanceGroupConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InstanceGroupConfig)(nil)).Elem()
}

func (o InstanceGroupConfigOutput) ToInstanceGroupConfigOutput() InstanceGroupConfigOutput {
	return o
}

func (o InstanceGroupConfigOutput) ToInstanceGroupConfigOutputWithContext(ctx context.Context) InstanceGroupConfigOutput {
	return o
}

func (o InstanceGroupConfigOutput) AutoScalingPolicy() InstanceGroupConfigAutoScalingPolicyPtrOutput {
	return o.ApplyT(func(v *InstanceGroupConfig) InstanceGroupConfigAutoScalingPolicyPtrOutput { return v.AutoScalingPolicy }).(InstanceGroupConfigAutoScalingPolicyPtrOutput)
}

func (o InstanceGroupConfigOutput) BidPrice() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InstanceGroupConfig) pulumi.StringPtrOutput { return v.BidPrice }).(pulumi.StringPtrOutput)
}

func (o InstanceGroupConfigOutput) Configurations() InstanceGroupConfigConfigurationArrayOutput {
	return o.ApplyT(func(v *InstanceGroupConfig) InstanceGroupConfigConfigurationArrayOutput { return v.Configurations }).(InstanceGroupConfigConfigurationArrayOutput)
}

func (o InstanceGroupConfigOutput) CustomAmiId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InstanceGroupConfig) pulumi.StringPtrOutput { return v.CustomAmiId }).(pulumi.StringPtrOutput)
}

func (o InstanceGroupConfigOutput) EbsConfiguration() InstanceGroupConfigEbsConfigurationPtrOutput {
	return o.ApplyT(func(v *InstanceGroupConfig) InstanceGroupConfigEbsConfigurationPtrOutput { return v.EbsConfiguration }).(InstanceGroupConfigEbsConfigurationPtrOutput)
}

func (o InstanceGroupConfigOutput) InstanceCount() pulumi.IntOutput {
	return o.ApplyT(func(v *InstanceGroupConfig) pulumi.IntOutput { return v.InstanceCount }).(pulumi.IntOutput)
}

func (o InstanceGroupConfigOutput) InstanceRole() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceGroupConfig) pulumi.StringOutput { return v.InstanceRole }).(pulumi.StringOutput)
}

func (o InstanceGroupConfigOutput) InstanceType() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceGroupConfig) pulumi.StringOutput { return v.InstanceType }).(pulumi.StringOutput)
}

func (o InstanceGroupConfigOutput) JobFlowId() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceGroupConfig) pulumi.StringOutput { return v.JobFlowId }).(pulumi.StringOutput)
}

func (o InstanceGroupConfigOutput) Market() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InstanceGroupConfig) pulumi.StringPtrOutput { return v.Market }).(pulumi.StringPtrOutput)
}

func (o InstanceGroupConfigOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InstanceGroupConfig) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceGroupConfigInput)(nil)).Elem(), &InstanceGroupConfig{})
	pulumi.RegisterOutputType(InstanceGroupConfigOutput{})
}
