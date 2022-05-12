// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AppMesh.Outputs
{

    [OutputType]
    public sealed class RouteGrpcRoute
    {
        public readonly Outputs.RouteGrpcRouteAction Action;
        public readonly Outputs.RouteGrpcRouteMatch Match;
        public readonly Outputs.RouteGrpcRetryPolicy? RetryPolicy;
        public readonly Outputs.RouteGrpcTimeout? Timeout;

        [OutputConstructor]
        private RouteGrpcRoute(
            Outputs.RouteGrpcRouteAction action,

            Outputs.RouteGrpcRouteMatch match,

            Outputs.RouteGrpcRetryPolicy? retryPolicy,

            Outputs.RouteGrpcTimeout? timeout)
        {
            Action = action;
            Match = match;
            RetryPolicy = retryPolicy;
            Timeout = timeout;
        }
    }
}
