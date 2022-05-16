// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.S3.Outputs
{

    /// <summary>
    /// Container for the transition rule that describes when noncurrent objects transition to the STANDARD_IA, ONEZONE_IA, INTELLIGENT_TIERING, GLACIER_IR, GLACIER, or DEEP_ARCHIVE storage class. If your bucket is versioning-enabled (or versioning is suspended), you can set this action to request that Amazon S3 transition noncurrent object versions to the STANDARD_IA, ONEZONE_IA, INTELLIGENT_TIERING, GLACIER_IR, GLACIER, or DEEP_ARCHIVE storage class at a specific period in the object's lifetime.
    /// </summary>
    [OutputType]
    public sealed class BucketNoncurrentVersionTransition
    {
        /// <summary>
        /// Specified the number of newer noncurrent and current versions that must exists before performing the associated action
        /// </summary>
        public readonly int? NewerNoncurrentVersions;
        /// <summary>
        /// The class of storage used to store the object.
        /// </summary>
        public readonly Pulumi.AwsNative.S3.BucketNoncurrentVersionTransitionStorageClass StorageClass;
        /// <summary>
        /// Specifies the number of days an object is noncurrent before Amazon S3 can perform the associated action.
        /// </summary>
        public readonly int TransitionInDays;

        [OutputConstructor]
        private BucketNoncurrentVersionTransition(
            int? newerNoncurrentVersions,

            Pulumi.AwsNative.S3.BucketNoncurrentVersionTransitionStorageClass storageClass,

            int transitionInDays)
        {
            NewerNoncurrentVersions = newerNoncurrentVersions;
            StorageClass = storageClass;
            TransitionInDays = transitionInDays;
        }
    }
}
