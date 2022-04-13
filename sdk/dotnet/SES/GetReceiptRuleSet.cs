// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SES
{
    public static class GetReceiptRuleSet
    {
        /// <summary>
        /// Resource Type definition for AWS::SES::ReceiptRuleSet
        /// </summary>
        public static Task<GetReceiptRuleSetResult> InvokeAsync(GetReceiptRuleSetArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetReceiptRuleSetResult>("aws-native:ses:getReceiptRuleSet", args ?? new GetReceiptRuleSetArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::SES::ReceiptRuleSet
        /// </summary>
        public static Output<GetReceiptRuleSetResult> Invoke(GetReceiptRuleSetInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetReceiptRuleSetResult>("aws-native:ses:getReceiptRuleSet", args ?? new GetReceiptRuleSetInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetReceiptRuleSetArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetReceiptRuleSetArgs()
        {
        }
    }

    public sealed class GetReceiptRuleSetInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetReceiptRuleSetInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetReceiptRuleSetResult
    {
        public readonly string? Id;

        [OutputConstructor]
        private GetReceiptRuleSetResult(string? id)
        {
            Id = id;
        }
    }
}
