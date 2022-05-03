// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::EC2::ClientVpnTargetNetworkAssociation
//
// Deprecated: ClientVpnTargetNetworkAssociation is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type ClientVpnTargetNetworkAssociation struct {
	pulumi.CustomResourceState

	ClientVpnEndpointId pulumi.StringOutput `pulumi:"clientVpnEndpointId"`
	SubnetId            pulumi.StringOutput `pulumi:"subnetId"`
}

// NewClientVpnTargetNetworkAssociation registers a new resource with the given unique name, arguments, and options.
func NewClientVpnTargetNetworkAssociation(ctx *pulumi.Context,
	name string, args *ClientVpnTargetNetworkAssociationArgs, opts ...pulumi.ResourceOption) (*ClientVpnTargetNetworkAssociation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClientVpnEndpointId == nil {
		return nil, errors.New("invalid value for required argument 'ClientVpnEndpointId'")
	}
	if args.SubnetId == nil {
		return nil, errors.New("invalid value for required argument 'SubnetId'")
	}
	var resource ClientVpnTargetNetworkAssociation
	err := ctx.RegisterResource("aws-native:ec2:ClientVpnTargetNetworkAssociation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetClientVpnTargetNetworkAssociation gets an existing ClientVpnTargetNetworkAssociation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetClientVpnTargetNetworkAssociation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClientVpnTargetNetworkAssociationState, opts ...pulumi.ResourceOption) (*ClientVpnTargetNetworkAssociation, error) {
	var resource ClientVpnTargetNetworkAssociation
	err := ctx.ReadResource("aws-native:ec2:ClientVpnTargetNetworkAssociation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ClientVpnTargetNetworkAssociation resources.
type clientVpnTargetNetworkAssociationState struct {
}

type ClientVpnTargetNetworkAssociationState struct {
}

func (ClientVpnTargetNetworkAssociationState) ElementType() reflect.Type {
	return reflect.TypeOf((*clientVpnTargetNetworkAssociationState)(nil)).Elem()
}

type clientVpnTargetNetworkAssociationArgs struct {
	ClientVpnEndpointId string `pulumi:"clientVpnEndpointId"`
	SubnetId            string `pulumi:"subnetId"`
}

// The set of arguments for constructing a ClientVpnTargetNetworkAssociation resource.
type ClientVpnTargetNetworkAssociationArgs struct {
	ClientVpnEndpointId pulumi.StringInput
	SubnetId            pulumi.StringInput
}

func (ClientVpnTargetNetworkAssociationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clientVpnTargetNetworkAssociationArgs)(nil)).Elem()
}

type ClientVpnTargetNetworkAssociationInput interface {
	pulumi.Input

	ToClientVpnTargetNetworkAssociationOutput() ClientVpnTargetNetworkAssociationOutput
	ToClientVpnTargetNetworkAssociationOutputWithContext(ctx context.Context) ClientVpnTargetNetworkAssociationOutput
}

func (*ClientVpnTargetNetworkAssociation) ElementType() reflect.Type {
	return reflect.TypeOf((**ClientVpnTargetNetworkAssociation)(nil)).Elem()
}

func (i *ClientVpnTargetNetworkAssociation) ToClientVpnTargetNetworkAssociationOutput() ClientVpnTargetNetworkAssociationOutput {
	return i.ToClientVpnTargetNetworkAssociationOutputWithContext(context.Background())
}

func (i *ClientVpnTargetNetworkAssociation) ToClientVpnTargetNetworkAssociationOutputWithContext(ctx context.Context) ClientVpnTargetNetworkAssociationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientVpnTargetNetworkAssociationOutput)
}

type ClientVpnTargetNetworkAssociationOutput struct{ *pulumi.OutputState }

func (ClientVpnTargetNetworkAssociationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClientVpnTargetNetworkAssociation)(nil)).Elem()
}

func (o ClientVpnTargetNetworkAssociationOutput) ToClientVpnTargetNetworkAssociationOutput() ClientVpnTargetNetworkAssociationOutput {
	return o
}

func (o ClientVpnTargetNetworkAssociationOutput) ToClientVpnTargetNetworkAssociationOutputWithContext(ctx context.Context) ClientVpnTargetNetworkAssociationOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ClientVpnTargetNetworkAssociationInput)(nil)).Elem(), &ClientVpnTargetNetworkAssociation{})
	pulumi.RegisterOutputType(ClientVpnTargetNetworkAssociationOutput{})
}
