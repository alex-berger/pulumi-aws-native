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
    public sealed class TopicRuleKinesisAction
    {
        public readonly string? PartitionKey;
        public readonly string RoleArn;
        public readonly string StreamName;

        [OutputConstructor]
        private TopicRuleKinesisAction(
            string? partitionKey,

            string roleArn,

            string streamName)
        {
            PartitionKey = partitionKey;
            RoleArn = roleArn;
            StreamName = streamName;
        }
    }
}