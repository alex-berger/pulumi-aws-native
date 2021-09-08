// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.GroundStation.Outputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-config-dataflowendpointconfig.html
    /// </summary>
    [OutputType]
    public sealed class ConfigDataflowEndpointConfig
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-config-dataflowendpointconfig.html#cfn-groundstation-config-dataflowendpointconfig-dataflowendpointname
        /// </summary>
        public readonly string? DataflowEndpointName;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-config-dataflowendpointconfig.html#cfn-groundstation-config-dataflowendpointconfig-dataflowendpointregion
        /// </summary>
        public readonly string? DataflowEndpointRegion;

        [OutputConstructor]
        private ConfigDataflowEndpointConfig(
            string? dataflowEndpointName,

            string? dataflowEndpointRegion)
        {
            DataflowEndpointName = dataflowEndpointName;
            DataflowEndpointRegion = dataflowEndpointRegion;
        }
    }
}