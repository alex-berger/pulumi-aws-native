// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoT
{
    public static class GetSecurityProfile
    {
        /// <summary>
        /// A security profile defines a set of expected behaviors for devices in your account.
        /// </summary>
        public static Task<GetSecurityProfileResult> InvokeAsync(GetSecurityProfileArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetSecurityProfileResult>("aws-native:iot:getSecurityProfile", args ?? new GetSecurityProfileArgs(), options.WithDefaults());

        /// <summary>
        /// A security profile defines a set of expected behaviors for devices in your account.
        /// </summary>
        public static Output<GetSecurityProfileResult> Invoke(GetSecurityProfileInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetSecurityProfileResult>("aws-native:iot:getSecurityProfile", args ?? new GetSecurityProfileInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSecurityProfileArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// A unique identifier for the security profile.
        /// </summary>
        [Input("securityProfileName", required: true)]
        public string SecurityProfileName { get; set; } = null!;

        public GetSecurityProfileArgs()
        {
        }
    }

    public sealed class GetSecurityProfileInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// A unique identifier for the security profile.
        /// </summary>
        [Input("securityProfileName", required: true)]
        public Input<string> SecurityProfileName { get; set; } = null!;

        public GetSecurityProfileInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetSecurityProfileResult
    {
        /// <summary>
        /// A list of metrics whose data is retained (stored). By default, data is retained for any metric used in the profile's behaviors, but it is also retained for any metric specified here.
        /// </summary>
        public readonly ImmutableArray<Outputs.SecurityProfileMetricToRetain> AdditionalMetricsToRetainV2;
        /// <summary>
        /// Specifies the destinations to which alerts are sent.
        /// </summary>
        public readonly object? AlertTargets;
        /// <summary>
        /// Specifies the behaviors that, when violated by a device (thing), cause an alert.
        /// </summary>
        public readonly ImmutableArray<Outputs.SecurityProfileBehavior> Behaviors;
        /// <summary>
        /// The ARN (Amazon resource name) of the created security profile.
        /// </summary>
        public readonly string? SecurityProfileArn;
        /// <summary>
        /// A description of the security profile.
        /// </summary>
        public readonly string? SecurityProfileDescription;
        /// <summary>
        /// Metadata that can be used to manage the security profile.
        /// </summary>
        public readonly ImmutableArray<Outputs.SecurityProfileTag> Tags;
        /// <summary>
        /// A set of target ARNs that the security profile is attached to.
        /// </summary>
        public readonly ImmutableArray<string> TargetArns;

        [OutputConstructor]
        private GetSecurityProfileResult(
            ImmutableArray<Outputs.SecurityProfileMetricToRetain> additionalMetricsToRetainV2,

            object? alertTargets,

            ImmutableArray<Outputs.SecurityProfileBehavior> behaviors,

            string? securityProfileArn,

            string? securityProfileDescription,

            ImmutableArray<Outputs.SecurityProfileTag> tags,

            ImmutableArray<string> targetArns)
        {
            AdditionalMetricsToRetainV2 = additionalMetricsToRetainV2;
            AlertTargets = alertTargets;
            Behaviors = behaviors;
            SecurityProfileArn = securityProfileArn;
            SecurityProfileDescription = securityProfileDescription;
            Tags = tags;
            TargetArns = targetArns;
        }
    }
}
