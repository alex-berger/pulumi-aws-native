// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.NetworkFirewall.Inputs
{

    public sealed class FirewallPolicyStatelessRuleGroupReferenceArgs : Pulumi.ResourceArgs
    {
        [Input("priority", required: true)]
        public Input<int> Priority { get; set; } = null!;

        [Input("resourceArn", required: true)]
        public Input<string> ResourceArn { get; set; } = null!;

        public FirewallPolicyStatelessRuleGroupReferenceArgs()
        {
        }
    }
}
