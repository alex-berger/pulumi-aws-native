// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.WAFv2.Inputs
{

    public sealed class WebACLRuleGroupReferenceStatementArgs : Pulumi.ResourceArgs
    {
        [Input("arn", required: true)]
        public Input<string> Arn { get; set; } = null!;

        [Input("excludedRules")]
        private InputList<Inputs.WebACLExcludedRuleArgs>? _excludedRules;
        public InputList<Inputs.WebACLExcludedRuleArgs> ExcludedRules
        {
            get => _excludedRules ?? (_excludedRules = new InputList<Inputs.WebACLExcludedRuleArgs>());
            set => _excludedRules = value;
        }

        public WebACLRuleGroupReferenceStatementArgs()
        {
        }
    }
}