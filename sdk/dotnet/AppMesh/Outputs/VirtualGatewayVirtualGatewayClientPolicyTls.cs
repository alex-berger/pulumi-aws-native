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
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualgateway-virtualgatewayclientpolicytls.html
    /// </summary>
    [OutputType]
    public sealed class VirtualGatewayVirtualGatewayClientPolicyTls
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualgateway-virtualgatewayclientpolicytls.html#cfn-appmesh-virtualgateway-virtualgatewayclientpolicytls-certificate
        /// </summary>
        public readonly Outputs.VirtualGatewayVirtualGatewayClientTlsCertificate? Certificate;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualgateway-virtualgatewayclientpolicytls.html#cfn-appmesh-virtualgateway-virtualgatewayclientpolicytls-enforce
        /// </summary>
        public readonly bool? Enforce;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualgateway-virtualgatewayclientpolicytls.html#cfn-appmesh-virtualgateway-virtualgatewayclientpolicytls-ports
        /// </summary>
        public readonly ImmutableArray<int> Ports;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualgateway-virtualgatewayclientpolicytls.html#cfn-appmesh-virtualgateway-virtualgatewayclientpolicytls-validation
        /// </summary>
        public readonly Outputs.VirtualGatewayVirtualGatewayTlsValidationContext Validation;

        [OutputConstructor]
        private VirtualGatewayVirtualGatewayClientPolicyTls(
            Outputs.VirtualGatewayVirtualGatewayClientTlsCertificate? certificate,

            bool? enforce,

            ImmutableArray<int> ports,

            Outputs.VirtualGatewayVirtualGatewayTlsValidationContext validation)
        {
            Certificate = certificate;
            Enforce = enforce;
            Ports = ports;
            Validation = validation;
        }
    }
}