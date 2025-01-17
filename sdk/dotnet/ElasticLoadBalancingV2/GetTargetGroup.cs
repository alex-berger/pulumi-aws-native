// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ElasticLoadBalancingV2
{
    public static class GetTargetGroup
    {
        /// <summary>
        /// Resource Type definition for AWS::ElasticLoadBalancingV2::TargetGroup
        /// </summary>
        public static Task<GetTargetGroupResult> InvokeAsync(GetTargetGroupArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetTargetGroupResult>("aws-native:elasticloadbalancingv2:getTargetGroup", args ?? new GetTargetGroupArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::ElasticLoadBalancingV2::TargetGroup
        /// </summary>
        public static Output<GetTargetGroupResult> Invoke(GetTargetGroupInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetTargetGroupResult>("aws-native:elasticloadbalancingv2:getTargetGroup", args ?? new GetTargetGroupInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetTargetGroupArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetTargetGroupArgs()
        {
        }
    }

    public sealed class GetTargetGroupInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetTargetGroupInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetTargetGroupResult
    {
        public readonly bool? HealthCheckEnabled;
        public readonly int? HealthCheckIntervalSeconds;
        public readonly string? HealthCheckPath;
        public readonly string? HealthCheckPort;
        public readonly string? HealthCheckProtocol;
        public readonly int? HealthCheckTimeoutSeconds;
        public readonly int? HealthyThresholdCount;
        public readonly string? Id;
        public readonly ImmutableArray<string> LoadBalancerArns;
        public readonly Outputs.TargetGroupMatcher? Matcher;
        public readonly ImmutableArray<Outputs.TargetGroupTag> Tags;
        public readonly ImmutableArray<Outputs.TargetGroupAttribute> TargetGroupAttributes;
        public readonly string? TargetGroupFullName;
        public readonly string? TargetGroupName;
        public readonly ImmutableArray<Outputs.TargetGroupTargetDescription> Targets;
        public readonly int? UnhealthyThresholdCount;

        [OutputConstructor]
        private GetTargetGroupResult(
            bool? healthCheckEnabled,

            int? healthCheckIntervalSeconds,

            string? healthCheckPath,

            string? healthCheckPort,

            string? healthCheckProtocol,

            int? healthCheckTimeoutSeconds,

            int? healthyThresholdCount,

            string? id,

            ImmutableArray<string> loadBalancerArns,

            Outputs.TargetGroupMatcher? matcher,

            ImmutableArray<Outputs.TargetGroupTag> tags,

            ImmutableArray<Outputs.TargetGroupAttribute> targetGroupAttributes,

            string? targetGroupFullName,

            string? targetGroupName,

            ImmutableArray<Outputs.TargetGroupTargetDescription> targets,

            int? unhealthyThresholdCount)
        {
            HealthCheckEnabled = healthCheckEnabled;
            HealthCheckIntervalSeconds = healthCheckIntervalSeconds;
            HealthCheckPath = healthCheckPath;
            HealthCheckPort = healthCheckPort;
            HealthCheckProtocol = healthCheckProtocol;
            HealthCheckTimeoutSeconds = healthCheckTimeoutSeconds;
            HealthyThresholdCount = healthyThresholdCount;
            Id = id;
            LoadBalancerArns = loadBalancerArns;
            Matcher = matcher;
            Tags = tags;
            TargetGroupAttributes = targetGroupAttributes;
            TargetGroupFullName = targetGroupFullName;
            TargetGroupName = targetGroupName;
            Targets = targets;
            UnhealthyThresholdCount = unhealthyThresholdCount;
        }
    }
}
