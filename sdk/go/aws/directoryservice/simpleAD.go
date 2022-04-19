// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package directoryservice

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::DirectoryService::SimpleAD
//
// Deprecated: SimpleAD is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type SimpleAD struct {
	pulumi.CustomResourceState

	Alias          pulumi.StringOutput       `pulumi:"alias"`
	CreateAlias    pulumi.BoolPtrOutput      `pulumi:"createAlias"`
	Description    pulumi.StringPtrOutput    `pulumi:"description"`
	DnsIpAddresses pulumi.StringArrayOutput  `pulumi:"dnsIpAddresses"`
	EnableSso      pulumi.BoolPtrOutput      `pulumi:"enableSso"`
	Name           pulumi.StringOutput       `pulumi:"name"`
	Password       pulumi.StringOutput       `pulumi:"password"`
	ShortName      pulumi.StringPtrOutput    `pulumi:"shortName"`
	Size           pulumi.StringOutput       `pulumi:"size"`
	VpcSettings    SimpleADVpcSettingsOutput `pulumi:"vpcSettings"`
}

// NewSimpleAD registers a new resource with the given unique name, arguments, and options.
func NewSimpleAD(ctx *pulumi.Context,
	name string, args *SimpleADArgs, opts ...pulumi.ResourceOption) (*SimpleAD, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Password == nil {
		return nil, errors.New("invalid value for required argument 'Password'")
	}
	if args.Size == nil {
		return nil, errors.New("invalid value for required argument 'Size'")
	}
	if args.VpcSettings == nil {
		return nil, errors.New("invalid value for required argument 'VpcSettings'")
	}
	var resource SimpleAD
	err := ctx.RegisterResource("aws-native:directoryservice:SimpleAD", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSimpleAD gets an existing SimpleAD resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSimpleAD(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SimpleADState, opts ...pulumi.ResourceOption) (*SimpleAD, error) {
	var resource SimpleAD
	err := ctx.ReadResource("aws-native:directoryservice:SimpleAD", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SimpleAD resources.
type simpleADState struct {
}

type SimpleADState struct {
}

func (SimpleADState) ElementType() reflect.Type {
	return reflect.TypeOf((*simpleADState)(nil)).Elem()
}

type simpleADArgs struct {
	CreateAlias *bool               `pulumi:"createAlias"`
	Description *string             `pulumi:"description"`
	EnableSso   *bool               `pulumi:"enableSso"`
	Name        *string             `pulumi:"name"`
	Password    string              `pulumi:"password"`
	ShortName   *string             `pulumi:"shortName"`
	Size        string              `pulumi:"size"`
	VpcSettings SimpleADVpcSettings `pulumi:"vpcSettings"`
}

// The set of arguments for constructing a SimpleAD resource.
type SimpleADArgs struct {
	CreateAlias pulumi.BoolPtrInput
	Description pulumi.StringPtrInput
	EnableSso   pulumi.BoolPtrInput
	Name        pulumi.StringPtrInput
	Password    pulumi.StringInput
	ShortName   pulumi.StringPtrInput
	Size        pulumi.StringInput
	VpcSettings SimpleADVpcSettingsInput
}

func (SimpleADArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*simpleADArgs)(nil)).Elem()
}

type SimpleADInput interface {
	pulumi.Input

	ToSimpleADOutput() SimpleADOutput
	ToSimpleADOutputWithContext(ctx context.Context) SimpleADOutput
}

func (*SimpleAD) ElementType() reflect.Type {
	return reflect.TypeOf((**SimpleAD)(nil)).Elem()
}

func (i *SimpleAD) ToSimpleADOutput() SimpleADOutput {
	return i.ToSimpleADOutputWithContext(context.Background())
}

func (i *SimpleAD) ToSimpleADOutputWithContext(ctx context.Context) SimpleADOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SimpleADOutput)
}

type SimpleADOutput struct{ *pulumi.OutputState }

func (SimpleADOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SimpleAD)(nil)).Elem()
}

func (o SimpleADOutput) ToSimpleADOutput() SimpleADOutput {
	return o
}

func (o SimpleADOutput) ToSimpleADOutputWithContext(ctx context.Context) SimpleADOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SimpleADInput)(nil)).Elem(), &SimpleAD{})
	pulumi.RegisterOutputType(SimpleADOutput{})
}
