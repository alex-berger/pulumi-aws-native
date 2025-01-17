// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package lambda

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Lambda::LayerVersionPermission
//
// Deprecated: LayerVersionPermission is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type LayerVersionPermission struct {
	pulumi.CustomResourceState

	Action          pulumi.StringOutput    `pulumi:"action"`
	LayerVersionArn pulumi.StringOutput    `pulumi:"layerVersionArn"`
	OrganizationId  pulumi.StringPtrOutput `pulumi:"organizationId"`
	Principal       pulumi.StringOutput    `pulumi:"principal"`
}

// NewLayerVersionPermission registers a new resource with the given unique name, arguments, and options.
func NewLayerVersionPermission(ctx *pulumi.Context,
	name string, args *LayerVersionPermissionArgs, opts ...pulumi.ResourceOption) (*LayerVersionPermission, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Action == nil {
		return nil, errors.New("invalid value for required argument 'Action'")
	}
	if args.LayerVersionArn == nil {
		return nil, errors.New("invalid value for required argument 'LayerVersionArn'")
	}
	if args.Principal == nil {
		return nil, errors.New("invalid value for required argument 'Principal'")
	}
	var resource LayerVersionPermission
	err := ctx.RegisterResource("aws-native:lambda:LayerVersionPermission", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLayerVersionPermission gets an existing LayerVersionPermission resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLayerVersionPermission(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LayerVersionPermissionState, opts ...pulumi.ResourceOption) (*LayerVersionPermission, error) {
	var resource LayerVersionPermission
	err := ctx.ReadResource("aws-native:lambda:LayerVersionPermission", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LayerVersionPermission resources.
type layerVersionPermissionState struct {
}

type LayerVersionPermissionState struct {
}

func (LayerVersionPermissionState) ElementType() reflect.Type {
	return reflect.TypeOf((*layerVersionPermissionState)(nil)).Elem()
}

type layerVersionPermissionArgs struct {
	Action          string  `pulumi:"action"`
	LayerVersionArn string  `pulumi:"layerVersionArn"`
	OrganizationId  *string `pulumi:"organizationId"`
	Principal       string  `pulumi:"principal"`
}

// The set of arguments for constructing a LayerVersionPermission resource.
type LayerVersionPermissionArgs struct {
	Action          pulumi.StringInput
	LayerVersionArn pulumi.StringInput
	OrganizationId  pulumi.StringPtrInput
	Principal       pulumi.StringInput
}

func (LayerVersionPermissionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*layerVersionPermissionArgs)(nil)).Elem()
}

type LayerVersionPermissionInput interface {
	pulumi.Input

	ToLayerVersionPermissionOutput() LayerVersionPermissionOutput
	ToLayerVersionPermissionOutputWithContext(ctx context.Context) LayerVersionPermissionOutput
}

func (*LayerVersionPermission) ElementType() reflect.Type {
	return reflect.TypeOf((**LayerVersionPermission)(nil)).Elem()
}

func (i *LayerVersionPermission) ToLayerVersionPermissionOutput() LayerVersionPermissionOutput {
	return i.ToLayerVersionPermissionOutputWithContext(context.Background())
}

func (i *LayerVersionPermission) ToLayerVersionPermissionOutputWithContext(ctx context.Context) LayerVersionPermissionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LayerVersionPermissionOutput)
}

type LayerVersionPermissionOutput struct{ *pulumi.OutputState }

func (LayerVersionPermissionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LayerVersionPermission)(nil)).Elem()
}

func (o LayerVersionPermissionOutput) ToLayerVersionPermissionOutput() LayerVersionPermissionOutput {
	return o
}

func (o LayerVersionPermissionOutput) ToLayerVersionPermissionOutputWithContext(ctx context.Context) LayerVersionPermissionOutput {
	return o
}

func (o LayerVersionPermissionOutput) Action() pulumi.StringOutput {
	return o.ApplyT(func(v *LayerVersionPermission) pulumi.StringOutput { return v.Action }).(pulumi.StringOutput)
}

func (o LayerVersionPermissionOutput) LayerVersionArn() pulumi.StringOutput {
	return o.ApplyT(func(v *LayerVersionPermission) pulumi.StringOutput { return v.LayerVersionArn }).(pulumi.StringOutput)
}

func (o LayerVersionPermissionOutput) OrganizationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LayerVersionPermission) pulumi.StringPtrOutput { return v.OrganizationId }).(pulumi.StringPtrOutput)
}

func (o LayerVersionPermissionOutput) Principal() pulumi.StringOutput {
	return o.ApplyT(func(v *LayerVersionPermission) pulumi.StringOutput { return v.Principal }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LayerVersionPermissionInput)(nil)).Elem(), &LayerVersionPermission{})
	pulumi.RegisterOutputType(LayerVersionPermissionOutput{})
}
