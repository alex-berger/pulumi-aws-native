// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.RoboMaker
{
    public static class GetSimulationApplication
    {
        /// <summary>
        /// AWS::RoboMaker::SimulationApplication resource creates an AWS RoboMaker SimulationApplication. Simulation application can be used in AWS RoboMaker Simulation Jobs.
        /// </summary>
        public static Task<GetSimulationApplicationResult> InvokeAsync(GetSimulationApplicationArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetSimulationApplicationResult>("aws-native:robomaker:getSimulationApplication", args ?? new GetSimulationApplicationArgs(), options.WithDefaults());

        /// <summary>
        /// AWS::RoboMaker::SimulationApplication resource creates an AWS RoboMaker SimulationApplication. Simulation application can be used in AWS RoboMaker Simulation Jobs.
        /// </summary>
        public static Output<GetSimulationApplicationResult> Invoke(GetSimulationApplicationInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetSimulationApplicationResult>("aws-native:robomaker:getSimulationApplication", args ?? new GetSimulationApplicationInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSimulationApplicationArgs : Pulumi.InvokeArgs
    {
        [Input("arn", required: true)]
        public string Arn { get; set; } = null!;

        public GetSimulationApplicationArgs()
        {
        }
    }

    public sealed class GetSimulationApplicationInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("arn", required: true)]
        public Input<string> Arn { get; set; } = null!;

        public GetSimulationApplicationInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetSimulationApplicationResult
    {
        public readonly string? Arn;
        /// <summary>
        /// The current revision id.
        /// </summary>
        public readonly string? CurrentRevisionId;
        /// <summary>
        /// The URI of the Docker image for the robot application.
        /// </summary>
        public readonly string? Environment;
        /// <summary>
        /// The rendering engine for the simulation application.
        /// </summary>
        public readonly Outputs.SimulationApplicationRenderingEngine? RenderingEngine;
        /// <summary>
        /// The robot software suite used by the simulation application.
        /// </summary>
        public readonly Outputs.SimulationApplicationRobotSoftwareSuite? RobotSoftwareSuite;
        /// <summary>
        /// The simulation software suite used by the simulation application.
        /// </summary>
        public readonly Outputs.SimulationApplicationSimulationSoftwareSuite? SimulationSoftwareSuite;
        /// <summary>
        /// The sources of the simulation application.
        /// </summary>
        public readonly ImmutableArray<Outputs.SimulationApplicationSourceConfig> Sources;
        public readonly Outputs.SimulationApplicationTags? Tags;

        [OutputConstructor]
        private GetSimulationApplicationResult(
            string? arn,

            string? currentRevisionId,

            string? environment,

            Outputs.SimulationApplicationRenderingEngine? renderingEngine,

            Outputs.SimulationApplicationRobotSoftwareSuite? robotSoftwareSuite,

            Outputs.SimulationApplicationSimulationSoftwareSuite? simulationSoftwareSuite,

            ImmutableArray<Outputs.SimulationApplicationSourceConfig> sources,

            Outputs.SimulationApplicationTags? tags)
        {
            Arn = arn;
            CurrentRevisionId = currentRevisionId;
            Environment = environment;
            RenderingEngine = renderingEngine;
            RobotSoftwareSuite = robotSoftwareSuite;
            SimulationSoftwareSuite = simulationSoftwareSuite;
            Sources = sources;
            Tags = tags;
        }
    }
}
