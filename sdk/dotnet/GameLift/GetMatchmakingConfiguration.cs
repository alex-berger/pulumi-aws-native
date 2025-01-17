// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.GameLift
{
    public static class GetMatchmakingConfiguration
    {
        /// <summary>
        /// Resource Type definition for AWS::GameLift::MatchmakingConfiguration
        /// </summary>
        public static Task<GetMatchmakingConfigurationResult> InvokeAsync(GetMatchmakingConfigurationArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetMatchmakingConfigurationResult>("aws-native:gamelift:getMatchmakingConfiguration", args ?? new GetMatchmakingConfigurationArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::GameLift::MatchmakingConfiguration
        /// </summary>
        public static Output<GetMatchmakingConfigurationResult> Invoke(GetMatchmakingConfigurationInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetMatchmakingConfigurationResult>("aws-native:gamelift:getMatchmakingConfiguration", args ?? new GetMatchmakingConfigurationInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetMatchmakingConfigurationArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetMatchmakingConfigurationArgs()
        {
        }
    }

    public sealed class GetMatchmakingConfigurationInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetMatchmakingConfigurationInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetMatchmakingConfigurationResult
    {
        public readonly bool? AcceptanceRequired;
        public readonly int? AcceptanceTimeoutSeconds;
        public readonly int? AdditionalPlayerCount;
        public readonly string? Arn;
        public readonly string? BackfillMode;
        public readonly string? CustomEventData;
        public readonly string? Description;
        public readonly string? FlexMatchMode;
        public readonly ImmutableArray<Outputs.MatchmakingConfigurationGameProperty> GameProperties;
        public readonly string? GameSessionData;
        public readonly ImmutableArray<string> GameSessionQueueArns;
        public readonly string? Id;
        public readonly string? NotificationTarget;
        public readonly int? RequestTimeoutSeconds;
        public readonly string? RuleSetName;
        public readonly ImmutableArray<Outputs.MatchmakingConfigurationTag> Tags;

        [OutputConstructor]
        private GetMatchmakingConfigurationResult(
            bool? acceptanceRequired,

            int? acceptanceTimeoutSeconds,

            int? additionalPlayerCount,

            string? arn,

            string? backfillMode,

            string? customEventData,

            string? description,

            string? flexMatchMode,

            ImmutableArray<Outputs.MatchmakingConfigurationGameProperty> gameProperties,

            string? gameSessionData,

            ImmutableArray<string> gameSessionQueueArns,

            string? id,

            string? notificationTarget,

            int? requestTimeoutSeconds,

            string? ruleSetName,

            ImmutableArray<Outputs.MatchmakingConfigurationTag> tags)
        {
            AcceptanceRequired = acceptanceRequired;
            AcceptanceTimeoutSeconds = acceptanceTimeoutSeconds;
            AdditionalPlayerCount = additionalPlayerCount;
            Arn = arn;
            BackfillMode = backfillMode;
            CustomEventData = customEventData;
            Description = description;
            FlexMatchMode = flexMatchMode;
            GameProperties = gameProperties;
            GameSessionData = gameSessionData;
            GameSessionQueueArns = gameSessionQueueArns;
            Id = id;
            NotificationTarget = notificationTarget;
            RequestTimeoutSeconds = requestTimeoutSeconds;
            RuleSetName = ruleSetName;
            Tags = tags;
        }
    }
}
