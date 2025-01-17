// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package apigateway

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::ApiGateway::RestApi
//
// Deprecated: RestApi is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type RestApi struct {
	pulumi.CustomResourceState

	ApiKeySourceType          pulumi.StringPtrOutput                `pulumi:"apiKeySourceType"`
	BinaryMediaTypes          pulumi.StringArrayOutput              `pulumi:"binaryMediaTypes"`
	Body                      pulumi.AnyOutput                      `pulumi:"body"`
	BodyS3Location            RestApiS3LocationPtrOutput            `pulumi:"bodyS3Location"`
	CloneFrom                 pulumi.StringPtrOutput                `pulumi:"cloneFrom"`
	Description               pulumi.StringPtrOutput                `pulumi:"description"`
	DisableExecuteApiEndpoint pulumi.BoolPtrOutput                  `pulumi:"disableExecuteApiEndpoint"`
	EndpointConfiguration     RestApiEndpointConfigurationPtrOutput `pulumi:"endpointConfiguration"`
	FailOnWarnings            pulumi.BoolPtrOutput                  `pulumi:"failOnWarnings"`
	MinimumCompressionSize    pulumi.IntPtrOutput                   `pulumi:"minimumCompressionSize"`
	Mode                      pulumi.StringPtrOutput                `pulumi:"mode"`
	Name                      pulumi.StringPtrOutput                `pulumi:"name"`
	Parameters                pulumi.AnyOutput                      `pulumi:"parameters"`
	Policy                    pulumi.AnyOutput                      `pulumi:"policy"`
	RootResourceId            pulumi.StringOutput                   `pulumi:"rootResourceId"`
	Tags                      RestApiTagArrayOutput                 `pulumi:"tags"`
}

// NewRestApi registers a new resource with the given unique name, arguments, and options.
func NewRestApi(ctx *pulumi.Context,
	name string, args *RestApiArgs, opts ...pulumi.ResourceOption) (*RestApi, error) {
	if args == nil {
		args = &RestApiArgs{}
	}

	var resource RestApi
	err := ctx.RegisterResource("aws-native:apigateway:RestApi", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRestApi gets an existing RestApi resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRestApi(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RestApiState, opts ...pulumi.ResourceOption) (*RestApi, error) {
	var resource RestApi
	err := ctx.ReadResource("aws-native:apigateway:RestApi", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RestApi resources.
type restApiState struct {
}

type RestApiState struct {
}

func (RestApiState) ElementType() reflect.Type {
	return reflect.TypeOf((*restApiState)(nil)).Elem()
}

type restApiArgs struct {
	ApiKeySourceType          *string                       `pulumi:"apiKeySourceType"`
	BinaryMediaTypes          []string                      `pulumi:"binaryMediaTypes"`
	Body                      interface{}                   `pulumi:"body"`
	BodyS3Location            *RestApiS3Location            `pulumi:"bodyS3Location"`
	CloneFrom                 *string                       `pulumi:"cloneFrom"`
	Description               *string                       `pulumi:"description"`
	DisableExecuteApiEndpoint *bool                         `pulumi:"disableExecuteApiEndpoint"`
	EndpointConfiguration     *RestApiEndpointConfiguration `pulumi:"endpointConfiguration"`
	FailOnWarnings            *bool                         `pulumi:"failOnWarnings"`
	MinimumCompressionSize    *int                          `pulumi:"minimumCompressionSize"`
	Mode                      *string                       `pulumi:"mode"`
	Name                      *string                       `pulumi:"name"`
	Parameters                interface{}                   `pulumi:"parameters"`
	Policy                    interface{}                   `pulumi:"policy"`
	Tags                      []RestApiTag                  `pulumi:"tags"`
}

// The set of arguments for constructing a RestApi resource.
type RestApiArgs struct {
	ApiKeySourceType          pulumi.StringPtrInput
	BinaryMediaTypes          pulumi.StringArrayInput
	Body                      pulumi.Input
	BodyS3Location            RestApiS3LocationPtrInput
	CloneFrom                 pulumi.StringPtrInput
	Description               pulumi.StringPtrInput
	DisableExecuteApiEndpoint pulumi.BoolPtrInput
	EndpointConfiguration     RestApiEndpointConfigurationPtrInput
	FailOnWarnings            pulumi.BoolPtrInput
	MinimumCompressionSize    pulumi.IntPtrInput
	Mode                      pulumi.StringPtrInput
	Name                      pulumi.StringPtrInput
	Parameters                pulumi.Input
	Policy                    pulumi.Input
	Tags                      RestApiTagArrayInput
}

func (RestApiArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*restApiArgs)(nil)).Elem()
}

type RestApiInput interface {
	pulumi.Input

	ToRestApiOutput() RestApiOutput
	ToRestApiOutputWithContext(ctx context.Context) RestApiOutput
}

func (*RestApi) ElementType() reflect.Type {
	return reflect.TypeOf((**RestApi)(nil)).Elem()
}

func (i *RestApi) ToRestApiOutput() RestApiOutput {
	return i.ToRestApiOutputWithContext(context.Background())
}

func (i *RestApi) ToRestApiOutputWithContext(ctx context.Context) RestApiOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RestApiOutput)
}

type RestApiOutput struct{ *pulumi.OutputState }

func (RestApiOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RestApi)(nil)).Elem()
}

func (o RestApiOutput) ToRestApiOutput() RestApiOutput {
	return o
}

func (o RestApiOutput) ToRestApiOutputWithContext(ctx context.Context) RestApiOutput {
	return o
}

func (o RestApiOutput) ApiKeySourceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RestApi) pulumi.StringPtrOutput { return v.ApiKeySourceType }).(pulumi.StringPtrOutput)
}

func (o RestApiOutput) BinaryMediaTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *RestApi) pulumi.StringArrayOutput { return v.BinaryMediaTypes }).(pulumi.StringArrayOutput)
}

func (o RestApiOutput) Body() pulumi.AnyOutput {
	return o.ApplyT(func(v *RestApi) pulumi.AnyOutput { return v.Body }).(pulumi.AnyOutput)
}

func (o RestApiOutput) BodyS3Location() RestApiS3LocationPtrOutput {
	return o.ApplyT(func(v *RestApi) RestApiS3LocationPtrOutput { return v.BodyS3Location }).(RestApiS3LocationPtrOutput)
}

func (o RestApiOutput) CloneFrom() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RestApi) pulumi.StringPtrOutput { return v.CloneFrom }).(pulumi.StringPtrOutput)
}

func (o RestApiOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RestApi) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o RestApiOutput) DisableExecuteApiEndpoint() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *RestApi) pulumi.BoolPtrOutput { return v.DisableExecuteApiEndpoint }).(pulumi.BoolPtrOutput)
}

func (o RestApiOutput) EndpointConfiguration() RestApiEndpointConfigurationPtrOutput {
	return o.ApplyT(func(v *RestApi) RestApiEndpointConfigurationPtrOutput { return v.EndpointConfiguration }).(RestApiEndpointConfigurationPtrOutput)
}

func (o RestApiOutput) FailOnWarnings() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *RestApi) pulumi.BoolPtrOutput { return v.FailOnWarnings }).(pulumi.BoolPtrOutput)
}

func (o RestApiOutput) MinimumCompressionSize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *RestApi) pulumi.IntPtrOutput { return v.MinimumCompressionSize }).(pulumi.IntPtrOutput)
}

func (o RestApiOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RestApi) pulumi.StringPtrOutput { return v.Mode }).(pulumi.StringPtrOutput)
}

func (o RestApiOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RestApi) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

func (o RestApiOutput) Parameters() pulumi.AnyOutput {
	return o.ApplyT(func(v *RestApi) pulumi.AnyOutput { return v.Parameters }).(pulumi.AnyOutput)
}

func (o RestApiOutput) Policy() pulumi.AnyOutput {
	return o.ApplyT(func(v *RestApi) pulumi.AnyOutput { return v.Policy }).(pulumi.AnyOutput)
}

func (o RestApiOutput) RootResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *RestApi) pulumi.StringOutput { return v.RootResourceId }).(pulumi.StringOutput)
}

func (o RestApiOutput) Tags() RestApiTagArrayOutput {
	return o.ApplyT(func(v *RestApi) RestApiTagArrayOutput { return v.Tags }).(RestApiTagArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RestApiInput)(nil)).Elem(), &RestApi{})
	pulumi.RegisterOutputType(RestApiOutput{})
}
