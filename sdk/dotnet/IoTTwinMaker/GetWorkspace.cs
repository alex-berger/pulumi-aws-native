// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoTTwinMaker
{
    public static class GetWorkspace
    {
        /// <summary>
        /// Resource schema for AWS::IoTTwinMaker::Workspace
        /// </summary>
        public static Task<GetWorkspaceResult> InvokeAsync(GetWorkspaceArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetWorkspaceResult>("aws-native:iottwinmaker:getWorkspace", args ?? new GetWorkspaceArgs(), options.WithDefaults());

        /// <summary>
        /// Resource schema for AWS::IoTTwinMaker::Workspace
        /// </summary>
        public static Output<GetWorkspaceResult> Invoke(GetWorkspaceInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetWorkspaceResult>("aws-native:iottwinmaker:getWorkspace", args ?? new GetWorkspaceInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetWorkspaceArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the workspace.
        /// </summary>
        [Input("workspaceId", required: true)]
        public string WorkspaceId { get; set; } = null!;

        public GetWorkspaceArgs()
        {
        }
    }

    public sealed class GetWorkspaceInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the workspace.
        /// </summary>
        [Input("workspaceId", required: true)]
        public Input<string> WorkspaceId { get; set; } = null!;

        public GetWorkspaceInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetWorkspaceResult
    {
        /// <summary>
        /// The ARN of the workspace.
        /// </summary>
        public readonly string? Arn;
        /// <summary>
        /// The date and time when the workspace was created.
        /// </summary>
        public readonly string? CreationDateTime;
        /// <summary>
        /// The description of the workspace.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// The ARN of the execution role associated with the workspace.
        /// </summary>
        public readonly string? Role;
        /// <summary>
        /// The ARN of the S3 bucket where resources associated with the workspace are stored.
        /// </summary>
        public readonly string? S3Location;
        /// <summary>
        /// A map of key-value pairs to associate with a resource.
        /// </summary>
        public readonly object? Tags;
        /// <summary>
        /// The date and time of the current update.
        /// </summary>
        public readonly string? UpdateDateTime;

        [OutputConstructor]
        private GetWorkspaceResult(
            string? arn,

            string? creationDateTime,

            string? description,

            string? role,

            string? s3Location,

            object? tags,

            string? updateDateTime)
        {
            Arn = arn;
            CreationDateTime = creationDateTime;
            Description = description;
            Role = role;
            S3Location = s3Location;
            Tags = tags;
            UpdateDateTime = updateDateTime;
        }
    }
}