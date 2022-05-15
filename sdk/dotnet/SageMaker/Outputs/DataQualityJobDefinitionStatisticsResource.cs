// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SageMaker.Outputs
{

    /// <summary>
    /// The baseline statistics resource for a monitoring job.
    /// </summary>
    [OutputType]
    public sealed class DataQualityJobDefinitionStatisticsResource
    {
        /// <summary>
        /// The Amazon S3 URI for the baseline statistics file in Amazon S3 that the current monitoring job should be validated against.
        /// </summary>
        public readonly string? S3Uri;

        [OutputConstructor]
        private DataQualityJobDefinitionStatisticsResource(string? s3Uri)
        {
            S3Uri = s3Uri;
        }
    }
}
