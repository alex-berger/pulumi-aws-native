// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AmazonMQ.Outputs
{

    [OutputType]
    public sealed class BrokerEncryptionOptions
    {
        public readonly string? KmsKeyId;
        public readonly bool UseAwsOwnedKey;

        [OutputConstructor]
        private BrokerEncryptionOptions(
            string? kmsKeyId,

            bool useAwsOwnedKey)
        {
            KmsKeyId = kmsKeyId;
            UseAwsOwnedKey = useAwsOwnedKey;
        }
    }
}
