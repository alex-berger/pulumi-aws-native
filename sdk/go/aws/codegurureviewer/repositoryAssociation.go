// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package codegurureviewer

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource schema represents the RepositoryAssociation resource in the Amazon CodeGuru Reviewer service.
type RepositoryAssociation struct {
	pulumi.CustomResourceState

	// The Amazon Resource Name (ARN) of the repository association.
	AssociationArn pulumi.StringOutput `pulumi:"associationArn"`
	// The name of the S3 bucket associated with an associated S3 repository. It must start with `codeguru-reviewer-`.
	BucketName pulumi.StringPtrOutput `pulumi:"bucketName"`
	// The Amazon Resource Name (ARN) of an AWS CodeStar Connections connection.
	ConnectionArn pulumi.StringPtrOutput `pulumi:"connectionArn"`
	// Name of the repository to be associated.
	Name pulumi.StringOutput `pulumi:"name"`
	// The owner of the repository. For a Bitbucket repository, this is the username for the account that owns the repository.
	Owner pulumi.StringPtrOutput `pulumi:"owner"`
	// The tags associated with a repository association.
	Tags RepositoryAssociationTagArrayOutput `pulumi:"tags"`
	// The type of repository to be associated.
	Type RepositoryAssociationTypeOutput `pulumi:"type"`
}

// NewRepositoryAssociation registers a new resource with the given unique name, arguments, and options.
func NewRepositoryAssociation(ctx *pulumi.Context,
	name string, args *RepositoryAssociationArgs, opts ...pulumi.ResourceOption) (*RepositoryAssociation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	var resource RepositoryAssociation
	err := ctx.RegisterResource("aws-native:codegurureviewer:RepositoryAssociation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRepositoryAssociation gets an existing RepositoryAssociation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRepositoryAssociation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RepositoryAssociationState, opts ...pulumi.ResourceOption) (*RepositoryAssociation, error) {
	var resource RepositoryAssociation
	err := ctx.ReadResource("aws-native:codegurureviewer:RepositoryAssociation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RepositoryAssociation resources.
type repositoryAssociationState struct {
}

type RepositoryAssociationState struct {
}

func (RepositoryAssociationState) ElementType() reflect.Type {
	return reflect.TypeOf((*repositoryAssociationState)(nil)).Elem()
}

type repositoryAssociationArgs struct {
	// The name of the S3 bucket associated with an associated S3 repository. It must start with `codeguru-reviewer-`.
	BucketName *string `pulumi:"bucketName"`
	// The Amazon Resource Name (ARN) of an AWS CodeStar Connections connection.
	ConnectionArn *string `pulumi:"connectionArn"`
	// Name of the repository to be associated.
	Name *string `pulumi:"name"`
	// The owner of the repository. For a Bitbucket repository, this is the username for the account that owns the repository.
	Owner *string `pulumi:"owner"`
	// The tags associated with a repository association.
	Tags []RepositoryAssociationTag `pulumi:"tags"`
	// The type of repository to be associated.
	Type RepositoryAssociationType `pulumi:"type"`
}

// The set of arguments for constructing a RepositoryAssociation resource.
type RepositoryAssociationArgs struct {
	// The name of the S3 bucket associated with an associated S3 repository. It must start with `codeguru-reviewer-`.
	BucketName pulumi.StringPtrInput
	// The Amazon Resource Name (ARN) of an AWS CodeStar Connections connection.
	ConnectionArn pulumi.StringPtrInput
	// Name of the repository to be associated.
	Name pulumi.StringPtrInput
	// The owner of the repository. For a Bitbucket repository, this is the username for the account that owns the repository.
	Owner pulumi.StringPtrInput
	// The tags associated with a repository association.
	Tags RepositoryAssociationTagArrayInput
	// The type of repository to be associated.
	Type RepositoryAssociationTypeInput
}

func (RepositoryAssociationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*repositoryAssociationArgs)(nil)).Elem()
}

type RepositoryAssociationInput interface {
	pulumi.Input

	ToRepositoryAssociationOutput() RepositoryAssociationOutput
	ToRepositoryAssociationOutputWithContext(ctx context.Context) RepositoryAssociationOutput
}

func (*RepositoryAssociation) ElementType() reflect.Type {
	return reflect.TypeOf((**RepositoryAssociation)(nil)).Elem()
}

func (i *RepositoryAssociation) ToRepositoryAssociationOutput() RepositoryAssociationOutput {
	return i.ToRepositoryAssociationOutputWithContext(context.Background())
}

func (i *RepositoryAssociation) ToRepositoryAssociationOutputWithContext(ctx context.Context) RepositoryAssociationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryAssociationOutput)
}

type RepositoryAssociationOutput struct{ *pulumi.OutputState }

func (RepositoryAssociationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RepositoryAssociation)(nil)).Elem()
}

func (o RepositoryAssociationOutput) ToRepositoryAssociationOutput() RepositoryAssociationOutput {
	return o
}

func (o RepositoryAssociationOutput) ToRepositoryAssociationOutputWithContext(ctx context.Context) RepositoryAssociationOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RepositoryAssociationInput)(nil)).Elem(), &RepositoryAssociation{})
	pulumi.RegisterOutputType(RepositoryAssociationOutput{})
}
