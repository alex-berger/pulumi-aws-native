// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Evidently.Inputs
{

    public sealed class FeatureEntityOverrideArgs : Pulumi.ResourceArgs
    {
        [Input("entityId")]
        public Input<string>? EntityId { get; set; }

        [Input("variation")]
        public Input<string>? Variation { get; set; }

        public FeatureEntityOverrideArgs()
        {
        }
    }
}
