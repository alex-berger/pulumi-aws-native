// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.MSK.Outputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-cluster-logginginfo.html
    /// </summary>
    [OutputType]
    public sealed class ClusterLoggingInfo
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-cluster-logginginfo.html#cfn-msk-cluster-logginginfo-brokerlogs
        /// </summary>
        public readonly Outputs.ClusterBrokerLogs BrokerLogs;

        [OutputConstructor]
        private ClusterLoggingInfo(Outputs.ClusterBrokerLogs brokerLogs)
        {
            BrokerLogs = brokerLogs;
        }
    }
}
