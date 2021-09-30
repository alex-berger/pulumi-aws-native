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
    public sealed class DataSourceServiceNowKnowledgeArticleConfiguration
    {
        public readonly bool? CrawlAttachments;
        public readonly string DocumentDataFieldName;
        public readonly string? DocumentTitleFieldName;
        public readonly ImmutableArray<string> ExcludeAttachmentFilePatterns;
        public readonly ImmutableArray<Outputs.DataSourceDataSourceToIndexFieldMapping> FieldMappings;
        public readonly string? FilterQuery;
        public readonly ImmutableArray<string> IncludeAttachmentFilePatterns;

        [OutputConstructor]
        private DataSourceServiceNowKnowledgeArticleConfiguration(
            bool? crawlAttachments,

            string documentDataFieldName,

            string? documentTitleFieldName,

            ImmutableArray<string> excludeAttachmentFilePatterns,

            ImmutableArray<Outputs.DataSourceDataSourceToIndexFieldMapping> fieldMappings,

            string? filterQuery,

            ImmutableArray<string> includeAttachmentFilePatterns)
        {
            CrawlAttachments = crawlAttachments;
            DocumentDataFieldName = documentDataFieldName;
            DocumentTitleFieldName = documentTitleFieldName;
            ExcludeAttachmentFilePatterns = excludeAttachmentFilePatterns;
            FieldMappings = fieldMappings;
            FilterQuery = filterQuery;
            IncludeAttachmentFilePatterns = includeAttachmentFilePatterns;
        }
    }
}
