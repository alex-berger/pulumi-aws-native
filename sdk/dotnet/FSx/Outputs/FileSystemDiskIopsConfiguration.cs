// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.FSx.Outputs
{

    [OutputType]
    public sealed class FileSystemDiskIopsConfiguration
    {
        public readonly int? Iops;
        public readonly string? Mode;

        [OutputConstructor]
        private FileSystemDiskIopsConfiguration(
            int? iops,

            string? mode)
        {
            Iops = iops;
            Mode = mode;
        }
    }
}
