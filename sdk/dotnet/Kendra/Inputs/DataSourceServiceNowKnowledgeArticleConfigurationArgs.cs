// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Kendra.Inputs
{

    public sealed class DataSourceServiceNowKnowledgeArticleConfigurationArgs : Pulumi.ResourceArgs
    {
        [Input("crawlAttachments")]
        public Input<bool>? CrawlAttachments { get; set; }

        [Input("documentDataFieldName", required: true)]
        public Input<string> DocumentDataFieldName { get; set; } = null!;

        [Input("documentTitleFieldName")]
        public Input<string>? DocumentTitleFieldName { get; set; }

        [Input("excludeAttachmentFilePatterns")]
        private InputList<string>? _excludeAttachmentFilePatterns;
        public InputList<string> ExcludeAttachmentFilePatterns
        {
            get => _excludeAttachmentFilePatterns ?? (_excludeAttachmentFilePatterns = new InputList<string>());
            set => _excludeAttachmentFilePatterns = value;
        }

        [Input("fieldMappings")]
        private InputList<Inputs.DataSourceToIndexFieldMappingArgs>? _fieldMappings;
        public InputList<Inputs.DataSourceToIndexFieldMappingArgs> FieldMappings
        {
            get => _fieldMappings ?? (_fieldMappings = new InputList<Inputs.DataSourceToIndexFieldMappingArgs>());
            set => _fieldMappings = value;
        }

        [Input("filterQuery")]
        public Input<string>? FilterQuery { get; set; }

        [Input("includeAttachmentFilePatterns")]
        private InputList<string>? _includeAttachmentFilePatterns;
        public InputList<string> IncludeAttachmentFilePatterns
        {
            get => _includeAttachmentFilePatterns ?? (_includeAttachmentFilePatterns = new InputList<string>());
            set => _includeAttachmentFilePatterns = value;
        }

        public DataSourceServiceNowKnowledgeArticleConfigurationArgs()
        {
        }
    }
}
