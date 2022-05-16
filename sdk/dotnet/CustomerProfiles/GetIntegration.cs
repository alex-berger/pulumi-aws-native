// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CustomerProfiles
{
    public static class GetIntegration
    {
        /// <summary>
        /// The resource schema for creating an Amazon Connect Customer Profiles Integration.
        /// </summary>
        public static Task<GetIntegrationResult> InvokeAsync(GetIntegrationArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetIntegrationResult>("aws-native:customerprofiles:getIntegration", args ?? new GetIntegrationArgs(), options.WithDefaults());

        /// <summary>
        /// The resource schema for creating an Amazon Connect Customer Profiles Integration.
        /// </summary>
        public static Output<GetIntegrationResult> Invoke(GetIntegrationInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetIntegrationResult>("aws-native:customerprofiles:getIntegration", args ?? new GetIntegrationInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetIntegrationArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The unique name of the domain.
        /// </summary>
        [Input("domainName", required: true)]
        public string DomainName { get; set; } = null!;

        /// <summary>
        /// The URI of the S3 bucket or any other type of data source.
        /// </summary>
        [Input("uri", required: true)]
        public string Uri { get; set; } = null!;

        public GetIntegrationArgs()
        {
        }
    }

    public sealed class GetIntegrationInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The unique name of the domain.
        /// </summary>
        [Input("domainName", required: true)]
        public Input<string> DomainName { get; set; } = null!;

        /// <summary>
        /// The URI of the S3 bucket or any other type of data source.
        /// </summary>
        [Input("uri", required: true)]
        public Input<string> Uri { get; set; } = null!;

        public GetIntegrationInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetIntegrationResult
    {
        /// <summary>
        /// The time of this integration got created
        /// </summary>
        public readonly string? CreatedAt;
        /// <summary>
        /// The time of this integration got last updated at
        /// </summary>
        public readonly string? LastUpdatedAt;
        /// <summary>
        /// The name of the ObjectType defined for the 3rd party data in Profile Service
        /// </summary>
        public readonly string? ObjectTypeName;
        /// <summary>
        /// The mapping between 3rd party event types and ObjectType names
        /// </summary>
        public readonly ImmutableArray<Outputs.IntegrationObjectTypeMapping> ObjectTypeNames;
        /// <summary>
        /// The tags (keys and values) associated with the integration
        /// </summary>
        public readonly ImmutableArray<Outputs.IntegrationTag> Tags;

        [OutputConstructor]
        private GetIntegrationResult(
            string? createdAt,

            string? lastUpdatedAt,

            string? objectTypeName,

            ImmutableArray<Outputs.IntegrationObjectTypeMapping> objectTypeNames,

            ImmutableArray<Outputs.IntegrationTag> tags)
        {
            CreatedAt = createdAt;
            LastUpdatedAt = lastUpdatedAt;
            ObjectTypeName = objectTypeName;
            ObjectTypeNames = objectTypeNames;
            Tags = tags;
        }
    }
}