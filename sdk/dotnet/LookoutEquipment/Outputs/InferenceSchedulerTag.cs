// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.LookoutEquipment.Outputs
{

    /// <summary>
    /// A tag is a key-value pair that can be added to a resource as metadata.
    /// </summary>
    [OutputType]
    public sealed class InferenceSchedulerTag
    {
        /// <summary>
        /// The key for the specified tag.
        /// </summary>
        public readonly string Key;
        /// <summary>
        /// The value for the specified tag.
        /// </summary>
        public readonly string Value;

        [OutputConstructor]
        private InferenceSchedulerTag(
            string key,

            string value)
        {
            Key = key;
            Value = value;
        }
    }
}
