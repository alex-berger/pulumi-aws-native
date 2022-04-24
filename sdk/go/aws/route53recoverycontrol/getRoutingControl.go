// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package route53recoverycontrol

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// AWS Route53 Recovery Control Routing Control resource schema .
func LookupRoutingControl(ctx *pulumi.Context, args *LookupRoutingControlArgs, opts ...pulumi.InvokeOption) (*LookupRoutingControlResult, error) {
	var rv LookupRoutingControlResult
	err := ctx.Invoke("aws-native:route53recoverycontrol:getRoutingControl", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRoutingControlArgs struct {
	// The Amazon Resource Name (ARN) of the routing control.
	RoutingControlArn string `pulumi:"routingControlArn"`
}

type LookupRoutingControlResult struct {
	// The name of the routing control. You can use any non-white space character in the name.
	Name *string `pulumi:"name"`
	// The Amazon Resource Name (ARN) of the routing control.
	RoutingControlArn *string `pulumi:"routingControlArn"`
	// The deployment status of the routing control. Status can be one of the following: PENDING, DEPLOYED, PENDING_DELETION.
	Status *RoutingControlStatus `pulumi:"status"`
}

func LookupRoutingControlOutput(ctx *pulumi.Context, args LookupRoutingControlOutputArgs, opts ...pulumi.InvokeOption) LookupRoutingControlResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRoutingControlResult, error) {
			args := v.(LookupRoutingControlArgs)
			r, err := LookupRoutingControl(ctx, &args, opts...)
			var s LookupRoutingControlResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRoutingControlResultOutput)
}

type LookupRoutingControlOutputArgs struct {
	// The Amazon Resource Name (ARN) of the routing control.
	RoutingControlArn pulumi.StringInput `pulumi:"routingControlArn"`
}

func (LookupRoutingControlOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRoutingControlArgs)(nil)).Elem()
}

type LookupRoutingControlResultOutput struct{ *pulumi.OutputState }

func (LookupRoutingControlResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRoutingControlResult)(nil)).Elem()
}

func (o LookupRoutingControlResultOutput) ToLookupRoutingControlResultOutput() LookupRoutingControlResultOutput {
	return o
}

func (o LookupRoutingControlResultOutput) ToLookupRoutingControlResultOutputWithContext(ctx context.Context) LookupRoutingControlResultOutput {
	return o
}

// The name of the routing control. You can use any non-white space character in the name.
func (o LookupRoutingControlResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRoutingControlResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// The Amazon Resource Name (ARN) of the routing control.
func (o LookupRoutingControlResultOutput) RoutingControlArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRoutingControlResult) *string { return v.RoutingControlArn }).(pulumi.StringPtrOutput)
}

// The deployment status of the routing control. Status can be one of the following: PENDING, DEPLOYED, PENDING_DELETION.
func (o LookupRoutingControlResultOutput) Status() RoutingControlStatusPtrOutput {
	return o.ApplyT(func(v LookupRoutingControlResult) *RoutingControlStatus { return v.Status }).(RoutingControlStatusPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRoutingControlResultOutput{})
}
