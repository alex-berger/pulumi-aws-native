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
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-route-tcprouteaction.html
    /// </summary>
    public sealed class RouteTcpRouteActionArgs : Pulumi.ResourceArgs
    {
        [Input("weightedTargets", required: true)]
        private InputList<Inputs.RouteWeightedTargetArgs>? _weightedTargets;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-route-tcprouteaction.html#cfn-appmesh-route-tcprouteaction-weightedtargets
        /// </summary>
        public InputList<Inputs.RouteWeightedTargetArgs> WeightedTargets
        {
            get => _weightedTargets ?? (_weightedTargets = new InputList<Inputs.RouteWeightedTargetArgs>());
            set => _weightedTargets = value;
        }

        public RouteTcpRouteActionArgs()
        {
        }
    }
}
