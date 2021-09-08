// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.KinesisAnalyticsV2.Outputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-applicationreferencedatasource-recordformat.html
    /// </summary>
    [OutputType]
    public sealed class ApplicationReferenceDataSourceRecordFormat
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-applicationreferencedatasource-recordformat.html#cfn-kinesisanalyticsv2-applicationreferencedatasource-recordformat-mappingparameters
        /// </summary>
        public readonly Outputs.ApplicationReferenceDataSourceMappingParameters? MappingParameters;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-applicationreferencedatasource-recordformat.html#cfn-kinesisanalyticsv2-applicationreferencedatasource-recordformat-recordformattype
        /// </summary>
        public readonly string RecordFormatType;

        [OutputConstructor]
        private ApplicationReferenceDataSourceRecordFormat(
            Outputs.ApplicationReferenceDataSourceMappingParameters? mappingParameters,

            string recordFormatType)
        {
            MappingParameters = mappingParameters;
            RecordFormatType = recordFormatType;
        }
    }
}
