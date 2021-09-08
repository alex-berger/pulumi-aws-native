// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.DataBrew.Outputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-job-jobsample.html
    /// </summary>
    [OutputType]
    public sealed class JobJobSample
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-job-jobsample.html#cfn-databrew-job-jobsample-mode
        /// </summary>
        public readonly string? Mode;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-job-jobsample.html#cfn-databrew-job-jobsample-size
        /// </summary>
        public readonly int? Size;

        [OutputConstructor]
        private JobJobSample(
            string? mode,

            int? size)
        {
            Mode = mode;
            Size = size;
        }
    }
}