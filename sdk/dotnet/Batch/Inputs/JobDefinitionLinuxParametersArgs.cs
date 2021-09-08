// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Batch.Inputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-containerproperties-linuxparameters.html
    /// </summary>
    public sealed class JobDefinitionLinuxParametersArgs : Pulumi.ResourceArgs
    {
        [Input("devices")]
        private InputList<Inputs.JobDefinitionDeviceArgs>? _devices;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-containerproperties-linuxparameters.html#cfn-batch-jobdefinition-containerproperties-linuxparameters-devices
        /// </summary>
        public InputList<Inputs.JobDefinitionDeviceArgs> Devices
        {
            get => _devices ?? (_devices = new InputList<Inputs.JobDefinitionDeviceArgs>());
            set => _devices = value;
        }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-containerproperties-linuxparameters.html#cfn-batch-jobdefinition-containerproperties-linuxparameters-initprocessenabled
        /// </summary>
        [Input("initProcessEnabled")]
        public Input<bool>? InitProcessEnabled { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-containerproperties-linuxparameters.html#cfn-batch-jobdefinition-containerproperties-linuxparameters-maxswap
        /// </summary>
        [Input("maxSwap")]
        public Input<int>? MaxSwap { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-containerproperties-linuxparameters.html#cfn-batch-jobdefinition-containerproperties-linuxparameters-sharedmemorysize
        /// </summary>
        [Input("sharedMemorySize")]
        public Input<int>? SharedMemorySize { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-containerproperties-linuxparameters.html#cfn-batch-jobdefinition-containerproperties-linuxparameters-swappiness
        /// </summary>
        [Input("swappiness")]
        public Input<int>? Swappiness { get; set; }

        [Input("tmpfs")]
        private InputList<Inputs.JobDefinitionTmpfsArgs>? _tmpfs;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-containerproperties-linuxparameters.html#cfn-batch-jobdefinition-containerproperties-linuxparameters-tmpfs
        /// </summary>
        public InputList<Inputs.JobDefinitionTmpfsArgs> Tmpfs
        {
            get => _tmpfs ?? (_tmpfs = new InputList<Inputs.JobDefinitionTmpfsArgs>());
            set => _tmpfs = value;
        }

        public JobDefinitionLinuxParametersArgs()
        {
        }
    }
}
