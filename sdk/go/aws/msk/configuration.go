// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package msk

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::MSK::Configuration
type Configuration struct {
	pulumi.CustomResourceState

	Arn               pulumi.StringOutput      `pulumi:"arn"`
	Description       pulumi.StringPtrOutput   `pulumi:"description"`
	KafkaVersionsList pulumi.StringArrayOutput `pulumi:"kafkaVersionsList"`
	Name              pulumi.StringPtrOutput   `pulumi:"name"`
	ServerProperties  pulumi.StringOutput      `pulumi:"serverProperties"`
}

// NewConfiguration registers a new resource with the given unique name, arguments, and options.
func NewConfiguration(ctx *pulumi.Context,
	name string, args *ConfigurationArgs, opts ...pulumi.ResourceOption) (*Configuration, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ServerProperties == nil {
		return nil, errors.New("invalid value for required argument 'ServerProperties'")
	}
	var resource Configuration
	err := ctx.RegisterResource("aws-native:msk:Configuration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConfiguration gets an existing Configuration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConfiguration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConfigurationState, opts ...pulumi.ResourceOption) (*Configuration, error) {
	var resource Configuration
	err := ctx.ReadResource("aws-native:msk:Configuration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Configuration resources.
type configurationState struct {
}

type ConfigurationState struct {
}

func (ConfigurationState) ElementType() reflect.Type {
	return reflect.TypeOf((*configurationState)(nil)).Elem()
}

type configurationArgs struct {
	Description       *string  `pulumi:"description"`
	KafkaVersionsList []string `pulumi:"kafkaVersionsList"`
	Name              *string  `pulumi:"name"`
	ServerProperties  string   `pulumi:"serverProperties"`
}

// The set of arguments for constructing a Configuration resource.
type ConfigurationArgs struct {
	Description       pulumi.StringPtrInput
	KafkaVersionsList pulumi.StringArrayInput
	Name              pulumi.StringPtrInput
	ServerProperties  pulumi.StringInput
}

func (ConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*configurationArgs)(nil)).Elem()
}

type ConfigurationInput interface {
	pulumi.Input

	ToConfigurationOutput() ConfigurationOutput
	ToConfigurationOutputWithContext(ctx context.Context) ConfigurationOutput
}

func (*Configuration) ElementType() reflect.Type {
	return reflect.TypeOf((**Configuration)(nil)).Elem()
}

func (i *Configuration) ToConfigurationOutput() ConfigurationOutput {
	return i.ToConfigurationOutputWithContext(context.Background())
}

func (i *Configuration) ToConfigurationOutputWithContext(ctx context.Context) ConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationOutput)
}

type ConfigurationOutput struct{ *pulumi.OutputState }

func (ConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Configuration)(nil)).Elem()
}

func (o ConfigurationOutput) ToConfigurationOutput() ConfigurationOutput {
	return o
}

func (o ConfigurationOutput) ToConfigurationOutputWithContext(ctx context.Context) ConfigurationOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ConfigurationInput)(nil)).Elem(), &Configuration{})
	pulumi.RegisterOutputType(ConfigurationOutput{})
}