// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package licensemanager

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::LicenseManager::License
type License struct {
	pulumi.CustomResourceState

	// Beneficiary of the license.
	Beneficiary              pulumi.StringPtrOutput                `pulumi:"beneficiary"`
	ConsumptionConfiguration LicenseConsumptionConfigurationOutput `pulumi:"consumptionConfiguration"`
	Entitlements             LicenseEntitlementArrayOutput         `pulumi:"entitlements"`
	// Home region for the created license.
	HomeRegion pulumi.StringOutput     `pulumi:"homeRegion"`
	Issuer     LicenseIssuerDataOutput `pulumi:"issuer"`
	// Amazon Resource Name is a unique name for each resource.
	LicenseArn      pulumi.StringOutput        `pulumi:"licenseArn"`
	LicenseMetadata LicenseMetadataArrayOutput `pulumi:"licenseMetadata"`
	// Name for the created license.
	LicenseName pulumi.StringOutput `pulumi:"licenseName"`
	// Product name for the created license.
	ProductName pulumi.StringOutput `pulumi:"productName"`
	// ProductSKU of the license.
	ProductSKU pulumi.StringPtrOutput          `pulumi:"productSKU"`
	Status     pulumi.StringPtrOutput          `pulumi:"status"`
	Validity   LicenseValidityDateFormatOutput `pulumi:"validity"`
	// The version of the license.
	Version pulumi.StringOutput `pulumi:"version"`
}

// NewLicense registers a new resource with the given unique name, arguments, and options.
func NewLicense(ctx *pulumi.Context,
	name string, args *LicenseArgs, opts ...pulumi.ResourceOption) (*License, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ConsumptionConfiguration == nil {
		return nil, errors.New("invalid value for required argument 'ConsumptionConfiguration'")
	}
	if args.Entitlements == nil {
		return nil, errors.New("invalid value for required argument 'Entitlements'")
	}
	if args.HomeRegion == nil {
		return nil, errors.New("invalid value for required argument 'HomeRegion'")
	}
	if args.Issuer == nil {
		return nil, errors.New("invalid value for required argument 'Issuer'")
	}
	if args.ProductName == nil {
		return nil, errors.New("invalid value for required argument 'ProductName'")
	}
	if args.Validity == nil {
		return nil, errors.New("invalid value for required argument 'Validity'")
	}
	var resource License
	err := ctx.RegisterResource("aws-native:licensemanager:License", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLicense gets an existing License resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLicense(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LicenseState, opts ...pulumi.ResourceOption) (*License, error) {
	var resource License
	err := ctx.ReadResource("aws-native:licensemanager:License", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering License resources.
type licenseState struct {
}

type LicenseState struct {
}

func (LicenseState) ElementType() reflect.Type {
	return reflect.TypeOf((*licenseState)(nil)).Elem()
}

type licenseArgs struct {
	// Beneficiary of the license.
	Beneficiary              *string                         `pulumi:"beneficiary"`
	ConsumptionConfiguration LicenseConsumptionConfiguration `pulumi:"consumptionConfiguration"`
	Entitlements             []LicenseEntitlement            `pulumi:"entitlements"`
	// Home region for the created license.
	HomeRegion      string            `pulumi:"homeRegion"`
	Issuer          LicenseIssuerData `pulumi:"issuer"`
	LicenseMetadata []LicenseMetadata `pulumi:"licenseMetadata"`
	// Name for the created license.
	LicenseName *string `pulumi:"licenseName"`
	// Product name for the created license.
	ProductName string `pulumi:"productName"`
	// ProductSKU of the license.
	ProductSKU *string                   `pulumi:"productSKU"`
	Status     *string                   `pulumi:"status"`
	Validity   LicenseValidityDateFormat `pulumi:"validity"`
}

// The set of arguments for constructing a License resource.
type LicenseArgs struct {
	// Beneficiary of the license.
	Beneficiary              pulumi.StringPtrInput
	ConsumptionConfiguration LicenseConsumptionConfigurationInput
	Entitlements             LicenseEntitlementArrayInput
	// Home region for the created license.
	HomeRegion      pulumi.StringInput
	Issuer          LicenseIssuerDataInput
	LicenseMetadata LicenseMetadataArrayInput
	// Name for the created license.
	LicenseName pulumi.StringPtrInput
	// Product name for the created license.
	ProductName pulumi.StringInput
	// ProductSKU of the license.
	ProductSKU pulumi.StringPtrInput
	Status     pulumi.StringPtrInput
	Validity   LicenseValidityDateFormatInput
}

func (LicenseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*licenseArgs)(nil)).Elem()
}

type LicenseInput interface {
	pulumi.Input

	ToLicenseOutput() LicenseOutput
	ToLicenseOutputWithContext(ctx context.Context) LicenseOutput
}

func (*License) ElementType() reflect.Type {
	return reflect.TypeOf((**License)(nil)).Elem()
}

func (i *License) ToLicenseOutput() LicenseOutput {
	return i.ToLicenseOutputWithContext(context.Background())
}

func (i *License) ToLicenseOutputWithContext(ctx context.Context) LicenseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LicenseOutput)
}

type LicenseOutput struct{ *pulumi.OutputState }

func (LicenseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**License)(nil)).Elem()
}

func (o LicenseOutput) ToLicenseOutput() LicenseOutput {
	return o
}

func (o LicenseOutput) ToLicenseOutputWithContext(ctx context.Context) LicenseOutput {
	return o
}

// Beneficiary of the license.
func (o LicenseOutput) Beneficiary() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *License) pulumi.StringPtrOutput { return v.Beneficiary }).(pulumi.StringPtrOutput)
}

func (o LicenseOutput) ConsumptionConfiguration() LicenseConsumptionConfigurationOutput {
	return o.ApplyT(func(v *License) LicenseConsumptionConfigurationOutput { return v.ConsumptionConfiguration }).(LicenseConsumptionConfigurationOutput)
}

func (o LicenseOutput) Entitlements() LicenseEntitlementArrayOutput {
	return o.ApplyT(func(v *License) LicenseEntitlementArrayOutput { return v.Entitlements }).(LicenseEntitlementArrayOutput)
}

// Home region for the created license.
func (o LicenseOutput) HomeRegion() pulumi.StringOutput {
	return o.ApplyT(func(v *License) pulumi.StringOutput { return v.HomeRegion }).(pulumi.StringOutput)
}

func (o LicenseOutput) Issuer() LicenseIssuerDataOutput {
	return o.ApplyT(func(v *License) LicenseIssuerDataOutput { return v.Issuer }).(LicenseIssuerDataOutput)
}

// Amazon Resource Name is a unique name for each resource.
func (o LicenseOutput) LicenseArn() pulumi.StringOutput {
	return o.ApplyT(func(v *License) pulumi.StringOutput { return v.LicenseArn }).(pulumi.StringOutput)
}

func (o LicenseOutput) LicenseMetadata() LicenseMetadataArrayOutput {
	return o.ApplyT(func(v *License) LicenseMetadataArrayOutput { return v.LicenseMetadata }).(LicenseMetadataArrayOutput)
}

// Name for the created license.
func (o LicenseOutput) LicenseName() pulumi.StringOutput {
	return o.ApplyT(func(v *License) pulumi.StringOutput { return v.LicenseName }).(pulumi.StringOutput)
}

// Product name for the created license.
func (o LicenseOutput) ProductName() pulumi.StringOutput {
	return o.ApplyT(func(v *License) pulumi.StringOutput { return v.ProductName }).(pulumi.StringOutput)
}

// ProductSKU of the license.
func (o LicenseOutput) ProductSKU() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *License) pulumi.StringPtrOutput { return v.ProductSKU }).(pulumi.StringPtrOutput)
}

func (o LicenseOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *License) pulumi.StringPtrOutput { return v.Status }).(pulumi.StringPtrOutput)
}

func (o LicenseOutput) Validity() LicenseValidityDateFormatOutput {
	return o.ApplyT(func(v *License) LicenseValidityDateFormatOutput { return v.Validity }).(LicenseValidityDateFormatOutput)
}

// The version of the license.
func (o LicenseOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v *License) pulumi.StringOutput { return v.Version }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LicenseInput)(nil)).Elem(), &License{})
	pulumi.RegisterOutputType(LicenseOutput{})
}
