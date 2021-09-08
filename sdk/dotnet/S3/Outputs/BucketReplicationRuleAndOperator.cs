// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.S3.Outputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-replicationruleandoperator.html
    /// </summary>
    [OutputType]
    public sealed class BucketReplicationRuleAndOperator
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-replicationruleandoperator.html#cfn-s3-bucket-replicationruleandoperator-prefix
        /// </summary>
        public readonly string? Prefix;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-replicationruleandoperator.html#cfn-s3-bucket-replicationruleandoperator-tagfilters
        /// </summary>
        public readonly ImmutableArray<Outputs.BucketTagFilter> TagFilters;

        [OutputConstructor]
        private BucketReplicationRuleAndOperator(
            string? prefix,

            ImmutableArray<Outputs.BucketTagFilter> tagFilters)
        {
            Prefix = prefix;
            TagFilters = tagFilters;
        }
    }
}
