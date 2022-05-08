// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ACMPCA.Inputs
{

    /// <summary>
    /// Your certificate authority can create and maintain a certificate revocation list (CRL). A CRL contains information about certificates that have been revoked.
    /// </summary>
    public sealed class CertificateAuthorityCrlConfigurationArgs : Pulumi.ResourceArgs
    {
        [Input("customCname")]
        public Input<string>? CustomCname { get; set; }

        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        [Input("expirationInDays")]
        public Input<int>? ExpirationInDays { get; set; }

        [Input("s3BucketName")]
        public Input<string>? S3BucketName { get; set; }

        [Input("s3ObjectAcl")]
        public Input<string>? S3ObjectAcl { get; set; }

        public CertificateAuthorityCrlConfigurationArgs()
        {
        }
    }
}
