// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apigateway

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::ApiGateway::Method
type Method struct {
	pulumi.CustomResourceState

	ApiKeyRequired      pulumi.BoolPtrOutput            `pulumi:"apiKeyRequired"`
	AuthorizationScopes pulumi.StringArrayOutput        `pulumi:"authorizationScopes"`
	AuthorizationType   pulumi.StringPtrOutput          `pulumi:"authorizationType"`
	AuthorizerId        pulumi.StringPtrOutput          `pulumi:"authorizerId"`
	HttpMethod          pulumi.StringOutput             `pulumi:"httpMethod"`
	Integration         MethodIntegrationPtrOutput      `pulumi:"integration"`
	MethodResponses     MethodMethodResponseArrayOutput `pulumi:"methodResponses"`
	OperationName       pulumi.StringPtrOutput          `pulumi:"operationName"`
	RequestModels       pulumi.AnyOutput                `pulumi:"requestModels"`
	RequestParameters   pulumi.AnyOutput                `pulumi:"requestParameters"`
	RequestValidatorId  pulumi.StringPtrOutput          `pulumi:"requestValidatorId"`
	ResourceId          pulumi.StringOutput             `pulumi:"resourceId"`
	RestApiId           pulumi.StringOutput             `pulumi:"restApiId"`
}

// NewMethod registers a new resource with the given unique name, arguments, and options.
func NewMethod(ctx *pulumi.Context,
	name string, args *MethodArgs, opts ...pulumi.ResourceOption) (*Method, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.HttpMethod == nil {
		return nil, errors.New("invalid value for required argument 'HttpMethod'")
	}
	if args.ResourceId == nil {
		return nil, errors.New("invalid value for required argument 'ResourceId'")
	}
	if args.RestApiId == nil {
		return nil, errors.New("invalid value for required argument 'RestApiId'")
	}
	var resource Method
	err := ctx.RegisterResource("aws-native:apigateway:Method", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMethod gets an existing Method resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMethod(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MethodState, opts ...pulumi.ResourceOption) (*Method, error) {
	var resource Method
	err := ctx.ReadResource("aws-native:apigateway:Method", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Method resources.
type methodState struct {
}

type MethodState struct {
}

func (MethodState) ElementType() reflect.Type {
	return reflect.TypeOf((*methodState)(nil)).Elem()
}

type methodArgs struct {
	ApiKeyRequired      *bool                  `pulumi:"apiKeyRequired"`
	AuthorizationScopes []string               `pulumi:"authorizationScopes"`
	AuthorizationType   *string                `pulumi:"authorizationType"`
	AuthorizerId        *string                `pulumi:"authorizerId"`
	HttpMethod          string                 `pulumi:"httpMethod"`
	Integration         *MethodIntegration     `pulumi:"integration"`
	MethodResponses     []MethodMethodResponse `pulumi:"methodResponses"`
	OperationName       *string                `pulumi:"operationName"`
	RequestModels       interface{}            `pulumi:"requestModels"`
	RequestParameters   interface{}            `pulumi:"requestParameters"`
	RequestValidatorId  *string                `pulumi:"requestValidatorId"`
	ResourceId          string                 `pulumi:"resourceId"`
	RestApiId           string                 `pulumi:"restApiId"`
}

// The set of arguments for constructing a Method resource.
type MethodArgs struct {
	ApiKeyRequired      pulumi.BoolPtrInput
	AuthorizationScopes pulumi.StringArrayInput
	AuthorizationType   pulumi.StringPtrInput
	AuthorizerId        pulumi.StringPtrInput
	HttpMethod          pulumi.StringInput
	Integration         MethodIntegrationPtrInput
	MethodResponses     MethodMethodResponseArrayInput
	OperationName       pulumi.StringPtrInput
	RequestModels       pulumi.Input
	RequestParameters   pulumi.Input
	RequestValidatorId  pulumi.StringPtrInput
	ResourceId          pulumi.StringInput
	RestApiId           pulumi.StringInput
}

func (MethodArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*methodArgs)(nil)).Elem()
}

type MethodInput interface {
	pulumi.Input

	ToMethodOutput() MethodOutput
	ToMethodOutputWithContext(ctx context.Context) MethodOutput
}

func (*Method) ElementType() reflect.Type {
	return reflect.TypeOf((*Method)(nil))
}

func (i *Method) ToMethodOutput() MethodOutput {
	return i.ToMethodOutputWithContext(context.Background())
}

func (i *Method) ToMethodOutputWithContext(ctx context.Context) MethodOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MethodOutput)
}

type MethodOutput struct{ *pulumi.OutputState }

func (MethodOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Method)(nil))
}

func (o MethodOutput) ToMethodOutput() MethodOutput {
	return o
}

func (o MethodOutput) ToMethodOutputWithContext(ctx context.Context) MethodOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(MethodOutput{})
}
