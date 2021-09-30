// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Kendra.Outputs
{

    [OutputType]
    public sealed class DataSourceWorkDocsConfiguration
    {
        public readonly bool? CrawlComments;
        public readonly ImmutableArray<string> ExclusionPatterns;
        public readonly ImmutableArray<Outputs.DataSourceDataSourceToIndexFieldMapping> FieldMappings;
        public readonly ImmutableArray<string> InclusionPatterns;
        public readonly string OrganizationId;
        public readonly bool? UseChangeLog;

        [OutputConstructor]
        private DataSourceWorkDocsConfiguration(
            bool? crawlComments,

            ImmutableArray<string> exclusionPatterns,

            ImmutableArray<Outputs.DataSourceDataSourceToIndexFieldMapping> fieldMappings,

            ImmutableArray<string> inclusionPatterns,

            string organizationId,

            bool? useChangeLog)
        {
            CrawlComments = crawlComments;
            ExclusionPatterns = exclusionPatterns;
            FieldMappings = fieldMappings;
            InclusionPatterns = inclusionPatterns;
            OrganizationId = organizationId;
            UseChangeLog = useChangeLog;
        }
    }
}
