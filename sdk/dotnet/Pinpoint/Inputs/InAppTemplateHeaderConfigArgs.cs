// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Pinpoint.Inputs
{

    public sealed class InAppTemplateHeaderConfigArgs : Pulumi.ResourceArgs
    {
        [Input("alignment")]
        public Input<Pulumi.AwsNative.Pinpoint.InAppTemplateAlignment>? Alignment { get; set; }

        [Input("header")]
        public Input<string>? Header { get; set; }

        [Input("textColor")]
        public Input<string>? TextColor { get; set; }

        public InAppTemplateHeaderConfigArgs()
        {
        }
    }
}
