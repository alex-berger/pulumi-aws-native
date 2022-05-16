// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ElastiCache.Outputs
{

    [OutputType]
    public sealed class ReplicationGroupLogDeliveryConfigurationRequest
    {
        public readonly Outputs.ReplicationGroupDestinationDetails DestinationDetails;
        public readonly string DestinationType;
        public readonly string LogFormat;
        public readonly string LogType;

        [OutputConstructor]
        private ReplicationGroupLogDeliveryConfigurationRequest(
            Outputs.ReplicationGroupDestinationDetails destinationDetails,

            string destinationType,

            string logFormat,

            string logType)
        {
            DestinationDetails = destinationDetails;
            DestinationType = destinationType;
            LogFormat = logFormat;
            LogType = logType;
        }
    }
}