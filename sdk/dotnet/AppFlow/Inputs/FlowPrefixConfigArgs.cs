// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AppFlow.Inputs
{

    public sealed class FlowPrefixConfigArgs : Pulumi.ResourceArgs
    {
        [Input("prefixFormat")]
        public Input<Pulumi.AwsNative.AppFlow.FlowPrefixFormat>? PrefixFormat { get; set; }

        [Input("prefixType")]
        public Input<Pulumi.AwsNative.AppFlow.FlowPrefixType>? PrefixType { get; set; }

        public FlowPrefixConfigArgs()
        {
        }
    }
}
