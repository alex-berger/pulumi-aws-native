// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.LookoutMetrics.Outputs
{

    [OutputType]
    public sealed class AnomalyDetectorS3SourceConfig
    {
        public readonly Outputs.AnomalyDetectorFileFormatDescriptor FileFormatDescriptor;
        public readonly ImmutableArray<string> HistoricalDataPathList;
        public readonly string RoleArn;
        public readonly ImmutableArray<string> TemplatedPathList;

        [OutputConstructor]
        private AnomalyDetectorS3SourceConfig(
            Outputs.AnomalyDetectorFileFormatDescriptor fileFormatDescriptor,

            ImmutableArray<string> historicalDataPathList,

            string roleArn,

            ImmutableArray<string> templatedPathList)
        {
            FileFormatDescriptor = fileFormatDescriptor;
            HistoricalDataPathList = historicalDataPathList;
            RoleArn = roleArn;
            TemplatedPathList = templatedPathList;
        }
    }
}
