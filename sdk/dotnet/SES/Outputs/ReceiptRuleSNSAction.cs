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
    public sealed class ReceiptRuleSNSAction
    {
        public readonly string? Encoding;
        public readonly string? TopicArn;

        [OutputConstructor]
        private ReceiptRuleSNSAction(
            string? encoding,

            string? topicArn)
        {
            Encoding = encoding;
            TopicArn = topicArn;
        }
    }
}
