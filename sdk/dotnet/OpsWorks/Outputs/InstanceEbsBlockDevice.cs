// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.OpsWorks.Outputs
{

    [OutputType]
    public sealed class InstanceEbsBlockDevice
    {
        public readonly bool? DeleteOnTermination;
        public readonly int? Iops;
        public readonly string? SnapshotId;
        public readonly int? VolumeSize;
        public readonly string? VolumeType;

        [OutputConstructor]
        private InstanceEbsBlockDevice(
            bool? deleteOnTermination,

            int? iops,

            string? snapshotId,

            int? volumeSize,

            string? volumeType)
        {
            DeleteOnTermination = deleteOnTermination;
            Iops = iops;
            SnapshotId = snapshotId;
            VolumeSize = volumeSize;
            VolumeType = volumeType;
        }
    }
}
