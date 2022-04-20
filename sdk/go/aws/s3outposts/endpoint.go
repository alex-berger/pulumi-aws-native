// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package s3outposts

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type Definition for AWS::S3Outposts::Endpoint
type Endpoint struct {
	pulumi.CustomResourceState

	// The type of access for the on-premise network connectivity for the Outpost endpoint. To access endpoint from an on-premises network, you must specify the access type and provide the customer owned Ipv4 pool.
	AccessType EndpointAccessTypePtrOutput `pulumi:"accessType"`
	// The Amazon Resource Name (ARN) of the endpoint.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The VPC CIDR committed by this endpoint.
	CidrBlock pulumi.StringOutput `pulumi:"cidrBlock"`
	// The time the endpoint was created.
	CreationTime pulumi.StringOutput `pulumi:"creationTime"`
	// The ID of the customer-owned IPv4 pool for the Endpoint. IP addresses will be allocated from this pool for the endpoint.
	CustomerOwnedIpv4Pool pulumi.StringPtrOutput `pulumi:"customerOwnedIpv4Pool"`
	// The network interfaces of the endpoint.
	NetworkInterfaces EndpointNetworkInterfaceArrayOutput `pulumi:"networkInterfaces"`
	// The id of the customer outpost on which the bucket resides.
	OutpostId pulumi.StringOutput `pulumi:"outpostId"`
	// The ID of the security group to use with the endpoint.
	SecurityGroupId pulumi.StringOutput  `pulumi:"securityGroupId"`
	Status          EndpointStatusOutput `pulumi:"status"`
	// The ID of the subnet in the selected VPC. The subnet must belong to the Outpost.
	SubnetId pulumi.StringOutput `pulumi:"subnetId"`
}

// NewEndpoint registers a new resource with the given unique name, arguments, and options.
func NewEndpoint(ctx *pulumi.Context,
	name string, args *EndpointArgs, opts ...pulumi.ResourceOption) (*Endpoint, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.OutpostId == nil {
		return nil, errors.New("invalid value for required argument 'OutpostId'")
	}
	if args.SecurityGroupId == nil {
		return nil, errors.New("invalid value for required argument 'SecurityGroupId'")
	}
	if args.SubnetId == nil {
		return nil, errors.New("invalid value for required argument 'SubnetId'")
	}
	var resource Endpoint
	err := ctx.RegisterResource("aws-native:s3outposts:Endpoint", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEndpoint gets an existing Endpoint resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEndpoint(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EndpointState, opts ...pulumi.ResourceOption) (*Endpoint, error) {
	var resource Endpoint
	err := ctx.ReadResource("aws-native:s3outposts:Endpoint", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Endpoint resources.
type endpointState struct {
}

type EndpointState struct {
}

func (EndpointState) ElementType() reflect.Type {
	return reflect.TypeOf((*endpointState)(nil)).Elem()
}

type endpointArgs struct {
	// The type of access for the on-premise network connectivity for the Outpost endpoint. To access endpoint from an on-premises network, you must specify the access type and provide the customer owned Ipv4 pool.
	AccessType *EndpointAccessType `pulumi:"accessType"`
	// The ID of the customer-owned IPv4 pool for the Endpoint. IP addresses will be allocated from this pool for the endpoint.
	CustomerOwnedIpv4Pool *string `pulumi:"customerOwnedIpv4Pool"`
	// The id of the customer outpost on which the bucket resides.
	OutpostId string `pulumi:"outpostId"`
	// The ID of the security group to use with the endpoint.
	SecurityGroupId string `pulumi:"securityGroupId"`
	// The ID of the subnet in the selected VPC. The subnet must belong to the Outpost.
	SubnetId string `pulumi:"subnetId"`
}

// The set of arguments for constructing a Endpoint resource.
type EndpointArgs struct {
	// The type of access for the on-premise network connectivity for the Outpost endpoint. To access endpoint from an on-premises network, you must specify the access type and provide the customer owned Ipv4 pool.
	AccessType EndpointAccessTypePtrInput
	// The ID of the customer-owned IPv4 pool for the Endpoint. IP addresses will be allocated from this pool for the endpoint.
	CustomerOwnedIpv4Pool pulumi.StringPtrInput
	// The id of the customer outpost on which the bucket resides.
	OutpostId pulumi.StringInput
	// The ID of the security group to use with the endpoint.
	SecurityGroupId pulumi.StringInput
	// The ID of the subnet in the selected VPC. The subnet must belong to the Outpost.
	SubnetId pulumi.StringInput
}

func (EndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*endpointArgs)(nil)).Elem()
}

type EndpointInput interface {
	pulumi.Input

	ToEndpointOutput() EndpointOutput
	ToEndpointOutputWithContext(ctx context.Context) EndpointOutput
}

func (*Endpoint) ElementType() reflect.Type {
	return reflect.TypeOf((**Endpoint)(nil)).Elem()
}

func (i *Endpoint) ToEndpointOutput() EndpointOutput {
	return i.ToEndpointOutputWithContext(context.Background())
}

func (i *Endpoint) ToEndpointOutputWithContext(ctx context.Context) EndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointOutput)
}

type EndpointOutput struct{ *pulumi.OutputState }

func (EndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Endpoint)(nil)).Elem()
}

func (o EndpointOutput) ToEndpointOutput() EndpointOutput {
	return o
}

func (o EndpointOutput) ToEndpointOutputWithContext(ctx context.Context) EndpointOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EndpointInput)(nil)).Elem(), &Endpoint{})
	pulumi.RegisterOutputType(EndpointOutput{})
}
