// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Pinpoint.Inputs
{

    public sealed class SegmentRecencyArgs : Pulumi.ResourceArgs
    {
        [Input("duration", required: true)]
        public Input<string> Duration { get; set; } = null!;

        [Input("recencyType", required: true)]
        public Input<string> RecencyType { get; set; } = null!;

        public SegmentRecencyArgs()
        {
        }
    }
}