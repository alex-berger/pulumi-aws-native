// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.EKS.Outputs
{

    /// <summary>
    /// The Kubernetes network configuration for the cluster.
    /// </summary>
    [OutputType]
    public sealed class ClusterKubernetesNetworkConfig
    {
        /// <summary>
        /// Ipv4 or Ipv6. You can only specify ipv6 for 1.21 and later clusters that use version 1.10.1 or later of the Amazon VPC CNI add-on
        /// </summary>
        public readonly Pulumi.AwsNative.EKS.ClusterKubernetesNetworkConfigIpFamily? IpFamily;
        /// <summary>
        /// The CIDR block to assign Kubernetes service IP addresses from. If you don't specify a block, Kubernetes assigns addresses from either the 10.100.0.0/16 or 172.20.0.0/16 CIDR blocks. We recommend that you specify a block that does not overlap with resources in other networks that are peered or connected to your VPC. 
        /// </summary>
        public readonly string? ServiceIpv4Cidr;
        /// <summary>
        /// The CIDR block to assign Kubernetes service IP addresses from.
        /// </summary>
        public readonly string? ServiceIpv6Cidr;

        [OutputConstructor]
        private ClusterKubernetesNetworkConfig(
            Pulumi.AwsNative.EKS.ClusterKubernetesNetworkConfigIpFamily? ipFamily,

            string? serviceIpv4Cidr,

            string? serviceIpv6Cidr)
        {
            IpFamily = ipFamily;
            ServiceIpv4Cidr = serviceIpv4Cidr;
            ServiceIpv6Cidr = serviceIpv6Cidr;
        }
    }
}
