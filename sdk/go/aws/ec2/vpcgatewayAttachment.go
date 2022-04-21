// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::EC2::VPCGatewayAttachment
//
// Deprecated: VPCGatewayAttachment is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type VPCGatewayAttachment struct {
	pulumi.CustomResourceState

	InternetGatewayId pulumi.StringPtrOutput `pulumi:"internetGatewayId"`
	VpcId             pulumi.StringOutput    `pulumi:"vpcId"`
	VpnGatewayId      pulumi.StringPtrOutput `pulumi:"vpnGatewayId"`
}

// NewVPCGatewayAttachment registers a new resource with the given unique name, arguments, and options.
func NewVPCGatewayAttachment(ctx *pulumi.Context,
	name string, args *VPCGatewayAttachmentArgs, opts ...pulumi.ResourceOption) (*VPCGatewayAttachment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.VpcId == nil {
		return nil, errors.New("invalid value for required argument 'VpcId'")
	}
	var resource VPCGatewayAttachment
	err := ctx.RegisterResource("aws-native:ec2:VPCGatewayAttachment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVPCGatewayAttachment gets an existing VPCGatewayAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVPCGatewayAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VPCGatewayAttachmentState, opts ...pulumi.ResourceOption) (*VPCGatewayAttachment, error) {
	var resource VPCGatewayAttachment
	err := ctx.ReadResource("aws-native:ec2:VPCGatewayAttachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VPCGatewayAttachment resources.
type vpcgatewayAttachmentState struct {
}

type VPCGatewayAttachmentState struct {
}

func (VPCGatewayAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcgatewayAttachmentState)(nil)).Elem()
}

type vpcgatewayAttachmentArgs struct {
	InternetGatewayId *string `pulumi:"internetGatewayId"`
	VpcId             string  `pulumi:"vpcId"`
	VpnGatewayId      *string `pulumi:"vpnGatewayId"`
}

// The set of arguments for constructing a VPCGatewayAttachment resource.
type VPCGatewayAttachmentArgs struct {
	InternetGatewayId pulumi.StringPtrInput
	VpcId             pulumi.StringInput
	VpnGatewayId      pulumi.StringPtrInput
}

func (VPCGatewayAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcgatewayAttachmentArgs)(nil)).Elem()
}

type VPCGatewayAttachmentInput interface {
	pulumi.Input

	ToVPCGatewayAttachmentOutput() VPCGatewayAttachmentOutput
	ToVPCGatewayAttachmentOutputWithContext(ctx context.Context) VPCGatewayAttachmentOutput
}

func (*VPCGatewayAttachment) ElementType() reflect.Type {
	return reflect.TypeOf((**VPCGatewayAttachment)(nil)).Elem()
}

func (i *VPCGatewayAttachment) ToVPCGatewayAttachmentOutput() VPCGatewayAttachmentOutput {
	return i.ToVPCGatewayAttachmentOutputWithContext(context.Background())
}

func (i *VPCGatewayAttachment) ToVPCGatewayAttachmentOutputWithContext(ctx context.Context) VPCGatewayAttachmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VPCGatewayAttachmentOutput)
}

type VPCGatewayAttachmentOutput struct{ *pulumi.OutputState }

func (VPCGatewayAttachmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VPCGatewayAttachment)(nil)).Elem()
}

func (o VPCGatewayAttachmentOutput) ToVPCGatewayAttachmentOutput() VPCGatewayAttachmentOutput {
	return o
}

func (o VPCGatewayAttachmentOutput) ToVPCGatewayAttachmentOutputWithContext(ctx context.Context) VPCGatewayAttachmentOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VPCGatewayAttachmentInput)(nil)).Elem(), &VPCGatewayAttachment{})
	pulumi.RegisterOutputType(VPCGatewayAttachmentOutput{})
}
