// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.EMR
{
    public static class GetStudio
    {
        /// <summary>
        /// Resource schema for AWS::EMR::Studio
        /// </summary>
        public static Task<GetStudioResult> InvokeAsync(GetStudioArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetStudioResult>("aws-native:emr:getStudio", args ?? new GetStudioArgs(), options.WithDefaults());

        /// <summary>
        /// Resource schema for AWS::EMR::Studio
        /// </summary>
        public static Output<GetStudioResult> Invoke(GetStudioInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetStudioResult>("aws-native:emr:getStudio", args ?? new GetStudioInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetStudioArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the EMR Studio.
        /// </summary>
        [Input("studioId", required: true)]
        public string StudioId { get; set; } = null!;

        public GetStudioArgs()
        {
        }
    }

    public sealed class GetStudioInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the EMR Studio.
        /// </summary>
        [Input("studioId", required: true)]
        public Input<string> StudioId { get; set; } = null!;

        public GetStudioInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetStudioResult
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the EMR Studio.
        /// </summary>
        public readonly string? Arn;
        /// <summary>
        /// The default Amazon S3 location to back up EMR Studio Workspaces and notebook files. A Studio user can select an alternative Amazon S3 location when creating a Workspace.
        /// </summary>
        public readonly string? DefaultS3Location;
        /// <summary>
        /// A detailed description of the Studio.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// Your identity provider's authentication endpoint. Amazon EMR Studio redirects federated users to this endpoint for authentication when logging in to a Studio with the Studio URL.
        /// </summary>
        public readonly string? IdpAuthUrl;
        /// <summary>
        /// The name of relay state parameter for external Identity Provider.
        /// </summary>
        public readonly string? IdpRelayStateParameterName;
        /// <summary>
        /// A descriptive name for the Amazon EMR Studio.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// The ID of the EMR Studio.
        /// </summary>
        public readonly string? StudioId;
        /// <summary>
        /// A list of up to 5 subnet IDs to associate with the Studio. The subnets must belong to the VPC specified by VpcId. Studio users can create a Workspace in any of the specified subnets.
        /// </summary>
        public readonly ImmutableArray<string> SubnetIds;
        /// <summary>
        /// A list of tags to associate with the Studio. Tags are user-defined key-value pairs that consist of a required key string with a maximum of 128 characters, and an optional value string with a maximum of 256 characters.
        /// </summary>
        public readonly ImmutableArray<Outputs.StudioTag> Tags;
        /// <summary>
        /// The unique Studio access URL.
        /// </summary>
        public readonly string? Url;

        [OutputConstructor]
        private GetStudioResult(
            string? arn,

            string? defaultS3Location,

            string? description,

            string? idpAuthUrl,

            string? idpRelayStateParameterName,

            string? name,

            string? studioId,

            ImmutableArray<string> subnetIds,

            ImmutableArray<Outputs.StudioTag> tags,

            string? url)
        {
            Arn = arn;
            DefaultS3Location = defaultS3Location;
            Description = description;
            IdpAuthUrl = idpAuthUrl;
            IdpRelayStateParameterName = idpRelayStateParameterName;
            Name = name;
            StudioId = studioId;
            SubnetIds = subnetIds;
            Tags = tags;
            Url = url;
        }
    }
}
