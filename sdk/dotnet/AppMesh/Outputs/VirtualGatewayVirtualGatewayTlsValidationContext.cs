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
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualgateway-virtualgatewaytlsvalidationcontext.html
    /// </summary>
    [OutputType]
    public sealed class VirtualGatewayVirtualGatewayTlsValidationContext
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualgateway-virtualgatewaytlsvalidationcontext.html#cfn-appmesh-virtualgateway-virtualgatewaytlsvalidationcontext-subjectalternativenames
        /// </summary>
        public readonly Outputs.VirtualGatewaySubjectAlternativeNames? SubjectAlternativeNames;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualgateway-virtualgatewaytlsvalidationcontext.html#cfn-appmesh-virtualgateway-virtualgatewaytlsvalidationcontext-trust
        /// </summary>
        public readonly Outputs.VirtualGatewayVirtualGatewayTlsValidationContextTrust Trust;

        [OutputConstructor]
        private VirtualGatewayVirtualGatewayTlsValidationContext(
            Outputs.VirtualGatewaySubjectAlternativeNames? subjectAlternativeNames,

            Outputs.VirtualGatewayVirtualGatewayTlsValidationContextTrust trust)
        {
            SubjectAlternativeNames = subjectAlternativeNames;
            Trust = trust;
        }
    }
}