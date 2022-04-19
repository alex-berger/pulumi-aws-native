// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoT.Outputs
{

    [OutputType]
    public sealed class TopicRuleCloudwatchLogsAction
    {
        public readonly string LogGroupName;
        public readonly string RoleArn;

        [OutputConstructor]
        private TopicRuleCloudwatchLogsAction(
            string logGroupName,

            string roleArn)
        {
            LogGroupName = logGroupName;
            RoleArn = roleArn;
        }
    }
}
