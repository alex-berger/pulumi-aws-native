// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IAM
{
    public static class GetServerCertificate
    {
        /// <summary>
        /// Resource Type definition for AWS::IAM::ServerCertificate
        /// </summary>
        public static Task<GetServerCertificateResult> InvokeAsync(GetServerCertificateArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetServerCertificateResult>("aws-native:iam:getServerCertificate", args ?? new GetServerCertificateArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::IAM::ServerCertificate
        /// </summary>
        public static Output<GetServerCertificateResult> Invoke(GetServerCertificateInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetServerCertificateResult>("aws-native:iam:getServerCertificate", args ?? new GetServerCertificateInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetServerCertificateArgs : Pulumi.InvokeArgs
    {
        [Input("serverCertificateName", required: true)]
        public string ServerCertificateName { get; set; } = null!;

        public GetServerCertificateArgs()
        {
        }
    }

    public sealed class GetServerCertificateInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("serverCertificateName", required: true)]
        public Input<string> ServerCertificateName { get; set; } = null!;

        public GetServerCertificateInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetServerCertificateResult
    {
        /// <summary>
        /// Amazon Resource Name (ARN) of the server certificate
        /// </summary>
        public readonly string? Arn;
        public readonly string? Path;
        public readonly ImmutableArray<Outputs.ServerCertificateTag> Tags;

        [OutputConstructor]
        private GetServerCertificateResult(
            string? arn,

            string? path,

            ImmutableArray<Outputs.ServerCertificateTag> tags)
        {
            Arn = arn;
            Path = path;
            Tags = tags;
        }
    }
}
