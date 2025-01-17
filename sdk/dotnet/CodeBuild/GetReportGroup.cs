// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CodeBuild
{
    public static class GetReportGroup
    {
        /// <summary>
        /// Resource Type definition for AWS::CodeBuild::ReportGroup
        /// </summary>
        public static Task<GetReportGroupResult> InvokeAsync(GetReportGroupArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetReportGroupResult>("aws-native:codebuild:getReportGroup", args ?? new GetReportGroupArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::CodeBuild::ReportGroup
        /// </summary>
        public static Output<GetReportGroupResult> Invoke(GetReportGroupInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetReportGroupResult>("aws-native:codebuild:getReportGroup", args ?? new GetReportGroupInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetReportGroupArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetReportGroupArgs()
        {
        }
    }

    public sealed class GetReportGroupInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetReportGroupInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetReportGroupResult
    {
        public readonly string? Arn;
        public readonly bool? DeleteReports;
        public readonly Outputs.ReportGroupReportExportConfig? ExportConfig;
        public readonly string? Id;
        public readonly ImmutableArray<Outputs.ReportGroupTag> Tags;

        [OutputConstructor]
        private GetReportGroupResult(
            string? arn,

            bool? deleteReports,

            Outputs.ReportGroupReportExportConfig? exportConfig,

            string? id,

            ImmutableArray<Outputs.ReportGroupTag> tags)
        {
            Arn = arn;
            DeleteReports = deleteReports;
            ExportConfig = exportConfig;
            Id = id;
            Tags = tags;
        }
    }
}
