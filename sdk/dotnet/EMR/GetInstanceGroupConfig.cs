// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.EMR
{
    public static class GetInstanceGroupConfig
    {
        /// <summary>
        /// Resource Type definition for AWS::EMR::InstanceGroupConfig
        /// </summary>
        public static Task<GetInstanceGroupConfigResult> InvokeAsync(GetInstanceGroupConfigArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetInstanceGroupConfigResult>("aws-native:emr:getInstanceGroupConfig", args ?? new GetInstanceGroupConfigArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::EMR::InstanceGroupConfig
        /// </summary>
        public static Output<GetInstanceGroupConfigResult> Invoke(GetInstanceGroupConfigInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetInstanceGroupConfigResult>("aws-native:emr:getInstanceGroupConfig", args ?? new GetInstanceGroupConfigInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetInstanceGroupConfigArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetInstanceGroupConfigArgs()
        {
        }
    }

    public sealed class GetInstanceGroupConfigInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetInstanceGroupConfigInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetInstanceGroupConfigResult
    {
        public readonly Outputs.InstanceGroupConfigAutoScalingPolicy? AutoScalingPolicy;
        public readonly string? Id;
        public readonly int? InstanceCount;

        [OutputConstructor]
        private GetInstanceGroupConfigResult(
            Outputs.InstanceGroupConfigAutoScalingPolicy? autoScalingPolicy,

            string? id,

            int? instanceCount)
        {
            AutoScalingPolicy = autoScalingPolicy;
            Id = id;
            InstanceCount = instanceCount;
        }
    }
}
