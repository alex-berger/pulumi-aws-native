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
    public sealed class UserPoolRiskConfigurationAttachmentCompromisedCredentialsRiskConfigurationType
    {
        public readonly Outputs.UserPoolRiskConfigurationAttachmentCompromisedCredentialsActionsType Actions;
        public readonly ImmutableArray<string> EventFilter;

        [OutputConstructor]
        private UserPoolRiskConfigurationAttachmentCompromisedCredentialsRiskConfigurationType(
            Outputs.UserPoolRiskConfigurationAttachmentCompromisedCredentialsActionsType actions,

            ImmutableArray<string> eventFilter)
        {
            Actions = actions;
            EventFilter = eventFilter;
        }
    }
}