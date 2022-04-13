// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package elasticache

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::ElastiCache::User
func LookupUser(ctx *pulumi.Context, args *LookupUserArgs, opts ...pulumi.InvokeOption) (*LookupUserResult, error) {
	var rv LookupUserResult
	err := ctx.Invoke("aws-native:elasticache:getUser", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupUserArgs struct {
	// The ID of the user.
	UserId string `pulumi:"userId"`
}

type LookupUserResult struct {
	// The Amazon Resource Name (ARN) of the user account.
	Arn *string `pulumi:"arn"`
	// Indicates the user status. Can be "active", "modifying" or "deleting".
	Status *string `pulumi:"status"`
}

func LookupUserOutput(ctx *pulumi.Context, args LookupUserOutputArgs, opts ...pulumi.InvokeOption) LookupUserResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupUserResult, error) {
			args := v.(LookupUserArgs)
			r, err := LookupUser(ctx, &args, opts...)
			var s LookupUserResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupUserResultOutput)
}

type LookupUserOutputArgs struct {
	// The ID of the user.
	UserId pulumi.StringInput `pulumi:"userId"`
}

func (LookupUserOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupUserArgs)(nil)).Elem()
}

type LookupUserResultOutput struct{ *pulumi.OutputState }

func (LookupUserResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupUserResult)(nil)).Elem()
}

func (o LookupUserResultOutput) ToLookupUserResultOutput() LookupUserResultOutput {
	return o
}

func (o LookupUserResultOutput) ToLookupUserResultOutputWithContext(ctx context.Context) LookupUserResultOutput {
	return o
}

// The Amazon Resource Name (ARN) of the user account.
func (o LookupUserResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUserResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

// Indicates the user status. Can be "active", "modifying" or "deleting".
func (o LookupUserResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUserResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupUserResultOutput{})
}
