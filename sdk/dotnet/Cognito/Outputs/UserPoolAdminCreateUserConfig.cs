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
    public sealed class UserPoolAdminCreateUserConfig
    {
        public readonly bool? AllowAdminCreateUserOnly;
        public readonly Outputs.UserPoolInviteMessageTemplate? InviteMessageTemplate;
        public readonly int? UnusedAccountValidityDays;

        [OutputConstructor]
        private UserPoolAdminCreateUserConfig(
            bool? allowAdminCreateUserOnly,

            Outputs.UserPoolInviteMessageTemplate? inviteMessageTemplate,

            int? unusedAccountValidityDays)
        {
            AllowAdminCreateUserOnly = allowAdminCreateUserOnly;
            InviteMessageTemplate = inviteMessageTemplate;
            UnusedAccountValidityDays = unusedAccountValidityDays;
        }
    }
}
