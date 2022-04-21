// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package finspace

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// An example resource schema demonstrating some basic constructs and validation rules.
func LookupEnvironment(ctx *pulumi.Context, args *LookupEnvironmentArgs, opts ...pulumi.InvokeOption) (*LookupEnvironmentResult, error) {
	var rv LookupEnvironmentResult
	err := ctx.Invoke("aws-native:finspace:getEnvironment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupEnvironmentArgs struct {
	// Unique identifier for representing FinSpace Environment
	EnvironmentId string `pulumi:"environmentId"`
}

type LookupEnvironmentResult struct {
	// AWS account ID associated with the Environment
	AwsAccountId *string `pulumi:"awsAccountId"`
	// ID for FinSpace created account used to store Environment artifacts
	DedicatedServiceAccountId *string `pulumi:"dedicatedServiceAccountId"`
	// Description of the Environment
	Description *string `pulumi:"description"`
	// ARN of the Environment
	EnvironmentArn *string `pulumi:"environmentArn"`
	// Unique identifier for representing FinSpace Environment
	EnvironmentId *string `pulumi:"environmentId"`
	// URL used to login to the Environment
	EnvironmentUrl *string `pulumi:"environmentUrl"`
	// Federation mode used with the Environment
	FederationMode       *EnvironmentFederationMode       `pulumi:"federationMode"`
	FederationParameters *EnvironmentFederationParameters `pulumi:"federationParameters"`
	// Name of the Environment
	Name *string `pulumi:"name"`
	// SageMaker Studio Domain URL associated with the Environment
	SageMakerStudioDomainUrl *string `pulumi:"sageMakerStudioDomainUrl"`
	// State of the Environment
	Status *EnvironmentStatus `pulumi:"status"`
}

func LookupEnvironmentOutput(ctx *pulumi.Context, args LookupEnvironmentOutputArgs, opts ...pulumi.InvokeOption) LookupEnvironmentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupEnvironmentResult, error) {
			args := v.(LookupEnvironmentArgs)
			r, err := LookupEnvironment(ctx, &args, opts...)
			var s LookupEnvironmentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupEnvironmentResultOutput)
}

type LookupEnvironmentOutputArgs struct {
	// Unique identifier for representing FinSpace Environment
	EnvironmentId pulumi.StringInput `pulumi:"environmentId"`
}

func (LookupEnvironmentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEnvironmentArgs)(nil)).Elem()
}

type LookupEnvironmentResultOutput struct{ *pulumi.OutputState }

func (LookupEnvironmentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEnvironmentResult)(nil)).Elem()
}

func (o LookupEnvironmentResultOutput) ToLookupEnvironmentResultOutput() LookupEnvironmentResultOutput {
	return o
}

func (o LookupEnvironmentResultOutput) ToLookupEnvironmentResultOutputWithContext(ctx context.Context) LookupEnvironmentResultOutput {
	return o
}

// AWS account ID associated with the Environment
func (o LookupEnvironmentResultOutput) AwsAccountId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEnvironmentResult) *string { return v.AwsAccountId }).(pulumi.StringPtrOutput)
}

// ID for FinSpace created account used to store Environment artifacts
func (o LookupEnvironmentResultOutput) DedicatedServiceAccountId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEnvironmentResult) *string { return v.DedicatedServiceAccountId }).(pulumi.StringPtrOutput)
}

// Description of the Environment
func (o LookupEnvironmentResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEnvironmentResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// ARN of the Environment
func (o LookupEnvironmentResultOutput) EnvironmentArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEnvironmentResult) *string { return v.EnvironmentArn }).(pulumi.StringPtrOutput)
}

// Unique identifier for representing FinSpace Environment
func (o LookupEnvironmentResultOutput) EnvironmentId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEnvironmentResult) *string { return v.EnvironmentId }).(pulumi.StringPtrOutput)
}

// URL used to login to the Environment
func (o LookupEnvironmentResultOutput) EnvironmentUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEnvironmentResult) *string { return v.EnvironmentUrl }).(pulumi.StringPtrOutput)
}

// Federation mode used with the Environment
func (o LookupEnvironmentResultOutput) FederationMode() EnvironmentFederationModePtrOutput {
	return o.ApplyT(func(v LookupEnvironmentResult) *EnvironmentFederationMode { return v.FederationMode }).(EnvironmentFederationModePtrOutput)
}

func (o LookupEnvironmentResultOutput) FederationParameters() EnvironmentFederationParametersPtrOutput {
	return o.ApplyT(func(v LookupEnvironmentResult) *EnvironmentFederationParameters { return v.FederationParameters }).(EnvironmentFederationParametersPtrOutput)
}

// Name of the Environment
func (o LookupEnvironmentResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEnvironmentResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// SageMaker Studio Domain URL associated with the Environment
func (o LookupEnvironmentResultOutput) SageMakerStudioDomainUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEnvironmentResult) *string { return v.SageMakerStudioDomainUrl }).(pulumi.StringPtrOutput)
}

// State of the Environment
func (o LookupEnvironmentResultOutput) Status() EnvironmentStatusPtrOutput {
	return o.ApplyT(func(v LookupEnvironmentResult) *EnvironmentStatus { return v.Status }).(EnvironmentStatusPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupEnvironmentResultOutput{})
}
