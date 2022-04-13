// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package appmesh

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::AppMesh::Mesh
func LookupMesh(ctx *pulumi.Context, args *LookupMeshArgs, opts ...pulumi.InvokeOption) (*LookupMeshResult, error) {
	var rv LookupMeshResult
	err := ctx.Invoke("aws-native:appmesh:getMesh", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupMeshArgs struct {
	Id string `pulumi:"id"`
}

type LookupMeshResult struct {
	Arn           *string   `pulumi:"arn"`
	Id            *string   `pulumi:"id"`
	MeshOwner     *string   `pulumi:"meshOwner"`
	ResourceOwner *string   `pulumi:"resourceOwner"`
	Spec          *MeshSpec `pulumi:"spec"`
	Tags          []MeshTag `pulumi:"tags"`
	Uid           *string   `pulumi:"uid"`
}

func LookupMeshOutput(ctx *pulumi.Context, args LookupMeshOutputArgs, opts ...pulumi.InvokeOption) LookupMeshResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupMeshResult, error) {
			args := v.(LookupMeshArgs)
			r, err := LookupMesh(ctx, &args, opts...)
			var s LookupMeshResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupMeshResultOutput)
}

type LookupMeshOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupMeshOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMeshArgs)(nil)).Elem()
}

type LookupMeshResultOutput struct{ *pulumi.OutputState }

func (LookupMeshResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMeshResult)(nil)).Elem()
}

func (o LookupMeshResultOutput) ToLookupMeshResultOutput() LookupMeshResultOutput {
	return o
}

func (o LookupMeshResultOutput) ToLookupMeshResultOutputWithContext(ctx context.Context) LookupMeshResultOutput {
	return o
}

func (o LookupMeshResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMeshResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

func (o LookupMeshResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMeshResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupMeshResultOutput) MeshOwner() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMeshResult) *string { return v.MeshOwner }).(pulumi.StringPtrOutput)
}

func (o LookupMeshResultOutput) ResourceOwner() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMeshResult) *string { return v.ResourceOwner }).(pulumi.StringPtrOutput)
}

func (o LookupMeshResultOutput) Spec() MeshSpecPtrOutput {
	return o.ApplyT(func(v LookupMeshResult) *MeshSpec { return v.Spec }).(MeshSpecPtrOutput)
}

func (o LookupMeshResultOutput) Tags() MeshTagArrayOutput {
	return o.ApplyT(func(v LookupMeshResult) []MeshTag { return v.Tags }).(MeshTagArrayOutput)
}

func (o LookupMeshResultOutput) Uid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMeshResult) *string { return v.Uid }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupMeshResultOutput{})
}
