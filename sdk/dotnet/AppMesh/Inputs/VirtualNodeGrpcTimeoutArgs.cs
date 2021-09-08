// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AppMesh.Inputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualnode-grpctimeout.html
    /// </summary>
    public sealed class VirtualNodeGrpcTimeoutArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualnode-grpctimeout.html#cfn-appmesh-virtualnode-grpctimeout-idle
        /// </summary>
        [Input("idle")]
        public Input<Inputs.VirtualNodeDurationArgs>? Idle { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualnode-grpctimeout.html#cfn-appmesh-virtualnode-grpctimeout-perrequest
        /// </summary>
        [Input("perRequest")]
        public Input<Inputs.VirtualNodeDurationArgs>? PerRequest { get; set; }

        public VirtualNodeGrpcTimeoutArgs()
        {
        }
    }
}
