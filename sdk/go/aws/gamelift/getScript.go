// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package gamelift

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::GameLift::Script
func LookupScript(ctx *pulumi.Context, args *LookupScriptArgs, opts ...pulumi.InvokeOption) (*LookupScriptResult, error) {
	var rv LookupScriptResult
	err := ctx.Invoke("aws-native:gamelift:getScript", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupScriptArgs struct {
	Id string `pulumi:"id"`
}

type LookupScriptResult struct {
	Arn             *string           `pulumi:"arn"`
	Id              *string           `pulumi:"id"`
	Name            *string           `pulumi:"name"`
	StorageLocation *ScriptS3Location `pulumi:"storageLocation"`
	Tags            []ScriptTag       `pulumi:"tags"`
	Version         *string           `pulumi:"version"`
}

func LookupScriptOutput(ctx *pulumi.Context, args LookupScriptOutputArgs, opts ...pulumi.InvokeOption) LookupScriptResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupScriptResult, error) {
			args := v.(LookupScriptArgs)
			r, err := LookupScript(ctx, &args, opts...)
			var s LookupScriptResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupScriptResultOutput)
}

type LookupScriptOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupScriptOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupScriptArgs)(nil)).Elem()
}

type LookupScriptResultOutput struct{ *pulumi.OutputState }

func (LookupScriptResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupScriptResult)(nil)).Elem()
}

func (o LookupScriptResultOutput) ToLookupScriptResultOutput() LookupScriptResultOutput {
	return o
}

func (o LookupScriptResultOutput) ToLookupScriptResultOutputWithContext(ctx context.Context) LookupScriptResultOutput {
	return o
}

func (o LookupScriptResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScriptResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

func (o LookupScriptResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScriptResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupScriptResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScriptResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupScriptResultOutput) StorageLocation() ScriptS3LocationPtrOutput {
	return o.ApplyT(func(v LookupScriptResult) *ScriptS3Location { return v.StorageLocation }).(ScriptS3LocationPtrOutput)
}

func (o LookupScriptResultOutput) Tags() ScriptTagArrayOutput {
	return o.ApplyT(func(v LookupScriptResult) []ScriptTag { return v.Tags }).(ScriptTagArrayOutput)
}

func (o LookupScriptResultOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScriptResult) *string { return v.Version }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupScriptResultOutput{})
}
