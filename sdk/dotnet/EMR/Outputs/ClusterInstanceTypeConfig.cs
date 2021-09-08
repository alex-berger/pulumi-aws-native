// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.EMR.Outputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-instancetypeconfig.html
    /// </summary>
    [OutputType]
    public sealed class ClusterInstanceTypeConfig
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-instancetypeconfig.html#cfn-elasticmapreduce-cluster-instancetypeconfig-bidprice
        /// </summary>
        public readonly string? BidPrice;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-instancetypeconfig.html#cfn-elasticmapreduce-cluster-instancetypeconfig-bidpriceaspercentageofondemandprice
        /// </summary>
        public readonly double? BidPriceAsPercentageOfOnDemandPrice;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-instancetypeconfig.html#cfn-elasticmapreduce-cluster-instancetypeconfig-configurations
        /// </summary>
        public readonly ImmutableArray<Outputs.ClusterConfiguration> Configurations;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-instancetypeconfig.html#cfn-elasticmapreduce-cluster-instancetypeconfig-ebsconfiguration
        /// </summary>
        public readonly Outputs.ClusterEbsConfiguration? EbsConfiguration;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-instancetypeconfig.html#cfn-elasticmapreduce-cluster-instancetypeconfig-instancetype
        /// </summary>
        public readonly string InstanceType;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-instancetypeconfig.html#cfn-elasticmapreduce-cluster-instancetypeconfig-weightedcapacity
        /// </summary>
        public readonly int? WeightedCapacity;

        [OutputConstructor]
        private ClusterInstanceTypeConfig(
            string? bidPrice,

            double? bidPriceAsPercentageOfOnDemandPrice,

            ImmutableArray<Outputs.ClusterConfiguration> configurations,

            Outputs.ClusterEbsConfiguration? ebsConfiguration,

            string instanceType,

            int? weightedCapacity)
        {
            BidPrice = bidPrice;
            BidPriceAsPercentageOfOnDemandPrice = bidPriceAsPercentageOfOnDemandPrice;
            Configurations = configurations;
            EbsConfiguration = ebsConfiguration;
            InstanceType = instanceType;
            WeightedCapacity = weightedCapacity;
        }
    }
}
