// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.MediaPackage
{
    public static class GetPackagingConfiguration
    {
        /// <summary>
        /// Resource schema for AWS::MediaPackage::PackagingConfiguration
        /// </summary>
        public static Task<GetPackagingConfigurationResult> InvokeAsync(GetPackagingConfigurationArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetPackagingConfigurationResult>("aws-native:mediapackage:getPackagingConfiguration", args ?? new GetPackagingConfigurationArgs(), options.WithDefaults());

        /// <summary>
        /// Resource schema for AWS::MediaPackage::PackagingConfiguration
        /// </summary>
        public static Output<GetPackagingConfigurationResult> Invoke(GetPackagingConfigurationInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetPackagingConfigurationResult>("aws-native:mediapackage:getPackagingConfiguration", args ?? new GetPackagingConfigurationInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetPackagingConfigurationArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the PackagingConfiguration.
        /// </summary>
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetPackagingConfigurationArgs()
        {
        }
    }

    public sealed class GetPackagingConfigurationInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the PackagingConfiguration.
        /// </summary>
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetPackagingConfigurationInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetPackagingConfigurationResult
    {
        /// <summary>
        /// The ARN of the PackagingConfiguration.
        /// </summary>
        public readonly string? Arn;
        /// <summary>
        /// A CMAF packaging configuration.
        /// </summary>
        public readonly Outputs.PackagingConfigurationCmafPackage? CmafPackage;
        /// <summary>
        /// A Dynamic Adaptive Streaming over HTTP (DASH) packaging configuration.
        /// </summary>
        public readonly Outputs.PackagingConfigurationDashPackage? DashPackage;
        /// <summary>
        /// An HTTP Live Streaming (HLS) packaging configuration.
        /// </summary>
        public readonly Outputs.PackagingConfigurationHlsPackage? HlsPackage;
        /// <summary>
        /// A Microsoft Smooth Streaming (MSS) PackagingConfiguration.
        /// </summary>
        public readonly Outputs.PackagingConfigurationMssPackage? MssPackage;
        /// <summary>
        /// The ID of a PackagingGroup.
        /// </summary>
        public readonly string? PackagingGroupId;
        /// <summary>
        /// A collection of tags associated with a resource
        /// </summary>
        public readonly ImmutableArray<Outputs.PackagingConfigurationTag> Tags;

        [OutputConstructor]
        private GetPackagingConfigurationResult(
            string? arn,

            Outputs.PackagingConfigurationCmafPackage? cmafPackage,

            Outputs.PackagingConfigurationDashPackage? dashPackage,

            Outputs.PackagingConfigurationHlsPackage? hlsPackage,

            Outputs.PackagingConfigurationMssPackage? mssPackage,

            string? packagingGroupId,

            ImmutableArray<Outputs.PackagingConfigurationTag> tags)
        {
            Arn = arn;
            CmafPackage = cmafPackage;
            DashPackage = dashPackage;
            HlsPackage = hlsPackage;
            MssPackage = mssPackage;
            PackagingGroupId = packagingGroupId;
            Tags = tags;
        }
    }
}
