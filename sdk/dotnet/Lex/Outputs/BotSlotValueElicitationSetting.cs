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
    /// Settings that you can use for eliciting a slot value.
    /// </summary>
    [OutputType]
    public sealed class BotSlotValueElicitationSetting
    {
        /// <summary>
        /// A list of default values for a slot.
        /// </summary>
        public readonly Outputs.BotSlotDefaultValueSpecification? DefaultValueSpecification;
        /// <summary>
        /// The prompt that Amazon Lex uses to elicit the slot value from the user.
        /// </summary>
        public readonly Outputs.BotPromptSpecification? PromptSpecification;
        /// <summary>
        /// If you know a specific pattern that users might respond to an Amazon Lex request for a slot value, you can provide those utterances to improve accuracy.
        /// </summary>
        public readonly ImmutableArray<Outputs.BotSampleUtterance> SampleUtterances;
        /// <summary>
        /// Specifies whether the slot is required or optional.
        /// </summary>
        public readonly Pulumi.AwsNative.Lex.BotSlotConstraint SlotConstraint;
        /// <summary>
        /// Specifies the prompts that Amazon Lex uses while a bot is waiting for customer input.
        /// </summary>
        public readonly Outputs.BotWaitAndContinueSpecification? WaitAndContinueSpecification;

        [OutputConstructor]
        private BotSlotValueElicitationSetting(
            Outputs.BotSlotDefaultValueSpecification? defaultValueSpecification,

            Outputs.BotPromptSpecification? promptSpecification,

            ImmutableArray<Outputs.BotSampleUtterance> sampleUtterances,

            Pulumi.AwsNative.Lex.BotSlotConstraint slotConstraint,

            Outputs.BotWaitAndContinueSpecification? waitAndContinueSpecification)
        {
            DefaultValueSpecification = defaultValueSpecification;
            PromptSpecification = promptSpecification;
            SampleUtterances = sampleUtterances;
            SlotConstraint = slotConstraint;
            WaitAndContinueSpecification = waitAndContinueSpecification;
        }
    }
}
