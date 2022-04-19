// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package qldb

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::QLDB::Ledger
//
// Deprecated: Ledger is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type Ledger struct {
	pulumi.CustomResourceState

	DeletionProtection pulumi.BoolPtrOutput   `pulumi:"deletionProtection"`
	KmsKey             pulumi.StringPtrOutput `pulumi:"kmsKey"`
	Name               pulumi.StringPtrOutput `pulumi:"name"`
	PermissionsMode    pulumi.StringOutput    `pulumi:"permissionsMode"`
	Tags               LedgerTagArrayOutput   `pulumi:"tags"`
}

// NewLedger registers a new resource with the given unique name, arguments, and options.
func NewLedger(ctx *pulumi.Context,
	name string, args *LedgerArgs, opts ...pulumi.ResourceOption) (*Ledger, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PermissionsMode == nil {
		return nil, errors.New("invalid value for required argument 'PermissionsMode'")
	}
	var resource Ledger
	err := ctx.RegisterResource("aws-native:qldb:Ledger", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLedger gets an existing Ledger resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLedger(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LedgerState, opts ...pulumi.ResourceOption) (*Ledger, error) {
	var resource Ledger
	err := ctx.ReadResource("aws-native:qldb:Ledger", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Ledger resources.
type ledgerState struct {
}

type LedgerState struct {
}

func (LedgerState) ElementType() reflect.Type {
	return reflect.TypeOf((*ledgerState)(nil)).Elem()
}

type ledgerArgs struct {
	DeletionProtection *bool       `pulumi:"deletionProtection"`
	KmsKey             *string     `pulumi:"kmsKey"`
	Name               *string     `pulumi:"name"`
	PermissionsMode    string      `pulumi:"permissionsMode"`
	Tags               []LedgerTag `pulumi:"tags"`
}

// The set of arguments for constructing a Ledger resource.
type LedgerArgs struct {
	DeletionProtection pulumi.BoolPtrInput
	KmsKey             pulumi.StringPtrInput
	Name               pulumi.StringPtrInput
	PermissionsMode    pulumi.StringInput
	Tags               LedgerTagArrayInput
}

func (LedgerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ledgerArgs)(nil)).Elem()
}

type LedgerInput interface {
	pulumi.Input

	ToLedgerOutput() LedgerOutput
	ToLedgerOutputWithContext(ctx context.Context) LedgerOutput
}

func (*Ledger) ElementType() reflect.Type {
	return reflect.TypeOf((**Ledger)(nil)).Elem()
}

func (i *Ledger) ToLedgerOutput() LedgerOutput {
	return i.ToLedgerOutputWithContext(context.Background())
}

func (i *Ledger) ToLedgerOutputWithContext(ctx context.Context) LedgerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LedgerOutput)
}

type LedgerOutput struct{ *pulumi.OutputState }

func (LedgerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Ledger)(nil)).Elem()
}

func (o LedgerOutput) ToLedgerOutput() LedgerOutput {
	return o
}

func (o LedgerOutput) ToLedgerOutputWithContext(ctx context.Context) LedgerOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LedgerInput)(nil)).Elem(), &Ledger{})
	pulumi.RegisterOutputType(LedgerOutput{})
}
