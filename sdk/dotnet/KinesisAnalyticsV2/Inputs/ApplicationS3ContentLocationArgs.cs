// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.KinesisAnalyticsV2.Inputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-s3contentlocation.html
    /// </summary>
    public sealed class ApplicationS3ContentLocationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-s3contentlocation.html#cfn-kinesisanalyticsv2-application-s3contentlocation-bucketarn
        /// </summary>
        [Input("bucketARN")]
        public Input<string>? BucketARN { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-s3contentlocation.html#cfn-kinesisanalyticsv2-application-s3contentlocation-filekey
        /// </summary>
        [Input("fileKey")]
        public Input<string>? FileKey { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-s3contentlocation.html#cfn-kinesisanalyticsv2-application-s3contentlocation-objectversion
        /// </summary>
        [Input("objectVersion")]
        public Input<string>? ObjectVersion { get; set; }

        public ApplicationS3ContentLocationArgs()
        {
        }
    }
}
