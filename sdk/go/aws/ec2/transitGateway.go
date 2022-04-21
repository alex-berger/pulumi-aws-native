// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::EC2::TransitGateway
type TransitGateway struct {
	pulumi.CustomResourceState

	AmazonSideAsn                  pulumi.IntPtrOutput          `pulumi:"amazonSideAsn"`
	AssociationDefaultRouteTableId pulumi.StringPtrOutput       `pulumi:"associationDefaultRouteTableId"`
	AutoAcceptSharedAttachments    pulumi.StringPtrOutput       `pulumi:"autoAcceptSharedAttachments"`
	DefaultRouteTableAssociation   pulumi.StringPtrOutput       `pulumi:"defaultRouteTableAssociation"`
	DefaultRouteTablePropagation   pulumi.StringPtrOutput       `pulumi:"defaultRouteTablePropagation"`
	Description                    pulumi.StringPtrOutput       `pulumi:"description"`
	DnsSupport                     pulumi.StringPtrOutput       `pulumi:"dnsSupport"`
	MulticastSupport               pulumi.StringPtrOutput       `pulumi:"multicastSupport"`
	PropagationDefaultRouteTableId pulumi.StringPtrOutput       `pulumi:"propagationDefaultRouteTableId"`
	Tags                           TransitGatewayTagArrayOutput `pulumi:"tags"`
	TransitGatewayCidrBlocks       pulumi.StringArrayOutput     `pulumi:"transitGatewayCidrBlocks"`
	VpnEcmpSupport                 pulumi.StringPtrOutput       `pulumi:"vpnEcmpSupport"`
}

// NewTransitGateway registers a new resource with the given unique name, arguments, and options.
func NewTransitGateway(ctx *pulumi.Context,
	name string, args *TransitGatewayArgs, opts ...pulumi.ResourceOption) (*TransitGateway, error) {
	if args == nil {
		args = &TransitGatewayArgs{}
	}

	var resource TransitGateway
	err := ctx.RegisterResource("aws-native:ec2:TransitGateway", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTransitGateway gets an existing TransitGateway resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTransitGateway(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TransitGatewayState, opts ...pulumi.ResourceOption) (*TransitGateway, error) {
	var resource TransitGateway
	err := ctx.ReadResource("aws-native:ec2:TransitGateway", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TransitGateway resources.
type transitGatewayState struct {
}

type TransitGatewayState struct {
}

func (TransitGatewayState) ElementType() reflect.Type {
	return reflect.TypeOf((*transitGatewayState)(nil)).Elem()
}

type transitGatewayArgs struct {
	AmazonSideAsn                  *int                `pulumi:"amazonSideAsn"`
	AssociationDefaultRouteTableId *string             `pulumi:"associationDefaultRouteTableId"`
	AutoAcceptSharedAttachments    *string             `pulumi:"autoAcceptSharedAttachments"`
	DefaultRouteTableAssociation   *string             `pulumi:"defaultRouteTableAssociation"`
	DefaultRouteTablePropagation   *string             `pulumi:"defaultRouteTablePropagation"`
	Description                    *string             `pulumi:"description"`
	DnsSupport                     *string             `pulumi:"dnsSupport"`
	MulticastSupport               *string             `pulumi:"multicastSupport"`
	PropagationDefaultRouteTableId *string             `pulumi:"propagationDefaultRouteTableId"`
	Tags                           []TransitGatewayTag `pulumi:"tags"`
	TransitGatewayCidrBlocks       []string            `pulumi:"transitGatewayCidrBlocks"`
	VpnEcmpSupport                 *string             `pulumi:"vpnEcmpSupport"`
}

// The set of arguments for constructing a TransitGateway resource.
type TransitGatewayArgs struct {
	AmazonSideAsn                  pulumi.IntPtrInput
	AssociationDefaultRouteTableId pulumi.StringPtrInput
	AutoAcceptSharedAttachments    pulumi.StringPtrInput
	DefaultRouteTableAssociation   pulumi.StringPtrInput
	DefaultRouteTablePropagation   pulumi.StringPtrInput
	Description                    pulumi.StringPtrInput
	DnsSupport                     pulumi.StringPtrInput
	MulticastSupport               pulumi.StringPtrInput
	PropagationDefaultRouteTableId pulumi.StringPtrInput
	Tags                           TransitGatewayTagArrayInput
	TransitGatewayCidrBlocks       pulumi.StringArrayInput
	VpnEcmpSupport                 pulumi.StringPtrInput
}

func (TransitGatewayArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*transitGatewayArgs)(nil)).Elem()
}

type TransitGatewayInput interface {
	pulumi.Input

	ToTransitGatewayOutput() TransitGatewayOutput
	ToTransitGatewayOutputWithContext(ctx context.Context) TransitGatewayOutput
}

func (*TransitGateway) ElementType() reflect.Type {
	return reflect.TypeOf((**TransitGateway)(nil)).Elem()
}

func (i *TransitGateway) ToTransitGatewayOutput() TransitGatewayOutput {
	return i.ToTransitGatewayOutputWithContext(context.Background())
}

func (i *TransitGateway) ToTransitGatewayOutputWithContext(ctx context.Context) TransitGatewayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TransitGatewayOutput)
}

type TransitGatewayOutput struct{ *pulumi.OutputState }

func (TransitGatewayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TransitGateway)(nil)).Elem()
}

func (o TransitGatewayOutput) ToTransitGatewayOutput() TransitGatewayOutput {
	return o
}

func (o TransitGatewayOutput) ToTransitGatewayOutputWithContext(ctx context.Context) TransitGatewayOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TransitGatewayInput)(nil)).Elem(), &TransitGateway{})
	pulumi.RegisterOutputType(TransitGatewayOutput{})
}
