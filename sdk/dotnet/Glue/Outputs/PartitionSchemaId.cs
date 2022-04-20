// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Glue.Outputs
{

    [OutputType]
    public sealed class PartitionSchemaId
    {
        public readonly string? RegistryName;
        public readonly string? SchemaArn;
        public readonly string? SchemaName;

        [OutputConstructor]
        private PartitionSchemaId(
            string? registryName,

            string? schemaArn,

            string? schemaName)
        {
            RegistryName = registryName;
            SchemaArn = schemaArn;
            SchemaName = schemaName;
        }
    }
}
