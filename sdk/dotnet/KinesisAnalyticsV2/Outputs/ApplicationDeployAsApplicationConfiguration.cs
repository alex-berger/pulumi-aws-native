// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.KinesisAnalyticsV2.Outputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-deployasapplicationconfiguration.html
    /// </summary>
    [OutputType]
    public sealed class ApplicationDeployAsApplicationConfiguration
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-deployasapplicationconfiguration.html#cfn-kinesisanalyticsv2-application-deployasapplicationconfiguration-s3contentlocation
        /// </summary>
        public readonly Outputs.ApplicationS3ContentBaseLocation S3ContentLocation;

        [OutputConstructor]
        private ApplicationDeployAsApplicationConfiguration(Outputs.ApplicationS3ContentBaseLocation s3ContentLocation)
        {
            S3ContentLocation = s3ContentLocation;
        }
    }
}
