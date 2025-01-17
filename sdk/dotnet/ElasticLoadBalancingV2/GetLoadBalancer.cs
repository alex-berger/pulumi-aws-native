// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ElasticLoadBalancingV2
{
    public static class GetLoadBalancer
    {
        /// <summary>
        /// Resource Type definition for AWS::ElasticLoadBalancingV2::LoadBalancer
        /// </summary>
        public static Task<GetLoadBalancerResult> InvokeAsync(GetLoadBalancerArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetLoadBalancerResult>("aws-native:elasticloadbalancingv2:getLoadBalancer", args ?? new GetLoadBalancerArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::ElasticLoadBalancingV2::LoadBalancer
        /// </summary>
        public static Output<GetLoadBalancerResult> Invoke(GetLoadBalancerInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetLoadBalancerResult>("aws-native:elasticloadbalancingv2:getLoadBalancer", args ?? new GetLoadBalancerInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetLoadBalancerArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetLoadBalancerArgs()
        {
        }
    }

    public sealed class GetLoadBalancerInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetLoadBalancerInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetLoadBalancerResult
    {
        public readonly string? CanonicalHostedZoneID;
        public readonly string? DNSName;
        public readonly string? Id;
        public readonly string? IpAddressType;
        public readonly ImmutableArray<Outputs.LoadBalancerAttribute> LoadBalancerAttributes;
        public readonly string? LoadBalancerFullName;
        public readonly string? LoadBalancerName;
        public readonly ImmutableArray<string> SecurityGroups;
        public readonly ImmutableArray<Outputs.LoadBalancerSubnetMapping> SubnetMappings;
        public readonly ImmutableArray<string> Subnets;
        public readonly ImmutableArray<Outputs.LoadBalancerTag> Tags;

        [OutputConstructor]
        private GetLoadBalancerResult(
            string? canonicalHostedZoneID,

            string? dNSName,

            string? id,

            string? ipAddressType,

            ImmutableArray<Outputs.LoadBalancerAttribute> loadBalancerAttributes,

            string? loadBalancerFullName,

            string? loadBalancerName,

            ImmutableArray<string> securityGroups,

            ImmutableArray<Outputs.LoadBalancerSubnetMapping> subnetMappings,

            ImmutableArray<string> subnets,

            ImmutableArray<Outputs.LoadBalancerTag> tags)
        {
            CanonicalHostedZoneID = canonicalHostedZoneID;
            DNSName = dNSName;
            Id = id;
            IpAddressType = ipAddressType;
            LoadBalancerAttributes = loadBalancerAttributes;
            LoadBalancerFullName = loadBalancerFullName;
            LoadBalancerName = loadBalancerName;
            SecurityGroups = securityGroups;
            SubnetMappings = subnetMappings;
            Subnets = subnets;
            Tags = tags;
        }
    }
}
