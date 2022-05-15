// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AppStream
{
    public static class GetDirectoryConfig
    {
        /// <summary>
        /// Resource Type definition for AWS::AppStream::DirectoryConfig
        /// </summary>
        public static Task<GetDirectoryConfigResult> InvokeAsync(GetDirectoryConfigArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetDirectoryConfigResult>("aws-native:appstream:getDirectoryConfig", args ?? new GetDirectoryConfigArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::AppStream::DirectoryConfig
        /// </summary>
        public static Output<GetDirectoryConfigResult> Invoke(GetDirectoryConfigInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetDirectoryConfigResult>("aws-native:appstream:getDirectoryConfig", args ?? new GetDirectoryConfigInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDirectoryConfigArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetDirectoryConfigArgs()
        {
        }
    }

    public sealed class GetDirectoryConfigInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetDirectoryConfigInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetDirectoryConfigResult
    {
        public readonly string? Id;
        public readonly ImmutableArray<string> OrganizationalUnitDistinguishedNames;
        public readonly Outputs.DirectoryConfigServiceAccountCredentials? ServiceAccountCredentials;

        [OutputConstructor]
        private GetDirectoryConfigResult(
            string? id,

            ImmutableArray<string> organizationalUnitDistinguishedNames,

            Outputs.DirectoryConfigServiceAccountCredentials? serviceAccountCredentials)
        {
            Id = id;
            OrganizationalUnitDistinguishedNames = organizationalUnitDistinguishedNames;
            ServiceAccountCredentials = serviceAccountCredentials;
        }
    }
}
