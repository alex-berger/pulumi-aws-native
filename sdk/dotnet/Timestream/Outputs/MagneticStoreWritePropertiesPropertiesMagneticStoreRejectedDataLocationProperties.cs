// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Timestream.Outputs
{

    /// <summary>
    /// Location to store information about records that were asynchronously rejected during magnetic store writes.
    /// </summary>
    [OutputType]
    public sealed class MagneticStoreWritePropertiesPropertiesMagneticStoreRejectedDataLocationProperties
    {
        /// <summary>
        /// S3 configuration for location to store rejections from magnetic store writes
        /// </summary>
        public readonly Outputs.MagneticStoreWritePropertiesPropertiesMagneticStoreRejectedDataLocationPropertiesS3ConfigurationProperties? S3Configuration;

        [OutputConstructor]
        private MagneticStoreWritePropertiesPropertiesMagneticStoreRejectedDataLocationProperties(Outputs.MagneticStoreWritePropertiesPropertiesMagneticStoreRejectedDataLocationPropertiesS3ConfigurationProperties? s3Configuration)
        {
            S3Configuration = s3Configuration;
        }
    }
}