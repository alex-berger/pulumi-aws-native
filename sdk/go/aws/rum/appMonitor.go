// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package rum

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::RUM::AppMonitor
type AppMonitor struct {
	pulumi.CustomResourceState

	AppMonitorConfiguration AppMonitorConfigurationPtrOutput `pulumi:"appMonitorConfiguration"`
	// Data collected by RUM is kept by RUM for 30 days and then deleted. This parameter specifies whether RUM sends a copy of this telemetry data to CWLlong in your account. This enables you to keep the telemetry data for more than 30 days, but it does incur CWLlong charges. If you omit this parameter, the default is false
	CwLogEnabled pulumi.BoolPtrOutput `pulumi:"cwLogEnabled"`
	// The top-level internet domain name for which your application has administrative authority.
	Domain pulumi.StringOutput `pulumi:"domain"`
	// A name for the app monitor
	Name pulumi.StringOutput      `pulumi:"name"`
	Tags AppMonitorTagArrayOutput `pulumi:"tags"`
}

// NewAppMonitor registers a new resource with the given unique name, arguments, and options.
func NewAppMonitor(ctx *pulumi.Context,
	name string, args *AppMonitorArgs, opts ...pulumi.ResourceOption) (*AppMonitor, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Domain == nil {
		return nil, errors.New("invalid value for required argument 'Domain'")
	}
	var resource AppMonitor
	err := ctx.RegisterResource("aws-native:rum:AppMonitor", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAppMonitor gets an existing AppMonitor resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAppMonitor(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AppMonitorState, opts ...pulumi.ResourceOption) (*AppMonitor, error) {
	var resource AppMonitor
	err := ctx.ReadResource("aws-native:rum:AppMonitor", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AppMonitor resources.
type appMonitorState struct {
}

type AppMonitorState struct {
}

func (AppMonitorState) ElementType() reflect.Type {
	return reflect.TypeOf((*appMonitorState)(nil)).Elem()
}

type appMonitorArgs struct {
	AppMonitorConfiguration *AppMonitorConfiguration `pulumi:"appMonitorConfiguration"`
	// Data collected by RUM is kept by RUM for 30 days and then deleted. This parameter specifies whether RUM sends a copy of this telemetry data to CWLlong in your account. This enables you to keep the telemetry data for more than 30 days, but it does incur CWLlong charges. If you omit this parameter, the default is false
	CwLogEnabled *bool `pulumi:"cwLogEnabled"`
	// The top-level internet domain name for which your application has administrative authority.
	Domain string `pulumi:"domain"`
	// A name for the app monitor
	Name *string         `pulumi:"name"`
	Tags []AppMonitorTag `pulumi:"tags"`
}

// The set of arguments for constructing a AppMonitor resource.
type AppMonitorArgs struct {
	AppMonitorConfiguration AppMonitorConfigurationPtrInput
	// Data collected by RUM is kept by RUM for 30 days and then deleted. This parameter specifies whether RUM sends a copy of this telemetry data to CWLlong in your account. This enables you to keep the telemetry data for more than 30 days, but it does incur CWLlong charges. If you omit this parameter, the default is false
	CwLogEnabled pulumi.BoolPtrInput
	// The top-level internet domain name for which your application has administrative authority.
	Domain pulumi.StringInput
	// A name for the app monitor
	Name pulumi.StringPtrInput
	Tags AppMonitorTagArrayInput
}

func (AppMonitorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*appMonitorArgs)(nil)).Elem()
}

type AppMonitorInput interface {
	pulumi.Input

	ToAppMonitorOutput() AppMonitorOutput
	ToAppMonitorOutputWithContext(ctx context.Context) AppMonitorOutput
}

func (*AppMonitor) ElementType() reflect.Type {
	return reflect.TypeOf((**AppMonitor)(nil)).Elem()
}

func (i *AppMonitor) ToAppMonitorOutput() AppMonitorOutput {
	return i.ToAppMonitorOutputWithContext(context.Background())
}

func (i *AppMonitor) ToAppMonitorOutputWithContext(ctx context.Context) AppMonitorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppMonitorOutput)
}

type AppMonitorOutput struct{ *pulumi.OutputState }

func (AppMonitorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AppMonitor)(nil)).Elem()
}

func (o AppMonitorOutput) ToAppMonitorOutput() AppMonitorOutput {
	return o
}

func (o AppMonitorOutput) ToAppMonitorOutputWithContext(ctx context.Context) AppMonitorOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AppMonitorInput)(nil)).Elem(), &AppMonitor{})
	pulumi.RegisterOutputType(AppMonitorOutput{})
}
