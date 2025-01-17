// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.EC2
{
    public static class GetKeyPair
    {
        /// <summary>
        /// The AWS::EC2::KeyPair creates an SSH key pair
        /// </summary>
        public static Task<GetKeyPairResult> InvokeAsync(GetKeyPairArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetKeyPairResult>("aws-native:ec2:getKeyPair", args ?? new GetKeyPairArgs(), options.WithDefaults());

        /// <summary>
        /// The AWS::EC2::KeyPair creates an SSH key pair
        /// </summary>
        public static Output<GetKeyPairResult> Invoke(GetKeyPairInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetKeyPairResult>("aws-native:ec2:getKeyPair", args ?? new GetKeyPairInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetKeyPairArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the SSH key pair
        /// </summary>
        [Input("keyName", required: true)]
        public string KeyName { get; set; } = null!;

        public GetKeyPairArgs()
        {
        }
    }

    public sealed class GetKeyPairInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the SSH key pair
        /// </summary>
        [Input("keyName", required: true)]
        public Input<string> KeyName { get; set; } = null!;

        public GetKeyPairInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetKeyPairResult
    {
        /// <summary>
        /// A short sequence of bytes used for public key verification
        /// </summary>
        public readonly string? KeyFingerprint;
        /// <summary>
        /// An AWS generated ID for the key pair
        /// </summary>
        public readonly string? KeyPairId;
        /// <summary>
        /// Plain text public key to import
        /// </summary>
        public readonly string? PublicKeyMaterial;
        /// <summary>
        /// An array of key-value pairs to apply to this resource.
        /// </summary>
        public readonly ImmutableArray<Outputs.KeyPairTag> Tags;

        [OutputConstructor]
        private GetKeyPairResult(
            string? keyFingerprint,

            string? keyPairId,

            string? publicKeyMaterial,

            ImmutableArray<Outputs.KeyPairTag> tags)
        {
            KeyFingerprint = keyFingerprint;
            KeyPairId = keyPairId;
            PublicKeyMaterial = publicKeyMaterial;
            Tags = tags;
        }
    }
}
