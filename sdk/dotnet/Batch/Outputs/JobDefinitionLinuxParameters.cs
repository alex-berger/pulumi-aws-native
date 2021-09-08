// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Batch.Outputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-containerproperties-linuxparameters.html
    /// </summary>
    [OutputType]
    public sealed class JobDefinitionLinuxParameters
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-containerproperties-linuxparameters.html#cfn-batch-jobdefinition-containerproperties-linuxparameters-devices
        /// </summary>
        public readonly ImmutableArray<Outputs.JobDefinitionDevice> Devices;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-containerproperties-linuxparameters.html#cfn-batch-jobdefinition-containerproperties-linuxparameters-initprocessenabled
        /// </summary>
        public readonly bool? InitProcessEnabled;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-containerproperties-linuxparameters.html#cfn-batch-jobdefinition-containerproperties-linuxparameters-maxswap
        /// </summary>
        public readonly int? MaxSwap;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-containerproperties-linuxparameters.html#cfn-batch-jobdefinition-containerproperties-linuxparameters-sharedmemorysize
        /// </summary>
        public readonly int? SharedMemorySize;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-containerproperties-linuxparameters.html#cfn-batch-jobdefinition-containerproperties-linuxparameters-swappiness
        /// </summary>
        public readonly int? Swappiness;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-containerproperties-linuxparameters.html#cfn-batch-jobdefinition-containerproperties-linuxparameters-tmpfs
        /// </summary>
        public readonly ImmutableArray<Outputs.JobDefinitionTmpfs> Tmpfs;

        [OutputConstructor]
        private JobDefinitionLinuxParameters(
            ImmutableArray<Outputs.JobDefinitionDevice> devices,

            bool? initProcessEnabled,

            int? maxSwap,

            int? sharedMemorySize,

            int? swappiness,

            ImmutableArray<Outputs.JobDefinitionTmpfs> tmpfs)
        {
            Devices = devices;
            InitProcessEnabled = initProcessEnabled;
            MaxSwap = maxSwap;
            SharedMemorySize = sharedMemorySize;
            Swappiness = swappiness;
            Tmpfs = tmpfs;
        }
    }
}
