// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Glue
{
    public static class GetSchemaVersionMetadata
    {
        /// <summary>
        /// This resource adds Key-Value metadata to a Schema version of Glue Schema Registry.
        /// </summary>
        public static Task<GetSchemaVersionMetadataResult> InvokeAsync(GetSchemaVersionMetadataArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetSchemaVersionMetadataResult>("aws-native:glue:getSchemaVersionMetadata", args ?? new GetSchemaVersionMetadataArgs(), options.WithDefaults());

        /// <summary>
        /// This resource adds Key-Value metadata to a Schema version of Glue Schema Registry.
        /// </summary>
        public static Output<GetSchemaVersionMetadataResult> Invoke(GetSchemaVersionMetadataInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetSchemaVersionMetadataResult>("aws-native:glue:getSchemaVersionMetadata", args ?? new GetSchemaVersionMetadataInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSchemaVersionMetadataArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Metadata key
        /// </summary>
        [Input("key", required: true)]
        public string Key { get; set; } = null!;

        /// <summary>
        /// Represents the version ID associated with the schema version.
        /// </summary>
        [Input("schemaVersionId", required: true)]
        public string SchemaVersionId { get; set; } = null!;

        /// <summary>
        /// Metadata value
        /// </summary>
        [Input("value", required: true)]
        public string Value { get; set; } = null!;

        public GetSchemaVersionMetadataArgs()
        {
        }
    }

    public sealed class GetSchemaVersionMetadataInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Metadata key
        /// </summary>
        [Input("key", required: true)]
        public Input<string> Key { get; set; } = null!;

        /// <summary>
        /// Represents the version ID associated with the schema version.
        /// </summary>
        [Input("schemaVersionId", required: true)]
        public Input<string> SchemaVersionId { get; set; } = null!;

        /// <summary>
        /// Metadata value
        /// </summary>
        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        public GetSchemaVersionMetadataInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetSchemaVersionMetadataResult
    {
        [OutputConstructor]
        private GetSchemaVersionMetadataResult()
        {
        }
    }
}
