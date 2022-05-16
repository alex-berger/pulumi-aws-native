// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.S3.Inputs
{

    /// <summary>
    /// Specifies the default server-side encryption configuration.
    /// </summary>
    public sealed class BucketServerSideEncryptionRuleArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies whether Amazon S3 should use an S3 Bucket Key with server-side encryption using KMS (SSE-KMS) for new objects in the bucket. Existing objects are not affected. Setting the BucketKeyEnabled element to true causes Amazon S3 to use an S3 Bucket Key. By default, S3 Bucket Key is not enabled.
        /// </summary>
        [Input("bucketKeyEnabled")]
        public Input<bool>? BucketKeyEnabled { get; set; }

        [Input("serverSideEncryptionByDefault")]
        public Input<Inputs.BucketServerSideEncryptionByDefaultArgs>? ServerSideEncryptionByDefault { get; set; }

        public BucketServerSideEncryptionRuleArgs()
        {
        }
    }
}