// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package customerprofiles

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A domain defined for 3rd party data source in Profile Service
func LookupDomain(ctx *pulumi.Context, args *LookupDomainArgs, opts ...pulumi.InvokeOption) (*LookupDomainResult, error) {
	var rv LookupDomainResult
	err := ctx.Invoke("aws-native:customerprofiles:getDomain", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDomainArgs struct {
	// The unique name of the domain.
	DomainName string `pulumi:"domainName"`
}

type LookupDomainResult struct {
	// The time of this integration got created
	CreatedAt *string `pulumi:"createdAt"`
	// The URL of the SQS dead letter queue
	DeadLetterQueueUrl *string `pulumi:"deadLetterQueueUrl"`
	// The default encryption key
	DefaultEncryptionKey *string `pulumi:"defaultEncryptionKey"`
	// The default number of days until the data within the domain expires.
	DefaultExpirationDays *int `pulumi:"defaultExpirationDays"`
	// The time of this integration got last updated at
	LastUpdatedAt *string `pulumi:"lastUpdatedAt"`
	// The tags (keys and values) associated with the domain
	Tags []DomainTag `pulumi:"tags"`
}

func LookupDomainOutput(ctx *pulumi.Context, args LookupDomainOutputArgs, opts ...pulumi.InvokeOption) LookupDomainResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDomainResult, error) {
			args := v.(LookupDomainArgs)
			r, err := LookupDomain(ctx, &args, opts...)
			var s LookupDomainResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDomainResultOutput)
}

type LookupDomainOutputArgs struct {
	// The unique name of the domain.
	DomainName pulumi.StringInput `pulumi:"domainName"`
}

func (LookupDomainOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDomainArgs)(nil)).Elem()
}

type LookupDomainResultOutput struct{ *pulumi.OutputState }

func (LookupDomainResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDomainResult)(nil)).Elem()
}

func (o LookupDomainResultOutput) ToLookupDomainResultOutput() LookupDomainResultOutput {
	return o
}

func (o LookupDomainResultOutput) ToLookupDomainResultOutputWithContext(ctx context.Context) LookupDomainResultOutput {
	return o
}

// The time of this integration got created
func (o LookupDomainResultOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDomainResult) *string { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

// The URL of the SQS dead letter queue
func (o LookupDomainResultOutput) DeadLetterQueueUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDomainResult) *string { return v.DeadLetterQueueUrl }).(pulumi.StringPtrOutput)
}

// The default encryption key
func (o LookupDomainResultOutput) DefaultEncryptionKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDomainResult) *string { return v.DefaultEncryptionKey }).(pulumi.StringPtrOutput)
}

// The default number of days until the data within the domain expires.
func (o LookupDomainResultOutput) DefaultExpirationDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupDomainResult) *int { return v.DefaultExpirationDays }).(pulumi.IntPtrOutput)
}

// The time of this integration got last updated at
func (o LookupDomainResultOutput) LastUpdatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDomainResult) *string { return v.LastUpdatedAt }).(pulumi.StringPtrOutput)
}

// The tags (keys and values) associated with the domain
func (o LookupDomainResultOutput) Tags() DomainTagArrayOutput {
	return o.ApplyT(func(v LookupDomainResult) []DomainTag { return v.Tags }).(DomainTagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDomainResultOutput{})
}
