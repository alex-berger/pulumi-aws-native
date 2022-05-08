// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package robomaker

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// AWS::RoboMaker::SimulationApplication resource creates an AWS RoboMaker SimulationApplication. Simulation application can be used in AWS RoboMaker Simulation Jobs.
func LookupSimulationApplication(ctx *pulumi.Context, args *LookupSimulationApplicationArgs, opts ...pulumi.InvokeOption) (*LookupSimulationApplicationResult, error) {
	var rv LookupSimulationApplicationResult
	err := ctx.Invoke("aws-native:robomaker:getSimulationApplication", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSimulationApplicationArgs struct {
	Arn string `pulumi:"arn"`
}

type LookupSimulationApplicationResult struct {
	Arn *string `pulumi:"arn"`
	// The current revision id.
	CurrentRevisionId *string `pulumi:"currentRevisionId"`
	// The URI of the Docker image for the robot application.
	Environment *string `pulumi:"environment"`
	// The rendering engine for the simulation application.
	RenderingEngine *SimulationApplicationRenderingEngine `pulumi:"renderingEngine"`
	// The robot software suite used by the simulation application.
	RobotSoftwareSuite *SimulationApplicationRobotSoftwareSuite `pulumi:"robotSoftwareSuite"`
	// The simulation software suite used by the simulation application.
	SimulationSoftwareSuite *SimulationApplicationSimulationSoftwareSuite `pulumi:"simulationSoftwareSuite"`
	// The sources of the simulation application.
	Sources []SimulationApplicationSourceConfig `pulumi:"sources"`
	Tags    *SimulationApplicationTags          `pulumi:"tags"`
}

func LookupSimulationApplicationOutput(ctx *pulumi.Context, args LookupSimulationApplicationOutputArgs, opts ...pulumi.InvokeOption) LookupSimulationApplicationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSimulationApplicationResult, error) {
			args := v.(LookupSimulationApplicationArgs)
			r, err := LookupSimulationApplication(ctx, &args, opts...)
			var s LookupSimulationApplicationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSimulationApplicationResultOutput)
}

type LookupSimulationApplicationOutputArgs struct {
	Arn pulumi.StringInput `pulumi:"arn"`
}

func (LookupSimulationApplicationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSimulationApplicationArgs)(nil)).Elem()
}

type LookupSimulationApplicationResultOutput struct{ *pulumi.OutputState }

func (LookupSimulationApplicationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSimulationApplicationResult)(nil)).Elem()
}

func (o LookupSimulationApplicationResultOutput) ToLookupSimulationApplicationResultOutput() LookupSimulationApplicationResultOutput {
	return o
}

func (o LookupSimulationApplicationResultOutput) ToLookupSimulationApplicationResultOutputWithContext(ctx context.Context) LookupSimulationApplicationResultOutput {
	return o
}

func (o LookupSimulationApplicationResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSimulationApplicationResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

// The current revision id.
func (o LookupSimulationApplicationResultOutput) CurrentRevisionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSimulationApplicationResult) *string { return v.CurrentRevisionId }).(pulumi.StringPtrOutput)
}

// The URI of the Docker image for the robot application.
func (o LookupSimulationApplicationResultOutput) Environment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSimulationApplicationResult) *string { return v.Environment }).(pulumi.StringPtrOutput)
}

// The rendering engine for the simulation application.
func (o LookupSimulationApplicationResultOutput) RenderingEngine() SimulationApplicationRenderingEnginePtrOutput {
	return o.ApplyT(func(v LookupSimulationApplicationResult) *SimulationApplicationRenderingEngine {
		return v.RenderingEngine
	}).(SimulationApplicationRenderingEnginePtrOutput)
}

// The robot software suite used by the simulation application.
func (o LookupSimulationApplicationResultOutput) RobotSoftwareSuite() SimulationApplicationRobotSoftwareSuitePtrOutput {
	return o.ApplyT(func(v LookupSimulationApplicationResult) *SimulationApplicationRobotSoftwareSuite {
		return v.RobotSoftwareSuite
	}).(SimulationApplicationRobotSoftwareSuitePtrOutput)
}

// The simulation software suite used by the simulation application.
func (o LookupSimulationApplicationResultOutput) SimulationSoftwareSuite() SimulationApplicationSimulationSoftwareSuitePtrOutput {
	return o.ApplyT(func(v LookupSimulationApplicationResult) *SimulationApplicationSimulationSoftwareSuite {
		return v.SimulationSoftwareSuite
	}).(SimulationApplicationSimulationSoftwareSuitePtrOutput)
}

// The sources of the simulation application.
func (o LookupSimulationApplicationResultOutput) Sources() SimulationApplicationSourceConfigArrayOutput {
	return o.ApplyT(func(v LookupSimulationApplicationResult) []SimulationApplicationSourceConfig { return v.Sources }).(SimulationApplicationSourceConfigArrayOutput)
}

func (o LookupSimulationApplicationResultOutput) Tags() SimulationApplicationTagsPtrOutput {
	return o.ApplyT(func(v LookupSimulationApplicationResult) *SimulationApplicationTags { return v.Tags }).(SimulationApplicationTagsPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSimulationApplicationResultOutput{})
}
