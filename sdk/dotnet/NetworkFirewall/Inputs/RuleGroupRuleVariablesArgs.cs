// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.NetworkFirewall.Inputs
{

    public sealed class RuleGroupRuleVariablesArgs : Pulumi.ResourceArgs
    {
        [Input("iPSets")]
        public Input<object>? IPSets { get; set; }

        [Input("portSets")]
        public Input<object>? PortSets { get; set; }

        public RuleGroupRuleVariablesArgs()
        {
        }
    }
}
