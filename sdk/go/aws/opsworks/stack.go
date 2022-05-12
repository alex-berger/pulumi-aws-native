// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package opsworks

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::OpsWorks::Stack
//
// Deprecated: Stack is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type Stack struct {
	pulumi.CustomResourceState

	AgentVersion              pulumi.StringPtrOutput             `pulumi:"agentVersion"`
	Attributes                pulumi.AnyOutput                   `pulumi:"attributes"`
	ChefConfiguration         StackChefConfigurationPtrOutput    `pulumi:"chefConfiguration"`
	CloneAppIds               pulumi.StringArrayOutput           `pulumi:"cloneAppIds"`
	ClonePermissions          pulumi.BoolPtrOutput               `pulumi:"clonePermissions"`
	ConfigurationManager      StackConfigurationManagerPtrOutput `pulumi:"configurationManager"`
	CustomCookbooksSource     StackSourcePtrOutput               `pulumi:"customCookbooksSource"`
	CustomJson                pulumi.AnyOutput                   `pulumi:"customJson"`
	DefaultAvailabilityZone   pulumi.StringPtrOutput             `pulumi:"defaultAvailabilityZone"`
	DefaultInstanceProfileArn pulumi.StringOutput                `pulumi:"defaultInstanceProfileArn"`
	DefaultOs                 pulumi.StringPtrOutput             `pulumi:"defaultOs"`
	DefaultRootDeviceType     pulumi.StringPtrOutput             `pulumi:"defaultRootDeviceType"`
	DefaultSshKeyName         pulumi.StringPtrOutput             `pulumi:"defaultSshKeyName"`
	DefaultSubnetId           pulumi.StringPtrOutput             `pulumi:"defaultSubnetId"`
	EcsClusterArn             pulumi.StringPtrOutput             `pulumi:"ecsClusterArn"`
	ElasticIps                StackElasticIpArrayOutput          `pulumi:"elasticIps"`
	HostnameTheme             pulumi.StringPtrOutput             `pulumi:"hostnameTheme"`
	Name                      pulumi.StringOutput                `pulumi:"name"`
	RdsDbInstances            StackRdsDbInstanceArrayOutput      `pulumi:"rdsDbInstances"`
	ServiceRoleArn            pulumi.StringOutput                `pulumi:"serviceRoleArn"`
	SourceStackId             pulumi.StringPtrOutput             `pulumi:"sourceStackId"`
	Tags                      StackTagArrayOutput                `pulumi:"tags"`
	UseCustomCookbooks        pulumi.BoolPtrOutput               `pulumi:"useCustomCookbooks"`
	UseOpsworksSecurityGroups pulumi.BoolPtrOutput               `pulumi:"useOpsworksSecurityGroups"`
	VpcId                     pulumi.StringPtrOutput             `pulumi:"vpcId"`
}

// NewStack registers a new resource with the given unique name, arguments, and options.
func NewStack(ctx *pulumi.Context,
	name string, args *StackArgs, opts ...pulumi.ResourceOption) (*Stack, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DefaultInstanceProfileArn == nil {
		return nil, errors.New("invalid value for required argument 'DefaultInstanceProfileArn'")
	}
	if args.ServiceRoleArn == nil {
		return nil, errors.New("invalid value for required argument 'ServiceRoleArn'")
	}
	var resource Stack
	err := ctx.RegisterResource("aws-native:opsworks:Stack", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStack gets an existing Stack resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStack(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StackState, opts ...pulumi.ResourceOption) (*Stack, error) {
	var resource Stack
	err := ctx.ReadResource("aws-native:opsworks:Stack", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Stack resources.
type stackState struct {
}

type StackState struct {
}

func (StackState) ElementType() reflect.Type {
	return reflect.TypeOf((*stackState)(nil)).Elem()
}

type stackArgs struct {
	AgentVersion              *string                    `pulumi:"agentVersion"`
	Attributes                interface{}                `pulumi:"attributes"`
	ChefConfiguration         *StackChefConfiguration    `pulumi:"chefConfiguration"`
	CloneAppIds               []string                   `pulumi:"cloneAppIds"`
	ClonePermissions          *bool                      `pulumi:"clonePermissions"`
	ConfigurationManager      *StackConfigurationManager `pulumi:"configurationManager"`
	CustomCookbooksSource     *StackSource               `pulumi:"customCookbooksSource"`
	CustomJson                interface{}                `pulumi:"customJson"`
	DefaultAvailabilityZone   *string                    `pulumi:"defaultAvailabilityZone"`
	DefaultInstanceProfileArn string                     `pulumi:"defaultInstanceProfileArn"`
	DefaultOs                 *string                    `pulumi:"defaultOs"`
	DefaultRootDeviceType     *string                    `pulumi:"defaultRootDeviceType"`
	DefaultSshKeyName         *string                    `pulumi:"defaultSshKeyName"`
	DefaultSubnetId           *string                    `pulumi:"defaultSubnetId"`
	EcsClusterArn             *string                    `pulumi:"ecsClusterArn"`
	ElasticIps                []StackElasticIp           `pulumi:"elasticIps"`
	HostnameTheme             *string                    `pulumi:"hostnameTheme"`
	Name                      *string                    `pulumi:"name"`
	RdsDbInstances            []StackRdsDbInstance       `pulumi:"rdsDbInstances"`
	ServiceRoleArn            string                     `pulumi:"serviceRoleArn"`
	SourceStackId             *string                    `pulumi:"sourceStackId"`
	Tags                      []StackTag                 `pulumi:"tags"`
	UseCustomCookbooks        *bool                      `pulumi:"useCustomCookbooks"`
	UseOpsworksSecurityGroups *bool                      `pulumi:"useOpsworksSecurityGroups"`
	VpcId                     *string                    `pulumi:"vpcId"`
}

// The set of arguments for constructing a Stack resource.
type StackArgs struct {
	AgentVersion              pulumi.StringPtrInput
	Attributes                pulumi.Input
	ChefConfiguration         StackChefConfigurationPtrInput
	CloneAppIds               pulumi.StringArrayInput
	ClonePermissions          pulumi.BoolPtrInput
	ConfigurationManager      StackConfigurationManagerPtrInput
	CustomCookbooksSource     StackSourcePtrInput
	CustomJson                pulumi.Input
	DefaultAvailabilityZone   pulumi.StringPtrInput
	DefaultInstanceProfileArn pulumi.StringInput
	DefaultOs                 pulumi.StringPtrInput
	DefaultRootDeviceType     pulumi.StringPtrInput
	DefaultSshKeyName         pulumi.StringPtrInput
	DefaultSubnetId           pulumi.StringPtrInput
	EcsClusterArn             pulumi.StringPtrInput
	ElasticIps                StackElasticIpArrayInput
	HostnameTheme             pulumi.StringPtrInput
	Name                      pulumi.StringPtrInput
	RdsDbInstances            StackRdsDbInstanceArrayInput
	ServiceRoleArn            pulumi.StringInput
	SourceStackId             pulumi.StringPtrInput
	Tags                      StackTagArrayInput
	UseCustomCookbooks        pulumi.BoolPtrInput
	UseOpsworksSecurityGroups pulumi.BoolPtrInput
	VpcId                     pulumi.StringPtrInput
}

func (StackArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*stackArgs)(nil)).Elem()
}

type StackInput interface {
	pulumi.Input

	ToStackOutput() StackOutput
	ToStackOutputWithContext(ctx context.Context) StackOutput
}

func (*Stack) ElementType() reflect.Type {
	return reflect.TypeOf((**Stack)(nil)).Elem()
}

func (i *Stack) ToStackOutput() StackOutput {
	return i.ToStackOutputWithContext(context.Background())
}

func (i *Stack) ToStackOutputWithContext(ctx context.Context) StackOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StackOutput)
}

type StackOutput struct{ *pulumi.OutputState }

func (StackOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Stack)(nil)).Elem()
}

func (o StackOutput) ToStackOutput() StackOutput {
	return o
}

func (o StackOutput) ToStackOutputWithContext(ctx context.Context) StackOutput {
	return o
}

func (o StackOutput) AgentVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Stack) pulumi.StringPtrOutput { return v.AgentVersion }).(pulumi.StringPtrOutput)
}

func (o StackOutput) Attributes() pulumi.AnyOutput {
	return o.ApplyT(func(v *Stack) pulumi.AnyOutput { return v.Attributes }).(pulumi.AnyOutput)
}

func (o StackOutput) ChefConfiguration() StackChefConfigurationPtrOutput {
	return o.ApplyT(func(v *Stack) StackChefConfigurationPtrOutput { return v.ChefConfiguration }).(StackChefConfigurationPtrOutput)
}

func (o StackOutput) CloneAppIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Stack) pulumi.StringArrayOutput { return v.CloneAppIds }).(pulumi.StringArrayOutput)
}

func (o StackOutput) ClonePermissions() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Stack) pulumi.BoolPtrOutput { return v.ClonePermissions }).(pulumi.BoolPtrOutput)
}

func (o StackOutput) ConfigurationManager() StackConfigurationManagerPtrOutput {
	return o.ApplyT(func(v *Stack) StackConfigurationManagerPtrOutput { return v.ConfigurationManager }).(StackConfigurationManagerPtrOutput)
}

func (o StackOutput) CustomCookbooksSource() StackSourcePtrOutput {
	return o.ApplyT(func(v *Stack) StackSourcePtrOutput { return v.CustomCookbooksSource }).(StackSourcePtrOutput)
}

func (o StackOutput) CustomJson() pulumi.AnyOutput {
	return o.ApplyT(func(v *Stack) pulumi.AnyOutput { return v.CustomJson }).(pulumi.AnyOutput)
}

func (o StackOutput) DefaultAvailabilityZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Stack) pulumi.StringPtrOutput { return v.DefaultAvailabilityZone }).(pulumi.StringPtrOutput)
}

func (o StackOutput) DefaultInstanceProfileArn() pulumi.StringOutput {
	return o.ApplyT(func(v *Stack) pulumi.StringOutput { return v.DefaultInstanceProfileArn }).(pulumi.StringOutput)
}

func (o StackOutput) DefaultOs() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Stack) pulumi.StringPtrOutput { return v.DefaultOs }).(pulumi.StringPtrOutput)
}

func (o StackOutput) DefaultRootDeviceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Stack) pulumi.StringPtrOutput { return v.DefaultRootDeviceType }).(pulumi.StringPtrOutput)
}

func (o StackOutput) DefaultSshKeyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Stack) pulumi.StringPtrOutput { return v.DefaultSshKeyName }).(pulumi.StringPtrOutput)
}

func (o StackOutput) DefaultSubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Stack) pulumi.StringPtrOutput { return v.DefaultSubnetId }).(pulumi.StringPtrOutput)
}

func (o StackOutput) EcsClusterArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Stack) pulumi.StringPtrOutput { return v.EcsClusterArn }).(pulumi.StringPtrOutput)
}

func (o StackOutput) ElasticIps() StackElasticIpArrayOutput {
	return o.ApplyT(func(v *Stack) StackElasticIpArrayOutput { return v.ElasticIps }).(StackElasticIpArrayOutput)
}

func (o StackOutput) HostnameTheme() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Stack) pulumi.StringPtrOutput { return v.HostnameTheme }).(pulumi.StringPtrOutput)
}

func (o StackOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Stack) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o StackOutput) RdsDbInstances() StackRdsDbInstanceArrayOutput {
	return o.ApplyT(func(v *Stack) StackRdsDbInstanceArrayOutput { return v.RdsDbInstances }).(StackRdsDbInstanceArrayOutput)
}

func (o StackOutput) ServiceRoleArn() pulumi.StringOutput {
	return o.ApplyT(func(v *Stack) pulumi.StringOutput { return v.ServiceRoleArn }).(pulumi.StringOutput)
}

func (o StackOutput) SourceStackId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Stack) pulumi.StringPtrOutput { return v.SourceStackId }).(pulumi.StringPtrOutput)
}

func (o StackOutput) Tags() StackTagArrayOutput {
	return o.ApplyT(func(v *Stack) StackTagArrayOutput { return v.Tags }).(StackTagArrayOutput)
}

func (o StackOutput) UseCustomCookbooks() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Stack) pulumi.BoolPtrOutput { return v.UseCustomCookbooks }).(pulumi.BoolPtrOutput)
}

func (o StackOutput) UseOpsworksSecurityGroups() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Stack) pulumi.BoolPtrOutput { return v.UseOpsworksSecurityGroups }).(pulumi.BoolPtrOutput)
}

func (o StackOutput) VpcId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Stack) pulumi.StringPtrOutput { return v.VpcId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*StackInput)(nil)).Elem(), &Stack{})
	pulumi.RegisterOutputType(StackOutput{})
}
