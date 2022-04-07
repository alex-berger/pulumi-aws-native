// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package devopsguru

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource schema represents the NotificationChannel resource in the Amazon DevOps Guru.
func LookupNotificationChannel(ctx *pulumi.Context, args *LookupNotificationChannelArgs, opts ...pulumi.InvokeOption) (*LookupNotificationChannelResult, error) {
	var rv LookupNotificationChannelResult
	err := ctx.Invoke("aws-native:devopsguru:getNotificationChannel", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupNotificationChannelArgs struct {
	// The ID of a notification channel.
	Id string `pulumi:"id"`
}

type LookupNotificationChannelResult struct {
	// The ID of a notification channel.
	Id *string `pulumi:"id"`
}

func LookupNotificationChannelOutput(ctx *pulumi.Context, args LookupNotificationChannelOutputArgs, opts ...pulumi.InvokeOption) LookupNotificationChannelResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupNotificationChannelResult, error) {
			args := v.(LookupNotificationChannelArgs)
			r, err := LookupNotificationChannel(ctx, &args, opts...)
			var s LookupNotificationChannelResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupNotificationChannelResultOutput)
}

type LookupNotificationChannelOutputArgs struct {
	// The ID of a notification channel.
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupNotificationChannelOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNotificationChannelArgs)(nil)).Elem()
}

type LookupNotificationChannelResultOutput struct{ *pulumi.OutputState }

func (LookupNotificationChannelResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNotificationChannelResult)(nil)).Elem()
}

func (o LookupNotificationChannelResultOutput) ToLookupNotificationChannelResultOutput() LookupNotificationChannelResultOutput {
	return o
}

func (o LookupNotificationChannelResultOutput) ToLookupNotificationChannelResultOutputWithContext(ctx context.Context) LookupNotificationChannelResultOutput {
	return o
}

// The ID of a notification channel.
func (o LookupNotificationChannelResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNotificationChannelResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNotificationChannelResultOutput{})
}
