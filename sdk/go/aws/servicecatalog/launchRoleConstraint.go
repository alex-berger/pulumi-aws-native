// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package servicecatalog

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-launchroleconstraint.html
type LaunchRoleConstraint struct {
	pulumi.CustomResourceState

	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-launchroleconstraint.html#cfn-servicecatalog-launchroleconstraint-acceptlanguage
	AcceptLanguage pulumi.StringPtrOutput `pulumi:"acceptLanguage"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-launchroleconstraint.html#cfn-servicecatalog-launchroleconstraint-description
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-launchroleconstraint.html#cfn-servicecatalog-launchroleconstraint-localrolename
	LocalRoleName pulumi.StringPtrOutput `pulumi:"localRoleName"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-launchroleconstraint.html#cfn-servicecatalog-launchroleconstraint-portfolioid
	PortfolioId pulumi.StringOutput `pulumi:"portfolioId"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-launchroleconstraint.html#cfn-servicecatalog-launchroleconstraint-productid
	ProductId pulumi.StringOutput `pulumi:"productId"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-launchroleconstraint.html#cfn-servicecatalog-launchroleconstraint-rolearn
	RoleArn pulumi.StringPtrOutput `pulumi:"roleArn"`
}

// NewLaunchRoleConstraint registers a new resource with the given unique name, arguments, and options.
func NewLaunchRoleConstraint(ctx *pulumi.Context,
	name string, args *LaunchRoleConstraintArgs, opts ...pulumi.ResourceOption) (*LaunchRoleConstraint, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PortfolioId == nil {
		return nil, errors.New("invalid value for required argument 'PortfolioId'")
	}
	if args.ProductId == nil {
		return nil, errors.New("invalid value for required argument 'ProductId'")
	}
	var resource LaunchRoleConstraint
	err := ctx.RegisterResource("aws-native:servicecatalog:LaunchRoleConstraint", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLaunchRoleConstraint gets an existing LaunchRoleConstraint resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLaunchRoleConstraint(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LaunchRoleConstraintState, opts ...pulumi.ResourceOption) (*LaunchRoleConstraint, error) {
	var resource LaunchRoleConstraint
	err := ctx.ReadResource("aws-native:servicecatalog:LaunchRoleConstraint", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LaunchRoleConstraint resources.
type launchRoleConstraintState struct {
}

type LaunchRoleConstraintState struct {
}

func (LaunchRoleConstraintState) ElementType() reflect.Type {
	return reflect.TypeOf((*launchRoleConstraintState)(nil)).Elem()
}

type launchRoleConstraintArgs struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-launchroleconstraint.html#cfn-servicecatalog-launchroleconstraint-acceptlanguage
	AcceptLanguage *string `pulumi:"acceptLanguage"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-launchroleconstraint.html#cfn-servicecatalog-launchroleconstraint-description
	Description *string `pulumi:"description"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-launchroleconstraint.html#cfn-servicecatalog-launchroleconstraint-localrolename
	LocalRoleName *string `pulumi:"localRoleName"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-launchroleconstraint.html#cfn-servicecatalog-launchroleconstraint-portfolioid
	PortfolioId string `pulumi:"portfolioId"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-launchroleconstraint.html#cfn-servicecatalog-launchroleconstraint-productid
	ProductId string `pulumi:"productId"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-launchroleconstraint.html#cfn-servicecatalog-launchroleconstraint-rolearn
	RoleArn *string `pulumi:"roleArn"`
}

// The set of arguments for constructing a LaunchRoleConstraint resource.
type LaunchRoleConstraintArgs struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-launchroleconstraint.html#cfn-servicecatalog-launchroleconstraint-acceptlanguage
	AcceptLanguage pulumi.StringPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-launchroleconstraint.html#cfn-servicecatalog-launchroleconstraint-description
	Description pulumi.StringPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-launchroleconstraint.html#cfn-servicecatalog-launchroleconstraint-localrolename
	LocalRoleName pulumi.StringPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-launchroleconstraint.html#cfn-servicecatalog-launchroleconstraint-portfolioid
	PortfolioId pulumi.StringInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-launchroleconstraint.html#cfn-servicecatalog-launchroleconstraint-productid
	ProductId pulumi.StringInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-launchroleconstraint.html#cfn-servicecatalog-launchroleconstraint-rolearn
	RoleArn pulumi.StringPtrInput
}

func (LaunchRoleConstraintArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*launchRoleConstraintArgs)(nil)).Elem()
}

type LaunchRoleConstraintInput interface {
	pulumi.Input

	ToLaunchRoleConstraintOutput() LaunchRoleConstraintOutput
	ToLaunchRoleConstraintOutputWithContext(ctx context.Context) LaunchRoleConstraintOutput
}

func (*LaunchRoleConstraint) ElementType() reflect.Type {
	return reflect.TypeOf((*LaunchRoleConstraint)(nil))
}

func (i *LaunchRoleConstraint) ToLaunchRoleConstraintOutput() LaunchRoleConstraintOutput {
	return i.ToLaunchRoleConstraintOutputWithContext(context.Background())
}

func (i *LaunchRoleConstraint) ToLaunchRoleConstraintOutputWithContext(ctx context.Context) LaunchRoleConstraintOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LaunchRoleConstraintOutput)
}

type LaunchRoleConstraintOutput struct{ *pulumi.OutputState }

func (LaunchRoleConstraintOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LaunchRoleConstraint)(nil))
}

func (o LaunchRoleConstraintOutput) ToLaunchRoleConstraintOutput() LaunchRoleConstraintOutput {
	return o
}

func (o LaunchRoleConstraintOutput) ToLaunchRoleConstraintOutputWithContext(ctx context.Context) LaunchRoleConstraintOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(LaunchRoleConstraintOutput{})
}