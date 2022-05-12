// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Cognito.Outputs
{

    [OutputType]
    public sealed class UserPoolRiskConfigurationAttachmentAccountTakeoverActionsType
    {
        public readonly Outputs.UserPoolRiskConfigurationAttachmentAccountTakeoverActionType? HighAction;
        public readonly Outputs.UserPoolRiskConfigurationAttachmentAccountTakeoverActionType? LowAction;
        public readonly Outputs.UserPoolRiskConfigurationAttachmentAccountTakeoverActionType? MediumAction;

        [OutputConstructor]
        private UserPoolRiskConfigurationAttachmentAccountTakeoverActionsType(
            Outputs.UserPoolRiskConfigurationAttachmentAccountTakeoverActionType? highAction,

            Outputs.UserPoolRiskConfigurationAttachmentAccountTakeoverActionType? lowAction,

            Outputs.UserPoolRiskConfigurationAttachmentAccountTakeoverActionType? mediumAction)
        {
            HighAction = highAction;
            LowAction = lowAction;
            MediumAction = mediumAction;
        }
    }
}
