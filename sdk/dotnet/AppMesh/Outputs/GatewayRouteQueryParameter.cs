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
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-gatewayroute-queryparameter.html
    /// </summary>
    [OutputType]
    public sealed class GatewayRouteQueryParameter
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-gatewayroute-queryparameter.html#cfn-appmesh-gatewayroute-queryparameter-match
        /// </summary>
        public readonly Outputs.GatewayRouteHttpQueryParameterMatch? Match;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-gatewayroute-queryparameter.html#cfn-appmesh-gatewayroute-queryparameter-name
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private GatewayRouteQueryParameter(
            Outputs.GatewayRouteHttpQueryParameterMatch? match,

            string name)
        {
            Match = match;
            Name = name;
        }
    }
}
