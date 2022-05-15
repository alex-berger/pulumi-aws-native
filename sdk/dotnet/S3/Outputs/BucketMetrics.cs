// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.S3.Outputs
{

    [OutputType]
    public sealed class BucketMetrics
    {
        public readonly Outputs.BucketReplicationTimeValue? EventThreshold;
        public readonly Pulumi.AwsNative.S3.BucketMetricsStatus Status;

        [OutputConstructor]
        private BucketMetrics(
            Outputs.BucketReplicationTimeValue? eventThreshold,

            Pulumi.AwsNative.S3.BucketMetricsStatus status)
        {
            EventThreshold = eventThreshold;
            Status = status;
        }
    }
}
