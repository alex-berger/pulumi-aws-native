// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.RoboMaker
{
    public static class GetRobotApplication
    {
        /// <summary>
        /// An example resource schema demonstrating some basic constructs and validation rules.
        /// </summary>
        public static Task<GetRobotApplicationResult> InvokeAsync(GetRobotApplicationArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetRobotApplicationResult>("aws-native:robomaker:getRobotApplication", args ?? new GetRobotApplicationArgs(), options.WithDefaults());

        /// <summary>
        /// An example resource schema demonstrating some basic constructs and validation rules.
        /// </summary>
        public static Output<GetRobotApplicationResult> Invoke(GetRobotApplicationInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetRobotApplicationResult>("aws-native:robomaker:getRobotApplication", args ?? new GetRobotApplicationInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetRobotApplicationArgs : Pulumi.InvokeArgs
    {
        [Input("arn", required: true)]
        public string Arn { get; set; } = null!;

        public GetRobotApplicationArgs()
        {
        }
    }

    public sealed class GetRobotApplicationInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("arn", required: true)]
        public Input<string> Arn { get; set; } = null!;

        public GetRobotApplicationInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetRobotApplicationResult
    {
        public readonly string? Arn;
        /// <summary>
        /// The revision ID of robot application.
        /// </summary>
        public readonly string? CurrentRevisionId;
        /// <summary>
        /// The URI of the Docker image for the robot application.
        /// </summary>
        public readonly string? Environment;
        public readonly Outputs.RobotApplicationRobotSoftwareSuite? RobotSoftwareSuite;
        /// <summary>
        /// The sources of the robot application.
        /// </summary>
        public readonly ImmutableArray<Outputs.RobotApplicationSourceConfig> Sources;
        public readonly Outputs.RobotApplicationTags? Tags;

        [OutputConstructor]
        private GetRobotApplicationResult(
            string? arn,

            string? currentRevisionId,

            string? environment,

            Outputs.RobotApplicationRobotSoftwareSuite? robotSoftwareSuite,

            ImmutableArray<Outputs.RobotApplicationSourceConfig> sources,

            Outputs.RobotApplicationTags? tags)
        {
            Arn = arn;
            CurrentRevisionId = currentRevisionId;
            Environment = environment;
            RobotSoftwareSuite = robotSoftwareSuite;
            Sources = sources;
            Tags = tags;
        }
    }
}
