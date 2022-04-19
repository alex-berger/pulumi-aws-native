// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.MediaConnect.Inputs
{

    /// <summary>
    /// Information about the encryption of the flow.
    /// </summary>
    public sealed class FlowEntitlementEncryptionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The type of algorithm that is used for the encryption (such as aes128, aes192, or aes256).
        /// </summary>
        [Input("algorithm", required: true)]
        public Input<Pulumi.AwsNative.MediaConnect.FlowEntitlementEncryptionAlgorithm> Algorithm { get; set; } = null!;

        /// <summary>
        /// A 128-bit, 16-byte hex value represented by a 32-character string, to be used with the key for encrypting content. This parameter is not valid for static key encryption.
        /// </summary>
        [Input("constantInitializationVector")]
        public Input<string>? ConstantInitializationVector { get; set; }

        /// <summary>
        /// The value of one of the devices that you configured with your digital rights management (DRM) platform key provider. This parameter is required for SPEKE encryption and is not valid for static key encryption.
        /// </summary>
        [Input("deviceId")]
        public Input<string>? DeviceId { get; set; }

        /// <summary>
        /// The type of key that is used for the encryption. If no keyType is provided, the service will use the default setting (static-key).
        /// </summary>
        [Input("keyType")]
        public Input<Pulumi.AwsNative.MediaConnect.FlowEntitlementEncryptionKeyType>? KeyType { get; set; }

        /// <summary>
        /// The AWS Region that the API Gateway proxy endpoint was created in. This parameter is required for SPEKE encryption and is not valid for static key encryption.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// An identifier for the content. The service sends this value to the key server to identify the current endpoint. The resource ID is also known as the content ID. This parameter is required for SPEKE encryption and is not valid for static key encryption.
        /// </summary>
        [Input("resourceId")]
        public Input<string>? ResourceId { get; set; }

        /// <summary>
        /// The ARN of the role that you created during setup (when you set up AWS Elemental MediaConnect as a trusted entity).
        /// </summary>
        [Input("roleArn", required: true)]
        public Input<string> RoleArn { get; set; } = null!;

        /// <summary>
        ///  The ARN of the secret that you created in AWS Secrets Manager to store the encryption key. This parameter is required for static key encryption and is not valid for SPEKE encryption.
        /// </summary>
        [Input("secretArn")]
        public Input<string>? SecretArn { get; set; }

        /// <summary>
        /// The URL from the API Gateway proxy that you set up to talk to your key server. This parameter is required for SPEKE encryption and is not valid for static key encryption.
        /// </summary>
        [Input("url")]
        public Input<string>? Url { get; set; }

        public FlowEntitlementEncryptionArgs()
        {
        }
    }
}
