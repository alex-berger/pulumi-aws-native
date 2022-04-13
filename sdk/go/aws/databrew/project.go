// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package databrew

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource schema for AWS::DataBrew::Project.
type Project struct {
	pulumi.CustomResourceState

	// Dataset name
	DatasetName pulumi.StringOutput `pulumi:"datasetName"`
	// Project name
	Name pulumi.StringOutput `pulumi:"name"`
	// Recipe name
	RecipeName pulumi.StringOutput `pulumi:"recipeName"`
	// Role arn
	RoleArn pulumi.StringOutput `pulumi:"roleArn"`
	// Sample
	Sample ProjectSamplePtrOutput `pulumi:"sample"`
	Tags   ProjectTagArrayOutput  `pulumi:"tags"`
}

// NewProject registers a new resource with the given unique name, arguments, and options.
func NewProject(ctx *pulumi.Context,
	name string, args *ProjectArgs, opts ...pulumi.ResourceOption) (*Project, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DatasetName == nil {
		return nil, errors.New("invalid value for required argument 'DatasetName'")
	}
	if args.RecipeName == nil {
		return nil, errors.New("invalid value for required argument 'RecipeName'")
	}
	if args.RoleArn == nil {
		return nil, errors.New("invalid value for required argument 'RoleArn'")
	}
	var resource Project
	err := ctx.RegisterResource("aws-native:databrew:Project", name, args, &resource, opts...)
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
	err := ctx.ReadResource("aws-native:databrew:Project", name, id, state, &resource, opts...)
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
	// Dataset name
	DatasetName string `pulumi:"datasetName"`
	// Project name
	Name *string `pulumi:"name"`
	// Recipe name
	RecipeName string `pulumi:"recipeName"`
	// Role arn
	RoleArn string `pulumi:"roleArn"`
	// Sample
	Sample *ProjectSample `pulumi:"sample"`
	Tags   []ProjectTag   `pulumi:"tags"`
}

// The set of arguments for constructing a Project resource.
type ProjectArgs struct {
	// Dataset name
	DatasetName pulumi.StringInput
	// Project name
	Name pulumi.StringPtrInput
	// Recipe name
	RecipeName pulumi.StringInput
	// Role arn
	RoleArn pulumi.StringInput
	// Sample
	Sample ProjectSamplePtrInput
	Tags   ProjectTagArrayInput
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

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectInput)(nil)).Elem(), &Project{})
	pulumi.RegisterOutputType(ProjectOutput{})
}
