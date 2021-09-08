// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoTAnalytics.Outputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-datastore-datastorepartition.html
    /// </summary>
    [OutputType]
    public sealed class DatastoreDatastorePartition
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-datastore-datastorepartition.html#cfn-iotanalytics-datastore-datastorepartition-partition
        /// </summary>
        public readonly Outputs.DatastorePartition? Partition;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-datastore-datastorepartition.html#cfn-iotanalytics-datastore-datastorepartition-timestamppartition
        /// </summary>
        public readonly Outputs.DatastoreTimestampPartition? TimestampPartition;

        [OutputConstructor]
        private DatastoreDatastorePartition(
            Outputs.DatastorePartition? partition,

            Outputs.DatastoreTimestampPartition? timestampPartition)
        {
            Partition = partition;
            TimestampPartition = timestampPartition;
        }
    }
}
