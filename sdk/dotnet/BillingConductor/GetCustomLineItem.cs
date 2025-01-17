// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.BillingConductor
{
    public static class GetCustomLineItem
    {
        /// <summary>
        /// A custom line item is an one time charge that is applied to a specific billing group's bill.
        /// </summary>
        public static Task<GetCustomLineItemResult> InvokeAsync(GetCustomLineItemArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetCustomLineItemResult>("aws-native:billingconductor:getCustomLineItem", args ?? new GetCustomLineItemArgs(), options.WithDefaults());

        /// <summary>
        /// A custom line item is an one time charge that is applied to a specific billing group's bill.
        /// </summary>
        public static Output<GetCustomLineItemResult> Invoke(GetCustomLineItemInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetCustomLineItemResult>("aws-native:billingconductor:getCustomLineItem", args ?? new GetCustomLineItemInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetCustomLineItemArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// ARN
        /// </summary>
        [Input("arn", required: true)]
        public string Arn { get; set; } = null!;

        public GetCustomLineItemArgs()
        {
        }
    }

    public sealed class GetCustomLineItemInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// ARN
        /// </summary>
        [Input("arn", required: true)]
        public Input<string> Arn { get; set; } = null!;

        public GetCustomLineItemInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetCustomLineItemResult
    {
        /// <summary>
        /// ARN
        /// </summary>
        public readonly string? Arn;
        /// <summary>
        /// Number of source values associated to this custom line item
        /// </summary>
        public readonly int? AssociationSize;
        public readonly Outputs.CustomLineItemBillingPeriodRange? BillingPeriodRange;
        /// <summary>
        /// Creation timestamp in UNIX epoch time format
        /// </summary>
        public readonly int? CreationTime;
        public readonly Pulumi.AwsNative.BillingConductor.CustomLineItemCurrencyCode? CurrencyCode;
        public readonly Outputs.CustomLineItemChargeDetails? CustomLineItemChargeDetails;
        public readonly string? Description;
        /// <summary>
        /// Latest modified timestamp in UNIX epoch time format
        /// </summary>
        public readonly int? LastModifiedTime;
        public readonly string? Name;
        public readonly string? ProductCode;
        public readonly ImmutableArray<Outputs.CustomLineItemTag> Tags;

        [OutputConstructor]
        private GetCustomLineItemResult(
            string? arn,

            int? associationSize,

            Outputs.CustomLineItemBillingPeriodRange? billingPeriodRange,

            int? creationTime,

            Pulumi.AwsNative.BillingConductor.CustomLineItemCurrencyCode? currencyCode,

            Outputs.CustomLineItemChargeDetails? customLineItemChargeDetails,

            string? description,

            int? lastModifiedTime,

            string? name,

            string? productCode,

            ImmutableArray<Outputs.CustomLineItemTag> tags)
        {
            Arn = arn;
            AssociationSize = associationSize;
            BillingPeriodRange = billingPeriodRange;
            CreationTime = creationTime;
            CurrencyCode = currencyCode;
            CustomLineItemChargeDetails = customLineItemChargeDetails;
            Description = description;
            LastModifiedTime = lastModifiedTime;
            Name = name;
            ProductCode = productCode;
            Tags = tags;
        }
    }
}
