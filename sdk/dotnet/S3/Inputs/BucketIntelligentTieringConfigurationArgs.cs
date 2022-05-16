// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.S3.Inputs
{

    public sealed class BucketIntelligentTieringConfigurationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID used to identify the S3 Intelligent-Tiering configuration.
        /// </summary>
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        /// <summary>
        /// An object key name prefix that identifies the subset of objects to which the rule applies.
        /// </summary>
        [Input("prefix")]
        public Input<string>? Prefix { get; set; }

        /// <summary>
        /// Specifies the status of the configuration.
        /// </summary>
        [Input("status", required: true)]
        public Input<Pulumi.AwsNative.S3.BucketIntelligentTieringConfigurationStatus> Status { get; set; } = null!;

        [Input("tagFilters")]
        private InputList<Inputs.BucketTagFilterArgs>? _tagFilters;

        /// <summary>
        /// A container for a key-value pair.
        /// </summary>
        public InputList<Inputs.BucketTagFilterArgs> TagFilters
        {
            get => _tagFilters ?? (_tagFilters = new InputList<Inputs.BucketTagFilterArgs>());
            set => _tagFilters = value;
        }

        [Input("tierings", required: true)]
        private InputList<Inputs.BucketTieringArgs>? _tierings;

        /// <summary>
        /// Specifies a list of S3 Intelligent-Tiering storage class tiers in the configuration. At least one tier must be defined in the list. At most, you can specify two tiers in the list, one for each available AccessTier: ARCHIVE_ACCESS and DEEP_ARCHIVE_ACCESS.
        /// </summary>
        public InputList<Inputs.BucketTieringArgs> Tierings
        {
            get => _tierings ?? (_tierings = new InputList<Inputs.BucketTieringArgs>());
            set => _tierings = value;
        }

        public BucketIntelligentTieringConfigurationArgs()
        {
        }
    }
}