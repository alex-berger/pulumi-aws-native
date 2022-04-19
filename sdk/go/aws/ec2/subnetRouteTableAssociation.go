// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::EC2::SubnetRouteTableAssociation
type SubnetRouteTableAssociation struct {
	pulumi.CustomResourceState

	RouteTableId pulumi.StringOutput `pulumi:"routeTableId"`
	SubnetId     pulumi.StringOutput `pulumi:"subnetId"`
}

// NewSubnetRouteTableAssociation registers a new resource with the given unique name, arguments, and options.
func NewSubnetRouteTableAssociation(ctx *pulumi.Context,
	name string, args *SubnetRouteTableAssociationArgs, opts ...pulumi.ResourceOption) (*SubnetRouteTableAssociation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RouteTableId == nil {
		return nil, errors.New("invalid value for required argument 'RouteTableId'")
	}
	if args.SubnetId == nil {
		return nil, errors.New("invalid value for required argument 'SubnetId'")
	}
	var resource SubnetRouteTableAssociation
	err := ctx.RegisterResource("aws-native:ec2:SubnetRouteTableAssociation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSubnetRouteTableAssociation gets an existing SubnetRouteTableAssociation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSubnetRouteTableAssociation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SubnetRouteTableAssociationState, opts ...pulumi.ResourceOption) (*SubnetRouteTableAssociation, error) {
	var resource SubnetRouteTableAssociation
	err := ctx.ReadResource("aws-native:ec2:SubnetRouteTableAssociation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SubnetRouteTableAssociation resources.
type subnetRouteTableAssociationState struct {
}

type SubnetRouteTableAssociationState struct {
}

func (SubnetRouteTableAssociationState) ElementType() reflect.Type {
	return reflect.TypeOf((*subnetRouteTableAssociationState)(nil)).Elem()
}

type subnetRouteTableAssociationArgs struct {
	RouteTableId string `pulumi:"routeTableId"`
	SubnetId     string `pulumi:"subnetId"`
}

// The set of arguments for constructing a SubnetRouteTableAssociation resource.
type SubnetRouteTableAssociationArgs struct {
	RouteTableId pulumi.StringInput
	SubnetId     pulumi.StringInput
}

func (SubnetRouteTableAssociationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*subnetRouteTableAssociationArgs)(nil)).Elem()
}

type SubnetRouteTableAssociationInput interface {
	pulumi.Input

	ToSubnetRouteTableAssociationOutput() SubnetRouteTableAssociationOutput
	ToSubnetRouteTableAssociationOutputWithContext(ctx context.Context) SubnetRouteTableAssociationOutput
}

func (*SubnetRouteTableAssociation) ElementType() reflect.Type {
	return reflect.TypeOf((**SubnetRouteTableAssociation)(nil)).Elem()
}

func (i *SubnetRouteTableAssociation) ToSubnetRouteTableAssociationOutput() SubnetRouteTableAssociationOutput {
	return i.ToSubnetRouteTableAssociationOutputWithContext(context.Background())
}

func (i *SubnetRouteTableAssociation) ToSubnetRouteTableAssociationOutputWithContext(ctx context.Context) SubnetRouteTableAssociationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubnetRouteTableAssociationOutput)
}

type SubnetRouteTableAssociationOutput struct{ *pulumi.OutputState }

func (SubnetRouteTableAssociationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SubnetRouteTableAssociation)(nil)).Elem()
}

func (o SubnetRouteTableAssociationOutput) ToSubnetRouteTableAssociationOutput() SubnetRouteTableAssociationOutput {
	return o
}

func (o SubnetRouteTableAssociationOutput) ToSubnetRouteTableAssociationOutputWithContext(ctx context.Context) SubnetRouteTableAssociationOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SubnetRouteTableAssociationInput)(nil)).Elem(), &SubnetRouteTableAssociation{})
	pulumi.RegisterOutputType(SubnetRouteTableAssociationOutput{})
}
