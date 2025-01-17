// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ElasticLoadBalancingV2.Inputs
{

    public sealed class ListenerActionArgs : Pulumi.ResourceArgs
    {
        [Input("authenticateCognitoConfig")]
        public Input<Inputs.ListenerAuthenticateCognitoConfigArgs>? AuthenticateCognitoConfig { get; set; }

        [Input("authenticateOidcConfig")]
        public Input<Inputs.ListenerAuthenticateOidcConfigArgs>? AuthenticateOidcConfig { get; set; }

        [Input("fixedResponseConfig")]
        public Input<Inputs.ListenerFixedResponseConfigArgs>? FixedResponseConfig { get; set; }

        [Input("forwardConfig")]
        public Input<Inputs.ListenerForwardConfigArgs>? ForwardConfig { get; set; }

        [Input("order")]
        public Input<int>? Order { get; set; }

        [Input("redirectConfig")]
        public Input<Inputs.ListenerRedirectConfigArgs>? RedirectConfig { get; set; }

        [Input("targetGroupArn")]
        public Input<string>? TargetGroupArn { get; set; }

        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public ListenerActionArgs()
        {
        }
    }
}
