// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.MSK
{
    public static class GetCluster
    {
        /// <summary>
        /// Resource Type definition for AWS::MSK::Cluster
        /// </summary>
        public static Task<GetClusterResult> InvokeAsync(GetClusterArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetClusterResult>("aws-native:msk:getCluster", args ?? new GetClusterArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::MSK::Cluster
        /// </summary>
        public static Output<GetClusterResult> Invoke(GetClusterInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetClusterResult>("aws-native:msk:getCluster", args ?? new GetClusterInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetClusterArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetClusterArgs()
        {
        }
    }

    public sealed class GetClusterInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetClusterInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetClusterResult
    {
        public readonly Outputs.ClusterBrokerNodeGroupInfo? BrokerNodeGroupInfo;
        public readonly Outputs.ClusterClientAuthentication? ClientAuthentication;
        public readonly Outputs.ClusterConfigurationInfo? ConfigurationInfo;
        public readonly Outputs.ClusterEncryptionInfo? EncryptionInfo;
        public readonly string? EnhancedMonitoring;
        public readonly string? Id;
        public readonly string? KafkaVersion;
        public readonly Outputs.ClusterLoggingInfo? LoggingInfo;
        public readonly int? NumberOfBrokerNodes;
        public readonly Outputs.ClusterOpenMonitoring? OpenMonitoring;

        [OutputConstructor]
        private GetClusterResult(
            Outputs.ClusterBrokerNodeGroupInfo? brokerNodeGroupInfo,

            Outputs.ClusterClientAuthentication? clientAuthentication,

            Outputs.ClusterConfigurationInfo? configurationInfo,

            Outputs.ClusterEncryptionInfo? encryptionInfo,

            string? enhancedMonitoring,

            string? id,

            string? kafkaVersion,

            Outputs.ClusterLoggingInfo? loggingInfo,

            int? numberOfBrokerNodes,

            Outputs.ClusterOpenMonitoring? openMonitoring)
        {
            BrokerNodeGroupInfo = brokerNodeGroupInfo;
            ClientAuthentication = clientAuthentication;
            ConfigurationInfo = configurationInfo;
            EncryptionInfo = encryptionInfo;
            EnhancedMonitoring = enhancedMonitoring;
            Id = id;
            KafkaVersion = kafkaVersion;
            LoggingInfo = loggingInfo;
            NumberOfBrokerNodes = numberOfBrokerNodes;
            OpenMonitoring = openMonitoring;
        }
    }
}