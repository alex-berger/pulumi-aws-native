// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package wafv2

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Contains a list of IP addresses. This can be either IPV4 or IPV6. The list will be mutually
type IPSet struct {
	pulumi.CustomResourceState

	// List of IPAddresses.
	Addresses        pulumi.StringArrayOutput    `pulumi:"addresses"`
	Arn              pulumi.StringOutput         `pulumi:"arn"`
	Description      pulumi.StringPtrOutput      `pulumi:"description"`
	IPAddressVersion IPSetIPAddressVersionOutput `pulumi:"iPAddressVersion"`
	Name             pulumi.StringPtrOutput      `pulumi:"name"`
	Scope            IPSetScopeOutput            `pulumi:"scope"`
	Tags             IPSetTagArrayOutput         `pulumi:"tags"`
}

// NewIPSet registers a new resource with the given unique name, arguments, and options.
func NewIPSet(ctx *pulumi.Context,
	name string, args *IPSetArgs, opts ...pulumi.ResourceOption) (*IPSet, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Addresses == nil {
		return nil, errors.New("invalid value for required argument 'Addresses'")
	}
	if args.IPAddressVersion == nil {
		return nil, errors.New("invalid value for required argument 'IPAddressVersion'")
	}
	if args.Scope == nil {
		return nil, errors.New("invalid value for required argument 'Scope'")
	}
	var resource IPSet
	err := ctx.RegisterResource("aws-native:wafv2:IPSet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIPSet gets an existing IPSet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIPSet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IPSetState, opts ...pulumi.ResourceOption) (*IPSet, error) {
	var resource IPSet
	err := ctx.ReadResource("aws-native:wafv2:IPSet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IPSet resources.
type ipsetState struct {
}

type IPSetState struct {
}

func (IPSetState) ElementType() reflect.Type {
	return reflect.TypeOf((*ipsetState)(nil)).Elem()
}

type ipsetArgs struct {
	// List of IPAddresses.
	Addresses        []string              `pulumi:"addresses"`
	Description      *string               `pulumi:"description"`
	IPAddressVersion IPSetIPAddressVersion `pulumi:"iPAddressVersion"`
	Name             *string               `pulumi:"name"`
	Scope            IPSetScope            `pulumi:"scope"`
	Tags             []IPSetTag            `pulumi:"tags"`
}

// The set of arguments for constructing a IPSet resource.
type IPSetArgs struct {
	// List of IPAddresses.
	Addresses        pulumi.StringArrayInput
	Description      pulumi.StringPtrInput
	IPAddressVersion IPSetIPAddressVersionInput
	Name             pulumi.StringPtrInput
	Scope            IPSetScopeInput
	Tags             IPSetTagArrayInput
}

func (IPSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ipsetArgs)(nil)).Elem()
}

type IPSetInput interface {
	pulumi.Input

	ToIPSetOutput() IPSetOutput
	ToIPSetOutputWithContext(ctx context.Context) IPSetOutput
}

func (*IPSet) ElementType() reflect.Type {
	return reflect.TypeOf((**IPSet)(nil)).Elem()
}

func (i *IPSet) ToIPSetOutput() IPSetOutput {
	return i.ToIPSetOutputWithContext(context.Background())
}

func (i *IPSet) ToIPSetOutputWithContext(ctx context.Context) IPSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IPSetOutput)
}

type IPSetOutput struct{ *pulumi.OutputState }

func (IPSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IPSet)(nil)).Elem()
}

func (o IPSetOutput) ToIPSetOutput() IPSetOutput {
	return o
}

func (o IPSetOutput) ToIPSetOutputWithContext(ctx context.Context) IPSetOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IPSetInput)(nil)).Elem(), &IPSet{})
	pulumi.RegisterOutputType(IPSetOutput{})
}
