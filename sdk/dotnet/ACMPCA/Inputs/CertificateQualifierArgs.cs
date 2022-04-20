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
    /// Structure that contains a X.509 policy qualifier.
    /// </summary>
    public sealed class CertificateQualifierArgs : Pulumi.ResourceArgs
    {
        [Input("cpsUri", required: true)]
        public Input<string> CpsUri { get; set; } = null!;

        public CertificateQualifierArgs()
        {
        }
    }
}
