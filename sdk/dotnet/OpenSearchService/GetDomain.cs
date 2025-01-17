// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.OpenSearchService
{
    public static class GetDomain
    {
        /// <summary>
        /// An example resource schema demonstrating some basic constructs and validation rules.
        /// </summary>
        public static Task<GetDomainResult> InvokeAsync(GetDomainArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetDomainResult>("aws-native:opensearchservice:getDomain", args ?? new GetDomainArgs(), options.WithDefaults());

        /// <summary>
        /// An example resource schema demonstrating some basic constructs and validation rules.
        /// </summary>
        public static Output<GetDomainResult> Invoke(GetDomainInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetDomainResult>("aws-native:opensearchservice:getDomain", args ?? new GetDomainInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDomainArgs : Pulumi.InvokeArgs
    {
        [Input("domainName", required: true)]
        public string DomainName { get; set; } = null!;

        public GetDomainArgs()
        {
        }
    }

    public sealed class GetDomainInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("domainName", required: true)]
        public Input<string> DomainName { get; set; } = null!;

        public GetDomainInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetDomainResult
    {
        public readonly object? AccessPolicies;
        public readonly object? AdvancedOptions;
        public readonly string? Arn;
        public readonly Outputs.DomainClusterConfig? ClusterConfig;
        public readonly Outputs.DomainCognitoOptions? CognitoOptions;
        public readonly string? DomainArn;
        public readonly string? DomainEndpoint;
        public readonly Outputs.DomainEndpointOptions? DomainEndpointOptions;
        public readonly object? DomainEndpoints;
        public readonly Outputs.DomainEBSOptions? EBSOptions;
        public readonly Outputs.DomainEncryptionAtRestOptions? EncryptionAtRestOptions;
        public readonly string? EngineVersion;
        public readonly string? Id;
        public readonly object? LogPublishingOptions;
        public readonly Outputs.DomainNodeToNodeEncryptionOptions? NodeToNodeEncryptionOptions;
        public readonly Outputs.DomainServiceSoftwareOptions? ServiceSoftwareOptions;
        public readonly Outputs.DomainSnapshotOptions? SnapshotOptions;
        /// <summary>
        /// An arbitrary set of tags (key-value pairs) for this Domain.
        /// </summary>
        public readonly ImmutableArray<Outputs.DomainTag> Tags;
        public readonly Outputs.DomainVPCOptions? VPCOptions;

        [OutputConstructor]
        private GetDomainResult(
            object? accessPolicies,

            object? advancedOptions,

            string? arn,

            Outputs.DomainClusterConfig? clusterConfig,

            Outputs.DomainCognitoOptions? cognitoOptions,

            string? domainArn,

            string? domainEndpoint,

            Outputs.DomainEndpointOptions? domainEndpointOptions,

            object? domainEndpoints,

            Outputs.DomainEBSOptions? eBSOptions,

            Outputs.DomainEncryptionAtRestOptions? encryptionAtRestOptions,

            string? engineVersion,

            string? id,

            object? logPublishingOptions,

            Outputs.DomainNodeToNodeEncryptionOptions? nodeToNodeEncryptionOptions,

            Outputs.DomainServiceSoftwareOptions? serviceSoftwareOptions,

            Outputs.DomainSnapshotOptions? snapshotOptions,

            ImmutableArray<Outputs.DomainTag> tags,

            Outputs.DomainVPCOptions? vPCOptions)
        {
            AccessPolicies = accessPolicies;
            AdvancedOptions = advancedOptions;
            Arn = arn;
            ClusterConfig = clusterConfig;
            CognitoOptions = cognitoOptions;
            DomainArn = domainArn;
            DomainEndpoint = domainEndpoint;
            DomainEndpointOptions = domainEndpointOptions;
            DomainEndpoints = domainEndpoints;
            EBSOptions = eBSOptions;
            EncryptionAtRestOptions = encryptionAtRestOptions;
            EngineVersion = engineVersion;
            Id = id;
            LogPublishingOptions = logPublishingOptions;
            NodeToNodeEncryptionOptions = nodeToNodeEncryptionOptions;
            ServiceSoftwareOptions = serviceSoftwareOptions;
            SnapshotOptions = snapshotOptions;
            Tags = tags;
            VPCOptions = vPCOptions;
        }
    }
}
