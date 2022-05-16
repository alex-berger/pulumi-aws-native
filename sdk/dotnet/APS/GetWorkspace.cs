// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.APS
{
    public static class GetWorkspace
    {
        /// <summary>
        /// Resource Type definition for AWS::APS::Workspace
        /// </summary>
        public static Task<GetWorkspaceResult> InvokeAsync(GetWorkspaceArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetWorkspaceResult>("aws-native:aps:getWorkspace", args ?? new GetWorkspaceArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::APS::Workspace
        /// </summary>
        public static Output<GetWorkspaceResult> Invoke(GetWorkspaceInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetWorkspaceResult>("aws-native:aps:getWorkspace", args ?? new GetWorkspaceInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetWorkspaceArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Workspace arn.
        /// </summary>
        [Input("arn", required: true)]
        public string Arn { get; set; } = null!;

        public GetWorkspaceArgs()
        {
        }
    }

    public sealed class GetWorkspaceInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Workspace arn.
        /// </summary>
        [Input("arn", required: true)]
        public Input<string> Arn { get; set; } = null!;

        public GetWorkspaceInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetWorkspaceResult
    {
        /// <summary>
        /// The AMP Workspace alert manager definition data
        /// </summary>
        public readonly string? AlertManagerDefinition;
        /// <summary>
        /// AMP Workspace alias.
        /// </summary>
        public readonly string? Alias;
        /// <summary>
        /// Workspace arn.
        /// </summary>
        public readonly string? Arn;
        /// <summary>
        /// AMP Workspace prometheus endpoint
        /// </summary>
        public readonly string? PrometheusEndpoint;
        /// <summary>
        /// An array of key-value pairs to apply to this resource.
        /// </summary>
        public readonly ImmutableArray<Outputs.WorkspaceTag> Tags;
        /// <summary>
        /// Required to identify a specific APS Workspace.
        /// </summary>
        public readonly string? WorkspaceId;

        [OutputConstructor]
        private GetWorkspaceResult(
            string? alertManagerDefinition,

            string? alias,

            string? arn,

            string? prometheusEndpoint,

            ImmutableArray<Outputs.WorkspaceTag> tags,

            string? workspaceId)
        {
            AlertManagerDefinition = alertManagerDefinition;
            Alias = alias;
            Arn = arn;
            PrometheusEndpoint = prometheusEndpoint;
            Tags = tags;
            WorkspaceId = workspaceId;
        }
    }
}