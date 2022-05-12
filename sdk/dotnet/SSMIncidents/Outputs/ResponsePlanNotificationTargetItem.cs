// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SSMIncidents.Outputs
{

    /// <summary>
    /// A notification target.
    /// </summary>
    [OutputType]
    public sealed class ResponsePlanNotificationTargetItem
    {
        public readonly string? SnsTopicArn;

        [OutputConstructor]
        private ResponsePlanNotificationTargetItem(string? snsTopicArn)
        {
            SnsTopicArn = snsTopicArn;
        }
    }
}
