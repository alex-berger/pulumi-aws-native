// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package emr

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::EMR::InstanceFleetConfig
//
// Deprecated: InstanceFleetConfig is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type InstanceFleetConfig struct {
	pulumi.CustomResourceState

	ClusterId              pulumi.StringOutput                                                 `pulumi:"clusterId"`
	InstanceFleetType      pulumi.StringOutput                                                 `pulumi:"instanceFleetType"`
	InstanceTypeConfigs    InstanceFleetConfigInstanceTypeConfigArrayOutput                    `pulumi:"instanceTypeConfigs"`
	LaunchSpecifications   InstanceFleetConfigInstanceFleetProvisioningSpecificationsPtrOutput `pulumi:"launchSpecifications"`
	Name                   pulumi.StringPtrOutput                                              `pulumi:"name"`
	TargetOnDemandCapacity pulumi.IntPtrOutput                                                 `pulumi:"targetOnDemandCapacity"`
	TargetSpotCapacity     pulumi.IntPtrOutput                                                 `pulumi:"targetSpotCapacity"`
}

// NewInstanceFleetConfig registers a new resource with the given unique name, arguments, and options.
func NewInstanceFleetConfig(ctx *pulumi.Context,
	name string, args *InstanceFleetConfigArgs, opts ...pulumi.ResourceOption) (*InstanceFleetConfig, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterId == nil {
		return nil, errors.New("invalid value for required argument 'ClusterId'")
	}
	if args.InstanceFleetType == nil {
		return nil, errors.New("invalid value for required argument 'InstanceFleetType'")
	}
	var resource InstanceFleetConfig
	err := ctx.RegisterResource("aws-native:emr:InstanceFleetConfig", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInstanceFleetConfig gets an existing InstanceFleetConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstanceFleetConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InstanceFleetConfigState, opts ...pulumi.ResourceOption) (*InstanceFleetConfig, error) {
	var resource InstanceFleetConfig
	err := ctx.ReadResource("aws-native:emr:InstanceFleetConfig", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering InstanceFleetConfig resources.
type instanceFleetConfigState struct {
}

type InstanceFleetConfigState struct {
}

func (InstanceFleetConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceFleetConfigState)(nil)).Elem()
}

type instanceFleetConfigArgs struct {
	ClusterId              string                                                      `pulumi:"clusterId"`
	InstanceFleetType      string                                                      `pulumi:"instanceFleetType"`
	InstanceTypeConfigs    []InstanceFleetConfigInstanceTypeConfig                     `pulumi:"instanceTypeConfigs"`
	LaunchSpecifications   *InstanceFleetConfigInstanceFleetProvisioningSpecifications `pulumi:"launchSpecifications"`
	Name                   *string                                                     `pulumi:"name"`
	TargetOnDemandCapacity *int                                                        `pulumi:"targetOnDemandCapacity"`
	TargetSpotCapacity     *int                                                        `pulumi:"targetSpotCapacity"`
}

// The set of arguments for constructing a InstanceFleetConfig resource.
type InstanceFleetConfigArgs struct {
	ClusterId              pulumi.StringInput
	InstanceFleetType      pulumi.StringInput
	InstanceTypeConfigs    InstanceFleetConfigInstanceTypeConfigArrayInput
	LaunchSpecifications   InstanceFleetConfigInstanceFleetProvisioningSpecificationsPtrInput
	Name                   pulumi.StringPtrInput
	TargetOnDemandCapacity pulumi.IntPtrInput
	TargetSpotCapacity     pulumi.IntPtrInput
}

func (InstanceFleetConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceFleetConfigArgs)(nil)).Elem()
}

type InstanceFleetConfigInput interface {
	pulumi.Input

	ToInstanceFleetConfigOutput() InstanceFleetConfigOutput
	ToInstanceFleetConfigOutputWithContext(ctx context.Context) InstanceFleetConfigOutput
}

func (*InstanceFleetConfig) ElementType() reflect.Type {
	return reflect.TypeOf((**InstanceFleetConfig)(nil)).Elem()
}

func (i *InstanceFleetConfig) ToInstanceFleetConfigOutput() InstanceFleetConfigOutput {
	return i.ToInstanceFleetConfigOutputWithContext(context.Background())
}

func (i *InstanceFleetConfig) ToInstanceFleetConfigOutputWithContext(ctx context.Context) InstanceFleetConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceFleetConfigOutput)
}

type InstanceFleetConfigOutput struct{ *pulumi.OutputState }

func (InstanceFleetConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InstanceFleetConfig)(nil)).Elem()
}

func (o InstanceFleetConfigOutput) ToInstanceFleetConfigOutput() InstanceFleetConfigOutput {
	return o
}

func (o InstanceFleetConfigOutput) ToInstanceFleetConfigOutputWithContext(ctx context.Context) InstanceFleetConfigOutput {
	return o
}

func (o InstanceFleetConfigOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceFleetConfig) pulumi.StringOutput { return v.ClusterId }).(pulumi.StringOutput)
}

func (o InstanceFleetConfigOutput) InstanceFleetType() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceFleetConfig) pulumi.StringOutput { return v.InstanceFleetType }).(pulumi.StringOutput)
}

func (o InstanceFleetConfigOutput) InstanceTypeConfigs() InstanceFleetConfigInstanceTypeConfigArrayOutput {
	return o.ApplyT(func(v *InstanceFleetConfig) InstanceFleetConfigInstanceTypeConfigArrayOutput {
		return v.InstanceTypeConfigs
	}).(InstanceFleetConfigInstanceTypeConfigArrayOutput)
}

func (o InstanceFleetConfigOutput) LaunchSpecifications() InstanceFleetConfigInstanceFleetProvisioningSpecificationsPtrOutput {
	return o.ApplyT(func(v *InstanceFleetConfig) InstanceFleetConfigInstanceFleetProvisioningSpecificationsPtrOutput {
		return v.LaunchSpecifications
	}).(InstanceFleetConfigInstanceFleetProvisioningSpecificationsPtrOutput)
}

func (o InstanceFleetConfigOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InstanceFleetConfig) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

func (o InstanceFleetConfigOutput) TargetOnDemandCapacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *InstanceFleetConfig) pulumi.IntPtrOutput { return v.TargetOnDemandCapacity }).(pulumi.IntPtrOutput)
}

func (o InstanceFleetConfigOutput) TargetSpotCapacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *InstanceFleetConfig) pulumi.IntPtrOutput { return v.TargetSpotCapacity }).(pulumi.IntPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceFleetConfigInput)(nil)).Elem(), &InstanceFleetConfig{})
	pulumi.RegisterOutputType(InstanceFleetConfigOutput{})
}
