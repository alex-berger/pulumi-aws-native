// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SES.Outputs
{

    [OutputType]
    public sealed class ReceiptRuleBounceAction
    {
        public readonly string Message;
        public readonly string Sender;
        public readonly string SmtpReplyCode;
        public readonly string? StatusCode;
        public readonly string? TopicArn;

        [OutputConstructor]
        private ReceiptRuleBounceAction(
            string message,

            string sender,

            string smtpReplyCode,

            string? statusCode,

            string? topicArn)
        {
            Message = message;
            Sender = sender;
            SmtpReplyCode = smtpReplyCode;
            StatusCode = statusCode;
            TopicArn = topicArn;
        }
    }
}
