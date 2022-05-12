// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sagemaker

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::SageMaker::Project
type Project struct {
	pulumi.CustomResourceState

	// The time at which the project was created.
	CreationTime       pulumi.StringOutput    `pulumi:"creationTime"`
	ProjectArn         pulumi.StringOutput    `pulumi:"projectArn"`
	ProjectDescription pulumi.StringPtrOutput `pulumi:"projectDescription"`
	ProjectId          pulumi.StringOutput    `pulumi:"projectId"`
	ProjectName        pulumi.StringOutput    `pulumi:"projectName"`
	// The status of a project.
	ProjectStatus ProjectStatusOutput `pulumi:"projectStatus"`
	// Provisioned ServiceCatalog  Details
	ServiceCatalogProvisionedProductDetails ServiceCatalogProvisionedProductDetailsPropertiesOutput `pulumi:"serviceCatalogProvisionedProductDetails"`
	// Input ServiceCatalog Provisioning Details
	ServiceCatalogProvisioningDetails ServiceCatalogProvisioningDetailsPropertiesOutput `pulumi:"serviceCatalogProvisioningDetails"`
	// An array of key-value pairs to apply to this resource.
	Tags ProjectTagArrayOutput `pulumi:"tags"`
}

// NewProject registers a new resource with the given unique name, arguments, and options.
func NewProject(ctx *pulumi.Context,
	name string, args *ProjectArgs, opts ...pulumi.ResourceOption) (*Project, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ServiceCatalogProvisioningDetails == nil {
		return nil, errors.New("invalid value for required argument 'ServiceCatalogProvisioningDetails'")
	}
	var resource Project
	err := ctx.RegisterResource("aws-native:sagemaker:Project", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProject gets an existing Project resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProject(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProjectState, opts ...pulumi.ResourceOption) (*Project, error) {
	var resource Project
	err := ctx.ReadResource("aws-native:sagemaker:Project", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Project resources.
type projectState struct {
}

type ProjectState struct {
}

func (ProjectState) ElementType() reflect.Type {
	return reflect.TypeOf((*projectState)(nil)).Elem()
}

type projectArgs struct {
	ProjectDescription *string `pulumi:"projectDescription"`
	ProjectName        *string `pulumi:"projectName"`
	// Input ServiceCatalog Provisioning Details
	ServiceCatalogProvisioningDetails ServiceCatalogProvisioningDetailsProperties `pulumi:"serviceCatalogProvisioningDetails"`
	// An array of key-value pairs to apply to this resource.
	Tags []ProjectTag `pulumi:"tags"`
}

// The set of arguments for constructing a Project resource.
type ProjectArgs struct {
	ProjectDescription pulumi.StringPtrInput
	ProjectName        pulumi.StringPtrInput
	// Input ServiceCatalog Provisioning Details
	ServiceCatalogProvisioningDetails ServiceCatalogProvisioningDetailsPropertiesInput
	// An array of key-value pairs to apply to this resource.
	Tags ProjectTagArrayInput
}

func (ProjectArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*projectArgs)(nil)).Elem()
}

type ProjectInput interface {
	pulumi.Input

	ToProjectOutput() ProjectOutput
	ToProjectOutputWithContext(ctx context.Context) ProjectOutput
}

func (*Project) ElementType() reflect.Type {
	return reflect.TypeOf((**Project)(nil)).Elem()
}

func (i *Project) ToProjectOutput() ProjectOutput {
	return i.ToProjectOutputWithContext(context.Background())
}

func (i *Project) ToProjectOutputWithContext(ctx context.Context) ProjectOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectOutput)
}

type ProjectOutput struct{ *pulumi.OutputState }

func (ProjectOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Project)(nil)).Elem()
}

func (o ProjectOutput) ToProjectOutput() ProjectOutput {
	return o
}

func (o ProjectOutput) ToProjectOutputWithContext(ctx context.Context) ProjectOutput {
	return o
}

// The time at which the project was created.
func (o ProjectOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Project) pulumi.StringOutput { return v.CreationTime }).(pulumi.StringOutput)
}

func (o ProjectOutput) ProjectArn() pulumi.StringOutput {
	return o.ApplyT(func(v *Project) pulumi.StringOutput { return v.ProjectArn }).(pulumi.StringOutput)
}

func (o ProjectOutput) ProjectDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Project) pulumi.StringPtrOutput { return v.ProjectDescription }).(pulumi.StringPtrOutput)
}

func (o ProjectOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *Project) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

func (o ProjectOutput) ProjectName() pulumi.StringOutput {
	return o.ApplyT(func(v *Project) pulumi.StringOutput { return v.ProjectName }).(pulumi.StringOutput)
}

// The status of a project.
func (o ProjectOutput) ProjectStatus() ProjectStatusOutput {
	return o.ApplyT(func(v *Project) ProjectStatusOutput { return v.ProjectStatus }).(ProjectStatusOutput)
}

// Provisioned ServiceCatalog  Details
func (o ProjectOutput) ServiceCatalogProvisionedProductDetails() ServiceCatalogProvisionedProductDetailsPropertiesOutput {
	return o.ApplyT(func(v *Project) ServiceCatalogProvisionedProductDetailsPropertiesOutput {
		return v.ServiceCatalogProvisionedProductDetails
	}).(ServiceCatalogProvisionedProductDetailsPropertiesOutput)
}

// Input ServiceCatalog Provisioning Details
func (o ProjectOutput) ServiceCatalogProvisioningDetails() ServiceCatalogProvisioningDetailsPropertiesOutput {
	return o.ApplyT(func(v *Project) ServiceCatalogProvisioningDetailsPropertiesOutput {
		return v.ServiceCatalogProvisioningDetails
	}).(ServiceCatalogProvisioningDetailsPropertiesOutput)
}

// An array of key-value pairs to apply to this resource.
func (o ProjectOutput) Tags() ProjectTagArrayOutput {
	return o.ApplyT(func(v *Project) ProjectTagArrayOutput { return v.Tags }).(ProjectTagArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectInput)(nil)).Elem(), &Project{})
	pulumi.RegisterOutputType(ProjectOutput{})
}
