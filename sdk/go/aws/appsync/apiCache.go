// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package appsync

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::AppSync::ApiCache
//
// Deprecated: ApiCache is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type ApiCache struct {
	pulumi.CustomResourceState

	ApiCachingBehavior       pulumi.StringOutput  `pulumi:"apiCachingBehavior"`
	ApiId                    pulumi.StringOutput  `pulumi:"apiId"`
	AtRestEncryptionEnabled  pulumi.BoolPtrOutput `pulumi:"atRestEncryptionEnabled"`
	TransitEncryptionEnabled pulumi.BoolPtrOutput `pulumi:"transitEncryptionEnabled"`
	Ttl                      pulumi.Float64Output `pulumi:"ttl"`
	Type                     pulumi.StringOutput  `pulumi:"type"`
}

// NewApiCache registers a new resource with the given unique name, arguments, and options.
func NewApiCache(ctx *pulumi.Context,
	name string, args *ApiCacheArgs, opts ...pulumi.ResourceOption) (*ApiCache, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApiCachingBehavior == nil {
		return nil, errors.New("invalid value for required argument 'ApiCachingBehavior'")
	}
	if args.ApiId == nil {
		return nil, errors.New("invalid value for required argument 'ApiId'")
	}
	if args.Ttl == nil {
		return nil, errors.New("invalid value for required argument 'Ttl'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	var resource ApiCache
	err := ctx.RegisterResource("aws-native:appsync:ApiCache", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApiCache gets an existing ApiCache resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApiCache(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApiCacheState, opts ...pulumi.ResourceOption) (*ApiCache, error) {
	var resource ApiCache
	err := ctx.ReadResource("aws-native:appsync:ApiCache", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApiCache resources.
type apiCacheState struct {
}

type ApiCacheState struct {
}

func (ApiCacheState) ElementType() reflect.Type {
	return reflect.TypeOf((*apiCacheState)(nil)).Elem()
}

type apiCacheArgs struct {
	ApiCachingBehavior       string  `pulumi:"apiCachingBehavior"`
	ApiId                    string  `pulumi:"apiId"`
	AtRestEncryptionEnabled  *bool   `pulumi:"atRestEncryptionEnabled"`
	TransitEncryptionEnabled *bool   `pulumi:"transitEncryptionEnabled"`
	Ttl                      float64 `pulumi:"ttl"`
	Type                     string  `pulumi:"type"`
}

// The set of arguments for constructing a ApiCache resource.
type ApiCacheArgs struct {
	ApiCachingBehavior       pulumi.StringInput
	ApiId                    pulumi.StringInput
	AtRestEncryptionEnabled  pulumi.BoolPtrInput
	TransitEncryptionEnabled pulumi.BoolPtrInput
	Ttl                      pulumi.Float64Input
	Type                     pulumi.StringInput
}

func (ApiCacheArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*apiCacheArgs)(nil)).Elem()
}

type ApiCacheInput interface {
	pulumi.Input

	ToApiCacheOutput() ApiCacheOutput
	ToApiCacheOutputWithContext(ctx context.Context) ApiCacheOutput
}

func (*ApiCache) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiCache)(nil)).Elem()
}

func (i *ApiCache) ToApiCacheOutput() ApiCacheOutput {
	return i.ToApiCacheOutputWithContext(context.Background())
}

func (i *ApiCache) ToApiCacheOutputWithContext(ctx context.Context) ApiCacheOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiCacheOutput)
}

type ApiCacheOutput struct{ *pulumi.OutputState }

func (ApiCacheOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiCache)(nil)).Elem()
}

func (o ApiCacheOutput) ToApiCacheOutput() ApiCacheOutput {
	return o
}

func (o ApiCacheOutput) ToApiCacheOutputWithContext(ctx context.Context) ApiCacheOutput {
	return o
}

func (o ApiCacheOutput) ApiCachingBehavior() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiCache) pulumi.StringOutput { return v.ApiCachingBehavior }).(pulumi.StringOutput)
}

func (o ApiCacheOutput) ApiId() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiCache) pulumi.StringOutput { return v.ApiId }).(pulumi.StringOutput)
}

func (o ApiCacheOutput) AtRestEncryptionEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ApiCache) pulumi.BoolPtrOutput { return v.AtRestEncryptionEnabled }).(pulumi.BoolPtrOutput)
}

func (o ApiCacheOutput) TransitEncryptionEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ApiCache) pulumi.BoolPtrOutput { return v.TransitEncryptionEnabled }).(pulumi.BoolPtrOutput)
}

func (o ApiCacheOutput) Ttl() pulumi.Float64Output {
	return o.ApplyT(func(v *ApiCache) pulumi.Float64Output { return v.Ttl }).(pulumi.Float64Output)
}

func (o ApiCacheOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiCache) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ApiCacheInput)(nil)).Elem(), &ApiCache{})
	pulumi.RegisterOutputType(ApiCacheOutput{})
}
