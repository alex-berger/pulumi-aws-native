// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ACMPCA.Outputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificateauthority-accessdescription.html
    /// </summary>
    [OutputType]
    public sealed class CertificateAuthorityAccessDescription
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificateauthority-accessdescription.html#cfn-acmpca-certificateauthority-accessdescription-accesslocation
        /// </summary>
        public readonly Outputs.CertificateAuthorityGeneralName AccessLocation;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificateauthority-accessdescription.html#cfn-acmpca-certificateauthority-accessdescription-accessmethod
        /// </summary>
        public readonly Outputs.CertificateAuthorityAccessMethod AccessMethod;

        [OutputConstructor]
        private CertificateAuthorityAccessDescription(
            Outputs.CertificateAuthorityGeneralName accessLocation,

            Outputs.CertificateAuthorityAccessMethod accessMethod)
        {
            AccessLocation = accessLocation;
            AccessMethod = accessMethod;
        }
    }
}