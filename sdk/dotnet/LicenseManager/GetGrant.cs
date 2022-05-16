// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.LicenseManager
{
    public static class GetGrant
    {
        /// <summary>
        /// An example resource schema demonstrating some basic constructs and validation rules.
        /// </summary>
        public static Task<GetGrantResult> InvokeAsync(GetGrantArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetGrantResult>("aws-native:licensemanager:getGrant", args ?? new GetGrantArgs(), options.WithDefaults());

        /// <summary>
        /// An example resource schema demonstrating some basic constructs and validation rules.
        /// </summary>
        public static Output<GetGrantResult> Invoke(GetGrantInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetGrantResult>("aws-native:licensemanager:getGrant", args ?? new GetGrantInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetGrantArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Arn of the grant.
        /// </summary>
        [Input("grantArn", required: true)]
        public string GrantArn { get; set; } = null!;

        public GetGrantArgs()
        {
        }
    }

    public sealed class GetGrantInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Arn of the grant.
        /// </summary>
        [Input("grantArn", required: true)]
        public Input<string> GrantArn { get; set; } = null!;

        public GetGrantInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetGrantResult
    {
        /// <summary>
        /// Arn of the grant.
        /// </summary>
        public readonly string? GrantArn;
        /// <summary>
        /// Name for the created Grant.
        /// </summary>
        public readonly string? GrantName;
        /// <summary>
        /// Home region for the created grant.
        /// </summary>
        public readonly string? HomeRegion;
        /// <summary>
        /// License Arn for the grant.
        /// </summary>
        public readonly string? LicenseArn;
        public readonly string? Status;
        /// <summary>
        /// The version of the grant.
        /// </summary>
        public readonly string? Version;

        [OutputConstructor]
        private GetGrantResult(
            string? grantArn,

            string? grantName,

            string? homeRegion,

            string? licenseArn,

            string? status,

            string? version)
        {
            GrantArn = grantArn;
            GrantName = grantName;
            HomeRegion = homeRegion;
            LicenseArn = licenseArn;
            Status = status;
            Version = version;
        }
    }
}