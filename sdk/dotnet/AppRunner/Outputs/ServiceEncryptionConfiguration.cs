// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AppRunner.Outputs
{

    /// <summary>
    /// Encryption configuration (KMS key)
    /// </summary>
    [OutputType]
    public sealed class ServiceEncryptionConfiguration
    {
        /// <summary>
        /// The KMS Key
        /// </summary>
        public readonly string KmsKey;

        [OutputConstructor]
        private ServiceEncryptionConfiguration(string kmsKey)
        {
            KmsKey = kmsKey;
        }
    }
}
