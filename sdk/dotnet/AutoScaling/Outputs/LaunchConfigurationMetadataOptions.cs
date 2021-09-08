// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AutoScaling.Outputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-launchconfig-metadataoptions.html
    /// </summary>
    [OutputType]
    public sealed class LaunchConfigurationMetadataOptions
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-launchconfig-metadataoptions.html#cfn-autoscaling-launchconfig-metadataoptions-httpendpoint
        /// </summary>
        public readonly string? HttpEndpoint;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-launchconfig-metadataoptions.html#cfn-autoscaling-launchconfig-metadataoptions-httpputresponsehoplimit
        /// </summary>
        public readonly int? HttpPutResponseHopLimit;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-launchconfig-metadataoptions.html#cfn-autoscaling-launchconfig-metadataoptions-httptokens
        /// </summary>
        public readonly string? HttpTokens;

        [OutputConstructor]
        private LaunchConfigurationMetadataOptions(
            string? httpEndpoint,

            int? httpPutResponseHopLimit,

            string? httpTokens)
        {
            HttpEndpoint = httpEndpoint;
            HttpPutResponseHopLimit = httpPutResponseHopLimit;
            HttpTokens = httpTokens;
        }
    }
}
