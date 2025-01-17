// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package eventschemas

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::EventSchemas::Discoverer
func LookupDiscoverer(ctx *pulumi.Context, args *LookupDiscovererArgs, opts ...pulumi.InvokeOption) (*LookupDiscovererResult, error) {
	var rv LookupDiscovererResult
	err := ctx.Invoke("aws-native:eventschemas:getDiscoverer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDiscovererArgs struct {
	DiscovererId string `pulumi:"discovererId"`
}

type LookupDiscovererResult struct {
	CrossAccount  *bool                 `pulumi:"crossAccount"`
	Description   *string               `pulumi:"description"`
	DiscovererArn *string               `pulumi:"discovererArn"`
	DiscovererId  *string               `pulumi:"discovererId"`
	Tags          []DiscovererTagsEntry `pulumi:"tags"`
}

func LookupDiscovererOutput(ctx *pulumi.Context, args LookupDiscovererOutputArgs, opts ...pulumi.InvokeOption) LookupDiscovererResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDiscovererResult, error) {
			args := v.(LookupDiscovererArgs)
			r, err := LookupDiscoverer(ctx, &args, opts...)
			var s LookupDiscovererResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDiscovererResultOutput)
}

type LookupDiscovererOutputArgs struct {
	DiscovererId pulumi.StringInput `pulumi:"discovererId"`
}

func (LookupDiscovererOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDiscovererArgs)(nil)).Elem()
}

type LookupDiscovererResultOutput struct{ *pulumi.OutputState }

func (LookupDiscovererResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDiscovererResult)(nil)).Elem()
}

func (o LookupDiscovererResultOutput) ToLookupDiscovererResultOutput() LookupDiscovererResultOutput {
	return o
}

func (o LookupDiscovererResultOutput) ToLookupDiscovererResultOutputWithContext(ctx context.Context) LookupDiscovererResultOutput {
	return o
}

func (o LookupDiscovererResultOutput) CrossAccount() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupDiscovererResult) *bool { return v.CrossAccount }).(pulumi.BoolPtrOutput)
}

func (o LookupDiscovererResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDiscovererResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupDiscovererResultOutput) DiscovererArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDiscovererResult) *string { return v.DiscovererArn }).(pulumi.StringPtrOutput)
}

func (o LookupDiscovererResultOutput) DiscovererId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDiscovererResult) *string { return v.DiscovererId }).(pulumi.StringPtrOutput)
}

func (o LookupDiscovererResultOutput) Tags() DiscovererTagsEntryArrayOutput {
	return o.ApplyT(func(v LookupDiscovererResult) []DiscovererTagsEntry { return v.Tags }).(DiscovererTagsEntryArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDiscovererResultOutput{})
}
