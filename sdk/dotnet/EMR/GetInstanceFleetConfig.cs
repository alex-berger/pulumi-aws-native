// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.EMR
{
    public static class GetInstanceFleetConfig
    {
        /// <summary>
        /// Resource Type definition for AWS::EMR::InstanceFleetConfig
        /// </summary>
        public static Task<GetInstanceFleetConfigResult> InvokeAsync(GetInstanceFleetConfigArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetInstanceFleetConfigResult>("aws-native:emr:getInstanceFleetConfig", args ?? new GetInstanceFleetConfigArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::EMR::InstanceFleetConfig
        /// </summary>
        public static Output<GetInstanceFleetConfigResult> Invoke(GetInstanceFleetConfigInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetInstanceFleetConfigResult>("aws-native:emr:getInstanceFleetConfig", args ?? new GetInstanceFleetConfigInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetInstanceFleetConfigArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetInstanceFleetConfigArgs()
        {
        }
    }

    public sealed class GetInstanceFleetConfigInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetInstanceFleetConfigInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetInstanceFleetConfigResult
    {
        public readonly string? Id;
        public readonly int? TargetOnDemandCapacity;
        public readonly int? TargetSpotCapacity;

        [OutputConstructor]
        private GetInstanceFleetConfigResult(
            string? id,

            int? targetOnDemandCapacity,

            int? targetSpotCapacity)
        {
            Id = id;
            TargetOnDemandCapacity = targetOnDemandCapacity;
            TargetSpotCapacity = targetSpotCapacity;
        }
    }
}
