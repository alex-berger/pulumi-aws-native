// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package datasync

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource schema for AWS::DataSync::LocationFSxOpenZFS.
type LocationFSxOpenZFS struct {
	pulumi.CustomResourceState

	// The Amazon Resource Name (ARN) for the FSx OpenZFS file system.
	FsxFilesystemArn pulumi.StringOutput `pulumi:"fsxFilesystemArn"`
	// The Amazon Resource Name (ARN) of the Amazon FSx OpenZFS file system location that is created.
	LocationArn pulumi.StringOutput `pulumi:"locationArn"`
	// The URL of the FSx OpenZFS that was described.
	LocationUri pulumi.StringOutput              `pulumi:"locationUri"`
	Protocol    LocationFSxOpenZFSProtocolOutput `pulumi:"protocol"`
	// The ARNs of the security groups that are to use to configure the FSx OpenZFS file system.
	SecurityGroupArns pulumi.StringArrayOutput `pulumi:"securityGroupArns"`
	// A subdirectory in the location's path.
	Subdirectory pulumi.StringPtrOutput `pulumi:"subdirectory"`
	// An array of key-value pairs to apply to this resource.
	Tags LocationFSxOpenZFSTagArrayOutput `pulumi:"tags"`
}

// NewLocationFSxOpenZFS registers a new resource with the given unique name, arguments, and options.
func NewLocationFSxOpenZFS(ctx *pulumi.Context,
	name string, args *LocationFSxOpenZFSArgs, opts ...pulumi.ResourceOption) (*LocationFSxOpenZFS, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FsxFilesystemArn == nil {
		return nil, errors.New("invalid value for required argument 'FsxFilesystemArn'")
	}
	if args.Protocol == nil {
		return nil, errors.New("invalid value for required argument 'Protocol'")
	}
	if args.SecurityGroupArns == nil {
		return nil, errors.New("invalid value for required argument 'SecurityGroupArns'")
	}
	var resource LocationFSxOpenZFS
	err := ctx.RegisterResource("aws-native:datasync:LocationFSxOpenZFS", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLocationFSxOpenZFS gets an existing LocationFSxOpenZFS resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLocationFSxOpenZFS(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LocationFSxOpenZFSState, opts ...pulumi.ResourceOption) (*LocationFSxOpenZFS, error) {
	var resource LocationFSxOpenZFS
	err := ctx.ReadResource("aws-native:datasync:LocationFSxOpenZFS", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LocationFSxOpenZFS resources.
type locationFSxOpenZFSState struct {
}

type LocationFSxOpenZFSState struct {
}

func (LocationFSxOpenZFSState) ElementType() reflect.Type {
	return reflect.TypeOf((*locationFSxOpenZFSState)(nil)).Elem()
}

type locationFSxOpenZFSArgs struct {
	// The Amazon Resource Name (ARN) for the FSx OpenZFS file system.
	FsxFilesystemArn string                     `pulumi:"fsxFilesystemArn"`
	Protocol         LocationFSxOpenZFSProtocol `pulumi:"protocol"`
	// The ARNs of the security groups that are to use to configure the FSx OpenZFS file system.
	SecurityGroupArns []string `pulumi:"securityGroupArns"`
	// A subdirectory in the location's path.
	Subdirectory *string `pulumi:"subdirectory"`
	// An array of key-value pairs to apply to this resource.
	Tags []LocationFSxOpenZFSTag `pulumi:"tags"`
}

// The set of arguments for constructing a LocationFSxOpenZFS resource.
type LocationFSxOpenZFSArgs struct {
	// The Amazon Resource Name (ARN) for the FSx OpenZFS file system.
	FsxFilesystemArn pulumi.StringInput
	Protocol         LocationFSxOpenZFSProtocolInput
	// The ARNs of the security groups that are to use to configure the FSx OpenZFS file system.
	SecurityGroupArns pulumi.StringArrayInput
	// A subdirectory in the location's path.
	Subdirectory pulumi.StringPtrInput
	// An array of key-value pairs to apply to this resource.
	Tags LocationFSxOpenZFSTagArrayInput
}

func (LocationFSxOpenZFSArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*locationFSxOpenZFSArgs)(nil)).Elem()
}

type LocationFSxOpenZFSInput interface {
	pulumi.Input

	ToLocationFSxOpenZFSOutput() LocationFSxOpenZFSOutput
	ToLocationFSxOpenZFSOutputWithContext(ctx context.Context) LocationFSxOpenZFSOutput
}

func (*LocationFSxOpenZFS) ElementType() reflect.Type {
	return reflect.TypeOf((**LocationFSxOpenZFS)(nil)).Elem()
}

func (i *LocationFSxOpenZFS) ToLocationFSxOpenZFSOutput() LocationFSxOpenZFSOutput {
	return i.ToLocationFSxOpenZFSOutputWithContext(context.Background())
}

func (i *LocationFSxOpenZFS) ToLocationFSxOpenZFSOutputWithContext(ctx context.Context) LocationFSxOpenZFSOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocationFSxOpenZFSOutput)
}

type LocationFSxOpenZFSOutput struct{ *pulumi.OutputState }

func (LocationFSxOpenZFSOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LocationFSxOpenZFS)(nil)).Elem()
}

func (o LocationFSxOpenZFSOutput) ToLocationFSxOpenZFSOutput() LocationFSxOpenZFSOutput {
	return o
}

func (o LocationFSxOpenZFSOutput) ToLocationFSxOpenZFSOutputWithContext(ctx context.Context) LocationFSxOpenZFSOutput {
	return o
}

// The Amazon Resource Name (ARN) for the FSx OpenZFS file system.
func (o LocationFSxOpenZFSOutput) FsxFilesystemArn() pulumi.StringOutput {
	return o.ApplyT(func(v *LocationFSxOpenZFS) pulumi.StringOutput { return v.FsxFilesystemArn }).(pulumi.StringOutput)
}

// The Amazon Resource Name (ARN) of the Amazon FSx OpenZFS file system location that is created.
func (o LocationFSxOpenZFSOutput) LocationArn() pulumi.StringOutput {
	return o.ApplyT(func(v *LocationFSxOpenZFS) pulumi.StringOutput { return v.LocationArn }).(pulumi.StringOutput)
}

// The URL of the FSx OpenZFS that was described.
func (o LocationFSxOpenZFSOutput) LocationUri() pulumi.StringOutput {
	return o.ApplyT(func(v *LocationFSxOpenZFS) pulumi.StringOutput { return v.LocationUri }).(pulumi.StringOutput)
}

func (o LocationFSxOpenZFSOutput) Protocol() LocationFSxOpenZFSProtocolOutput {
	return o.ApplyT(func(v *LocationFSxOpenZFS) LocationFSxOpenZFSProtocolOutput { return v.Protocol }).(LocationFSxOpenZFSProtocolOutput)
}

// The ARNs of the security groups that are to use to configure the FSx OpenZFS file system.
func (o LocationFSxOpenZFSOutput) SecurityGroupArns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *LocationFSxOpenZFS) pulumi.StringArrayOutput { return v.SecurityGroupArns }).(pulumi.StringArrayOutput)
}

// A subdirectory in the location's path.
func (o LocationFSxOpenZFSOutput) Subdirectory() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocationFSxOpenZFS) pulumi.StringPtrOutput { return v.Subdirectory }).(pulumi.StringPtrOutput)
}

// An array of key-value pairs to apply to this resource.
func (o LocationFSxOpenZFSOutput) Tags() LocationFSxOpenZFSTagArrayOutput {
	return o.ApplyT(func(v *LocationFSxOpenZFS) LocationFSxOpenZFSTagArrayOutput { return v.Tags }).(LocationFSxOpenZFSTagArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LocationFSxOpenZFSInput)(nil)).Elem(), &LocationFSxOpenZFS{})
	pulumi.RegisterOutputType(LocationFSxOpenZFSOutput{})
}
