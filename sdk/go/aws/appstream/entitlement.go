// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package appstream

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::AppStream::Entitlement
type Entitlement struct {
	pulumi.CustomResourceState

	AppVisibility    pulumi.StringOutput             `pulumi:"appVisibility"`
	Attributes       EntitlementAttributeArrayOutput `pulumi:"attributes"`
	CreatedTime      pulumi.StringOutput             `pulumi:"createdTime"`
	Description      pulumi.StringPtrOutput          `pulumi:"description"`
	LastModifiedTime pulumi.StringOutput             `pulumi:"lastModifiedTime"`
	Name             pulumi.StringOutput             `pulumi:"name"`
	StackName        pulumi.StringOutput             `pulumi:"stackName"`
}

// NewEntitlement registers a new resource with the given unique name, arguments, and options.
func NewEntitlement(ctx *pulumi.Context,
	name string, args *EntitlementArgs, opts ...pulumi.ResourceOption) (*Entitlement, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AppVisibility == nil {
		return nil, errors.New("invalid value for required argument 'AppVisibility'")
	}
	if args.Attributes == nil {
		return nil, errors.New("invalid value for required argument 'Attributes'")
	}
	if args.StackName == nil {
		return nil, errors.New("invalid value for required argument 'StackName'")
	}
	var resource Entitlement
	err := ctx.RegisterResource("aws-native:appstream:Entitlement", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEntitlement gets an existing Entitlement resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEntitlement(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EntitlementState, opts ...pulumi.ResourceOption) (*Entitlement, error) {
	var resource Entitlement
	err := ctx.ReadResource("aws-native:appstream:Entitlement", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Entitlement resources.
type entitlementState struct {
}

type EntitlementState struct {
}

func (EntitlementState) ElementType() reflect.Type {
	return reflect.TypeOf((*entitlementState)(nil)).Elem()
}

type entitlementArgs struct {
	AppVisibility string                 `pulumi:"appVisibility"`
	Attributes    []EntitlementAttribute `pulumi:"attributes"`
	Description   *string                `pulumi:"description"`
	Name          *string                `pulumi:"name"`
	StackName     string                 `pulumi:"stackName"`
}

// The set of arguments for constructing a Entitlement resource.
type EntitlementArgs struct {
	AppVisibility pulumi.StringInput
	Attributes    EntitlementAttributeArrayInput
	Description   pulumi.StringPtrInput
	Name          pulumi.StringPtrInput
	StackName     pulumi.StringInput
}

func (EntitlementArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*entitlementArgs)(nil)).Elem()
}

type EntitlementInput interface {
	pulumi.Input

	ToEntitlementOutput() EntitlementOutput
	ToEntitlementOutputWithContext(ctx context.Context) EntitlementOutput
}

func (*Entitlement) ElementType() reflect.Type {
	return reflect.TypeOf((**Entitlement)(nil)).Elem()
}

func (i *Entitlement) ToEntitlementOutput() EntitlementOutput {
	return i.ToEntitlementOutputWithContext(context.Background())
}

func (i *Entitlement) ToEntitlementOutputWithContext(ctx context.Context) EntitlementOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EntitlementOutput)
}

type EntitlementOutput struct{ *pulumi.OutputState }

func (EntitlementOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Entitlement)(nil)).Elem()
}

func (o EntitlementOutput) ToEntitlementOutput() EntitlementOutput {
	return o
}

func (o EntitlementOutput) ToEntitlementOutputWithContext(ctx context.Context) EntitlementOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EntitlementInput)(nil)).Elem(), &Entitlement{})
	pulumi.RegisterOutputType(EntitlementOutput{})
}
