// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Pinpoint.Outputs
{

    [OutputType]
    public sealed class CampaignEmailMessage
    {
        public readonly string? Body;
        public readonly string? FromAddress;
        public readonly string? HtmlBody;
        public readonly string? Title;

        [OutputConstructor]
        private CampaignEmailMessage(
            string? body,

            string? fromAddress,

            string? htmlBody,

            string? title)
        {
            Body = body;
            FromAddress = fromAddress;
            HtmlBody = htmlBody;
            Title = title;
        }
    }
}
