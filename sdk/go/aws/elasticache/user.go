// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package elasticache

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::ElastiCache::User
type User struct {
	pulumi.CustomResourceState

	// Access permissions string used for this user account.
	AccessString pulumi.StringPtrOutput `pulumi:"accessString"`
	// The Amazon Resource Name (ARN) of the user account.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Must be redis.
	Engine UserEngineOutput `pulumi:"engine"`
	// Indicates a password is not required for this user account.
	NoPasswordRequired pulumi.BoolPtrOutput `pulumi:"noPasswordRequired"`
	// Passwords used for this user account. You can create up to two passwords for each user.
	Passwords pulumi.StringArrayOutput `pulumi:"passwords"`
	// Indicates the user status. Can be "active", "modifying" or "deleting".
	Status pulumi.StringOutput `pulumi:"status"`
	// The ID of the user.
	UserId pulumi.StringOutput `pulumi:"userId"`
	// The username of the user.
	UserName pulumi.StringOutput `pulumi:"userName"`
}

// NewUser registers a new resource with the given unique name, arguments, and options.
func NewUser(ctx *pulumi.Context,
	name string, args *UserArgs, opts ...pulumi.ResourceOption) (*User, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Engine == nil {
		return nil, errors.New("invalid value for required argument 'Engine'")
	}
	if args.UserId == nil {
		return nil, errors.New("invalid value for required argument 'UserId'")
	}
	var resource User
	err := ctx.RegisterResource("aws-native:elasticache:User", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUser gets an existing User resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUser(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UserState, opts ...pulumi.ResourceOption) (*User, error) {
	var resource User
	err := ctx.ReadResource("aws-native:elasticache:User", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering User resources.
type userState struct {
}

type UserState struct {
}

func (UserState) ElementType() reflect.Type {
	return reflect.TypeOf((*userState)(nil)).Elem()
}

type userArgs struct {
	// Access permissions string used for this user account.
	AccessString *string `pulumi:"accessString"`
	// Must be redis.
	Engine UserEngine `pulumi:"engine"`
	// Indicates a password is not required for this user account.
	NoPasswordRequired *bool `pulumi:"noPasswordRequired"`
	// Passwords used for this user account. You can create up to two passwords for each user.
	Passwords []string `pulumi:"passwords"`
	// The ID of the user.
	UserId string `pulumi:"userId"`
	// The username of the user.
	UserName *string `pulumi:"userName"`
}

// The set of arguments for constructing a User resource.
type UserArgs struct {
	// Access permissions string used for this user account.
	AccessString pulumi.StringPtrInput
	// Must be redis.
	Engine UserEngineInput
	// Indicates a password is not required for this user account.
	NoPasswordRequired pulumi.BoolPtrInput
	// Passwords used for this user account. You can create up to two passwords for each user.
	Passwords pulumi.StringArrayInput
	// The ID of the user.
	UserId pulumi.StringInput
	// The username of the user.
	UserName pulumi.StringPtrInput
}

func (UserArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*userArgs)(nil)).Elem()
}

type UserInput interface {
	pulumi.Input

	ToUserOutput() UserOutput
	ToUserOutputWithContext(ctx context.Context) UserOutput
}

func (*User) ElementType() reflect.Type {
	return reflect.TypeOf((**User)(nil)).Elem()
}

func (i *User) ToUserOutput() UserOutput {
	return i.ToUserOutputWithContext(context.Background())
}

func (i *User) ToUserOutputWithContext(ctx context.Context) UserOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserOutput)
}

type UserOutput struct{ *pulumi.OutputState }

func (UserOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**User)(nil)).Elem()
}

func (o UserOutput) ToUserOutput() UserOutput {
	return o
}

func (o UserOutput) ToUserOutputWithContext(ctx context.Context) UserOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*UserInput)(nil)).Elem(), &User{})
	pulumi.RegisterOutputType(UserOutput{})
}
