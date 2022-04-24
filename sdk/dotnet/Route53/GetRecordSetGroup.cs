// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Route53
{
    public static class GetRecordSetGroup
    {
        /// <summary>
        /// Resource Type definition for AWS::Route53::RecordSetGroup
        /// </summary>
        public static Task<GetRecordSetGroupResult> InvokeAsync(GetRecordSetGroupArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetRecordSetGroupResult>("aws-native:route53:getRecordSetGroup", args ?? new GetRecordSetGroupArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::Route53::RecordSetGroup
        /// </summary>
        public static Output<GetRecordSetGroupResult> Invoke(GetRecordSetGroupInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetRecordSetGroupResult>("aws-native:route53:getRecordSetGroup", args ?? new GetRecordSetGroupInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetRecordSetGroupArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetRecordSetGroupArgs()
        {
        }
    }

    public sealed class GetRecordSetGroupInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetRecordSetGroupInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetRecordSetGroupResult
    {
        public readonly string? Comment;
        public readonly string? Id;
        public readonly ImmutableArray<Outputs.RecordSetGroupRecordSet> RecordSets;

        [OutputConstructor]
        private GetRecordSetGroupResult(
            string? comment,

            string? id,

            ImmutableArray<Outputs.RecordSetGroupRecordSet> recordSets)
        {
            Comment = comment;
            Id = id;
            RecordSets = recordSets;
        }
    }
}
