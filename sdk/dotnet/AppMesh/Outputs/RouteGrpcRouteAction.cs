// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AppMesh.Outputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-route-grpcrouteaction.html
    /// </summary>
    [OutputType]
    public sealed class RouteGrpcRouteAction
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-route-grpcrouteaction.html#cfn-appmesh-route-grpcrouteaction-weightedtargets
        /// </summary>
        public readonly ImmutableArray<Outputs.RouteWeightedTarget> WeightedTargets;

        [OutputConstructor]
        private RouteGrpcRouteAction(ImmutableArray<Outputs.RouteWeightedTarget> weightedTargets)
        {
            WeightedTargets = weightedTargets;
        }
    }
}
