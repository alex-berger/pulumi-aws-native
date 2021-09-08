// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package appsync

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-resolver.html
type Resolver struct {
	pulumi.CustomResourceState

	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-resolver.html#cfn-appsync-resolver-apiid
	ApiId pulumi.StringOutput `pulumi:"apiId"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-resolver.html#cfn-appsync-resolver-cachingconfig
	CachingConfig ResolverCachingConfigPtrOutput `pulumi:"cachingConfig"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-resolver.html#cfn-appsync-resolver-datasourcename
	DataSourceName pulumi.StringPtrOutput `pulumi:"dataSourceName"`
	FieldName      pulumi.StringOutput    `pulumi:"fieldName"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-resolver.html#cfn-appsync-resolver-kind
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-resolver.html#cfn-appsync-resolver-pipelineconfig
	PipelineConfig ResolverPipelineConfigPtrOutput `pulumi:"pipelineConfig"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-resolver.html#cfn-appsync-resolver-requestmappingtemplate
	RequestMappingTemplate pulumi.StringPtrOutput `pulumi:"requestMappingTemplate"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-resolver.html#cfn-appsync-resolver-requestmappingtemplates3location
	RequestMappingTemplateS3Location pulumi.StringPtrOutput `pulumi:"requestMappingTemplateS3Location"`
	ResolverArn                      pulumi.StringOutput    `pulumi:"resolverArn"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-resolver.html#cfn-appsync-resolver-responsemappingtemplate
	ResponseMappingTemplate pulumi.StringPtrOutput `pulumi:"responseMappingTemplate"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-resolver.html#cfn-appsync-resolver-responsemappingtemplates3location
	ResponseMappingTemplateS3Location pulumi.StringPtrOutput `pulumi:"responseMappingTemplateS3Location"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-resolver.html#cfn-appsync-resolver-syncconfig
	SyncConfig ResolverSyncConfigPtrOutput `pulumi:"syncConfig"`
	TypeName   pulumi.StringOutput         `pulumi:"typeName"`
}

// NewResolver registers a new resource with the given unique name, arguments, and options.
func NewResolver(ctx *pulumi.Context,
	name string, args *ResolverArgs, opts ...pulumi.ResourceOption) (*Resolver, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApiId == nil {
		return nil, errors.New("invalid value for required argument 'ApiId'")
	}
	if args.FieldName == nil {
		return nil, errors.New("invalid value for required argument 'FieldName'")
	}
	if args.TypeName == nil {
		return nil, errors.New("invalid value for required argument 'TypeName'")
	}
	var resource Resolver
	err := ctx.RegisterResource("aws-native:appsync:Resolver", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetResolver gets an existing Resolver resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetResolver(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ResolverState, opts ...pulumi.ResourceOption) (*Resolver, error) {
	var resource Resolver
	err := ctx.ReadResource("aws-native:appsync:Resolver", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Resolver resources.
type resolverState struct {
}

type ResolverState struct {
}

func (ResolverState) ElementType() reflect.Type {
	return reflect.TypeOf((*resolverState)(nil)).Elem()
}

type resolverArgs struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-resolver.html#cfn-appsync-resolver-apiid
	ApiId string `pulumi:"apiId"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-resolver.html#cfn-appsync-resolver-cachingconfig
	CachingConfig *ResolverCachingConfig `pulumi:"cachingConfig"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-resolver.html#cfn-appsync-resolver-datasourcename
	DataSourceName *string `pulumi:"dataSourceName"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-resolver.html#cfn-appsync-resolver-fieldname
	FieldName string `pulumi:"fieldName"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-resolver.html#cfn-appsync-resolver-kind
	Kind *string `pulumi:"kind"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-resolver.html#cfn-appsync-resolver-pipelineconfig
	PipelineConfig *ResolverPipelineConfig `pulumi:"pipelineConfig"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-resolver.html#cfn-appsync-resolver-requestmappingtemplate
	RequestMappingTemplate *string `pulumi:"requestMappingTemplate"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-resolver.html#cfn-appsync-resolver-requestmappingtemplates3location
	RequestMappingTemplateS3Location *string `pulumi:"requestMappingTemplateS3Location"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-resolver.html#cfn-appsync-resolver-responsemappingtemplate
	ResponseMappingTemplate *string `pulumi:"responseMappingTemplate"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-resolver.html#cfn-appsync-resolver-responsemappingtemplates3location
	ResponseMappingTemplateS3Location *string `pulumi:"responseMappingTemplateS3Location"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-resolver.html#cfn-appsync-resolver-syncconfig
	SyncConfig *ResolverSyncConfig `pulumi:"syncConfig"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-resolver.html#cfn-appsync-resolver-typename
	TypeName string `pulumi:"typeName"`
}

// The set of arguments for constructing a Resolver resource.
type ResolverArgs struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-resolver.html#cfn-appsync-resolver-apiid
	ApiId pulumi.StringInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-resolver.html#cfn-appsync-resolver-cachingconfig
	CachingConfig ResolverCachingConfigPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-resolver.html#cfn-appsync-resolver-datasourcename
	DataSourceName pulumi.StringPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-resolver.html#cfn-appsync-resolver-fieldname
	FieldName pulumi.StringInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-resolver.html#cfn-appsync-resolver-kind
	Kind pulumi.StringPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-resolver.html#cfn-appsync-resolver-pipelineconfig
	PipelineConfig ResolverPipelineConfigPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-resolver.html#cfn-appsync-resolver-requestmappingtemplate
	RequestMappingTemplate pulumi.StringPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-resolver.html#cfn-appsync-resolver-requestmappingtemplates3location
	RequestMappingTemplateS3Location pulumi.StringPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-resolver.html#cfn-appsync-resolver-responsemappingtemplate
	ResponseMappingTemplate pulumi.StringPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-resolver.html#cfn-appsync-resolver-responsemappingtemplates3location
	ResponseMappingTemplateS3Location pulumi.StringPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-resolver.html#cfn-appsync-resolver-syncconfig
	SyncConfig ResolverSyncConfigPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-resolver.html#cfn-appsync-resolver-typename
	TypeName pulumi.StringInput
}

func (ResolverArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*resolverArgs)(nil)).Elem()
}

type ResolverInput interface {
	pulumi.Input

	ToResolverOutput() ResolverOutput
	ToResolverOutputWithContext(ctx context.Context) ResolverOutput
}

func (*Resolver) ElementType() reflect.Type {
	return reflect.TypeOf((*Resolver)(nil))
}

func (i *Resolver) ToResolverOutput() ResolverOutput {
	return i.ToResolverOutputWithContext(context.Background())
}

func (i *Resolver) ToResolverOutputWithContext(ctx context.Context) ResolverOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResolverOutput)
}

type ResolverOutput struct{ *pulumi.OutputState }

func (ResolverOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Resolver)(nil))
}

func (o ResolverOutput) ToResolverOutput() ResolverOutput {
	return o
}

func (o ResolverOutput) ToResolverOutputWithContext(ctx context.Context) ResolverOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ResolverOutput{})
}