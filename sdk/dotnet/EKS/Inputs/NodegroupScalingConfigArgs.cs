// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.EKS.Inputs
{

    /// <summary>
    /// An object representing a auto scaling group specification for AWS EKS Nodegroup.
    /// </summary>
    public sealed class NodegroupScalingConfigArgs : Pulumi.ResourceArgs
    {
        [Input("desiredSize")]
        public Input<int>? DesiredSize { get; set; }

        [Input("maxSize")]
        public Input<int>? MaxSize { get; set; }

        [Input("minSize")]
        public Input<int>? MinSize { get; set; }

        public NodegroupScalingConfigArgs()
        {
        }
    }
}
