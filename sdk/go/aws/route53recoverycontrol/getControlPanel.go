// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package route53recoverycontrol

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// AWS Route53 Recovery Control Control Panel resource schema .
func LookupControlPanel(ctx *pulumi.Context, args *LookupControlPanelArgs, opts ...pulumi.InvokeOption) (*LookupControlPanelResult, error) {
	var rv LookupControlPanelResult
	err := ctx.Invoke("aws-native:route53recoverycontrol:getControlPanel", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupControlPanelArgs struct {
	// The Amazon Resource Name (ARN) of the cluster.
	ControlPanelArn string `pulumi:"controlPanelArn"`
}

type LookupControlPanelResult struct {
	// The Amazon Resource Name (ARN) of the cluster.
	ControlPanelArn *string `pulumi:"controlPanelArn"`
	// A flag that Amazon Route 53 Application Recovery Controller sets to true to designate the default control panel for a cluster. When you create a cluster, Amazon Route 53 Application Recovery Controller creates a control panel, and sets this flag for that control panel. If you create a control panel yourself, this flag is set to false.
	DefaultControlPanel *bool `pulumi:"defaultControlPanel"`
	// The name of the control panel. You can use any non-white space character in the name.
	Name *string `pulumi:"name"`
	// Count of associated routing controls
	RoutingControlCount *int `pulumi:"routingControlCount"`
	// The deployment status of control panel. Status can be one of the following: PENDING, DEPLOYED, PENDING_DELETION.
	Status *ControlPanelStatus `pulumi:"status"`
}

func LookupControlPanelOutput(ctx *pulumi.Context, args LookupControlPanelOutputArgs, opts ...pulumi.InvokeOption) LookupControlPanelResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupControlPanelResult, error) {
			args := v.(LookupControlPanelArgs)
			r, err := LookupControlPanel(ctx, &args, opts...)
			var s LookupControlPanelResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupControlPanelResultOutput)
}

type LookupControlPanelOutputArgs struct {
	// The Amazon Resource Name (ARN) of the cluster.
	ControlPanelArn pulumi.StringInput `pulumi:"controlPanelArn"`
}

func (LookupControlPanelOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupControlPanelArgs)(nil)).Elem()
}

type LookupControlPanelResultOutput struct{ *pulumi.OutputState }

func (LookupControlPanelResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupControlPanelResult)(nil)).Elem()
}

func (o LookupControlPanelResultOutput) ToLookupControlPanelResultOutput() LookupControlPanelResultOutput {
	return o
}

func (o LookupControlPanelResultOutput) ToLookupControlPanelResultOutputWithContext(ctx context.Context) LookupControlPanelResultOutput {
	return o
}

// The Amazon Resource Name (ARN) of the cluster.
func (o LookupControlPanelResultOutput) ControlPanelArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupControlPanelResult) *string { return v.ControlPanelArn }).(pulumi.StringPtrOutput)
}

// A flag that Amazon Route 53 Application Recovery Controller sets to true to designate the default control panel for a cluster. When you create a cluster, Amazon Route 53 Application Recovery Controller creates a control panel, and sets this flag for that control panel. If you create a control panel yourself, this flag is set to false.
func (o LookupControlPanelResultOutput) DefaultControlPanel() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupControlPanelResult) *bool { return v.DefaultControlPanel }).(pulumi.BoolPtrOutput)
}

// The name of the control panel. You can use any non-white space character in the name.
func (o LookupControlPanelResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupControlPanelResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// Count of associated routing controls
func (o LookupControlPanelResultOutput) RoutingControlCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupControlPanelResult) *int { return v.RoutingControlCount }).(pulumi.IntPtrOutput)
}

// The deployment status of control panel. Status can be one of the following: PENDING, DEPLOYED, PENDING_DELETION.
func (o LookupControlPanelResultOutput) Status() ControlPanelStatusPtrOutput {
	return o.ApplyT(func(v LookupControlPanelResult) *ControlPanelStatus { return v.Status }).(ControlPanelStatusPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupControlPanelResultOutput{})
}
