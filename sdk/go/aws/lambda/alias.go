// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package lambda

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Lambda::Alias
//
// Deprecated: Alias is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type Alias struct {
	pulumi.CustomResourceState

	Description                  pulumi.StringPtrOutput                            `pulumi:"description"`
	FunctionName                 pulumi.StringOutput                               `pulumi:"functionName"`
	FunctionVersion              pulumi.StringOutput                               `pulumi:"functionVersion"`
	Name                         pulumi.StringOutput                               `pulumi:"name"`
	ProvisionedConcurrencyConfig AliasProvisionedConcurrencyConfigurationPtrOutput `pulumi:"provisionedConcurrencyConfig"`
	RoutingConfig                AliasRoutingConfigurationPtrOutput                `pulumi:"routingConfig"`
}

// NewAlias registers a new resource with the given unique name, arguments, and options.
func NewAlias(ctx *pulumi.Context,
	name string, args *AliasArgs, opts ...pulumi.ResourceOption) (*Alias, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FunctionName == nil {
		return nil, errors.New("invalid value for required argument 'FunctionName'")
	}
	if args.FunctionVersion == nil {
		return nil, errors.New("invalid value for required argument 'FunctionVersion'")
	}
	var resource Alias
	err := ctx.RegisterResource("aws-native:lambda:Alias", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAlias gets an existing Alias resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAlias(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AliasState, opts ...pulumi.ResourceOption) (*Alias, error) {
	var resource Alias
	err := ctx.ReadResource("aws-native:lambda:Alias", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Alias resources.
type aliasState struct {
}

type AliasState struct {
}

func (AliasState) ElementType() reflect.Type {
	return reflect.TypeOf((*aliasState)(nil)).Elem()
}

type aliasArgs struct {
	Description                  *string                                   `pulumi:"description"`
	FunctionName                 string                                    `pulumi:"functionName"`
	FunctionVersion              string                                    `pulumi:"functionVersion"`
	Name                         *string                                   `pulumi:"name"`
	ProvisionedConcurrencyConfig *AliasProvisionedConcurrencyConfiguration `pulumi:"provisionedConcurrencyConfig"`
	RoutingConfig                *AliasRoutingConfiguration                `pulumi:"routingConfig"`
}

// The set of arguments for constructing a Alias resource.
type AliasArgs struct {
	Description                  pulumi.StringPtrInput
	FunctionName                 pulumi.StringInput
	FunctionVersion              pulumi.StringInput
	Name                         pulumi.StringPtrInput
	ProvisionedConcurrencyConfig AliasProvisionedConcurrencyConfigurationPtrInput
	RoutingConfig                AliasRoutingConfigurationPtrInput
}

func (AliasArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*aliasArgs)(nil)).Elem()
}

type AliasInput interface {
	pulumi.Input

	ToAliasOutput() AliasOutput
	ToAliasOutputWithContext(ctx context.Context) AliasOutput
}

func (*Alias) ElementType() reflect.Type {
	return reflect.TypeOf((**Alias)(nil)).Elem()
}

func (i *Alias) ToAliasOutput() AliasOutput {
	return i.ToAliasOutputWithContext(context.Background())
}

func (i *Alias) ToAliasOutputWithContext(ctx context.Context) AliasOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AliasOutput)
}

type AliasOutput struct{ *pulumi.OutputState }

func (AliasOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Alias)(nil)).Elem()
}

func (o AliasOutput) ToAliasOutput() AliasOutput {
	return o
}

func (o AliasOutput) ToAliasOutputWithContext(ctx context.Context) AliasOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AliasInput)(nil)).Elem(), &Alias{})
	pulumi.RegisterOutputType(AliasOutput{})
}
