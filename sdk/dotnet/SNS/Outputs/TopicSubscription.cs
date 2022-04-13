// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SNS.Outputs
{

    [OutputType]
    public sealed class TopicSubscription
    {
        public readonly string Endpoint;
        public readonly string Protocol;

        [OutputConstructor]
        private TopicSubscription(
            string endpoint,

            string protocol)
        {
            Endpoint = endpoint;
            Protocol = protocol;
        }
    }
}
