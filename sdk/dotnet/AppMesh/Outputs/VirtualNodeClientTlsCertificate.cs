// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AppMesh.Outputs
{

    [OutputType]
    public sealed class VirtualNodeClientTlsCertificate
    {
        public readonly Outputs.VirtualNodeListenerTlsFileCertificate? File;
        public readonly Outputs.VirtualNodeListenerTlsSdsCertificate? SDS;

        [OutputConstructor]
        private VirtualNodeClientTlsCertificate(
            Outputs.VirtualNodeListenerTlsFileCertificate? file,

            Outputs.VirtualNodeListenerTlsSdsCertificate? sDS)
        {
            File = file;
            SDS = sDS;
        }
    }
}
