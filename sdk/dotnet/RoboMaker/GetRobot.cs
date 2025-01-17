// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.RoboMaker
{
    public static class GetRobot
    {
        /// <summary>
        /// AWS::RoboMaker::Robot resource creates an AWS RoboMaker Robot.
        /// </summary>
        public static Task<GetRobotResult> InvokeAsync(GetRobotArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetRobotResult>("aws-native:robomaker:getRobot", args ?? new GetRobotArgs(), options.WithDefaults());

        /// <summary>
        /// AWS::RoboMaker::Robot resource creates an AWS RoboMaker Robot.
        /// </summary>
        public static Output<GetRobotResult> Invoke(GetRobotInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetRobotResult>("aws-native:robomaker:getRobot", args ?? new GetRobotInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetRobotArgs : Pulumi.InvokeArgs
    {
        [Input("arn", required: true)]
        public string Arn { get; set; } = null!;

        public GetRobotArgs()
        {
        }
    }

    public sealed class GetRobotInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("arn", required: true)]
        public Input<string> Arn { get; set; } = null!;

        public GetRobotInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetRobotResult
    {
        public readonly string? Arn;
        public readonly Outputs.RobotTags? Tags;

        [OutputConstructor]
        private GetRobotResult(
            string? arn,

            Outputs.RobotTags? tags)
        {
            Arn = arn;
            Tags = tags;
        }
    }
}
