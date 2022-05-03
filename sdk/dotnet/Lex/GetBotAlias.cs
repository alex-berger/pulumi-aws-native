// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Lex
{
    public static class GetBotAlias
    {
        /// <summary>
        /// A Bot Alias enables you to change the version of a bot without updating applications that use the bot
        /// </summary>
        public static Task<GetBotAliasResult> InvokeAsync(GetBotAliasArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetBotAliasResult>("aws-native:lex:getBotAlias", args ?? new GetBotAliasArgs(), options.WithDefaults());

        /// <summary>
        /// A Bot Alias enables you to change the version of a bot without updating applications that use the bot
        /// </summary>
        public static Output<GetBotAliasResult> Invoke(GetBotAliasInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetBotAliasResult>("aws-native:lex:getBotAlias", args ?? new GetBotAliasInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetBotAliasArgs : Pulumi.InvokeArgs
    {
        [Input("botAliasId", required: true)]
        public string BotAliasId { get; set; } = null!;

        [Input("botId", required: true)]
        public string BotId { get; set; } = null!;

        public GetBotAliasArgs()
        {
        }
    }

    public sealed class GetBotAliasInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("botAliasId", required: true)]
        public Input<string> BotAliasId { get; set; } = null!;

        [Input("botId", required: true)]
        public Input<string> BotId { get; set; } = null!;

        public GetBotAliasInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetBotAliasResult
    {
        public readonly string? Arn;
        public readonly string? BotAliasId;
        public readonly ImmutableArray<Outputs.BotAliasLocaleSettingsItem> BotAliasLocaleSettings;
        public readonly string? BotAliasName;
        public readonly Pulumi.AwsNative.Lex.BotAliasStatus? BotAliasStatus;
        public readonly string? BotVersion;
        public readonly Outputs.BotAliasConversationLogSettings? ConversationLogSettings;
        public readonly string? Description;
        /// <summary>
        /// Determines whether Amazon Lex will use Amazon Comprehend to detect the sentiment of user utterances.
        /// </summary>
        public readonly Outputs.SentimentAnalysisSettingsProperties? SentimentAnalysisSettings;

        [OutputConstructor]
        private GetBotAliasResult(
            string? arn,

            string? botAliasId,

            ImmutableArray<Outputs.BotAliasLocaleSettingsItem> botAliasLocaleSettings,

            string? botAliasName,

            Pulumi.AwsNative.Lex.BotAliasStatus? botAliasStatus,

            string? botVersion,

            Outputs.BotAliasConversationLogSettings? conversationLogSettings,

            string? description,

            Outputs.SentimentAnalysisSettingsProperties? sentimentAnalysisSettings)
        {
            Arn = arn;
            BotAliasId = botAliasId;
            BotAliasLocaleSettings = botAliasLocaleSettings;
            BotAliasName = botAliasName;
            BotAliasStatus = botAliasStatus;
            BotVersion = botVersion;
            ConversationLogSettings = conversationLogSettings;
            Description = description;
            SentimentAnalysisSettings = sentimentAnalysisSettings;
        }
    }
}
