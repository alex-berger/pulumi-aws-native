// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.S3Outposts.Inputs
{

    /// <summary>
    /// Specifies lifecycle rules for an Amazon S3Outposts bucket. You must specify at least one of the following: AbortIncompleteMultipartUpload, ExpirationDate, ExpirationInDays.
    /// </summary>
    public sealed class BucketRuleArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies a lifecycle rule that stops incomplete multipart uploads to an Amazon S3Outposts bucket.
        /// </summary>
        [Input("abortIncompleteMultipartUpload")]
        public Input<Inputs.BucketAbortIncompleteMultipartUploadArgs>? AbortIncompleteMultipartUpload { get; set; }

        /// <summary>
        /// Indicates when objects are deleted from Amazon S3Outposts. The date value must be in ISO 8601 format. The time is always midnight UTC.
        /// </summary>
        [Input("expirationDate")]
        public Input<string>? ExpirationDate { get; set; }

        /// <summary>
        /// Indicates the number of days after creation when objects are deleted from Amazon S3Outposts.
        /// </summary>
        [Input("expirationInDays")]
        public Input<int>? ExpirationInDays { get; set; }

        /// <summary>
        /// The container for the filter of the lifecycle rule.
        /// </summary>
        [Input("filter")]
        public Input<Inputs.BucketRuleFilterPropertiesArgs>? Filter { get; set; }

        /// <summary>
        /// Unique identifier for the lifecycle rule. The value can't be longer than 255 characters.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        [Input("status")]
        public Input<Pulumi.AwsNative.S3Outposts.BucketRuleStatus>? Status { get; set; }

        public BucketRuleArgs()
        {
        }
    }
}
