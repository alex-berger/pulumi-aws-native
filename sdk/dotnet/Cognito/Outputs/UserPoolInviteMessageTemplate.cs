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
    public sealed class UserPoolInviteMessageTemplate
    {
        public readonly string? EmailMessage;
        public readonly string? EmailSubject;
        public readonly string? SMSMessage;

        [OutputConstructor]
        private UserPoolInviteMessageTemplate(
            string? emailMessage,

            string? emailSubject,

            string? sMSMessage)
        {
            EmailMessage = emailMessage;
            EmailSubject = emailSubject;
            SMSMessage = sMSMessage;
        }
    }
}