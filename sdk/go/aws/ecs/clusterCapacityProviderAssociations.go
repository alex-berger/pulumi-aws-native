// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ecs

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-clustercapacityproviderassociations.html
type ClusterCapacityProviderAssociations struct {
	pulumi.CustomResourceState

	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-clustercapacityproviderassociations.html#cfn-ecs-clustercapacityproviderassociations-capacityproviders
	CapacityProviders pulumi.StringArrayOutput `pulumi:"capacityProviders"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-clustercapacityproviderassociations.html#cfn-ecs-clustercapacityproviderassociations-cluster
	Cluster pulumi.StringOutput `pulumi:"cluster"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-clustercapacityproviderassociations.html#cfn-ecs-clustercapacityproviderassociations-defaultcapacityproviderstrategy
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
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-clustercapacityproviderassociations.html#cfn-ecs-clustercapacityproviderassociations-capacityproviders
	CapacityProviders []string `pulumi:"capacityProviders"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-clustercapacityproviderassociations.html#cfn-ecs-clustercapacityproviderassociations-cluster
	Cluster string `pulumi:"cluster"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-clustercapacityproviderassociations.html#cfn-ecs-clustercapacityproviderassociations-defaultcapacityproviderstrategy
	DefaultCapacityProviderStrategy []ClusterCapacityProviderAssociationsCapacityProviderStrategy `pulumi:"defaultCapacityProviderStrategy"`
}

// The set of arguments for constructing a ClusterCapacityProviderAssociations resource.
type ClusterCapacityProviderAssociationsArgs struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-clustercapacityproviderassociations.html#cfn-ecs-clustercapacityproviderassociations-capacityproviders
	CapacityProviders pulumi.StringArrayInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-clustercapacityproviderassociations.html#cfn-ecs-clustercapacityproviderassociations-cluster
	Cluster pulumi.StringInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-clustercapacityproviderassociations.html#cfn-ecs-clustercapacityproviderassociations-defaultcapacityproviderstrategy
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
	return reflect.TypeOf((*ClusterCapacityProviderAssociations)(nil))
}

func (i *ClusterCapacityProviderAssociations) ToClusterCapacityProviderAssociationsOutput() ClusterCapacityProviderAssociationsOutput {
	return i.ToClusterCapacityProviderAssociationsOutputWithContext(context.Background())
}

func (i *ClusterCapacityProviderAssociations) ToClusterCapacityProviderAssociationsOutputWithContext(ctx context.Context) ClusterCapacityProviderAssociationsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterCapacityProviderAssociationsOutput)
}

type ClusterCapacityProviderAssociationsOutput struct{ *pulumi.OutputState }

func (ClusterCapacityProviderAssociationsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterCapacityProviderAssociations)(nil))
}

func (o ClusterCapacityProviderAssociationsOutput) ToClusterCapacityProviderAssociationsOutput() ClusterCapacityProviderAssociationsOutput {
	return o
}

func (o ClusterCapacityProviderAssociationsOutput) ToClusterCapacityProviderAssociationsOutputWithContext(ctx context.Context) ClusterCapacityProviderAssociationsOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ClusterCapacityProviderAssociationsOutput{})
}