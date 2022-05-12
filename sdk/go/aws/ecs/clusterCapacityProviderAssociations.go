// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ecs

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Associate a set of ECS Capacity Providers with a specified ECS Cluster
type ClusterCapacityProviderAssociations struct {
	pulumi.CustomResourceState

	CapacityProviders               pulumi.StringArrayOutput                                               `pulumi:"capacityProviders"`
	Cluster                         pulumi.StringOutput                                                    `pulumi:"cluster"`
	DefaultCapacityProviderStrategy ClusterCapacityProviderAssociationsCapacityProviderStrategyArrayOutput `pulumi:"defaultCapacityProviderStrategy"`
}

// NewClusterCapacityProviderAssociations registers a new resource with the given unique name, arguments, and options.
func NewClusterCapacityProviderAssociations(ctx *pulumi.Context,
	name string, args *ClusterCapacityProviderAssociationsArgs, opts ...pulumi.ResourceOption) (*ClusterCapacityProviderAssociations, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CapacityProviders == nil {
		return nil, errors.New("invalid value for required argument 'CapacityProviders'")
	}
	if args.Cluster == nil {
		return nil, errors.New("invalid value for required argument 'Cluster'")
	}
	if args.DefaultCapacityProviderStrategy == nil {
		return nil, errors.New("invalid value for required argument 'DefaultCapacityProviderStrategy'")
	}
	var resource ClusterCapacityProviderAssociations
	err := ctx.RegisterResource("aws-native:ecs:ClusterCapacityProviderAssociations", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetClusterCapacityProviderAssociations gets an existing ClusterCapacityProviderAssociations resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetClusterCapacityProviderAssociations(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClusterCapacityProviderAssociationsState, opts ...pulumi.ResourceOption) (*ClusterCapacityProviderAssociations, error) {
	var resource ClusterCapacityProviderAssociations
	err := ctx.ReadResource("aws-native:ecs:ClusterCapacityProviderAssociations", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ClusterCapacityProviderAssociations resources.
type clusterCapacityProviderAssociationsState struct {
}

type ClusterCapacityProviderAssociationsState struct {
}

func (ClusterCapacityProviderAssociationsState) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterCapacityProviderAssociationsState)(nil)).Elem()
}

type clusterCapacityProviderAssociationsArgs struct {
	CapacityProviders               []string                                                      `pulumi:"capacityProviders"`
	Cluster                         string                                                        `pulumi:"cluster"`
	DefaultCapacityProviderStrategy []ClusterCapacityProviderAssociationsCapacityProviderStrategy `pulumi:"defaultCapacityProviderStrategy"`
}

// The set of arguments for constructing a ClusterCapacityProviderAssociations resource.
type ClusterCapacityProviderAssociationsArgs struct {
	CapacityProviders               pulumi.StringArrayInput
	Cluster                         pulumi.StringInput
	DefaultCapacityProviderStrategy ClusterCapacityProviderAssociationsCapacityProviderStrategyArrayInput
}

func (ClusterCapacityProviderAssociationsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterCapacityProviderAssociationsArgs)(nil)).Elem()
}

type ClusterCapacityProviderAssociationsInput interface {
	pulumi.Input

	ToClusterCapacityProviderAssociationsOutput() ClusterCapacityProviderAssociationsOutput
	ToClusterCapacityProviderAssociationsOutputWithContext(ctx context.Context) ClusterCapacityProviderAssociationsOutput
}

func (*ClusterCapacityProviderAssociations) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterCapacityProviderAssociations)(nil)).Elem()
}

func (i *ClusterCapacityProviderAssociations) ToClusterCapacityProviderAssociationsOutput() ClusterCapacityProviderAssociationsOutput {
	return i.ToClusterCapacityProviderAssociationsOutputWithContext(context.Background())
}

func (i *ClusterCapacityProviderAssociations) ToClusterCapacityProviderAssociationsOutputWithContext(ctx context.Context) ClusterCapacityProviderAssociationsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterCapacityProviderAssociationsOutput)
}

type ClusterCapacityProviderAssociationsOutput struct{ *pulumi.OutputState }

func (ClusterCapacityProviderAssociationsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterCapacityProviderAssociations)(nil)).Elem()
}

func (o ClusterCapacityProviderAssociationsOutput) ToClusterCapacityProviderAssociationsOutput() ClusterCapacityProviderAssociationsOutput {
	return o
}

func (o ClusterCapacityProviderAssociationsOutput) ToClusterCapacityProviderAssociationsOutputWithContext(ctx context.Context) ClusterCapacityProviderAssociationsOutput {
	return o
}

func (o ClusterCapacityProviderAssociationsOutput) CapacityProviders() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ClusterCapacityProviderAssociations) pulumi.StringArrayOutput { return v.CapacityProviders }).(pulumi.StringArrayOutput)
}

func (o ClusterCapacityProviderAssociationsOutput) Cluster() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterCapacityProviderAssociations) pulumi.StringOutput { return v.Cluster }).(pulumi.StringOutput)
}

func (o ClusterCapacityProviderAssociationsOutput) DefaultCapacityProviderStrategy() ClusterCapacityProviderAssociationsCapacityProviderStrategyArrayOutput {
	return o.ApplyT(func(v *ClusterCapacityProviderAssociations) ClusterCapacityProviderAssociationsCapacityProviderStrategyArrayOutput {
		return v.DefaultCapacityProviderStrategy
	}).(ClusterCapacityProviderAssociationsCapacityProviderStrategyArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterCapacityProviderAssociationsInput)(nil)).Elem(), &ClusterCapacityProviderAssociations{})
	pulumi.RegisterOutputType(ClusterCapacityProviderAssociationsOutput{})
}
