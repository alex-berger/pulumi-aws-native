// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.XRay
{
    public static class GetGroup
    {
        /// <summary>
        /// This schema provides construct and validation rules for AWS-XRay Group resource parameters.
        /// </summary>
        public static Task<GetGroupResult> InvokeAsync(GetGroupArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetGroupResult>("aws-native:xray:getGroup", args ?? new GetGroupArgs(), options.WithDefaults());

        /// <summary>
        /// This schema provides construct and validation rules for AWS-XRay Group resource parameters.
        /// </summary>
        public static Output<GetGroupResult> Invoke(GetGroupInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetGroupResult>("aws-native:xray:getGroup", args ?? new GetGroupInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetGroupArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ARN of the group that was generated on creation.
        /// </summary>
        [Input("groupARN", required: true)]
        public string GroupARN { get; set; } = null!;

        public GetGroupArgs()
        {
        }
    }

    public sealed class GetGroupInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ARN of the group that was generated on creation.
        /// </summary>
        [Input("groupARN", required: true)]
        public Input<string> GroupARN { get; set; } = null!;

        public GetGroupInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetGroupResult
    {
        /// <summary>
        /// The filter expression defining criteria by which to group traces.
        /// </summary>
        public readonly string? FilterExpression;
        /// <summary>
        /// The ARN of the group that was generated on creation.
        /// </summary>
        public readonly string? GroupARN;
        /// <summary>
        /// The case-sensitive name of the new group. Names must be unique.
        /// </summary>
        public readonly string? GroupName;
        public readonly Outputs.GroupInsightsConfiguration? InsightsConfiguration;
        public readonly ImmutableArray<Outputs.TagsItemProperties> Tags;

        [OutputConstructor]
        private GetGroupResult(
            string? filterExpression,

            string? groupARN,

            string? groupName,

            Outputs.GroupInsightsConfiguration? insightsConfiguration,

            ImmutableArray<Outputs.TagsItemProperties> tags)
        {
            FilterExpression = filterExpression;
            GroupARN = groupARN;
            GroupName = groupName;
            InsightsConfiguration = insightsConfiguration;
            Tags = tags;
        }
    }
}
