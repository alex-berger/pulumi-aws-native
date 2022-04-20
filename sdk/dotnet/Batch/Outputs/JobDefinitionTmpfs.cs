// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Batch.Outputs
{

    [OutputType]
    public sealed class JobDefinitionTmpfs
    {
        public readonly string ContainerPath;
        public readonly ImmutableArray<string> MountOptions;
        public readonly int Size;

        [OutputConstructor]
        private JobDefinitionTmpfs(
            string containerPath,

            ImmutableArray<string> mountOptions,

            int size)
        {
            ContainerPath = containerPath;
            MountOptions = mountOptions;
            Size = size;
        }
    }
}
