// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Glue.Outputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-job-notificationproperty.html
    /// </summary>
    [OutputType]
    public sealed class JobNotificationProperty
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-job-notificationproperty.html#cfn-glue-job-notificationproperty-notifydelayafter
        /// </summary>
        public readonly int? NotifyDelayAfter;

        [OutputConstructor]
        private JobNotificationProperty(int? notifyDelayAfter)
        {
            NotifyDelayAfter = notifyDelayAfter;
        }
    }
}