// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package eks

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource schema for AWS::EKS::Nodegroup
type Nodegroup struct {
	pulumi.CustomResourceState

	// The AMI type for your node group.
	AmiType pulumi.StringPtrOutput `pulumi:"amiType"`
	Arn     pulumi.StringOutput    `pulumi:"arn"`
	// The capacity type of your managed node group.
	CapacityType pulumi.StringPtrOutput `pulumi:"capacityType"`
	// Name of the cluster to create the node group in.
	ClusterName pulumi.StringOutput `pulumi:"clusterName"`
	// The root device disk size (in GiB) for your node group instances.
	DiskSize pulumi.IntPtrOutput `pulumi:"diskSize"`
	// Force the update if the existing node group's pods are unable to be drained due to a pod disruption budget issue.
	ForceUpdateEnabled pulumi.BoolPtrOutput `pulumi:"forceUpdateEnabled"`
	// Specify the instance types for a node group.
	InstanceTypes pulumi.StringArrayOutput `pulumi:"instanceTypes"`
	// The Kubernetes labels to be applied to the nodes in the node group when they are created.
	Labels pulumi.AnyOutput `pulumi:"labels"`
	// An object representing a node group's launch template specification.
	LaunchTemplate NodegroupLaunchTemplateSpecificationPtrOutput `pulumi:"launchTemplate"`
	// The Amazon Resource Name (ARN) of the IAM role to associate with your node group.
	NodeRole pulumi.StringOutput `pulumi:"nodeRole"`
	// The unique name to give your node group.
	NodegroupName pulumi.StringPtrOutput `pulumi:"nodegroupName"`
	// The AMI version of the Amazon EKS-optimized AMI to use with your node group.
	ReleaseVersion pulumi.StringPtrOutput `pulumi:"releaseVersion"`
	// The remote access (SSH) configuration to use with your node group.
	RemoteAccess NodegroupRemoteAccessPtrOutput `pulumi:"remoteAccess"`
	// The scaling configuration details for the Auto Scaling group that is created for your node group.
	ScalingConfig NodegroupScalingConfigPtrOutput `pulumi:"scalingConfig"`
	// The subnets to use for the Auto Scaling group that is created for your node group.
	Subnets pulumi.StringArrayOutput `pulumi:"subnets"`
	// The metadata, as key-value pairs, to apply to the node group to assist with categorization and organization. Follows same schema as Labels for consistency.
	Tags pulumi.AnyOutput `pulumi:"tags"`
	// The Kubernetes taints to be applied to the nodes in the node group when they are created.
	Taints NodegroupTaintArrayOutput `pulumi:"taints"`
	// The node group update configuration.
	UpdateConfig NodegroupUpdateConfigPtrOutput `pulumi:"updateConfig"`
	// The Kubernetes version to use for your managed nodes.
	Version pulumi.StringPtrOutput `pulumi:"version"`
}

// NewNodegroup registers a new resource with the given unique name, arguments, and options.
func NewNodegroup(ctx *pulumi.Context,
	name string, args *NodegroupArgs, opts ...pulumi.ResourceOption) (*Nodegroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterName == nil {
		return nil, errors.New("invalid value for required argument 'ClusterName'")
	}
	if args.NodeRole == nil {
		return nil, errors.New("invalid value for required argument 'NodeRole'")
	}
	if args.Subnets == nil {
		return nil, errors.New("invalid value for required argument 'Subnets'")
	}
	var resource Nodegroup
	err := ctx.RegisterResource("aws-native:eks:Nodegroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNodegroup gets an existing Nodegroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNodegroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NodegroupState, opts ...pulumi.ResourceOption) (*Nodegroup, error) {
	var resource Nodegroup
	err := ctx.ReadResource("aws-native:eks:Nodegroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Nodegroup resources.
type nodegroupState struct {
}

type NodegroupState struct {
}

func (NodegroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*nodegroupState)(nil)).Elem()
}

type nodegroupArgs struct {
	// The AMI type for your node group.
	AmiType *string `pulumi:"amiType"`
	// The capacity type of your managed node group.
	CapacityType *string `pulumi:"capacityType"`
	// Name of the cluster to create the node group in.
	ClusterName string `pulumi:"clusterName"`
	// The root device disk size (in GiB) for your node group instances.
	DiskSize *int `pulumi:"diskSize"`
	// Force the update if the existing node group's pods are unable to be drained due to a pod disruption budget issue.
	ForceUpdateEnabled *bool `pulumi:"forceUpdateEnabled"`
	// Specify the instance types for a node group.
	InstanceTypes []string `pulumi:"instanceTypes"`
	// The Kubernetes labels to be applied to the nodes in the node group when they are created.
	Labels interface{} `pulumi:"labels"`
	// An object representing a node group's launch template specification.
	LaunchTemplate *NodegroupLaunchTemplateSpecification `pulumi:"launchTemplate"`
	// The Amazon Resource Name (ARN) of the IAM role to associate with your node group.
	NodeRole string `pulumi:"nodeRole"`
	// The unique name to give your node group.
	NodegroupName *string `pulumi:"nodegroupName"`
	// The AMI version of the Amazon EKS-optimized AMI to use with your node group.
	ReleaseVersion *string `pulumi:"releaseVersion"`
	// The remote access (SSH) configuration to use with your node group.
	RemoteAccess *NodegroupRemoteAccess `pulumi:"remoteAccess"`
	// The scaling configuration details for the Auto Scaling group that is created for your node group.
	ScalingConfig *NodegroupScalingConfig `pulumi:"scalingConfig"`
	// The subnets to use for the Auto Scaling group that is created for your node group.
	Subnets []string `pulumi:"subnets"`
	// The metadata, as key-value pairs, to apply to the node group to assist with categorization and organization. Follows same schema as Labels for consistency.
	Tags interface{} `pulumi:"tags"`
	// The Kubernetes taints to be applied to the nodes in the node group when they are created.
	Taints []NodegroupTaint `pulumi:"taints"`
	// The node group update configuration.
	UpdateConfig *NodegroupUpdateConfig `pulumi:"updateConfig"`
	// The Kubernetes version to use for your managed nodes.
	Version *string `pulumi:"version"`
}

// The set of arguments for constructing a Nodegroup resource.
type NodegroupArgs struct {
	// The AMI type for your node group.
	AmiType pulumi.StringPtrInput
	// The capacity type of your managed node group.
	CapacityType pulumi.StringPtrInput
	// Name of the cluster to create the node group in.
	ClusterName pulumi.StringInput
	// The root device disk size (in GiB) for your node group instances.
	DiskSize pulumi.IntPtrInput
	// Force the update if the existing node group's pods are unable to be drained due to a pod disruption budget issue.
	ForceUpdateEnabled pulumi.BoolPtrInput
	// Specify the instance types for a node group.
	InstanceTypes pulumi.StringArrayInput
	// The Kubernetes labels to be applied to the nodes in the node group when they are created.
	Labels pulumi.Input
	// An object representing a node group's launch template specification.
	LaunchTemplate NodegroupLaunchTemplateSpecificationPtrInput
	// The Amazon Resource Name (ARN) of the IAM role to associate with your node group.
	NodeRole pulumi.StringInput
	// The unique name to give your node group.
	NodegroupName pulumi.StringPtrInput
	// The AMI version of the Amazon EKS-optimized AMI to use with your node group.
	ReleaseVersion pulumi.StringPtrInput
	// The remote access (SSH) configuration to use with your node group.
	RemoteAccess NodegroupRemoteAccessPtrInput
	// The scaling configuration details for the Auto Scaling group that is created for your node group.
	ScalingConfig NodegroupScalingConfigPtrInput
	// The subnets to use for the Auto Scaling group that is created for your node group.
	Subnets pulumi.StringArrayInput
	// The metadata, as key-value pairs, to apply to the node group to assist with categorization and organization. Follows same schema as Labels for consistency.
	Tags pulumi.Input
	// The Kubernetes taints to be applied to the nodes in the node group when they are created.
	Taints NodegroupTaintArrayInput
	// The node group update configuration.
	UpdateConfig NodegroupUpdateConfigPtrInput
	// The Kubernetes version to use for your managed nodes.
	Version pulumi.StringPtrInput
}

func (NodegroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*nodegroupArgs)(nil)).Elem()
}

type NodegroupInput interface {
	pulumi.Input

	ToNodegroupOutput() NodegroupOutput
	ToNodegroupOutputWithContext(ctx context.Context) NodegroupOutput
}

func (*Nodegroup) ElementType() reflect.Type {
	return reflect.TypeOf((**Nodegroup)(nil)).Elem()
}

func (i *Nodegroup) ToNodegroupOutput() NodegroupOutput {
	return i.ToNodegroupOutputWithContext(context.Background())
}

func (i *Nodegroup) ToNodegroupOutputWithContext(ctx context.Context) NodegroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NodegroupOutput)
}

type NodegroupOutput struct{ *pulumi.OutputState }

func (NodegroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Nodegroup)(nil)).Elem()
}

func (o NodegroupOutput) ToNodegroupOutput() NodegroupOutput {
	return o
}

func (o NodegroupOutput) ToNodegroupOutputWithContext(ctx context.Context) NodegroupOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NodegroupInput)(nil)).Elem(), &Nodegroup{})
	pulumi.RegisterOutputType(NodegroupOutput{})
}
