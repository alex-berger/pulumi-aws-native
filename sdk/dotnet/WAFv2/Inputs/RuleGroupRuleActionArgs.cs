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
    /// Action taken when Rule matches its condition.
    /// </summary>
    public sealed class RuleGroupRuleActionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Allow traffic towards application.
        /// </summary>
        [Input("allow")]
        public Input<Inputs.RuleGroupRuleActionAllowPropertiesArgs>? Allow { get; set; }

        /// <summary>
        /// Block traffic towards application.
        /// </summary>
        [Input("block")]
        public Input<Inputs.RuleGroupRuleActionBlockPropertiesArgs>? Block { get; set; }

        /// <summary>
        /// Checks valid token exists with request.
        /// </summary>
        [Input("captcha")]
        public Input<Inputs.RuleGroupRuleActionCaptchaPropertiesArgs>? Captcha { get; set; }

        /// <summary>
        /// Count traffic towards application.
        /// </summary>
        [Input("count")]
        public Input<Inputs.RuleGroupRuleActionCountPropertiesArgs>? Count { get; set; }

        public RuleGroupRuleActionArgs()
        {
        }
    }
}
