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
    /// The properties that determine whether magnetic store writes are enabled.
    /// </summary>
    [OutputType]
    public sealed class MagneticStoreWritePropertiesProperties
    {
        /// <summary>
        /// Boolean flag indicating whether magnetic store writes are enabled.
        /// </summary>
        public readonly bool EnableMagneticStoreWrites;
        /// <summary>
        /// Location to store information about records that were asynchronously rejected during magnetic store writes.
        /// </summary>
        public readonly Outputs.MagneticStoreWritePropertiesPropertiesMagneticStoreRejectedDataLocationProperties? MagneticStoreRejectedDataLocation;

        [OutputConstructor]
        private MagneticStoreWritePropertiesProperties(
            bool enableMagneticStoreWrites,

            Outputs.MagneticStoreWritePropertiesPropertiesMagneticStoreRejectedDataLocationProperties? magneticStoreRejectedDataLocation)
        {
            EnableMagneticStoreWrites = enableMagneticStoreWrites;
            MagneticStoreRejectedDataLocation = magneticStoreRejectedDataLocation;
        }
    }
}
