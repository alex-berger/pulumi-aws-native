// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apigateway

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-method.html
type Method struct {
	pulumi.CustomResourceState

	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-method.html#cfn-apigateway-method-apikeyrequired
	ApiKeyRequired pulumi.BoolPtrOutput `pulumi:"apiKeyRequired"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-method.html#cfn-apigateway-method-authorizationscopes
	AuthorizationScopes pulumi.StringArrayOutput `pulumi:"authorizationScopes"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-method.html#cfn-apigateway-method-authorizationtype
	AuthorizationType pulumi.StringPtrOutput `pulumi:"authorizationType"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-method.html#cfn-apigateway-method-authorizerid
	AuthorizerId pulumi.StringPtrOutput `pulumi:"authorizerId"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-method.html#cfn-apigateway-method-httpmethod
	HttpMethod pulumi.StringOutput `pulumi:"httpMethod"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-method.html#cfn-apigateway-method-integration
	Integration MethodIntegrationPtrOutput `pulumi:"integration"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-method.html#cfn-apigateway-method-methodresponses
	MethodResponses MethodMethodResponseArrayOutput `pulumi:"methodResponses"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-method.html#cfn-apigateway-method-operationname
	OperationName pulumi.StringPtrOutput `pulumi:"operationName"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-method.html#cfn-apigateway-method-requestmodels
	RequestModels pulumi.StringMapOutput `pulumi:"requestModels"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-method.html#cfn-apigateway-method-requestparameters
	RequestParameters pulumi.BoolMapOutput `pulumi:"requestParameters"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-method.html#cfn-apigateway-method-requestvalidatorid
	RequestValidatorId pulumi.StringPtrOutput `pulumi:"requestValidatorId"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-method.html#cfn-apigateway-method-resourceid
	ResourceId pulumi.StringOutput `pulumi:"resourceId"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-method.html#cfn-apigateway-method-restapiid
	RestApiId pulumi.StringOutput `pulumi:"restApiId"`
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
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-method.html#cfn-apigateway-method-apikeyrequired
	ApiKeyRequired *bool `pulumi:"apiKeyRequired"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-method.html#cfn-apigateway-method-authorizationscopes
	AuthorizationScopes []string `pulumi:"authorizationScopes"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-method.html#cfn-apigateway-method-authorizationtype
	AuthorizationType *string `pulumi:"authorizationType"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-method.html#cfn-apigateway-method-authorizerid
	AuthorizerId *string `pulumi:"authorizerId"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-method.html#cfn-apigateway-method-httpmethod
	HttpMethod string `pulumi:"httpMethod"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-method.html#cfn-apigateway-method-integration
	Integration *MethodIntegration `pulumi:"integration"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-method.html#cfn-apigateway-method-methodresponses
	MethodResponses []MethodMethodResponse `pulumi:"methodResponses"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-method.html#cfn-apigateway-method-operationname
	OperationName *string `pulumi:"operationName"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-method.html#cfn-apigateway-method-requestmodels
	RequestModels map[string]string `pulumi:"requestModels"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-method.html#cfn-apigateway-method-requestparameters
	RequestParameters map[string]bool `pulumi:"requestParameters"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-method.html#cfn-apigateway-method-requestvalidatorid
	RequestValidatorId *string `pulumi:"requestValidatorId"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-method.html#cfn-apigateway-method-resourceid
	ResourceId string `pulumi:"resourceId"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-method.html#cfn-apigateway-method-restapiid
	RestApiId string `pulumi:"restApiId"`
}

// The set of arguments for constructing a Method resource.
type MethodArgs struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-method.html#cfn-apigateway-method-apikeyrequired
	ApiKeyRequired pulumi.BoolPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-method.html#cfn-apigateway-method-authorizationscopes
	AuthorizationScopes pulumi.StringArrayInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-method.html#cfn-apigateway-method-authorizationtype
	AuthorizationType pulumi.StringPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-method.html#cfn-apigateway-method-authorizerid
	AuthorizerId pulumi.StringPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-method.html#cfn-apigateway-method-httpmethod
	HttpMethod pulumi.StringInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-method.html#cfn-apigateway-method-integration
	Integration MethodIntegrationPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-method.html#cfn-apigateway-method-methodresponses
	MethodResponses MethodMethodResponseArrayInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-method.html#cfn-apigateway-method-operationname
	OperationName pulumi.StringPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-method.html#cfn-apigateway-method-requestmodels
	RequestModels pulumi.StringMapInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-method.html#cfn-apigateway-method-requestparameters
	RequestParameters pulumi.BoolMapInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-method.html#cfn-apigateway-method-requestvalidatorid
	RequestValidatorId pulumi.StringPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-method.html#cfn-apigateway-method-resourceid
	ResourceId pulumi.StringInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-method.html#cfn-apigateway-method-restapiid
	RestApiId pulumi.StringInput
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