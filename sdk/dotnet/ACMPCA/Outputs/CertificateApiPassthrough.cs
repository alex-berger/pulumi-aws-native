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
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificate-apipassthrough.html
    /// </summary>
    [OutputType]
    public sealed class CertificateApiPassthrough
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificate-apipassthrough.html#cfn-acmpca-certificate-apipassthrough-extensions
        /// </summary>
        public readonly Outputs.CertificateExtensions? Extensions;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificate-apipassthrough.html#cfn-acmpca-certificate-apipassthrough-subject
        /// </summary>
        public readonly Outputs.CertificateSubject? Subject;

        [OutputConstructor]
        private CertificateApiPassthrough(
            Outputs.CertificateExtensions? extensions,

            Outputs.CertificateSubject? subject)
        {
            Extensions = extensions;
            Subject = subject;
        }
    }
}