// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.DMS.Outputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-elasticsearchsettings.html
    /// </summary>
    [OutputType]
    public sealed class EndpointElasticsearchSettings
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-elasticsearchsettings.html#cfn-dms-endpoint-elasticsearchsettings-endpointuri
        /// </summary>
        public readonly string? EndpointUri;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-elasticsearchsettings.html#cfn-dms-endpoint-elasticsearchsettings-errorretryduration
        /// </summary>
        public readonly int? ErrorRetryDuration;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-elasticsearchsettings.html#cfn-dms-endpoint-elasticsearchsettings-fullloaderrorpercentage
        /// </summary>
        public readonly int? FullLoadErrorPercentage;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-elasticsearchsettings.html#cfn-dms-endpoint-elasticsearchsettings-serviceaccessrolearn
        /// </summary>
        public readonly string? ServiceAccessRoleArn;

        [OutputConstructor]
        private EndpointElasticsearchSettings(
            string? endpointUri,

            int? errorRetryDuration,

            int? fullLoadErrorPercentage,

            string? serviceAccessRoleArn)
        {
            EndpointUri = endpointUri;
            ErrorRetryDuration = errorRetryDuration;
            FullLoadErrorPercentage = fullLoadErrorPercentage;
            ServiceAccessRoleArn = serviceAccessRoleArn;
        }
    }
}
