// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package networkmanager

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The AWS::NetworkManager::CustomerGatewayAssociation type associates a customer gateway with a device and optionally, with a link.
type CustomerGatewayAssociation struct {
	pulumi.CustomResourceState

	// The Amazon Resource Name (ARN) of the customer gateway.
	CustomerGatewayArn pulumi.StringOutput `pulumi:"customerGatewayArn"`
	// The ID of the device
	DeviceId pulumi.StringOutput `pulumi:"deviceId"`
	// The ID of the global network.
	GlobalNetworkId pulumi.StringOutput `pulumi:"globalNetworkId"`
	// The ID of the link
	LinkId pulumi.StringPtrOutput `pulumi:"linkId"`
}

// NewCustomerGatewayAssociation registers a new resource with the given unique name, arguments, and options.
func NewCustomerGatewayAssociation(ctx *pulumi.Context,
	name string, args *CustomerGatewayAssociationArgs, opts ...pulumi.ResourceOption) (*CustomerGatewayAssociation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CustomerGatewayArn == nil {
		return nil, errors.New("invalid value for required argument 'CustomerGatewayArn'")
	}
	if args.DeviceId == nil {
		return nil, errors.New("invalid value for required argument 'DeviceId'")
	}
	if args.GlobalNetworkId == nil {
		return nil, errors.New("invalid value for required argument 'GlobalNetworkId'")
	}
	var resource CustomerGatewayAssociation
	err := ctx.RegisterResource("aws-native:networkmanager:CustomerGatewayAssociation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCustomerGatewayAssociation gets an existing CustomerGatewayAssociation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCustomerGatewayAssociation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CustomerGatewayAssociationState, opts ...pulumi.ResourceOption) (*CustomerGatewayAssociation, error) {
	var resource CustomerGatewayAssociation
	err := ctx.ReadResource("aws-native:networkmanager:CustomerGatewayAssociation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CustomerGatewayAssociation resources.
type customerGatewayAssociationState struct {
}

type CustomerGatewayAssociationState struct {
}

func (CustomerGatewayAssociationState) ElementType() reflect.Type {
	return reflect.TypeOf((*customerGatewayAssociationState)(nil)).Elem()
}

type customerGatewayAssociationArgs struct {
	// The Amazon Resource Name (ARN) of the customer gateway.
	CustomerGatewayArn string `pulumi:"customerGatewayArn"`
	// The ID of the device
	DeviceId string `pulumi:"deviceId"`
	// The ID of the global network.
	GlobalNetworkId string `pulumi:"globalNetworkId"`
	// The ID of the link
	LinkId *string `pulumi:"linkId"`
}

// The set of arguments for constructing a CustomerGatewayAssociation resource.
type CustomerGatewayAssociationArgs struct {
	// The Amazon Resource Name (ARN) of the customer gateway.
	CustomerGatewayArn pulumi.StringInput
	// The ID of the device
	DeviceId pulumi.StringInput
	// The ID of the global network.
	GlobalNetworkId pulumi.StringInput
	// The ID of the link
	LinkId pulumi.StringPtrInput
}

func (CustomerGatewayAssociationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*customerGatewayAssociationArgs)(nil)).Elem()
}

type CustomerGatewayAssociationInput interface {
	pulumi.Input

	ToCustomerGatewayAssociationOutput() CustomerGatewayAssociationOutput
	ToCustomerGatewayAssociationOutputWithContext(ctx context.Context) CustomerGatewayAssociationOutput
}

func (*CustomerGatewayAssociation) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomerGatewayAssociation)(nil)).Elem()
}

func (i *CustomerGatewayAssociation) ToCustomerGatewayAssociationOutput() CustomerGatewayAssociationOutput {
	return i.ToCustomerGatewayAssociationOutputWithContext(context.Background())
}

func (i *CustomerGatewayAssociation) ToCustomerGatewayAssociationOutputWithContext(ctx context.Context) CustomerGatewayAssociationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomerGatewayAssociationOutput)
}

type CustomerGatewayAssociationOutput struct{ *pulumi.OutputState }

func (CustomerGatewayAssociationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomerGatewayAssociation)(nil)).Elem()
}

func (o CustomerGatewayAssociationOutput) ToCustomerGatewayAssociationOutput() CustomerGatewayAssociationOutput {
	return o
}

func (o CustomerGatewayAssociationOutput) ToCustomerGatewayAssociationOutputWithContext(ctx context.Context) CustomerGatewayAssociationOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CustomerGatewayAssociationInput)(nil)).Elem(), &CustomerGatewayAssociation{})
	pulumi.RegisterOutputType(CustomerGatewayAssociationOutput{})
}
