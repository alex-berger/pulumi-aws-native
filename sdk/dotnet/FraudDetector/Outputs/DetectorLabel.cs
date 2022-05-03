// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.FraudDetector.Outputs
{

    [OutputType]
    public sealed class DetectorLabel
    {
        public readonly string? Arn;
        /// <summary>
        /// The time when the label was created.
        /// </summary>
        public readonly string? CreatedTime;
        /// <summary>
        /// The description.
        /// </summary>
        public readonly string? Description;
        public readonly bool? Inline;
        /// <summary>
        /// The time when the label was last updated.
        /// </summary>
        public readonly string? LastUpdatedTime;
        public readonly string? Name;
        /// <summary>
        /// Tags associated with this label.
        /// </summary>
        public readonly ImmutableArray<Outputs.DetectorTag> Tags;

        [OutputConstructor]
        private DetectorLabel(
            string? arn,

            string? createdTime,

            string? description,

            bool? inline,

            string? lastUpdatedTime,

            string? name,

            ImmutableArray<Outputs.DetectorTag> tags)
        {
            Arn = arn;
            CreatedTime = createdTime;
            Description = description;
            Inline = inline;
            LastUpdatedTime = lastUpdatedTime;
            Name = name;
            Tags = tags;
        }
    }
}
