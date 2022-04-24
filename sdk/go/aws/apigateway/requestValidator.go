// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apigateway

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::ApiGateway::RequestValidator
type RequestValidator struct {
	pulumi.CustomResourceState

	// Name of the request validator.
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// ID of the request validator.
	RequestValidatorId pulumi.StringOutput `pulumi:"requestValidatorId"`
	// The identifier of the targeted API entity.
	RestApiId pulumi.StringOutput `pulumi:"restApiId"`
	// Indicates whether to validate the request body according to the configured schema for the targeted API and method.
	ValidateRequestBody pulumi.BoolPtrOutput `pulumi:"validateRequestBody"`
	// Indicates whether to validate request parameters.
	ValidateRequestParameters pulumi.BoolPtrOutput `pulumi:"validateRequestParameters"`
}

// NewRequestValidator registers a new resource with the given unique name, arguments, and options.
func NewRequestValidator(ctx *pulumi.Context,
	name string, args *RequestValidatorArgs, opts ...pulumi.ResourceOption) (*RequestValidator, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RestApiId == nil {
		return nil, errors.New("invalid value for required argument 'RestApiId'")
	}
	var resource RequestValidator
	err := ctx.RegisterResource("aws-native:apigateway:RequestValidator", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRequestValidator gets an existing RequestValidator resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRequestValidator(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RequestValidatorState, opts ...pulumi.ResourceOption) (*RequestValidator, error) {
	var resource RequestValidator
	err := ctx.ReadResource("aws-native:apigateway:RequestValidator", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RequestValidator resources.
type requestValidatorState struct {
}

type RequestValidatorState struct {
}

func (RequestValidatorState) ElementType() reflect.Type {
	return reflect.TypeOf((*requestValidatorState)(nil)).Elem()
}

type requestValidatorArgs struct {
	// Name of the request validator.
	Name *string `pulumi:"name"`
	// The identifier of the targeted API entity.
	RestApiId string `pulumi:"restApiId"`
	// Indicates whether to validate the request body according to the configured schema for the targeted API and method.
	ValidateRequestBody *bool `pulumi:"validateRequestBody"`
	// Indicates whether to validate request parameters.
	ValidateRequestParameters *bool `pulumi:"validateRequestParameters"`
}

// The set of arguments for constructing a RequestValidator resource.
type RequestValidatorArgs struct {
	// Name of the request validator.
	Name pulumi.StringPtrInput
	// The identifier of the targeted API entity.
	RestApiId pulumi.StringInput
	// Indicates whether to validate the request body according to the configured schema for the targeted API and method.
	ValidateRequestBody pulumi.BoolPtrInput
	// Indicates whether to validate request parameters.
	ValidateRequestParameters pulumi.BoolPtrInput
}

func (RequestValidatorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*requestValidatorArgs)(nil)).Elem()
}

type RequestValidatorInput interface {
	pulumi.Input

	ToRequestValidatorOutput() RequestValidatorOutput
	ToRequestValidatorOutputWithContext(ctx context.Context) RequestValidatorOutput
}

func (*RequestValidator) ElementType() reflect.Type {
	return reflect.TypeOf((**RequestValidator)(nil)).Elem()
}

func (i *RequestValidator) ToRequestValidatorOutput() RequestValidatorOutput {
	return i.ToRequestValidatorOutputWithContext(context.Background())
}

func (i *RequestValidator) ToRequestValidatorOutputWithContext(ctx context.Context) RequestValidatorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RequestValidatorOutput)
}

type RequestValidatorOutput struct{ *pulumi.OutputState }

func (RequestValidatorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RequestValidator)(nil)).Elem()
}

func (o RequestValidatorOutput) ToRequestValidatorOutput() RequestValidatorOutput {
	return o
}

func (o RequestValidatorOutput) ToRequestValidatorOutputWithContext(ctx context.Context) RequestValidatorOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RequestValidatorInput)(nil)).Elem(), &RequestValidator{})
	pulumi.RegisterOutputType(RequestValidatorOutput{})
}
