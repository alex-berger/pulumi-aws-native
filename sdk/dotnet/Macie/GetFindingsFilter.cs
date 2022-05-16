// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Macie
{
    public static class GetFindingsFilter
    {
        /// <summary>
        /// Macie FindingsFilter resource schema.
        /// </summary>
        public static Task<GetFindingsFilterResult> InvokeAsync(GetFindingsFilterArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetFindingsFilterResult>("aws-native:macie:getFindingsFilter", args ?? new GetFindingsFilterArgs(), options.WithDefaults());

        /// <summary>
        /// Macie FindingsFilter resource schema.
        /// </summary>
        public static Output<GetFindingsFilterResult> Invoke(GetFindingsFilterInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetFindingsFilterResult>("aws-native:macie:getFindingsFilter", args ?? new GetFindingsFilterInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetFindingsFilterArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Findings filter ID.
        /// </summary>
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetFindingsFilterArgs()
        {
        }
    }

    public sealed class GetFindingsFilterInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Findings filter ID.
        /// </summary>
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetFindingsFilterInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetFindingsFilterResult
    {
        /// <summary>
        /// Findings filter action.
        /// </summary>
        public readonly Pulumi.AwsNative.Macie.FindingsFilterFindingFilterAction? Action;
        /// <summary>
        /// Findings filter ARN.
        /// </summary>
        public readonly string? Arn;
        /// <summary>
        /// Findings filter description
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// Findings filter criteria.
        /// </summary>
        public readonly Outputs.FindingsFilterFindingCriteria? FindingCriteria;
        /// <summary>
        /// Findings filters list.
        /// </summary>
        public readonly ImmutableArray<Outputs.FindingsFilterListItem> FindingsFilterListItems;
        /// <summary>
        /// Findings filter ID.
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// Findings filter name
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// Findings filter position.
        /// </summary>
        public readonly int? Position;

        [OutputConstructor]
        private GetFindingsFilterResult(
            Pulumi.AwsNative.Macie.FindingsFilterFindingFilterAction? action,

            string? arn,

            string? description,

            Outputs.FindingsFilterFindingCriteria? findingCriteria,

            ImmutableArray<Outputs.FindingsFilterListItem> findingsFilterListItems,

            string? id,

            string? name,

            int? position)
        {
            Action = action;
            Arn = arn;
            Description = description;
            FindingCriteria = findingCriteria;
            FindingsFilterListItems = findingsFilterListItems;
            Id = id;
            Name = name;
            Position = position;
        }
    }
}