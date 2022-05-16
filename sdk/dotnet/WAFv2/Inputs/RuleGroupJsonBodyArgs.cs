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
    /// Inspect the request body as JSON. The request body immediately follows the request headers.
    /// </summary>
    public sealed class RuleGroupJsonBodyArgs : Pulumi.ResourceArgs
    {
        [Input("invalidFallbackBehavior")]
        public Input<Pulumi.AwsNative.WAFv2.RuleGroupBodyParsingFallbackBehavior>? InvalidFallbackBehavior { get; set; }

        [Input("matchPattern", required: true)]
        public Input<Inputs.RuleGroupJsonMatchPatternArgs> MatchPattern { get; set; } = null!;

        [Input("matchScope", required: true)]
        public Input<Pulumi.AwsNative.WAFv2.RuleGroupJsonMatchScope> MatchScope { get; set; } = null!;

        [Input("oversizeHandling")]
        public Input<Pulumi.AwsNative.WAFv2.RuleGroupOversizeHandling>? OversizeHandling { get; set; }

        public RuleGroupJsonBodyArgs()
        {
        }
    }
}
