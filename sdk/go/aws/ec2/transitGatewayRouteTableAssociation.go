// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewayroutetableassociation.html
type TransitGatewayRouteTableAssociation struct {
	pulumi.CustomResourceState

	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewayroutetableassociation.html#cfn-ec2-transitgatewayroutetableassociation-transitgatewayattachmentid
	TransitGatewayAttachmentId pulumi.StringOutput `pulumi:"transitGatewayAttachmentId"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewayroutetableassociation.html#cfn-ec2-transitgatewayroutetableassociation-transitgatewayroutetableid
	TransitGatewayRouteTableId pulumi.StringOutput `pulumi:"transitGatewayRouteTableId"`
}

// NewTransitGatewayRouteTableAssociation registers a new resource with the given unique name, arguments, and options.
func NewTransitGatewayRouteTableAssociation(ctx *pulumi.Context,
	name string, args *TransitGatewayRouteTableAssociationArgs, opts ...pulumi.ResourceOption) (*TransitGatewayRouteTableAssociation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.TransitGatewayAttachmentId == nil {
		return nil, errors.New("invalid value for required argument 'TransitGatewayAttachmentId'")
	}
	if args.TransitGatewayRouteTableId == nil {
		return nil, errors.New("invalid value for required argument 'TransitGatewayRouteTableId'")
	}
	var resource TransitGatewayRouteTableAssociation
	err := ctx.RegisterResource("aws-native:ec2:TransitGatewayRouteTableAssociation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTransitGatewayRouteTableAssociation gets an existing TransitGatewayRouteTableAssociation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTransitGatewayRouteTableAssociation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TransitGatewayRouteTableAssociationState, opts ...pulumi.ResourceOption) (*TransitGatewayRouteTableAssociation, error) {
	var resource TransitGatewayRouteTableAssociation
	err := ctx.ReadResource("aws-native:ec2:TransitGatewayRouteTableAssociation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TransitGatewayRouteTableAssociation resources.
type transitGatewayRouteTableAssociationState struct {
}

type TransitGatewayRouteTableAssociationState struct {
}

func (TransitGatewayRouteTableAssociationState) ElementType() reflect.Type {
	return reflect.TypeOf((*transitGatewayRouteTableAssociationState)(nil)).Elem()
}

type transitGatewayRouteTableAssociationArgs struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewayroutetableassociation.html#cfn-ec2-transitgatewayroutetableassociation-transitgatewayattachmentid
	TransitGatewayAttachmentId string `pulumi:"transitGatewayAttachmentId"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewayroutetableassociation.html#cfn-ec2-transitgatewayroutetableassociation-transitgatewayroutetableid
	TransitGatewayRouteTableId string `pulumi:"transitGatewayRouteTableId"`
}

// The set of arguments for constructing a TransitGatewayRouteTableAssociation resource.
type TransitGatewayRouteTableAssociationArgs struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewayroutetableassociation.html#cfn-ec2-transitgatewayroutetableassociation-transitgatewayattachmentid
	TransitGatewayAttachmentId pulumi.StringInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewayroutetableassociation.html#cfn-ec2-transitgatewayroutetableassociation-transitgatewayroutetableid
	TransitGatewayRouteTableId pulumi.StringInput
}

func (TransitGatewayRouteTableAssociationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*transitGatewayRouteTableAssociationArgs)(nil)).Elem()
}

type TransitGatewayRouteTableAssociationInput interface {
	pulumi.Input

	ToTransitGatewayRouteTableAssociationOutput() TransitGatewayRouteTableAssociationOutput
	ToTransitGatewayRouteTableAssociationOutputWithContext(ctx context.Context) TransitGatewayRouteTableAssociationOutput
}

func (*TransitGatewayRouteTableAssociation) ElementType() reflect.Type {
	return reflect.TypeOf((*TransitGatewayRouteTableAssociation)(nil))
}

func (i *TransitGatewayRouteTableAssociation) ToTransitGatewayRouteTableAssociationOutput() TransitGatewayRouteTableAssociationOutput {
	return i.ToTransitGatewayRouteTableAssociationOutputWithContext(context.Background())
}

func (i *TransitGatewayRouteTableAssociation) ToTransitGatewayRouteTableAssociationOutputWithContext(ctx context.Context) TransitGatewayRouteTableAssociationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TransitGatewayRouteTableAssociationOutput)
}

type TransitGatewayRouteTableAssociationOutput struct{ *pulumi.OutputState }

func (TransitGatewayRouteTableAssociationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TransitGatewayRouteTableAssociation)(nil))
}

func (o TransitGatewayRouteTableAssociationOutput) ToTransitGatewayRouteTableAssociationOutput() TransitGatewayRouteTableAssociationOutput {
	return o
}

func (o TransitGatewayRouteTableAssociationOutput) ToTransitGatewayRouteTableAssociationOutputWithContext(ctx context.Context) TransitGatewayRouteTableAssociationOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(TransitGatewayRouteTableAssociationOutput{})
}