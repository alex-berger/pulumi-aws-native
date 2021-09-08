// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.S3.Inputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-notificationconfiguration-config-filter-s3key.html
    /// </summary>
    public sealed class BucketS3KeyFilterArgs : Pulumi.ResourceArgs
    {
        [Input("rules", required: true)]
        private InputList<Inputs.BucketFilterRuleArgs>? _rules;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-notificationconfiguration-config-filter-s3key.html#cfn-s3-bucket-notificationconfiguraiton-config-filter-s3key-rules
        /// </summary>
        public InputList<Inputs.BucketFilterRuleArgs> Rules
        {
            get => _rules ?? (_rules = new InputList<Inputs.BucketFilterRuleArgs>());
            set => _rules = value;
        }

        public BucketS3KeyFilterArgs()
        {
        }
    }
}