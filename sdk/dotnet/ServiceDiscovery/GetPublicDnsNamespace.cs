// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ServiceDiscovery
{
    public static class GetPublicDnsNamespace
    {
        /// <summary>
        /// Resource Type definition for AWS::ServiceDiscovery::PublicDnsNamespace
        /// </summary>
        public static Task<GetPublicDnsNamespaceResult> InvokeAsync(GetPublicDnsNamespaceArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetPublicDnsNamespaceResult>("aws-native:servicediscovery:getPublicDnsNamespace", args ?? new GetPublicDnsNamespaceArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::ServiceDiscovery::PublicDnsNamespace
        /// </summary>
        public static Output<GetPublicDnsNamespaceResult> Invoke(GetPublicDnsNamespaceInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetPublicDnsNamespaceResult>("aws-native:servicediscovery:getPublicDnsNamespace", args ?? new GetPublicDnsNamespaceInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetPublicDnsNamespaceArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetPublicDnsNamespaceArgs()
        {
        }
    }

    public sealed class GetPublicDnsNamespaceInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetPublicDnsNamespaceInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetPublicDnsNamespaceResult
    {
        public readonly string? Arn;
        public readonly string? Description;
        public readonly string? HostedZoneId;
        public readonly string? Id;
        public readonly Outputs.PublicDnsNamespaceProperties? Properties;
        public readonly ImmutableArray<Outputs.PublicDnsNamespaceTag> Tags;

        [OutputConstructor]
        private GetPublicDnsNamespaceResult(
            string? arn,

            string? description,

            string? hostedZoneId,

            string? id,

            Outputs.PublicDnsNamespaceProperties? properties,

            ImmutableArray<Outputs.PublicDnsNamespaceTag> tags)
        {
            Arn = arn;
            Description = description;
            HostedZoneId = hostedZoneId;
            Id = id;
            Properties = properties;
            Tags = tags;
        }
    }
}
