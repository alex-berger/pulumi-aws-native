// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoT.Inputs
{

    public sealed class TopicRuleKinesisActionArgs : Pulumi.ResourceArgs
    {
        [Input("partitionKey")]
        public Input<string>? PartitionKey { get; set; }

        [Input("roleArn", required: true)]
        public Input<string> RoleArn { get; set; } = null!;

        [Input("streamName", required: true)]
        public Input<string> StreamName { get; set; } = null!;

        public TopicRuleKinesisActionArgs()
        {
        }
    }
}
