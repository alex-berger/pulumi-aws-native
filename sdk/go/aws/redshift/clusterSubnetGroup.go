// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package redshift

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Redshift::ClusterSubnetGroup
//
// Deprecated: ClusterSubnetGroup is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type ClusterSubnetGroup struct {
	pulumi.CustomResourceState

	Description pulumi.StringOutput              `pulumi:"description"`
	SubnetIds   pulumi.StringArrayOutput         `pulumi:"subnetIds"`
	Tags        ClusterSubnetGroupTagArrayOutput `pulumi:"tags"`
}

// NewClusterSubnetGroup registers a new resource with the given unique name, arguments, and options.
func NewClusterSubnetGroup(ctx *pulumi.Context,
	name string, args *ClusterSubnetGroupArgs, opts ...pulumi.ResourceOption) (*ClusterSubnetGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Description == nil {
		return nil, errors.New("invalid value for required argument 'Description'")
	}
	if args.SubnetIds == nil {
		return nil, errors.New("invalid value for required argument 'SubnetIds'")
	}
	var resource ClusterSubnetGroup
	err := ctx.RegisterResource("aws-native:redshift:ClusterSubnetGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetClusterSubnetGroup gets an existing ClusterSubnetGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetClusterSubnetGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClusterSubnetGroupState, opts ...pulumi.ResourceOption) (*ClusterSubnetGroup, error) {
	var resource ClusterSubnetGroup
	err := ctx.ReadResource("aws-native:redshift:ClusterSubnetGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ClusterSubnetGroup resources.
type clusterSubnetGroupState struct {
}

type ClusterSubnetGroupState struct {
}

func (ClusterSubnetGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterSubnetGroupState)(nil)).Elem()
}

type clusterSubnetGroupArgs struct {
	Description string                  `pulumi:"description"`
	SubnetIds   []string                `pulumi:"subnetIds"`
	Tags        []ClusterSubnetGroupTag `pulumi:"tags"`
}

// The set of arguments for constructing a ClusterSubnetGroup resource.
type ClusterSubnetGroupArgs struct {
	Description pulumi.StringInput
	SubnetIds   pulumi.StringArrayInput
	Tags        ClusterSubnetGroupTagArrayInput
}

func (ClusterSubnetGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterSubnetGroupArgs)(nil)).Elem()
}

type ClusterSubnetGroupInput interface {
	pulumi.Input

	ToClusterSubnetGroupOutput() ClusterSubnetGroupOutput
	ToClusterSubnetGroupOutputWithContext(ctx context.Context) ClusterSubnetGroupOutput
}

func (*ClusterSubnetGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterSubnetGroup)(nil)).Elem()
}

func (i *ClusterSubnetGroup) ToClusterSubnetGroupOutput() ClusterSubnetGroupOutput {
	return i.ToClusterSubnetGroupOutputWithContext(context.Background())
}

func (i *ClusterSubnetGroup) ToClusterSubnetGroupOutputWithContext(ctx context.Context) ClusterSubnetGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterSubnetGroupOutput)
}

type ClusterSubnetGroupOutput struct{ *pulumi.OutputState }

func (ClusterSubnetGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterSubnetGroup)(nil)).Elem()
}

func (o ClusterSubnetGroupOutput) ToClusterSubnetGroupOutput() ClusterSubnetGroupOutput {
	return o
}

func (o ClusterSubnetGroupOutput) ToClusterSubnetGroupOutputWithContext(ctx context.Context) ClusterSubnetGroupOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterSubnetGroupInput)(nil)).Elem(), &ClusterSubnetGroup{})
	pulumi.RegisterOutputType(ClusterSubnetGroupOutput{})
}
