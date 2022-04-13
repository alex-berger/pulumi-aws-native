// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.HealthLake.Outputs
{

    /// <summary>
    /// The time that a Data Store was created.
    /// </summary>
    [OutputType]
    public sealed class FHIRDatastoreCreatedAt
    {
        /// <summary>
        /// Nanoseconds.
        /// </summary>
        public readonly int Nanos;
        /// <summary>
        /// Seconds since epoch.
        /// </summary>
        public readonly string Seconds;

        [OutputConstructor]
        private FHIRDatastoreCreatedAt(
            int nanos,

            string seconds)
        {
            Nanos = nanos;
            Seconds = seconds;
        }
    }
}
