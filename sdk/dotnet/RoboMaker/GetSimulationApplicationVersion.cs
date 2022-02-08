// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.RoboMaker
{
    public static class GetSimulationApplicationVersion
    {
        /// <summary>
        /// An example resource schema demonstrating some basic constructs and validation rules.
        /// </summary>
        public static Task<GetSimulationApplicationVersionResult> InvokeAsync(GetSimulationApplicationVersionArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetSimulationApplicationVersionResult>("aws-native:robomaker:getSimulationApplicationVersion", args ?? new GetSimulationApplicationVersionArgs(), options.WithDefaults());

        /// <summary>
        /// An example resource schema demonstrating some basic constructs and validation rules.
        /// </summary>
        public static Output<GetSimulationApplicationVersionResult> Invoke(GetSimulationApplicationVersionInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetSimulationApplicationVersionResult>("aws-native:robomaker:getSimulationApplicationVersion", args ?? new GetSimulationApplicationVersionInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSimulationApplicationVersionArgs : Pulumi.InvokeArgs
    {
        [Input("arn", required: true)]
        public string Arn { get; set; } = null!;

        public GetSimulationApplicationVersionArgs()
        {
        }
    }

    public sealed class GetSimulationApplicationVersionInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("arn", required: true)]
        public Input<string> Arn { get; set; } = null!;

        public GetSimulationApplicationVersionInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetSimulationApplicationVersionResult
    {
        public readonly string? ApplicationVersion;
        public readonly string? Arn;

        [OutputConstructor]
        private GetSimulationApplicationVersionResult(
            string? applicationVersion,

            string? arn)
        {
            ApplicationVersion = applicationVersion;
            Arn = arn;
        }
    }
}