// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Lex
{
    /// <summary>
    /// Amazon Lex conversational bot performing automated tasks such as ordering a pizza, booking a hotel, and so on.
    /// </summary>
    [AwsNativeResourceType("aws-native:lex:Bot")]
    public partial class Bot : Pulumi.CustomResource
    {
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// Specifies whether to build the bot locales after bot creation completes.
        /// </summary>
        [Output("autoBuildBotLocales")]
        public Output<bool?> AutoBuildBotLocales { get; private set; } = null!;

        [Output("botFileS3Location")]
        public Output<Outputs.BotS3Location?> BotFileS3Location { get; private set; } = null!;

        /// <summary>
        /// List of bot locales
        /// </summary>
        [Output("botLocales")]
        public Output<ImmutableArray<Outputs.BotLocale>> BotLocales { get; private set; } = null!;

        /// <summary>
        /// A list of tags to add to the bot, which can only be added at bot creation.
        /// </summary>
        [Output("botTags")]
        public Output<ImmutableArray<Outputs.BotTag>> BotTags { get; private set; } = null!;

        /// <summary>
        /// Data privacy setting of the Bot.
        /// </summary>
        [Output("dataPrivacy")]
        public Output<Outputs.DataPrivacyProperties> DataPrivacy { get; private set; } = null!;

        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// IdleSessionTTLInSeconds of the resource
        /// </summary>
        [Output("idleSessionTTLInSeconds")]
        public Output<int> IdleSessionTTLInSeconds { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("roleArn")]
        public Output<string> RoleArn { get; private set; } = null!;

        [Output("testBotAliasSettings")]
        public Output<Outputs.BotTestBotAliasSettings?> TestBotAliasSettings { get; private set; } = null!;

        /// <summary>
        /// A list of tags to add to the test alias for a bot, , which can only be added at bot/bot alias creation.
        /// </summary>
        [Output("testBotAliasTags")]
        public Output<ImmutableArray<Outputs.BotTag>> TestBotAliasTags { get; private set; } = null!;


        /// <summary>
        /// Create a Bot resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Bot(string name, BotArgs args, CustomResourceOptions? options = null)
            : base("aws-native:lex:Bot", name, args ?? new BotArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Bot(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:lex:Bot", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Bot resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Bot Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Bot(name, id, options);
        }
    }

    public sealed class BotArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies whether to build the bot locales after bot creation completes.
        /// </summary>
        [Input("autoBuildBotLocales")]
        public Input<bool>? AutoBuildBotLocales { get; set; }

        [Input("botFileS3Location")]
        public Input<Inputs.BotS3LocationArgs>? BotFileS3Location { get; set; }

        [Input("botLocales")]
        private InputList<Inputs.BotLocaleArgs>? _botLocales;

        /// <summary>
        /// List of bot locales
        /// </summary>
        public InputList<Inputs.BotLocaleArgs> BotLocales
        {
            get => _botLocales ?? (_botLocales = new InputList<Inputs.BotLocaleArgs>());
            set => _botLocales = value;
        }

        [Input("botTags")]
        private InputList<Inputs.BotTagArgs>? _botTags;

        /// <summary>
        /// A list of tags to add to the bot, which can only be added at bot creation.
        /// </summary>
        public InputList<Inputs.BotTagArgs> BotTags
        {
            get => _botTags ?? (_botTags = new InputList<Inputs.BotTagArgs>());
            set => _botTags = value;
        }

        /// <summary>
        /// Data privacy setting of the Bot.
        /// </summary>
        [Input("dataPrivacy", required: true)]
        public Input<Inputs.DataPrivacyPropertiesArgs> DataPrivacy { get; set; } = null!;

        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// IdleSessionTTLInSeconds of the resource
        /// </summary>
        [Input("idleSessionTTLInSeconds", required: true)]
        public Input<int> IdleSessionTTLInSeconds { get; set; } = null!;

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("roleArn", required: true)]
        public Input<string> RoleArn { get; set; } = null!;

        [Input("testBotAliasSettings")]
        public Input<Inputs.BotTestBotAliasSettingsArgs>? TestBotAliasSettings { get; set; }

        [Input("testBotAliasTags")]
        private InputList<Inputs.BotTagArgs>? _testBotAliasTags;

        /// <summary>
        /// A list of tags to add to the test alias for a bot, , which can only be added at bot/bot alias creation.
        /// </summary>
        public InputList<Inputs.BotTagArgs> TestBotAliasTags
        {
            get => _testBotAliasTags ?? (_testBotAliasTags = new InputList<Inputs.BotTagArgs>());
            set => _testBotAliasTags = value;
        }

        public BotArgs()
        {
        }
    }
}
