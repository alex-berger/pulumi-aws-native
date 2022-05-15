// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.OpenSearchService.Outputs
{

    [OutputType]
    public sealed class DomainEncryptionAtRestOptions
    {
        public readonly bool? Enabled;
        public readonly string? KmsKeyId;

        [OutputConstructor]
        private DomainEncryptionAtRestOptions(
            bool? enabled,

            string? kmsKeyId)
        {
            Enabled = enabled;
            KmsKeyId = kmsKeyId;
        }
    }
}
