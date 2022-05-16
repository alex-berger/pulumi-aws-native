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
    public sealed class WebACLJsonBodyArgs : Pulumi.ResourceArgs
    {
        [Input("invalidFallbackBehavior")]
        public Input<Pulumi.AwsNative.WAFv2.WebACLBodyParsingFallbackBehavior>? InvalidFallbackBehavior { get; set; }

        [Input("matchPattern", required: true)]
        public Input<Inputs.WebACLJsonMatchPatternArgs> MatchPattern { get; set; } = null!;

        [Input("matchScope", required: true)]
        public Input<Pulumi.AwsNative.WAFv2.WebACLJsonMatchScope> MatchScope { get; set; } = null!;

        [Input("oversizeHandling")]
        public Input<Pulumi.AwsNative.WAFv2.WebACLOversizeHandling>? OversizeHandling { get; set; }

        public WebACLJsonBodyArgs()
        {
        }
    }
}
