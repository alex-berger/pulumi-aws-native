// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sagemaker

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::SageMaker::ModelExplainabilityJobDefinition
type ModelExplainabilityJobDefinition struct {
	pulumi.CustomResourceState

	// The time at which the job definition was created.
	CreationTime pulumi.StringOutput `pulumi:"creationTime"`
	// The Amazon Resource Name (ARN) of job definition.
	JobDefinitionArn                    pulumi.StringOutput                                                        `pulumi:"jobDefinitionArn"`
	JobDefinitionName                   pulumi.StringPtrOutput                                                     `pulumi:"jobDefinitionName"`
	JobResources                        ModelExplainabilityJobDefinitionMonitoringResourcesOutput                  `pulumi:"jobResources"`
	ModelExplainabilityAppSpecification ModelExplainabilityJobDefinitionModelExplainabilityAppSpecificationOutput  `pulumi:"modelExplainabilityAppSpecification"`
	ModelExplainabilityBaselineConfig   ModelExplainabilityJobDefinitionModelExplainabilityBaselineConfigPtrOutput `pulumi:"modelExplainabilityBaselineConfig"`
	ModelExplainabilityJobInput         ModelExplainabilityJobDefinitionModelExplainabilityJobInputOutput          `pulumi:"modelExplainabilityJobInput"`
	ModelExplainabilityJobOutputConfig  ModelExplainabilityJobDefinitionMonitoringOutputConfigOutput               `pulumi:"modelExplainabilityJobOutputConfig"`
	NetworkConfig                       ModelExplainabilityJobDefinitionNetworkConfigPtrOutput                     `pulumi:"networkConfig"`
	// The Amazon Resource Name (ARN) of an IAM role that Amazon SageMaker can assume to perform tasks on your behalf.
	RoleArn           pulumi.StringOutput                                        `pulumi:"roleArn"`
	StoppingCondition ModelExplainabilityJobDefinitionStoppingConditionPtrOutput `pulumi:"stoppingCondition"`
	// An array of key-value pairs to apply to this resource.
	Tags ModelExplainabilityJobDefinitionTagArrayOutput `pulumi:"tags"`
}

// NewModelExplainabilityJobDefinition registers a new resource with the given unique name, arguments, and options.
func NewModelExplainabilityJobDefinition(ctx *pulumi.Context,
	name string, args *ModelExplainabilityJobDefinitionArgs, opts ...pulumi.ResourceOption) (*ModelExplainabilityJobDefinition, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.JobResources == nil {
		return nil, errors.New("invalid value for required argument 'JobResources'")
	}
	if args.ModelExplainabilityAppSpecification == nil {
		return nil, errors.New("invalid value for required argument 'ModelExplainabilityAppSpecification'")
	}
	if args.ModelExplainabilityJobInput == nil {
		return nil, errors.New("invalid value for required argument 'ModelExplainabilityJobInput'")
	}
	if args.ModelExplainabilityJobOutputConfig == nil {
		return nil, errors.New("invalid value for required argument 'ModelExplainabilityJobOutputConfig'")
	}
	if args.RoleArn == nil {
		return nil, errors.New("invalid value for required argument 'RoleArn'")
	}
	var resource ModelExplainabilityJobDefinition
	err := ctx.RegisterResource("aws-native:sagemaker:ModelExplainabilityJobDefinition", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetModelExplainabilityJobDefinition gets an existing ModelExplainabilityJobDefinition resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetModelExplainabilityJobDefinition(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ModelExplainabilityJobDefinitionState, opts ...pulumi.ResourceOption) (*ModelExplainabilityJobDefinition, error) {
	var resource ModelExplainabilityJobDefinition
	err := ctx.ReadResource("aws-native:sagemaker:ModelExplainabilityJobDefinition", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ModelExplainabilityJobDefinition resources.
type modelExplainabilityJobDefinitionState struct {
}

type ModelExplainabilityJobDefinitionState struct {
}

func (ModelExplainabilityJobDefinitionState) ElementType() reflect.Type {
	return reflect.TypeOf((*modelExplainabilityJobDefinitionState)(nil)).Elem()
}

type modelExplainabilityJobDefinitionArgs struct {
	JobDefinitionName                   *string                                                             `pulumi:"jobDefinitionName"`
	JobResources                        ModelExplainabilityJobDefinitionMonitoringResources                 `pulumi:"jobResources"`
	ModelExplainabilityAppSpecification ModelExplainabilityJobDefinitionModelExplainabilityAppSpecification `pulumi:"modelExplainabilityAppSpecification"`
	ModelExplainabilityBaselineConfig   *ModelExplainabilityJobDefinitionModelExplainabilityBaselineConfig  `pulumi:"modelExplainabilityBaselineConfig"`
	ModelExplainabilityJobInput         ModelExplainabilityJobDefinitionModelExplainabilityJobInput         `pulumi:"modelExplainabilityJobInput"`
	ModelExplainabilityJobOutputConfig  ModelExplainabilityJobDefinitionMonitoringOutputConfig              `pulumi:"modelExplainabilityJobOutputConfig"`
	NetworkConfig                       *ModelExplainabilityJobDefinitionNetworkConfig                      `pulumi:"networkConfig"`
	// The Amazon Resource Name (ARN) of an IAM role that Amazon SageMaker can assume to perform tasks on your behalf.
	RoleArn           string                                             `pulumi:"roleArn"`
	StoppingCondition *ModelExplainabilityJobDefinitionStoppingCondition `pulumi:"stoppingCondition"`
	// An array of key-value pairs to apply to this resource.
	Tags []ModelExplainabilityJobDefinitionTag `pulumi:"tags"`
}

// The set of arguments for constructing a ModelExplainabilityJobDefinition resource.
type ModelExplainabilityJobDefinitionArgs struct {
	JobDefinitionName                   pulumi.StringPtrInput
	JobResources                        ModelExplainabilityJobDefinitionMonitoringResourcesInput
	ModelExplainabilityAppSpecification ModelExplainabilityJobDefinitionModelExplainabilityAppSpecificationInput
	ModelExplainabilityBaselineConfig   ModelExplainabilityJobDefinitionModelExplainabilityBaselineConfigPtrInput
	ModelExplainabilityJobInput         ModelExplainabilityJobDefinitionModelExplainabilityJobInputInput
	ModelExplainabilityJobOutputConfig  ModelExplainabilityJobDefinitionMonitoringOutputConfigInput
	NetworkConfig                       ModelExplainabilityJobDefinitionNetworkConfigPtrInput
	// The Amazon Resource Name (ARN) of an IAM role that Amazon SageMaker can assume to perform tasks on your behalf.
	RoleArn           pulumi.StringInput
	StoppingCondition ModelExplainabilityJobDefinitionStoppingConditionPtrInput
	// An array of key-value pairs to apply to this resource.
	Tags ModelExplainabilityJobDefinitionTagArrayInput
}

func (ModelExplainabilityJobDefinitionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*modelExplainabilityJobDefinitionArgs)(nil)).Elem()
}

type ModelExplainabilityJobDefinitionInput interface {
	pulumi.Input

	ToModelExplainabilityJobDefinitionOutput() ModelExplainabilityJobDefinitionOutput
	ToModelExplainabilityJobDefinitionOutputWithContext(ctx context.Context) ModelExplainabilityJobDefinitionOutput
}

func (*ModelExplainabilityJobDefinition) ElementType() reflect.Type {
	return reflect.TypeOf((**ModelExplainabilityJobDefinition)(nil)).Elem()
}

func (i *ModelExplainabilityJobDefinition) ToModelExplainabilityJobDefinitionOutput() ModelExplainabilityJobDefinitionOutput {
	return i.ToModelExplainabilityJobDefinitionOutputWithContext(context.Background())
}

func (i *ModelExplainabilityJobDefinition) ToModelExplainabilityJobDefinitionOutputWithContext(ctx context.Context) ModelExplainabilityJobDefinitionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ModelExplainabilityJobDefinitionOutput)
}

type ModelExplainabilityJobDefinitionOutput struct{ *pulumi.OutputState }

func (ModelExplainabilityJobDefinitionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ModelExplainabilityJobDefinition)(nil)).Elem()
}

func (o ModelExplainabilityJobDefinitionOutput) ToModelExplainabilityJobDefinitionOutput() ModelExplainabilityJobDefinitionOutput {
	return o
}

func (o ModelExplainabilityJobDefinitionOutput) ToModelExplainabilityJobDefinitionOutputWithContext(ctx context.Context) ModelExplainabilityJobDefinitionOutput {
	return o
}

// The time at which the job definition was created.
func (o ModelExplainabilityJobDefinitionOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v *ModelExplainabilityJobDefinition) pulumi.StringOutput { return v.CreationTime }).(pulumi.StringOutput)
}

// The Amazon Resource Name (ARN) of job definition.
func (o ModelExplainabilityJobDefinitionOutput) JobDefinitionArn() pulumi.StringOutput {
	return o.ApplyT(func(v *ModelExplainabilityJobDefinition) pulumi.StringOutput { return v.JobDefinitionArn }).(pulumi.StringOutput)
}

func (o ModelExplainabilityJobDefinitionOutput) JobDefinitionName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ModelExplainabilityJobDefinition) pulumi.StringPtrOutput { return v.JobDefinitionName }).(pulumi.StringPtrOutput)
}

func (o ModelExplainabilityJobDefinitionOutput) JobResources() ModelExplainabilityJobDefinitionMonitoringResourcesOutput {
	return o.ApplyT(func(v *ModelExplainabilityJobDefinition) ModelExplainabilityJobDefinitionMonitoringResourcesOutput {
		return v.JobResources
	}).(ModelExplainabilityJobDefinitionMonitoringResourcesOutput)
}

func (o ModelExplainabilityJobDefinitionOutput) ModelExplainabilityAppSpecification() ModelExplainabilityJobDefinitionModelExplainabilityAppSpecificationOutput {
	return o.ApplyT(func(v *ModelExplainabilityJobDefinition) ModelExplainabilityJobDefinitionModelExplainabilityAppSpecificationOutput {
		return v.ModelExplainabilityAppSpecification
	}).(ModelExplainabilityJobDefinitionModelExplainabilityAppSpecificationOutput)
}

func (o ModelExplainabilityJobDefinitionOutput) ModelExplainabilityBaselineConfig() ModelExplainabilityJobDefinitionModelExplainabilityBaselineConfigPtrOutput {
	return o.ApplyT(func(v *ModelExplainabilityJobDefinition) ModelExplainabilityJobDefinitionModelExplainabilityBaselineConfigPtrOutput {
		return v.ModelExplainabilityBaselineConfig
	}).(ModelExplainabilityJobDefinitionModelExplainabilityBaselineConfigPtrOutput)
}

func (o ModelExplainabilityJobDefinitionOutput) ModelExplainabilityJobInput() ModelExplainabilityJobDefinitionModelExplainabilityJobInputOutput {
	return o.ApplyT(func(v *ModelExplainabilityJobDefinition) ModelExplainabilityJobDefinitionModelExplainabilityJobInputOutput {
		return v.ModelExplainabilityJobInput
	}).(ModelExplainabilityJobDefinitionModelExplainabilityJobInputOutput)
}

func (o ModelExplainabilityJobDefinitionOutput) ModelExplainabilityJobOutputConfig() ModelExplainabilityJobDefinitionMonitoringOutputConfigOutput {
	return o.ApplyT(func(v *ModelExplainabilityJobDefinition) ModelExplainabilityJobDefinitionMonitoringOutputConfigOutput {
		return v.ModelExplainabilityJobOutputConfig
	}).(ModelExplainabilityJobDefinitionMonitoringOutputConfigOutput)
}

func (o ModelExplainabilityJobDefinitionOutput) NetworkConfig() ModelExplainabilityJobDefinitionNetworkConfigPtrOutput {
	return o.ApplyT(func(v *ModelExplainabilityJobDefinition) ModelExplainabilityJobDefinitionNetworkConfigPtrOutput {
		return v.NetworkConfig
	}).(ModelExplainabilityJobDefinitionNetworkConfigPtrOutput)
}

// The Amazon Resource Name (ARN) of an IAM role that Amazon SageMaker can assume to perform tasks on your behalf.
func (o ModelExplainabilityJobDefinitionOutput) RoleArn() pulumi.StringOutput {
	return o.ApplyT(func(v *ModelExplainabilityJobDefinition) pulumi.StringOutput { return v.RoleArn }).(pulumi.StringOutput)
}

func (o ModelExplainabilityJobDefinitionOutput) StoppingCondition() ModelExplainabilityJobDefinitionStoppingConditionPtrOutput {
	return o.ApplyT(func(v *ModelExplainabilityJobDefinition) ModelExplainabilityJobDefinitionStoppingConditionPtrOutput {
		return v.StoppingCondition
	}).(ModelExplainabilityJobDefinitionStoppingConditionPtrOutput)
}

// An array of key-value pairs to apply to this resource.
func (o ModelExplainabilityJobDefinitionOutput) Tags() ModelExplainabilityJobDefinitionTagArrayOutput {
	return o.ApplyT(func(v *ModelExplainabilityJobDefinition) ModelExplainabilityJobDefinitionTagArrayOutput {
		return v.Tags
	}).(ModelExplainabilityJobDefinitionTagArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ModelExplainabilityJobDefinitionInput)(nil)).Elem(), &ModelExplainabilityJobDefinition{})
	pulumi.RegisterOutputType(ModelExplainabilityJobDefinitionOutput{})
}
