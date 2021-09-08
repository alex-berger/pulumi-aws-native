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
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-route-httprouteheader.html
    /// </summary>
    [OutputType]
    public sealed class RouteHttpRouteHeader
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-route-httprouteheader.html#cfn-appmesh-route-httprouteheader-invert
        /// </summary>
        public readonly bool? Invert;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-route-httprouteheader.html#cfn-appmesh-route-httprouteheader-match
        /// </summary>
        public readonly Outputs.RouteHeaderMatchMethod? Match;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-route-httprouteheader.html#cfn-appmesh-route-httprouteheader-name
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private RouteHttpRouteHeader(
            bool? invert,

            Outputs.RouteHeaderMatchMethod? match,

            string name)
        {
            Invert = invert;
            Match = match;
            Name = name;
        }
    }
}