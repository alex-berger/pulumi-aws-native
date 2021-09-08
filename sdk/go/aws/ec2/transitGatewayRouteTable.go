// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewayroutetable.html
type TransitGatewayRouteTable struct {
	pulumi.CustomResourceState

	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewayroutetable.html#cfn-ec2-transitgatewayroutetable-tags
	Tags aws.TagArrayOutput `pulumi:"tags"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewayroutetable.html#cfn-ec2-transitgatewayroutetable-transitgatewayid
	TransitGatewayId pulumi.StringOutput `pulumi:"transitGatewayId"`
}

// NewTransitGatewayRouteTable registers a new resource with the given unique name, arguments, and options.
func NewTransitGatewayRouteTable(ctx *pulumi.Context,
	name string, args *TransitGatewayRouteTableArgs, opts ...pulumi.ResourceOption) (*TransitGatewayRouteTable, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.TransitGatewayId == nil {
		return nil, errors.New("invalid value for required argument 'TransitGatewayId'")
	}
	var resource TransitGatewayRouteTable
	err := ctx.RegisterResource("aws-native:ec2:TransitGatewayRouteTable", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTransitGatewayRouteTable gets an existing TransitGatewayRouteTable resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTransitGatewayRouteTable(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TransitGatewayRouteTableState, opts ...pulumi.ResourceOption) (*TransitGatewayRouteTable, error) {
	var resource TransitGatewayRouteTable
	err := ctx.ReadResource("aws-native:ec2:TransitGatewayRouteTable", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TransitGatewayRouteTable resources.
type transitGatewayRouteTableState struct {
}

type TransitGatewayRouteTableState struct {
}

func (TransitGatewayRouteTableState) ElementType() reflect.Type {
	return reflect.TypeOf((*transitGatewayRouteTableState)(nil)).Elem()
}

type transitGatewayRouteTableArgs struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewayroutetable.html#cfn-ec2-transitgatewayroutetable-tags
	Tags []aws.Tag `pulumi:"tags"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewayroutetable.html#cfn-ec2-transitgatewayroutetable-transitgatewayid
	TransitGatewayId string `pulumi:"transitGatewayId"`
}

// The set of arguments for constructing a TransitGatewayRouteTable resource.
type TransitGatewayRouteTableArgs struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewayroutetable.html#cfn-ec2-transitgatewayroutetable-tags
	Tags aws.TagArrayInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewayroutetable.html#cfn-ec2-transitgatewayroutetable-transitgatewayid
	TransitGatewayId pulumi.StringInput
}

func (TransitGatewayRouteTableArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*transitGatewayRouteTableArgs)(nil)).Elem()
}

type TransitGatewayRouteTableInput interface {
	pulumi.Input

	ToTransitGatewayRouteTableOutput() TransitGatewayRouteTableOutput
	ToTransitGatewayRouteTableOutputWithContext(ctx context.Context) TransitGatewayRouteTableOutput
}

func (*TransitGatewayRouteTable) ElementType() reflect.Type {
	return reflect.TypeOf((*TransitGatewayRouteTable)(nil))
}

func (i *TransitGatewayRouteTable) ToTransitGatewayRouteTableOutput() TransitGatewayRouteTableOutput {
	return i.ToTransitGatewayRouteTableOutputWithContext(context.Background())
}

func (i *TransitGatewayRouteTable) ToTransitGatewayRouteTableOutputWithContext(ctx context.Context) TransitGatewayRouteTableOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TransitGatewayRouteTableOutput)
}

type TransitGatewayRouteTableOutput struct{ *pulumi.OutputState }

func (TransitGatewayRouteTableOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TransitGatewayRouteTable)(nil))
}

func (o TransitGatewayRouteTableOutput) ToTransitGatewayRouteTableOutput() TransitGatewayRouteTableOutput {
	return o
}

func (o TransitGatewayRouteTableOutput) ToTransitGatewayRouteTableOutputWithContext(ctx context.Context) TransitGatewayRouteTableOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(TransitGatewayRouteTableOutput{})
}