// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.DLM.Inputs
{

    public sealed class LifecyclePolicyCrossRegionCopyRetainRuleArgs : Pulumi.ResourceArgs
    {
        [Input("interval", required: true)]
        public Input<int> Interval { get; set; } = null!;

        [Input("intervalUnit", required: true)]
        public Input<string> IntervalUnit { get; set; } = null!;

        public LifecyclePolicyCrossRegionCopyRetainRuleArgs()
        {
        }
    }
}
