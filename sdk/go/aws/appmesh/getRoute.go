// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package appmesh

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::AppMesh::Route
func LookupRoute(ctx *pulumi.Context, args *LookupRouteArgs, opts ...pulumi.InvokeOption) (*LookupRouteResult, error) {
	var rv LookupRouteResult
	err := ctx.Invoke("aws-native:appmesh:getRoute", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRouteArgs struct {
	Id string `pulumi:"id"`
}

type LookupRouteResult struct {
	Arn           *string    `pulumi:"arn"`
	Id            *string    `pulumi:"id"`
	ResourceOwner *string    `pulumi:"resourceOwner"`
	Spec          *RouteSpec `pulumi:"spec"`
	Tags          []RouteTag `pulumi:"tags"`
	Uid           *string    `pulumi:"uid"`
}

func LookupRouteOutput(ctx *pulumi.Context, args LookupRouteOutputArgs, opts ...pulumi.InvokeOption) LookupRouteResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRouteResult, error) {
			args := v.(LookupRouteArgs)
			r, err := LookupRoute(ctx, &args, opts...)
			var s LookupRouteResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRouteResultOutput)
}

type LookupRouteOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupRouteOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRouteArgs)(nil)).Elem()
}

type LookupRouteResultOutput struct{ *pulumi.OutputState }

func (LookupRouteResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRouteResult)(nil)).Elem()
}

func (o LookupRouteResultOutput) ToLookupRouteResultOutput() LookupRouteResultOutput {
	return o
}

func (o LookupRouteResultOutput) ToLookupRouteResultOutputWithContext(ctx context.Context) LookupRouteResultOutput {
	return o
}

func (o LookupRouteResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRouteResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

func (o LookupRouteResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRouteResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupRouteResultOutput) ResourceOwner() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRouteResult) *string { return v.ResourceOwner }).(pulumi.StringPtrOutput)
}

func (o LookupRouteResultOutput) Spec() RouteSpecPtrOutput {
	return o.ApplyT(func(v LookupRouteResult) *RouteSpec { return v.Spec }).(RouteSpecPtrOutput)
}

func (o LookupRouteResultOutput) Tags() RouteTagArrayOutput {
	return o.ApplyT(func(v LookupRouteResult) []RouteTag { return v.Tags }).(RouteTagArrayOutput)
}

func (o LookupRouteResultOutput) Uid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRouteResult) *string { return v.Uid }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRouteResultOutput{})
}
