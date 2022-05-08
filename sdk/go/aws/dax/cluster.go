// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dax

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::DAX::Cluster
//
// Deprecated: Cluster is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type Cluster struct {
	pulumi.CustomResourceState

	Arn                           pulumi.StringOutput              `pulumi:"arn"`
	AvailabilityZones             pulumi.StringArrayOutput         `pulumi:"availabilityZones"`
	ClusterDiscoveryEndpoint      pulumi.StringOutput              `pulumi:"clusterDiscoveryEndpoint"`
	ClusterDiscoveryEndpointURL   pulumi.StringOutput              `pulumi:"clusterDiscoveryEndpointURL"`
	ClusterEndpointEncryptionType pulumi.StringPtrOutput           `pulumi:"clusterEndpointEncryptionType"`
	ClusterName                   pulumi.StringPtrOutput           `pulumi:"clusterName"`
	Description                   pulumi.StringPtrOutput           `pulumi:"description"`
	IAMRoleARN                    pulumi.StringOutput              `pulumi:"iAMRoleARN"`
	NodeType                      pulumi.StringOutput              `pulumi:"nodeType"`
	NotificationTopicARN          pulumi.StringPtrOutput           `pulumi:"notificationTopicARN"`
	ParameterGroupName            pulumi.StringPtrOutput           `pulumi:"parameterGroupName"`
	PreferredMaintenanceWindow    pulumi.StringPtrOutput           `pulumi:"preferredMaintenanceWindow"`
	ReplicationFactor             pulumi.IntOutput                 `pulumi:"replicationFactor"`
	SSESpecification              ClusterSSESpecificationPtrOutput `pulumi:"sSESpecification"`
	SecurityGroupIds              pulumi.StringArrayOutput         `pulumi:"securityGroupIds"`
	SubnetGroupName               pulumi.StringPtrOutput           `pulumi:"subnetGroupName"`
	Tags                          pulumi.AnyOutput                 `pulumi:"tags"`
}

// NewCluster registers a new resource with the given unique name, arguments, and options.
func NewCluster(ctx *pulumi.Context,
	name string, args *ClusterArgs, opts ...pulumi.ResourceOption) (*Cluster, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.IAMRoleARN == nil {
		return nil, errors.New("invalid value for required argument 'IAMRoleARN'")
	}
	if args.NodeType == nil {
		return nil, errors.New("invalid value for required argument 'NodeType'")
	}
	if args.ReplicationFactor == nil {
		return nil, errors.New("invalid value for required argument 'ReplicationFactor'")
	}
	var resource Cluster
	err := ctx.RegisterResource("aws-native:dax:Cluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCluster gets an existing Cluster resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClusterState, opts ...pulumi.ResourceOption) (*Cluster, error) {
	var resource Cluster
	err := ctx.ReadResource("aws-native:dax:Cluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Cluster resources.
type clusterState struct {
}

type ClusterState struct {
}

func (ClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterState)(nil)).Elem()
}

type clusterArgs struct {
	AvailabilityZones             []string                 `pulumi:"availabilityZones"`
	ClusterEndpointEncryptionType *string                  `pulumi:"clusterEndpointEncryptionType"`
	ClusterName                   *string                  `pulumi:"clusterName"`
	Description                   *string                  `pulumi:"description"`
	IAMRoleARN                    string                   `pulumi:"iAMRoleARN"`
	NodeType                      string                   `pulumi:"nodeType"`
	NotificationTopicARN          *string                  `pulumi:"notificationTopicARN"`
	ParameterGroupName            *string                  `pulumi:"parameterGroupName"`
	PreferredMaintenanceWindow    *string                  `pulumi:"preferredMaintenanceWindow"`
	ReplicationFactor             int                      `pulumi:"replicationFactor"`
	SSESpecification              *ClusterSSESpecification `pulumi:"sSESpecification"`
	SecurityGroupIds              []string                 `pulumi:"securityGroupIds"`
	SubnetGroupName               *string                  `pulumi:"subnetGroupName"`
	Tags                          interface{}              `pulumi:"tags"`
}

// The set of arguments for constructing a Cluster resource.
type ClusterArgs struct {
	AvailabilityZones             pulumi.StringArrayInput
	ClusterEndpointEncryptionType pulumi.StringPtrInput
	ClusterName                   pulumi.StringPtrInput
	Description                   pulumi.StringPtrInput
	IAMRoleARN                    pulumi.StringInput
	NodeType                      pulumi.StringInput
	NotificationTopicARN          pulumi.StringPtrInput
	ParameterGroupName            pulumi.StringPtrInput
	PreferredMaintenanceWindow    pulumi.StringPtrInput
	ReplicationFactor             pulumi.IntInput
	SSESpecification              ClusterSSESpecificationPtrInput
	SecurityGroupIds              pulumi.StringArrayInput
	SubnetGroupName               pulumi.StringPtrInput
	Tags                          pulumi.Input
}

func (ClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterArgs)(nil)).Elem()
}

type ClusterInput interface {
	pulumi.Input

	ToClusterOutput() ClusterOutput
	ToClusterOutputWithContext(ctx context.Context) ClusterOutput
}

func (*Cluster) ElementType() reflect.Type {
	return reflect.TypeOf((**Cluster)(nil)).Elem()
}

func (i *Cluster) ToClusterOutput() ClusterOutput {
	return i.ToClusterOutputWithContext(context.Background())
}

func (i *Cluster) ToClusterOutputWithContext(ctx context.Context) ClusterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterOutput)
}

type ClusterOutput struct{ *pulumi.OutputState }

func (ClusterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Cluster)(nil)).Elem()
}

func (o ClusterOutput) ToClusterOutput() ClusterOutput {
	return o
}

func (o ClusterOutput) ToClusterOutputWithContext(ctx context.Context) ClusterOutput {
	return o
}

func (o ClusterOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

func (o ClusterOutput) AvailabilityZones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringArrayOutput { return v.AvailabilityZones }).(pulumi.StringArrayOutput)
}

func (o ClusterOutput) ClusterDiscoveryEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.ClusterDiscoveryEndpoint }).(pulumi.StringOutput)
}

func (o ClusterOutput) ClusterDiscoveryEndpointURL() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.ClusterDiscoveryEndpointURL }).(pulumi.StringOutput)
}

func (o ClusterOutput) ClusterEndpointEncryptionType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringPtrOutput { return v.ClusterEndpointEncryptionType }).(pulumi.StringPtrOutput)
}

func (o ClusterOutput) ClusterName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringPtrOutput { return v.ClusterName }).(pulumi.StringPtrOutput)
}

func (o ClusterOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o ClusterOutput) IAMRoleARN() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.IAMRoleARN }).(pulumi.StringOutput)
}

func (o ClusterOutput) NodeType() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.NodeType }).(pulumi.StringOutput)
}

func (o ClusterOutput) NotificationTopicARN() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringPtrOutput { return v.NotificationTopicARN }).(pulumi.StringPtrOutput)
}

func (o ClusterOutput) ParameterGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringPtrOutput { return v.ParameterGroupName }).(pulumi.StringPtrOutput)
}

func (o ClusterOutput) PreferredMaintenanceWindow() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringPtrOutput { return v.PreferredMaintenanceWindow }).(pulumi.StringPtrOutput)
}

func (o ClusterOutput) ReplicationFactor() pulumi.IntOutput {
	return o.ApplyT(func(v *Cluster) pulumi.IntOutput { return v.ReplicationFactor }).(pulumi.IntOutput)
}

func (o ClusterOutput) SSESpecification() ClusterSSESpecificationPtrOutput {
	return o.ApplyT(func(v *Cluster) ClusterSSESpecificationPtrOutput { return v.SSESpecification }).(ClusterSSESpecificationPtrOutput)
}

func (o ClusterOutput) SecurityGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringArrayOutput { return v.SecurityGroupIds }).(pulumi.StringArrayOutput)
}

func (o ClusterOutput) SubnetGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringPtrOutput { return v.SubnetGroupName }).(pulumi.StringPtrOutput)
}

func (o ClusterOutput) Tags() pulumi.AnyOutput {
	return o.ApplyT(func(v *Cluster) pulumi.AnyOutput { return v.Tags }).(pulumi.AnyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterInput)(nil)).Elem(), &Cluster{})
	pulumi.RegisterOutputType(ClusterOutput{})
}
