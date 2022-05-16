// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Lex.Outputs
{

    /// <summary>
    /// A list of message groups that Amazon Lex uses to respond the user input.
    /// </summary>
    [OutputType]
    public sealed class BotResponseSpecification
    {
        /// <summary>
        /// Indicates whether the user can interrupt a speech prompt from the bot.
        /// </summary>
        public readonly bool? AllowInterrupt;
        public readonly ImmutableArray<Outputs.BotMessageGroup> MessageGroupsList;

        [OutputConstructor]
        private BotResponseSpecification(
            bool? allowInterrupt,

            ImmutableArray<Outputs.BotMessageGroup> messageGroupsList)
        {
            AllowInterrupt = allowInterrupt;
            MessageGroupsList = messageGroupsList;
        }
    }
}