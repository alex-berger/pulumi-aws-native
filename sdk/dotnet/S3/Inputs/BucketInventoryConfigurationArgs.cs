// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.S3.Inputs
{

    public sealed class BucketInventoryConfigurationArgs : Pulumi.ResourceArgs
    {
        [Input("destination", required: true)]
        public Input<Inputs.BucketDestinationArgs> Destination { get; set; } = null!;

        /// <summary>
        /// Specifies whether the inventory is enabled or disabled.
        /// </summary>
        [Input("enabled", required: true)]
        public Input<bool> Enabled { get; set; } = null!;

        /// <summary>
        /// The ID used to identify the inventory configuration.
        /// </summary>
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        /// <summary>
        /// Object versions to include in the inventory list.
        /// </summary>
        [Input("includedObjectVersions", required: true)]
        public Input<Pulumi.AwsNative.S3.BucketInventoryConfigurationIncludedObjectVersions> IncludedObjectVersions { get; set; } = null!;

        [Input("optionalFields")]
        private InputList<Pulumi.AwsNative.S3.BucketInventoryConfigurationOptionalFieldsItem>? _optionalFields;

        /// <summary>
        /// Contains the optional fields that are included in the inventory results.
        /// </summary>
        public InputList<Pulumi.AwsNative.S3.BucketInventoryConfigurationOptionalFieldsItem> OptionalFields
        {
            get => _optionalFields ?? (_optionalFields = new InputList<Pulumi.AwsNative.S3.BucketInventoryConfigurationOptionalFieldsItem>());
            set => _optionalFields = value;
        }

        /// <summary>
        /// The prefix that is prepended to all inventory results.
        /// </summary>
        [Input("prefix")]
        public Input<string>? Prefix { get; set; }

        /// <summary>
        /// Specifies the schedule for generating inventory results.
        /// </summary>
        [Input("scheduleFrequency", required: true)]
        public Input<Pulumi.AwsNative.S3.BucketInventoryConfigurationScheduleFrequency> ScheduleFrequency { get; set; } = null!;

        public BucketInventoryConfigurationArgs()
        {
        }
    }
}
