// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ssmincidents

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource type definition for AWS::SSMIncidents::ReplicationSet
type ReplicationSet struct {
	pulumi.CustomResourceState

	// The ARN of the ReplicationSet.
	Arn               pulumi.StringOutput  `pulumi:"arn"`
	DeletionProtected pulumi.BoolPtrOutput `pulumi:"deletionProtected"`
	// The ReplicationSet configuration.
	Regions ReplicationSetReplicationRegionArrayOutput `pulumi:"regions"`
}

// NewReplicationSet registers a new resource with the given unique name, arguments, and options.
func NewReplicationSet(ctx *pulumi.Context,
	name string, args *ReplicationSetArgs, opts ...pulumi.ResourceOption) (*ReplicationSet, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Regions == nil {
		return nil, errors.New("invalid value for required argument 'Regions'")
	}
	var resource ReplicationSet
	err := ctx.RegisterResource("aws-native:ssmincidents:ReplicationSet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetReplicationSet gets an existing ReplicationSet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetReplicationSet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ReplicationSetState, opts ...pulumi.ResourceOption) (*ReplicationSet, error) {
	var resource ReplicationSet
	err := ctx.ReadResource("aws-native:ssmincidents:ReplicationSet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ReplicationSet resources.
type replicationSetState struct {
}

type ReplicationSetState struct {
}

func (ReplicationSetState) ElementType() reflect.Type {
	return reflect.TypeOf((*replicationSetState)(nil)).Elem()
}

type replicationSetArgs struct {
	DeletionProtected *bool `pulumi:"deletionProtected"`
	// The ReplicationSet configuration.
	Regions []ReplicationSetReplicationRegion `pulumi:"regions"`
}

// The set of arguments for constructing a ReplicationSet resource.
type ReplicationSetArgs struct {
	DeletionProtected pulumi.BoolPtrInput
	// The ReplicationSet configuration.
	Regions ReplicationSetReplicationRegionArrayInput
}

func (ReplicationSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*replicationSetArgs)(nil)).Elem()
}

type ReplicationSetInput interface {
	pulumi.Input

	ToReplicationSetOutput() ReplicationSetOutput
	ToReplicationSetOutputWithContext(ctx context.Context) ReplicationSetOutput
}

func (*ReplicationSet) ElementType() reflect.Type {
	return reflect.TypeOf((**ReplicationSet)(nil)).Elem()
}

func (i *ReplicationSet) ToReplicationSetOutput() ReplicationSetOutput {
	return i.ToReplicationSetOutputWithContext(context.Background())
}

func (i *ReplicationSet) ToReplicationSetOutputWithContext(ctx context.Context) ReplicationSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplicationSetOutput)
}

type ReplicationSetOutput struct{ *pulumi.OutputState }

func (ReplicationSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReplicationSet)(nil)).Elem()
}

func (o ReplicationSetOutput) ToReplicationSetOutput() ReplicationSetOutput {
	return o
}

func (o ReplicationSetOutput) ToReplicationSetOutputWithContext(ctx context.Context) ReplicationSetOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ReplicationSetInput)(nil)).Elem(), &ReplicationSet{})
	pulumi.RegisterOutputType(ReplicationSetOutput{})
}
