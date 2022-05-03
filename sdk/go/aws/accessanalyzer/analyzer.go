// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package accessanalyzer

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The AWS::AccessAnalyzer::Analyzer type specifies an analyzer of the user's account
type Analyzer struct {
	pulumi.CustomResourceState

	// Analyzer name
	AnalyzerName pulumi.StringPtrOutput         `pulumi:"analyzerName"`
	ArchiveRules AnalyzerArchiveRuleArrayOutput `pulumi:"archiveRules"`
	// Amazon Resource Name (ARN) of the analyzer
	Arn pulumi.StringOutput `pulumi:"arn"`
	// An array of key-value pairs to apply to this resource.
	Tags AnalyzerTagArrayOutput `pulumi:"tags"`
	// The type of the analyzer, must be ACCOUNT or ORGANIZATION
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewAnalyzer registers a new resource with the given unique name, arguments, and options.
func NewAnalyzer(ctx *pulumi.Context,
	name string, args *AnalyzerArgs, opts ...pulumi.ResourceOption) (*Analyzer, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	var resource Analyzer
	err := ctx.RegisterResource("aws-native:accessanalyzer:Analyzer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAnalyzer gets an existing Analyzer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAnalyzer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AnalyzerState, opts ...pulumi.ResourceOption) (*Analyzer, error) {
	var resource Analyzer
	err := ctx.ReadResource("aws-native:accessanalyzer:Analyzer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Analyzer resources.
type analyzerState struct {
}

type AnalyzerState struct {
}

func (AnalyzerState) ElementType() reflect.Type {
	return reflect.TypeOf((*analyzerState)(nil)).Elem()
}

type analyzerArgs struct {
	// Analyzer name
	AnalyzerName *string               `pulumi:"analyzerName"`
	ArchiveRules []AnalyzerArchiveRule `pulumi:"archiveRules"`
	// An array of key-value pairs to apply to this resource.
	Tags []AnalyzerTag `pulumi:"tags"`
	// The type of the analyzer, must be ACCOUNT or ORGANIZATION
	Type string `pulumi:"type"`
}

// The set of arguments for constructing a Analyzer resource.
type AnalyzerArgs struct {
	// Analyzer name
	AnalyzerName pulumi.StringPtrInput
	ArchiveRules AnalyzerArchiveRuleArrayInput
	// An array of key-value pairs to apply to this resource.
	Tags AnalyzerTagArrayInput
	// The type of the analyzer, must be ACCOUNT or ORGANIZATION
	Type pulumi.StringInput
}

func (AnalyzerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*analyzerArgs)(nil)).Elem()
}

type AnalyzerInput interface {
	pulumi.Input

	ToAnalyzerOutput() AnalyzerOutput
	ToAnalyzerOutputWithContext(ctx context.Context) AnalyzerOutput
}

func (*Analyzer) ElementType() reflect.Type {
	return reflect.TypeOf((**Analyzer)(nil)).Elem()
}

func (i *Analyzer) ToAnalyzerOutput() AnalyzerOutput {
	return i.ToAnalyzerOutputWithContext(context.Background())
}

func (i *Analyzer) ToAnalyzerOutputWithContext(ctx context.Context) AnalyzerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AnalyzerOutput)
}

type AnalyzerOutput struct{ *pulumi.OutputState }

func (AnalyzerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Analyzer)(nil)).Elem()
}

func (o AnalyzerOutput) ToAnalyzerOutput() AnalyzerOutput {
	return o
}

func (o AnalyzerOutput) ToAnalyzerOutputWithContext(ctx context.Context) AnalyzerOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AnalyzerInput)(nil)).Elem(), &Analyzer{})
	pulumi.RegisterOutputType(AnalyzerOutput{})
}
