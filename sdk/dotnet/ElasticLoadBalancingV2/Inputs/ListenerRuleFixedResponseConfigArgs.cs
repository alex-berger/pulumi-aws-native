// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ElasticLoadBalancingV2.Inputs
{

    public sealed class ListenerRuleFixedResponseConfigArgs : Pulumi.ResourceArgs
    {
        [Input("contentType")]
        public Input<string>? ContentType { get; set; }

        [Input("messageBody")]
        public Input<string>? MessageBody { get; set; }

        [Input("statusCode", required: true)]
        public Input<string> StatusCode { get; set; } = null!;

        public ListenerRuleFixedResponseConfigArgs()
        {
        }
    }
}
