// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Lex.Inputs
{

    /// <summary>
    /// Prompts that Amazon Lex sends to the user to confirm the completion of an intent.
    /// </summary>
    public sealed class BotIntentConfirmationSettingArgs : Pulumi.ResourceArgs
    {
        [Input("declinationResponse", required: true)]
        public Input<Inputs.BotResponseSpecificationArgs> DeclinationResponse { get; set; } = null!;

        [Input("isActive")]
        public Input<bool>? IsActive { get; set; }

        [Input("promptSpecification", required: true)]
        public Input<Inputs.BotPromptSpecificationArgs> PromptSpecification { get; set; } = null!;

        public BotIntentConfirmationSettingArgs()
        {
        }
    }
}
