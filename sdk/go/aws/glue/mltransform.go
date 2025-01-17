// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package glue

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Glue::MLTransform
//
// Deprecated: MLTransform is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type MLTransform struct {
	pulumi.CustomResourceState

	Description         pulumi.StringPtrOutput                  `pulumi:"description"`
	GlueVersion         pulumi.StringPtrOutput                  `pulumi:"glueVersion"`
	InputRecordTables   MLTransformInputRecordTablesOutput      `pulumi:"inputRecordTables"`
	MaxCapacity         pulumi.Float64PtrOutput                 `pulumi:"maxCapacity"`
	MaxRetries          pulumi.IntPtrOutput                     `pulumi:"maxRetries"`
	Name                pulumi.StringPtrOutput                  `pulumi:"name"`
	NumberOfWorkers     pulumi.IntPtrOutput                     `pulumi:"numberOfWorkers"`
	Role                pulumi.StringOutput                     `pulumi:"role"`
	Tags                pulumi.AnyOutput                        `pulumi:"tags"`
	Timeout             pulumi.IntPtrOutput                     `pulumi:"timeout"`
	TransformEncryption MLTransformTransformEncryptionPtrOutput `pulumi:"transformEncryption"`
	TransformParameters MLTransformTransformParametersOutput    `pulumi:"transformParameters"`
	WorkerType          pulumi.StringPtrOutput                  `pulumi:"workerType"`
}

// NewMLTransform registers a new resource with the given unique name, arguments, and options.
func NewMLTransform(ctx *pulumi.Context,
	name string, args *MLTransformArgs, opts ...pulumi.ResourceOption) (*MLTransform, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InputRecordTables == nil {
		return nil, errors.New("invalid value for required argument 'InputRecordTables'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	if args.TransformParameters == nil {
		return nil, errors.New("invalid value for required argument 'TransformParameters'")
	}
	var resource MLTransform
	err := ctx.RegisterResource("aws-native:glue:MLTransform", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMLTransform gets an existing MLTransform resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMLTransform(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MLTransformState, opts ...pulumi.ResourceOption) (*MLTransform, error) {
	var resource MLTransform
	err := ctx.ReadResource("aws-native:glue:MLTransform", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MLTransform resources.
type mltransformState struct {
}

type MLTransformState struct {
}

func (MLTransformState) ElementType() reflect.Type {
	return reflect.TypeOf((*mltransformState)(nil)).Elem()
}

type mltransformArgs struct {
	Description         *string                         `pulumi:"description"`
	GlueVersion         *string                         `pulumi:"glueVersion"`
	InputRecordTables   MLTransformInputRecordTables    `pulumi:"inputRecordTables"`
	MaxCapacity         *float64                        `pulumi:"maxCapacity"`
	MaxRetries          *int                            `pulumi:"maxRetries"`
	Name                *string                         `pulumi:"name"`
	NumberOfWorkers     *int                            `pulumi:"numberOfWorkers"`
	Role                string                          `pulumi:"role"`
	Tags                interface{}                     `pulumi:"tags"`
	Timeout             *int                            `pulumi:"timeout"`
	TransformEncryption *MLTransformTransformEncryption `pulumi:"transformEncryption"`
	TransformParameters MLTransformTransformParameters  `pulumi:"transformParameters"`
	WorkerType          *string                         `pulumi:"workerType"`
}

// The set of arguments for constructing a MLTransform resource.
type MLTransformArgs struct {
	Description         pulumi.StringPtrInput
	GlueVersion         pulumi.StringPtrInput
	InputRecordTables   MLTransformInputRecordTablesInput
	MaxCapacity         pulumi.Float64PtrInput
	MaxRetries          pulumi.IntPtrInput
	Name                pulumi.StringPtrInput
	NumberOfWorkers     pulumi.IntPtrInput
	Role                pulumi.StringInput
	Tags                pulumi.Input
	Timeout             pulumi.IntPtrInput
	TransformEncryption MLTransformTransformEncryptionPtrInput
	TransformParameters MLTransformTransformParametersInput
	WorkerType          pulumi.StringPtrInput
}

func (MLTransformArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*mltransformArgs)(nil)).Elem()
}

type MLTransformInput interface {
	pulumi.Input

	ToMLTransformOutput() MLTransformOutput
	ToMLTransformOutputWithContext(ctx context.Context) MLTransformOutput
}

func (*MLTransform) ElementType() reflect.Type {
	return reflect.TypeOf((**MLTransform)(nil)).Elem()
}

func (i *MLTransform) ToMLTransformOutput() MLTransformOutput {
	return i.ToMLTransformOutputWithContext(context.Background())
}

func (i *MLTransform) ToMLTransformOutputWithContext(ctx context.Context) MLTransformOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MLTransformOutput)
}

type MLTransformOutput struct{ *pulumi.OutputState }

func (MLTransformOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MLTransform)(nil)).Elem()
}

func (o MLTransformOutput) ToMLTransformOutput() MLTransformOutput {
	return o
}

func (o MLTransformOutput) ToMLTransformOutputWithContext(ctx context.Context) MLTransformOutput {
	return o
}

func (o MLTransformOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MLTransform) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o MLTransformOutput) GlueVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MLTransform) pulumi.StringPtrOutput { return v.GlueVersion }).(pulumi.StringPtrOutput)
}

func (o MLTransformOutput) InputRecordTables() MLTransformInputRecordTablesOutput {
	return o.ApplyT(func(v *MLTransform) MLTransformInputRecordTablesOutput { return v.InputRecordTables }).(MLTransformInputRecordTablesOutput)
}

func (o MLTransformOutput) MaxCapacity() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *MLTransform) pulumi.Float64PtrOutput { return v.MaxCapacity }).(pulumi.Float64PtrOutput)
}

func (o MLTransformOutput) MaxRetries() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *MLTransform) pulumi.IntPtrOutput { return v.MaxRetries }).(pulumi.IntPtrOutput)
}

func (o MLTransformOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MLTransform) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

func (o MLTransformOutput) NumberOfWorkers() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *MLTransform) pulumi.IntPtrOutput { return v.NumberOfWorkers }).(pulumi.IntPtrOutput)
}

func (o MLTransformOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *MLTransform) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

func (o MLTransformOutput) Tags() pulumi.AnyOutput {
	return o.ApplyT(func(v *MLTransform) pulumi.AnyOutput { return v.Tags }).(pulumi.AnyOutput)
}

func (o MLTransformOutput) Timeout() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *MLTransform) pulumi.IntPtrOutput { return v.Timeout }).(pulumi.IntPtrOutput)
}

func (o MLTransformOutput) TransformEncryption() MLTransformTransformEncryptionPtrOutput {
	return o.ApplyT(func(v *MLTransform) MLTransformTransformEncryptionPtrOutput { return v.TransformEncryption }).(MLTransformTransformEncryptionPtrOutput)
}

func (o MLTransformOutput) TransformParameters() MLTransformTransformParametersOutput {
	return o.ApplyT(func(v *MLTransform) MLTransformTransformParametersOutput { return v.TransformParameters }).(MLTransformTransformParametersOutput)
}

func (o MLTransformOutput) WorkerType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MLTransform) pulumi.StringPtrOutput { return v.WorkerType }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MLTransformInput)(nil)).Elem(), &MLTransform{})
	pulumi.RegisterOutputType(MLTransformOutput{})
}
