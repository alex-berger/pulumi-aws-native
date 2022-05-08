// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudformation

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A resource that has been registered in the CloudFormation Registry.
type ResourceVersion struct {
	pulumi.CustomResourceState

	// The Amazon Resource Name (ARN) of the type, here the ResourceVersion. This is used to uniquely identify a ResourceVersion resource
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The Amazon Resource Name (ARN) of the IAM execution role to use to register the type. If your resource type calls AWS APIs in any of its handlers, you must create an IAM execution role that includes the necessary permissions to call those AWS APIs, and provision that execution role in your account. CloudFormation then assumes that execution role to provide your resource type with the appropriate credentials.
	ExecutionRoleArn pulumi.StringPtrOutput `pulumi:"executionRoleArn"`
	// Indicates if this type version is the current default version
	IsDefaultVersion pulumi.BoolOutput `pulumi:"isDefaultVersion"`
	// Specifies logging configuration information for a type.
	LoggingConfig ResourceVersionLoggingConfigPtrOutput `pulumi:"loggingConfig"`
	// The provisioning behavior of the type. AWS CloudFormation determines the provisioning type during registration, based on the types of handlers in the schema handler package submitted.
	ProvisioningType ResourceVersionProvisioningTypeOutput `pulumi:"provisioningType"`
	// A url to the S3 bucket containing the schema handler package that contains the schema, event handlers, and associated files for the type you want to register.
	//
	// For information on generating a schema handler package for the type you want to register, see submit in the CloudFormation CLI User Guide.
	SchemaHandlerPackage pulumi.StringOutput `pulumi:"schemaHandlerPackage"`
	// The Amazon Resource Name (ARN) of the type without the versionID.
	TypeArn pulumi.StringOutput `pulumi:"typeArn"`
	// The name of the type being registered.
	//
	// We recommend that type names adhere to the following pattern: company_or_organization::service::type.
	TypeName pulumi.StringOutput `pulumi:"typeName"`
	// The ID of the version of the type represented by this resource instance.
	VersionId pulumi.StringOutput `pulumi:"versionId"`
	// The scope at which the type is visible and usable in CloudFormation operations.
	//
	// Valid values include:
	//
	// PRIVATE: The type is only visible and usable within the account in which it is registered. Currently, AWS CloudFormation marks any types you register as PRIVATE.
	//
	// PUBLIC: The type is publically visible and usable within any Amazon account.
	Visibility ResourceVersionVisibilityOutput `pulumi:"visibility"`
}

// NewResourceVersion registers a new resource with the given unique name, arguments, and options.
func NewResourceVersion(ctx *pulumi.Context,
	name string, args *ResourceVersionArgs, opts ...pulumi.ResourceOption) (*ResourceVersion, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.SchemaHandlerPackage == nil {
		return nil, errors.New("invalid value for required argument 'SchemaHandlerPackage'")
	}
	if args.TypeName == nil {
		return nil, errors.New("invalid value for required argument 'TypeName'")
	}
	var resource ResourceVersion
	err := ctx.RegisterResource("aws-native:cloudformation:ResourceVersion", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetResourceVersion gets an existing ResourceVersion resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetResourceVersion(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ResourceVersionState, opts ...pulumi.ResourceOption) (*ResourceVersion, error) {
	var resource ResourceVersion
	err := ctx.ReadResource("aws-native:cloudformation:ResourceVersion", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ResourceVersion resources.
type resourceVersionState struct {
}

type ResourceVersionState struct {
}

func (ResourceVersionState) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceVersionState)(nil)).Elem()
}

type resourceVersionArgs struct {
	// The Amazon Resource Name (ARN) of the IAM execution role to use to register the type. If your resource type calls AWS APIs in any of its handlers, you must create an IAM execution role that includes the necessary permissions to call those AWS APIs, and provision that execution role in your account. CloudFormation then assumes that execution role to provide your resource type with the appropriate credentials.
	ExecutionRoleArn *string `pulumi:"executionRoleArn"`
	// Specifies logging configuration information for a type.
	LoggingConfig *ResourceVersionLoggingConfig `pulumi:"loggingConfig"`
	// A url to the S3 bucket containing the schema handler package that contains the schema, event handlers, and associated files for the type you want to register.
	//
	// For information on generating a schema handler package for the type you want to register, see submit in the CloudFormation CLI User Guide.
	SchemaHandlerPackage string `pulumi:"schemaHandlerPackage"`
	// The name of the type being registered.
	//
	// We recommend that type names adhere to the following pattern: company_or_organization::service::type.
	TypeName string `pulumi:"typeName"`
}

// The set of arguments for constructing a ResourceVersion resource.
type ResourceVersionArgs struct {
	// The Amazon Resource Name (ARN) of the IAM execution role to use to register the type. If your resource type calls AWS APIs in any of its handlers, you must create an IAM execution role that includes the necessary permissions to call those AWS APIs, and provision that execution role in your account. CloudFormation then assumes that execution role to provide your resource type with the appropriate credentials.
	ExecutionRoleArn pulumi.StringPtrInput
	// Specifies logging configuration information for a type.
	LoggingConfig ResourceVersionLoggingConfigPtrInput
	// A url to the S3 bucket containing the schema handler package that contains the schema, event handlers, and associated files for the type you want to register.
	//
	// For information on generating a schema handler package for the type you want to register, see submit in the CloudFormation CLI User Guide.
	SchemaHandlerPackage pulumi.StringInput
	// The name of the type being registered.
	//
	// We recommend that type names adhere to the following pattern: company_or_organization::service::type.
	TypeName pulumi.StringInput
}

func (ResourceVersionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceVersionArgs)(nil)).Elem()
}

type ResourceVersionInput interface {
	pulumi.Input

	ToResourceVersionOutput() ResourceVersionOutput
	ToResourceVersionOutputWithContext(ctx context.Context) ResourceVersionOutput
}

func (*ResourceVersion) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceVersion)(nil)).Elem()
}

func (i *ResourceVersion) ToResourceVersionOutput() ResourceVersionOutput {
	return i.ToResourceVersionOutputWithContext(context.Background())
}

func (i *ResourceVersion) ToResourceVersionOutputWithContext(ctx context.Context) ResourceVersionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceVersionOutput)
}

type ResourceVersionOutput struct{ *pulumi.OutputState }

func (ResourceVersionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceVersion)(nil)).Elem()
}

func (o ResourceVersionOutput) ToResourceVersionOutput() ResourceVersionOutput {
	return o
}

func (o ResourceVersionOutput) ToResourceVersionOutputWithContext(ctx context.Context) ResourceVersionOutput {
	return o
}

// The Amazon Resource Name (ARN) of the type, here the ResourceVersion. This is used to uniquely identify a ResourceVersion resource
func (o ResourceVersionOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *ResourceVersion) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// The Amazon Resource Name (ARN) of the IAM execution role to use to register the type. If your resource type calls AWS APIs in any of its handlers, you must create an IAM execution role that includes the necessary permissions to call those AWS APIs, and provision that execution role in your account. CloudFormation then assumes that execution role to provide your resource type with the appropriate credentials.
func (o ResourceVersionOutput) ExecutionRoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceVersion) pulumi.StringPtrOutput { return v.ExecutionRoleArn }).(pulumi.StringPtrOutput)
}

// Indicates if this type version is the current default version
func (o ResourceVersionOutput) IsDefaultVersion() pulumi.BoolOutput {
	return o.ApplyT(func(v *ResourceVersion) pulumi.BoolOutput { return v.IsDefaultVersion }).(pulumi.BoolOutput)
}

// Specifies logging configuration information for a type.
func (o ResourceVersionOutput) LoggingConfig() ResourceVersionLoggingConfigPtrOutput {
	return o.ApplyT(func(v *ResourceVersion) ResourceVersionLoggingConfigPtrOutput { return v.LoggingConfig }).(ResourceVersionLoggingConfigPtrOutput)
}

// The provisioning behavior of the type. AWS CloudFormation determines the provisioning type during registration, based on the types of handlers in the schema handler package submitted.
func (o ResourceVersionOutput) ProvisioningType() ResourceVersionProvisioningTypeOutput {
	return o.ApplyT(func(v *ResourceVersion) ResourceVersionProvisioningTypeOutput { return v.ProvisioningType }).(ResourceVersionProvisioningTypeOutput)
}

// A url to the S3 bucket containing the schema handler package that contains the schema, event handlers, and associated files for the type you want to register.
//
// For information on generating a schema handler package for the type you want to register, see submit in the CloudFormation CLI User Guide.
func (o ResourceVersionOutput) SchemaHandlerPackage() pulumi.StringOutput {
	return o.ApplyT(func(v *ResourceVersion) pulumi.StringOutput { return v.SchemaHandlerPackage }).(pulumi.StringOutput)
}

// The Amazon Resource Name (ARN) of the type without the versionID.
func (o ResourceVersionOutput) TypeArn() pulumi.StringOutput {
	return o.ApplyT(func(v *ResourceVersion) pulumi.StringOutput { return v.TypeArn }).(pulumi.StringOutput)
}

// The name of the type being registered.
//
// We recommend that type names adhere to the following pattern: company_or_organization::service::type.
func (o ResourceVersionOutput) TypeName() pulumi.StringOutput {
	return o.ApplyT(func(v *ResourceVersion) pulumi.StringOutput { return v.TypeName }).(pulumi.StringOutput)
}

// The ID of the version of the type represented by this resource instance.
func (o ResourceVersionOutput) VersionId() pulumi.StringOutput {
	return o.ApplyT(func(v *ResourceVersion) pulumi.StringOutput { return v.VersionId }).(pulumi.StringOutput)
}

// The scope at which the type is visible and usable in CloudFormation operations.
//
// Valid values include:
//
// PRIVATE: The type is only visible and usable within the account in which it is registered. Currently, AWS CloudFormation marks any types you register as PRIVATE.
//
// PUBLIC: The type is publically visible and usable within any Amazon account.
func (o ResourceVersionOutput) Visibility() ResourceVersionVisibilityOutput {
	return o.ApplyT(func(v *ResourceVersion) ResourceVersionVisibilityOutput { return v.Visibility }).(ResourceVersionVisibilityOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceVersionInput)(nil)).Elem(), &ResourceVersion{})
	pulumi.RegisterOutputType(ResourceVersionOutput{})
}
