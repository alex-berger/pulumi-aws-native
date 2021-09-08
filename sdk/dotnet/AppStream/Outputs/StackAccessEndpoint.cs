// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AppStream.Outputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appstream-stack-accessendpoint.html
    /// </summary>
    [OutputType]
    public sealed class StackAccessEndpoint
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appstream-stack-accessendpoint.html#cfn-appstream-stack-accessendpoint-endpointtype
        /// </summary>
        public readonly string EndpointType;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appstream-stack-accessendpoint.html#cfn-appstream-stack-accessendpoint-vpceid
        /// </summary>
        public readonly string VpceId;

        [OutputConstructor]
        private StackAccessEndpoint(
            string endpointType,

            string vpceId)
        {
            EndpointType = endpointType;
            VpceId = vpceId;
        }
    }
}