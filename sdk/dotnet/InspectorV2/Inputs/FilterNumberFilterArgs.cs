// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.InspectorV2.Inputs
{

    public sealed class FilterNumberFilterArgs : Pulumi.ResourceArgs
    {
        [Input("lowerInclusive")]
        public Input<double>? LowerInclusive { get; set; }

        [Input("upperInclusive")]
        public Input<double>? UpperInclusive { get; set; }

        public FilterNumberFilterArgs()
        {
        }
    }
}
