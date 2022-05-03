// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package configuration

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A conformance pack is a collection of AWS Config rules and remediation actions that can be easily deployed as a single entity in an account and a region or across an entire AWS Organization.
type ConformancePack struct {
	pulumi.CustomResourceState

	// A list of ConformancePackInputParameter objects.
	ConformancePackInputParameters ConformancePackInputParameterArrayOutput `pulumi:"conformancePackInputParameters"`
	// Name of the conformance pack which will be assigned as the unique identifier.
	ConformancePackName pulumi.StringOutput `pulumi:"conformancePackName"`
	// AWS Config stores intermediate files while processing conformance pack template.
	DeliveryS3Bucket pulumi.StringPtrOutput `pulumi:"deliveryS3Bucket"`
	// The prefix for delivery S3 bucket.
	DeliveryS3KeyPrefix pulumi.StringPtrOutput `pulumi:"deliveryS3KeyPrefix"`
	// A string containing full conformance pack template body. You can only specify one of the template body or template S3Uri fields.
	TemplateBody pulumi.StringPtrOutput `pulumi:"templateBody"`
	// Location of file containing the template body which points to the conformance pack template that is located in an Amazon S3 bucket. You can only specify one of the template body or template S3Uri fields.
	TemplateS3Uri pulumi.StringPtrOutput `pulumi:"templateS3Uri"`
}

// NewConformancePack registers a new resource with the given unique name, arguments, and options.
func NewConformancePack(ctx *pulumi.Context,
	name string, args *ConformancePackArgs, opts ...pulumi.ResourceOption) (*ConformancePack, error) {
	if args == nil {
		args = &ConformancePackArgs{}
	}

	var resource ConformancePack
	err := ctx.RegisterResource("aws-native:configuration:ConformancePack", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConformancePack gets an existing ConformancePack resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConformancePack(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConformancePackState, opts ...pulumi.ResourceOption) (*ConformancePack, error) {
	var resource ConformancePack
	err := ctx.ReadResource("aws-native:configuration:ConformancePack", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ConformancePack resources.
type conformancePackState struct {
}

type ConformancePackState struct {
}

func (ConformancePackState) ElementType() reflect.Type {
	return reflect.TypeOf((*conformancePackState)(nil)).Elem()
}

type conformancePackArgs struct {
	// A list of ConformancePackInputParameter objects.
	ConformancePackInputParameters []ConformancePackInputParameter `pulumi:"conformancePackInputParameters"`
	// Name of the conformance pack which will be assigned as the unique identifier.
	ConformancePackName *string `pulumi:"conformancePackName"`
	// AWS Config stores intermediate files while processing conformance pack template.
	DeliveryS3Bucket *string `pulumi:"deliveryS3Bucket"`
	// The prefix for delivery S3 bucket.
	DeliveryS3KeyPrefix *string `pulumi:"deliveryS3KeyPrefix"`
	// A string containing full conformance pack template body. You can only specify one of the template body or template S3Uri fields.
	TemplateBody *string `pulumi:"templateBody"`
	// Location of file containing the template body which points to the conformance pack template that is located in an Amazon S3 bucket. You can only specify one of the template body or template S3Uri fields.
	TemplateS3Uri *string `pulumi:"templateS3Uri"`
}

// The set of arguments for constructing a ConformancePack resource.
type ConformancePackArgs struct {
	// A list of ConformancePackInputParameter objects.
	ConformancePackInputParameters ConformancePackInputParameterArrayInput
	// Name of the conformance pack which will be assigned as the unique identifier.
	ConformancePackName pulumi.StringPtrInput
	// AWS Config stores intermediate files while processing conformance pack template.
	DeliveryS3Bucket pulumi.StringPtrInput
	// The prefix for delivery S3 bucket.
	DeliveryS3KeyPrefix pulumi.StringPtrInput
	// A string containing full conformance pack template body. You can only specify one of the template body or template S3Uri fields.
	TemplateBody pulumi.StringPtrInput
	// Location of file containing the template body which points to the conformance pack template that is located in an Amazon S3 bucket. You can only specify one of the template body or template S3Uri fields.
	TemplateS3Uri pulumi.StringPtrInput
}

func (ConformancePackArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*conformancePackArgs)(nil)).Elem()
}

type ConformancePackInput interface {
	pulumi.Input

	ToConformancePackOutput() ConformancePackOutput
	ToConformancePackOutputWithContext(ctx context.Context) ConformancePackOutput
}

func (*ConformancePack) ElementType() reflect.Type {
	return reflect.TypeOf((**ConformancePack)(nil)).Elem()
}

func (i *ConformancePack) ToConformancePackOutput() ConformancePackOutput {
	return i.ToConformancePackOutputWithContext(context.Background())
}

func (i *ConformancePack) ToConformancePackOutputWithContext(ctx context.Context) ConformancePackOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConformancePackOutput)
}

type ConformancePackOutput struct{ *pulumi.OutputState }

func (ConformancePackOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConformancePack)(nil)).Elem()
}

func (o ConformancePackOutput) ToConformancePackOutput() ConformancePackOutput {
	return o
}

func (o ConformancePackOutput) ToConformancePackOutputWithContext(ctx context.Context) ConformancePackOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ConformancePackInput)(nil)).Elem(), &ConformancePack{})
	pulumi.RegisterOutputType(ConformancePackOutput{})
}
