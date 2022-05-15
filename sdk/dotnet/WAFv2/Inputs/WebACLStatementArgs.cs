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
    /// First level statement that contains conditions, such as ByteMatch, SizeConstraint, etc
    /// </summary>
    public sealed class WebACLStatementArgs : Pulumi.ResourceArgs
    {
        [Input("andStatement")]
        public Input<Inputs.WebACLAndStatementArgs>? AndStatement { get; set; }

        [Input("byteMatchStatement")]
        public Input<Inputs.WebACLByteMatchStatementArgs>? ByteMatchStatement { get; set; }

        [Input("geoMatchStatement")]
        public Input<Inputs.WebACLGeoMatchStatementArgs>? GeoMatchStatement { get; set; }

        [Input("iPSetReferenceStatement")]
        public Input<Inputs.WebACLIPSetReferenceStatementArgs>? IPSetReferenceStatement { get; set; }

        [Input("labelMatchStatement")]
        public Input<Inputs.WebACLLabelMatchStatementArgs>? LabelMatchStatement { get; set; }

        [Input("managedRuleGroupStatement")]
        public Input<Inputs.WebACLManagedRuleGroupStatementArgs>? ManagedRuleGroupStatement { get; set; }

        [Input("notStatement")]
        public Input<Inputs.WebACLNotStatementArgs>? NotStatement { get; set; }

        [Input("orStatement")]
        public Input<Inputs.WebACLOrStatementArgs>? OrStatement { get; set; }

        [Input("rateBasedStatement")]
        public Input<Inputs.WebACLRateBasedStatementArgs>? RateBasedStatement { get; set; }

        [Input("regexMatchStatement")]
        public Input<Inputs.WebACLRegexMatchStatementArgs>? RegexMatchStatement { get; set; }

        [Input("regexPatternSetReferenceStatement")]
        public Input<Inputs.WebACLRegexPatternSetReferenceStatementArgs>? RegexPatternSetReferenceStatement { get; set; }

        [Input("ruleGroupReferenceStatement")]
        public Input<Inputs.WebACLRuleGroupReferenceStatementArgs>? RuleGroupReferenceStatement { get; set; }

        [Input("sizeConstraintStatement")]
        public Input<Inputs.WebACLSizeConstraintStatementArgs>? SizeConstraintStatement { get; set; }

        [Input("sqliMatchStatement")]
        public Input<Inputs.WebACLSqliMatchStatementArgs>? SqliMatchStatement { get; set; }

        [Input("xssMatchStatement")]
        public Input<Inputs.WebACLXssMatchStatementArgs>? XssMatchStatement { get; set; }

        public WebACLStatementArgs()
        {
        }
    }
}
