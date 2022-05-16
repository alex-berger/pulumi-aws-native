// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package redshift

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Redshift::ClusterSecurityGroup
//
// Deprecated: ClusterSecurityGroup is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type ClusterSecurityGroup struct {
	pulumi.CustomResourceState

	Description pulumi.StringOutput                `pulumi:"description"`
	Tags        ClusterSecurityGroupTagArrayOutput `pulumi:"tags"`
}

// NewClusterSecurityGroup registers a new resource with the given unique name, arguments, and options.
func NewClusterSecurityGroup(ctx *pulumi.Context,
	name string, args *ClusterSecurityGroupArgs, opts ...pulumi.ResourceOption) (*ClusterSecurityGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Description == nil {
		return nil, errors.New("invalid value for required argument 'Description'")
	}
	var resource ClusterSecurityGroup
	err := ctx.RegisterResource("aws-native:redshift:ClusterSecurityGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetClusterSecurityGroup gets an existing ClusterSecurityGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetClusterSecurityGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClusterSecurityGroupState, opts ...pulumi.ResourceOption) (*ClusterSecurityGroup, error) {
	var resource ClusterSecurityGroup
	err := ctx.ReadResource("aws-native:redshift:ClusterSecurityGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ClusterSecurityGroup resources.
type clusterSecurityGroupState struct {
}

type ClusterSecurityGroupState struct {
}

func (ClusterSecurityGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterSecurityGroupState)(nil)).Elem()
}

type clusterSecurityGroupArgs struct {
	Description string                    `pulumi:"description"`
	Tags        []ClusterSecurityGroupTag `pulumi:"tags"`
}

// The set of arguments for constructing a ClusterSecurityGroup resource.
type ClusterSecurityGroupArgs struct {
	Description pulumi.StringInput
	Tags        ClusterSecurityGroupTagArrayInput
}

func (ClusterSecurityGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterSecurityGroupArgs)(nil)).Elem()
}

type ClusterSecurityGroupInput interface {
	pulumi.Input

	ToClusterSecurityGroupOutput() ClusterSecurityGroupOutput
	ToClusterSecurityGroupOutputWithContext(ctx context.Context) ClusterSecurityGroupOutput
}

func (*ClusterSecurityGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterSecurityGroup)(nil)).Elem()
}

func (i *ClusterSecurityGroup) ToClusterSecurityGroupOutput() ClusterSecurityGroupOutput {
	return i.ToClusterSecurityGroupOutputWithContext(context.Background())
}

func (i *ClusterSecurityGroup) ToClusterSecurityGroupOutputWithContext(ctx context.Context) ClusterSecurityGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterSecurityGroupOutput)
}

type ClusterSecurityGroupOutput struct{ *pulumi.OutputState }

func (ClusterSecurityGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterSecurityGroup)(nil)).Elem()
}

func (o ClusterSecurityGroupOutput) ToClusterSecurityGroupOutput() ClusterSecurityGroupOutput {
	return o
}

func (o ClusterSecurityGroupOutput) ToClusterSecurityGroupOutputWithContext(ctx context.Context) ClusterSecurityGroupOutput {
	return o
}

func (o ClusterSecurityGroupOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterSecurityGroup) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

func (o ClusterSecurityGroupOutput) Tags() ClusterSecurityGroupTagArrayOutput {
	return o.ApplyT(func(v *ClusterSecurityGroup) ClusterSecurityGroupTagArrayOutput { return v.Tags }).(ClusterSecurityGroupTagArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterSecurityGroupInput)(nil)).Elem(), &ClusterSecurityGroup{})
	pulumi.RegisterOutputType(ClusterSecurityGroupOutput{})
}
