// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ACMPCA
{
    /// <summary>
    /// A certificate issued via a private certificate authority
    /// </summary>
    [AwsNativeResourceType("aws-native:acmpca:Certificate")]
    public partial class Certificate : Pulumi.CustomResource
    {
        /// <summary>
        /// These are fields to be overridden in a certificate at the time of issuance. These requires an API_Passthrough template be used or they will be ignored.
        /// </summary>
        [Output("apiPassthrough")]
        public Output<Outputs.CertificateApiPassthrough?> ApiPassthrough { get; private set; } = null!;

        /// <summary>
        /// The ARN of the issued certificate.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The issued certificate in base 64 PEM-encoded format.
        /// </summary>
        [Output("certificate")]
        public Output<string> CertificateValue { get; private set; } = null!;

        /// <summary>
        /// The Amazon Resource Name (ARN) for the private CA to issue the certificate.
        /// </summary>
        [Output("certificateAuthorityArn")]
        public Output<string> CertificateAuthorityArn { get; private set; } = null!;

        /// <summary>
        /// The certificate signing request (CSR) for the Certificate.
        /// </summary>
        [Output("certificateSigningRequest")]
        public Output<string> CertificateSigningRequest { get; private set; } = null!;

        /// <summary>
        /// The name of the algorithm that will be used to sign the Certificate.
        /// </summary>
        [Output("signingAlgorithm")]
        public Output<string> SigningAlgorithm { get; private set; } = null!;

        /// <summary>
        /// Specifies a custom configuration template to use when issuing a certificate. If this parameter is not provided, ACM Private CA defaults to the EndEntityCertificate/V1 template.
        /// </summary>
        [Output("templateArn")]
        public Output<string?> TemplateArn { get; private set; } = null!;

        /// <summary>
        /// The time before which the Certificate will be valid.
        /// </summary>
        [Output("validity")]
        public Output<Outputs.CertificateValidity> Validity { get; private set; } = null!;

        /// <summary>
        /// The time after which the Certificate will be valid.
        /// </summary>
        [Output("validityNotBefore")]
        public Output<Outputs.CertificateValidity?> ValidityNotBefore { get; private set; } = null!;


        /// <summary>
        /// Create a Certificate resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Certificate(string name, CertificateArgs args, CustomResourceOptions? options = null)
            : base("aws-native:acmpca:Certificate", name, args ?? new CertificateArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Certificate(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:acmpca:Certificate", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Certificate resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Certificate Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Certificate(name, id, options);
        }
    }

    public sealed class CertificateArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// These are fields to be overridden in a certificate at the time of issuance. These requires an API_Passthrough template be used or they will be ignored.
        /// </summary>
        [Input("apiPassthrough")]
        public Input<Inputs.CertificateApiPassthroughArgs>? ApiPassthrough { get; set; }

        /// <summary>
        /// The Amazon Resource Name (ARN) for the private CA to issue the certificate.
        /// </summary>
        [Input("certificateAuthorityArn", required: true)]
        public Input<string> CertificateAuthorityArn { get; set; } = null!;

        /// <summary>
        /// The certificate signing request (CSR) for the Certificate.
        /// </summary>
        [Input("certificateSigningRequest", required: true)]
        public Input<string> CertificateSigningRequest { get; set; } = null!;

        /// <summary>
        /// The name of the algorithm that will be used to sign the Certificate.
        /// </summary>
        [Input("signingAlgorithm", required: true)]
        public Input<string> SigningAlgorithm { get; set; } = null!;

        /// <summary>
        /// Specifies a custom configuration template to use when issuing a certificate. If this parameter is not provided, ACM Private CA defaults to the EndEntityCertificate/V1 template.
        /// </summary>
        [Input("templateArn")]
        public Input<string>? TemplateArn { get; set; }

        /// <summary>
        /// The time before which the Certificate will be valid.
        /// </summary>
        [Input("validity", required: true)]
        public Input<Inputs.CertificateValidityArgs> Validity { get; set; } = null!;

        /// <summary>
        /// The time after which the Certificate will be valid.
        /// </summary>
        [Input("validityNotBefore")]
        public Input<Inputs.CertificateValidityArgs>? ValidityNotBefore { get; set; }

        public CertificateArgs()
        {
        }
    }
}
