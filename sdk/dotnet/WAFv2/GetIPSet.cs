// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.WAFv2
{
    public static class GetIPSet
    {
        /// <summary>
        /// Contains a list of IP addresses. This can be either IPV4 or IPV6. The list will be mutually
        /// </summary>
        public static Task<GetIPSetResult> InvokeAsync(GetIPSetArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetIPSetResult>("aws-native:wafv2:getIPSet", args ?? new GetIPSetArgs(), options.WithDefaults());

        /// <summary>
        /// Contains a list of IP addresses. This can be either IPV4 or IPV6. The list will be mutually
        /// </summary>
        public static Output<GetIPSetResult> Invoke(GetIPSetInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetIPSetResult>("aws-native:wafv2:getIPSet", args ?? new GetIPSetInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetIPSetArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        [Input("scope", required: true)]
        public Pulumi.AwsNative.WAFv2.IPSetScope Scope { get; set; }

        public GetIPSetArgs()
        {
        }
    }

    public sealed class GetIPSetInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("scope", required: true)]
        public Input<Pulumi.AwsNative.WAFv2.IPSetScope> Scope { get; set; } = null!;

        public GetIPSetInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetIPSetResult
    {
        /// <summary>
        /// List of IPAddresses.
        /// </summary>
        public readonly ImmutableArray<string> Addresses;
        public readonly string? Arn;
        public readonly string? Description;
        public readonly Pulumi.AwsNative.WAFv2.IPSetIPAddressVersion? IPAddressVersion;
        public readonly string? Id;
        public readonly ImmutableArray<Outputs.IPSetTag> Tags;

        [OutputConstructor]
        private GetIPSetResult(
            ImmutableArray<string> addresses,

            string? arn,

            string? description,

            Pulumi.AwsNative.WAFv2.IPSetIPAddressVersion? iPAddressVersion,

            string? id,

            ImmutableArray<Outputs.IPSetTag> tags)
        {
            Addresses = addresses;
            Arn = arn;
            Description = description;
            IPAddressVersion = iPAddressVersion;
            Id = id;
            Tags = tags;
        }
    }
}
