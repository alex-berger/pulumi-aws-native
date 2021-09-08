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
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-cluster-brokernodegroupinfo.html
    /// </summary>
    [OutputType]
    public sealed class ClusterBrokerNodeGroupInfo
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-cluster-brokernodegroupinfo.html#cfn-msk-cluster-brokernodegroupinfo-brokerazdistribution
        /// </summary>
        public readonly string? BrokerAZDistribution;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-cluster-brokernodegroupinfo.html#cfn-msk-cluster-brokernodegroupinfo-clientsubnets
        /// </summary>
        public readonly ImmutableArray<string> ClientSubnets;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-cluster-brokernodegroupinfo.html#cfn-msk-cluster-brokernodegroupinfo-instancetype
        /// </summary>
        public readonly string InstanceType;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-cluster-brokernodegroupinfo.html#cfn-msk-cluster-brokernodegroupinfo-securitygroups
        /// </summary>
        public readonly ImmutableArray<string> SecurityGroups;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-cluster-brokernodegroupinfo.html#cfn-msk-cluster-brokernodegroupinfo-storageinfo
        /// </summary>
        public readonly Outputs.ClusterStorageInfo? StorageInfo;

        [OutputConstructor]
        private ClusterBrokerNodeGroupInfo(
            string? brokerAZDistribution,

            ImmutableArray<string> clientSubnets,

            string instanceType,

            ImmutableArray<string> securityGroups,

            Outputs.ClusterStorageInfo? storageInfo)
        {
            BrokerAZDistribution = brokerAZDistribution;
            ClientSubnets = clientSubnets;
            InstanceType = instanceType;
            SecurityGroups = securityGroups;
            StorageInfo = storageInfo;
        }
    }
}
