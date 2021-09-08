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
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificateauthority-csrextensions.html
    /// </summary>
    [OutputType]
    public sealed class CertificateAuthorityCsrExtensions
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificateauthority-csrextensions.html#cfn-acmpca-certificateauthority-csrextensions-keyusage
        /// </summary>
        public readonly Outputs.CertificateAuthorityKeyUsage? KeyUsage;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificateauthority-csrextensions.html#cfn-acmpca-certificateauthority-csrextensions-subjectinformationaccess
        /// </summary>
        public readonly ImmutableArray<Outputs.CertificateAuthorityAccessDescription> SubjectInformationAccess;

        [OutputConstructor]
        private CertificateAuthorityCsrExtensions(
            Outputs.CertificateAuthorityKeyUsage? keyUsage,

            ImmutableArray<Outputs.CertificateAuthorityAccessDescription> subjectInformationAccess)
        {
            KeyUsage = keyUsage;
            SubjectInformationAccess = subjectInformationAccess;
        }
    }
}