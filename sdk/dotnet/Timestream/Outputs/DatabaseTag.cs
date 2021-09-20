// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Timestream.Outputs
{

    /// <summary>
    /// You can use the Resource Tags property to apply tags to resources, which can help you identify and categorize those resources.
    /// </summary>
    [OutputType]
    public sealed class DatabaseTag
    {
        public readonly string? Key;
        public readonly string? Value;

        [OutputConstructor]
        private DatabaseTag(
            string? key,

            string? value)
        {
            Key = key;
            Value = value;
        }
    }
}