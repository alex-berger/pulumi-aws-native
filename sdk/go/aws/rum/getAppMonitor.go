// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package rum

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::RUM::AppMonitor
func LookupAppMonitor(ctx *pulumi.Context, args *LookupAppMonitorArgs, opts ...pulumi.InvokeOption) (*LookupAppMonitorResult, error) {
	var rv LookupAppMonitorResult
	err := ctx.Invoke("aws-native:rum:getAppMonitor", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAppMonitorArgs struct {
	// A name for the app monitor
	Name string `pulumi:"name"`
}

type LookupAppMonitorResult struct {
	AppMonitorConfiguration *AppMonitorConfiguration `pulumi:"appMonitorConfiguration"`
	// Data collected by RUM is kept by RUM for 30 days and then deleted. This parameter specifies whether RUM sends a copy of this telemetry data to CWLlong in your account. This enables you to keep the telemetry data for more than 30 days, but it does incur CWLlong charges. If you omit this parameter, the default is false
	CwLogEnabled *bool `pulumi:"cwLogEnabled"`
	// The top-level internet domain name for which your application has administrative authority.
	Domain *string         `pulumi:"domain"`
	Tags   []AppMonitorTag `pulumi:"tags"`
}

func LookupAppMonitorOutput(ctx *pulumi.Context, args LookupAppMonitorOutputArgs, opts ...pulumi.InvokeOption) LookupAppMonitorResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAppMonitorResult, error) {
			args := v.(LookupAppMonitorArgs)
			r, err := LookupAppMonitor(ctx, &args, opts...)
			var s LookupAppMonitorResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAppMonitorResultOutput)
}

type LookupAppMonitorOutputArgs struct {
	// A name for the app monitor
	Name pulumi.StringInput `pulumi:"name"`
}

func (LookupAppMonitorOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAppMonitorArgs)(nil)).Elem()
}

type LookupAppMonitorResultOutput struct{ *pulumi.OutputState }

func (LookupAppMonitorResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAppMonitorResult)(nil)).Elem()
}

func (o LookupAppMonitorResultOutput) ToLookupAppMonitorResultOutput() LookupAppMonitorResultOutput {
	return o
}

func (o LookupAppMonitorResultOutput) ToLookupAppMonitorResultOutputWithContext(ctx context.Context) LookupAppMonitorResultOutput {
	return o
}

func (o LookupAppMonitorResultOutput) AppMonitorConfiguration() AppMonitorConfigurationPtrOutput {
	return o.ApplyT(func(v LookupAppMonitorResult) *AppMonitorConfiguration { return v.AppMonitorConfiguration }).(AppMonitorConfigurationPtrOutput)
}

// Data collected by RUM is kept by RUM for 30 days and then deleted. This parameter specifies whether RUM sends a copy of this telemetry data to CWLlong in your account. This enables you to keep the telemetry data for more than 30 days, but it does incur CWLlong charges. If you omit this parameter, the default is false
func (o LookupAppMonitorResultOutput) CwLogEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupAppMonitorResult) *bool { return v.CwLogEnabled }).(pulumi.BoolPtrOutput)
}

// The top-level internet domain name for which your application has administrative authority.
func (o LookupAppMonitorResultOutput) Domain() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAppMonitorResult) *string { return v.Domain }).(pulumi.StringPtrOutput)
}

func (o LookupAppMonitorResultOutput) Tags() AppMonitorTagArrayOutput {
	return o.ApplyT(func(v LookupAppMonitorResult) []AppMonitorTag { return v.Tags }).(AppMonitorTagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAppMonitorResultOutput{})
}
