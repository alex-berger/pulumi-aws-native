// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SageMaker.Inputs
{

    /// <summary>
    /// Configuration for the cluster used to run model monitoring jobs.
    /// </summary>
    public sealed class ModelQualityJobDefinitionClusterConfigArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The number of ML compute instances to use in the model monitoring job. For distributed processing jobs, specify a value greater than 1. The default value is 1.
        /// </summary>
        [Input("instanceCount", required: true)]
        public Input<int> InstanceCount { get; set; } = null!;

        /// <summary>
        /// The ML compute instance type for the processing job.
        /// </summary>
        [Input("instanceType", required: true)]
        public Input<string> InstanceType { get; set; } = null!;

        /// <summary>
        /// The AWS Key Management Service (AWS KMS) key that Amazon SageMaker uses to encrypt data on the storage volume attached to the ML compute instance(s) that run the model monitoring job.
        /// </summary>
        [Input("volumeKmsKeyId")]
        public Input<string>? VolumeKmsKeyId { get; set; }

        /// <summary>
        /// The size of the ML storage volume, in gigabytes, that you want to provision. You must specify sufficient ML storage for your scenario.
        /// </summary>
        [Input("volumeSizeInGB", required: true)]
        public Input<int> VolumeSizeInGB { get; set; } = null!;

        public ModelQualityJobDefinitionClusterConfigArgs()
        {
        }
    }
}
