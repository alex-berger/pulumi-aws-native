// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CE.Outputs
{

    /// <summary>
    /// A key-value pair to associate with a resource.
    /// </summary>
    [OutputType]
    public sealed class AnomalySubscriptionResourceTag
    {
        /// <summary>
        /// The key name for the tag.
        /// </summary>
        public readonly string Key;
        /// <summary>
        /// The value for the tag.
        /// </summary>
        public readonly string Value;

        [OutputConstructor]
        private AnomalySubscriptionResourceTag(
            string key,

            string value)
        {
            Key = key;
            Value = value;
        }
    }
}