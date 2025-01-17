// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloud9

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Cloud9::EnvironmentEC2
func LookupEnvironmentEC2(ctx *pulumi.Context, args *LookupEnvironmentEC2Args, opts ...pulumi.InvokeOption) (*LookupEnvironmentEC2Result, error) {
	var rv LookupEnvironmentEC2Result
	err := ctx.Invoke("aws-native:cloud9:getEnvironmentEC2", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupEnvironmentEC2Args struct {
	Id string `pulumi:"id"`
}

type LookupEnvironmentEC2Result struct {
	Arn         *string             `pulumi:"arn"`
	Description *string             `pulumi:"description"`
	Id          *string             `pulumi:"id"`
	Name        *string             `pulumi:"name"`
	Tags        []EnvironmentEC2Tag `pulumi:"tags"`
}

func LookupEnvironmentEC2Output(ctx *pulumi.Context, args LookupEnvironmentEC2OutputArgs, opts ...pulumi.InvokeOption) LookupEnvironmentEC2ResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupEnvironmentEC2Result, error) {
			args := v.(LookupEnvironmentEC2Args)
			r, err := LookupEnvironmentEC2(ctx, &args, opts...)
			var s LookupEnvironmentEC2Result
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupEnvironmentEC2ResultOutput)
}

type LookupEnvironmentEC2OutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupEnvironmentEC2OutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEnvironmentEC2Args)(nil)).Elem()
}

type LookupEnvironmentEC2ResultOutput struct{ *pulumi.OutputState }

func (LookupEnvironmentEC2ResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEnvironmentEC2Result)(nil)).Elem()
}

func (o LookupEnvironmentEC2ResultOutput) ToLookupEnvironmentEC2ResultOutput() LookupEnvironmentEC2ResultOutput {
	return o
}

func (o LookupEnvironmentEC2ResultOutput) ToLookupEnvironmentEC2ResultOutputWithContext(ctx context.Context) LookupEnvironmentEC2ResultOutput {
	return o
}

func (o LookupEnvironmentEC2ResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEnvironmentEC2Result) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

func (o LookupEnvironmentEC2ResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEnvironmentEC2Result) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupEnvironmentEC2ResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEnvironmentEC2Result) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupEnvironmentEC2ResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEnvironmentEC2Result) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupEnvironmentEC2ResultOutput) Tags() EnvironmentEC2TagArrayOutput {
	return o.ApplyT(func(v LookupEnvironmentEC2Result) []EnvironmentEC2Tag { return v.Tags }).(EnvironmentEC2TagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupEnvironmentEC2ResultOutput{})
}
