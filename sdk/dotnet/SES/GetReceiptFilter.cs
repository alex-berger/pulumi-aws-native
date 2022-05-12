// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SES
{
    public static class GetReceiptFilter
    {
        /// <summary>
        /// Resource Type definition for AWS::SES::ReceiptFilter
        /// </summary>
        public static Task<GetReceiptFilterResult> InvokeAsync(GetReceiptFilterArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetReceiptFilterResult>("aws-native:ses:getReceiptFilter", args ?? new GetReceiptFilterArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::SES::ReceiptFilter
        /// </summary>
        public static Output<GetReceiptFilterResult> Invoke(GetReceiptFilterInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetReceiptFilterResult>("aws-native:ses:getReceiptFilter", args ?? new GetReceiptFilterInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetReceiptFilterArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetReceiptFilterArgs()
        {
        }
    }

    public sealed class GetReceiptFilterInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetReceiptFilterInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetReceiptFilterResult
    {
        public readonly string? Id;

        [OutputConstructor]
        private GetReceiptFilterResult(string? id)
        {
            Id = id;
        }
    }
}
