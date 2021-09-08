// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.EventSchemas.Outputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eventschemas-discoverer-tagsentry.html
    /// </summary>
    [OutputType]
    public sealed class DiscovererTagsEntry
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eventschemas-discoverer-tagsentry.html#cfn-eventschemas-discoverer-tagsentry-key
        /// </summary>
        public readonly string Key;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eventschemas-discoverer-tagsentry.html#cfn-eventschemas-discoverer-tagsentry-value
        /// </summary>
        public readonly string Value;

        [OutputConstructor]
        private DiscovererTagsEntry(
            string key,

            string value)
        {
            Key = key;
            Value = value;
        }
    }
}