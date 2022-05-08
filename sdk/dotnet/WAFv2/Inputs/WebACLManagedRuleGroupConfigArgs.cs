// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.WAFv2.Inputs
{

    /// <summary>
    /// ManagedRuleGroupConfig.
    /// </summary>
    public sealed class WebACLManagedRuleGroupConfigArgs : Pulumi.ResourceArgs
    {
        [Input("loginPath")]
        public Input<string>? LoginPath { get; set; }

        [Input("passwordField")]
        public Input<Inputs.WebACLFieldIdentifierArgs>? PasswordField { get; set; }

        [Input("payloadType")]
        public Input<Pulumi.AwsNative.WAFv2.WebACLManagedRuleGroupConfigPayloadType>? PayloadType { get; set; }

        [Input("usernameField")]
        public Input<Inputs.WebACLFieldIdentifierArgs>? UsernameField { get; set; }

        public WebACLManagedRuleGroupConfigArgs()
        {
        }
    }
}
