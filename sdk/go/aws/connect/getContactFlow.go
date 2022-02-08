// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package connect

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Connect::ContactFlow
func LookupContactFlow(ctx *pulumi.Context, args *LookupContactFlowArgs, opts ...pulumi.InvokeOption) (*LookupContactFlowResult, error) {
	var rv LookupContactFlowResult
	err := ctx.Invoke("aws-native:connect:getContactFlow", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupContactFlowArgs struct {
	// The identifier of the contact flow (ARN).
	ContactFlowArn string `pulumi:"contactFlowArn"`
}

type LookupContactFlowResult struct {
	// The identifier of the contact flow (ARN).
	ContactFlowArn *string `pulumi:"contactFlowArn"`
	// The content of the contact flow in JSON format.
	Content *string `pulumi:"content"`
	// The description of the contact flow.
	Description *string `pulumi:"description"`
	// The identifier of the Amazon Connect instance (ARN).
	InstanceArn *string `pulumi:"instanceArn"`
	// The name of the contact flow.
	Name *string `pulumi:"name"`
	// The state of the contact flow.
	State *ContactFlowStateEnum `pulumi:"state"`
	// One or more tags.
	Tags []ContactFlowTag `pulumi:"tags"`
}

func LookupContactFlowOutput(ctx *pulumi.Context, args LookupContactFlowOutputArgs, opts ...pulumi.InvokeOption) LookupContactFlowResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupContactFlowResult, error) {
			args := v.(LookupContactFlowArgs)
			r, err := LookupContactFlow(ctx, &args, opts...)
			return *r, err
		}).(LookupContactFlowResultOutput)
}

type LookupContactFlowOutputArgs struct {
	// The identifier of the contact flow (ARN).
	ContactFlowArn pulumi.StringInput `pulumi:"contactFlowArn"`
}

func (LookupContactFlowOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupContactFlowArgs)(nil)).Elem()
}

type LookupContactFlowResultOutput struct{ *pulumi.OutputState }

func (LookupContactFlowResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupContactFlowResult)(nil)).Elem()
}

func (o LookupContactFlowResultOutput) ToLookupContactFlowResultOutput() LookupContactFlowResultOutput {
	return o
}

func (o LookupContactFlowResultOutput) ToLookupContactFlowResultOutputWithContext(ctx context.Context) LookupContactFlowResultOutput {
	return o
}

// The identifier of the contact flow (ARN).
func (o LookupContactFlowResultOutput) ContactFlowArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupContactFlowResult) *string { return v.ContactFlowArn }).(pulumi.StringPtrOutput)
}

// The content of the contact flow in JSON format.
func (o LookupContactFlowResultOutput) Content() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupContactFlowResult) *string { return v.Content }).(pulumi.StringPtrOutput)
}

// The description of the contact flow.
func (o LookupContactFlowResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupContactFlowResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// The identifier of the Amazon Connect instance (ARN).
func (o LookupContactFlowResultOutput) InstanceArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupContactFlowResult) *string { return v.InstanceArn }).(pulumi.StringPtrOutput)
}

// The name of the contact flow.
func (o LookupContactFlowResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupContactFlowResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// The state of the contact flow.
func (o LookupContactFlowResultOutput) State() ContactFlowStateEnumPtrOutput {
	return o.ApplyT(func(v LookupContactFlowResult) *ContactFlowStateEnum { return v.State }).(ContactFlowStateEnumPtrOutput)
}

// One or more tags.
func (o LookupContactFlowResultOutput) Tags() ContactFlowTagArrayOutput {
	return o.ApplyT(func(v LookupContactFlowResult) []ContactFlowTag { return v.Tags }).(ContactFlowTagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupContactFlowResultOutput{})
}