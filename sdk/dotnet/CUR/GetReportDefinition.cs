// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CUR
{
    public static class GetReportDefinition
    {
        /// <summary>
        /// The AWS::CUR::ReportDefinition resource creates a Cost &amp; Usage Report with user-defined settings. You can use this resource to define settings like time granularity (hourly, daily, monthly), file format (Parquet, CSV), and S3 bucket for delivery of these reports.
        /// </summary>
        public static Task<GetReportDefinitionResult> InvokeAsync(GetReportDefinitionArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetReportDefinitionResult>("aws-native:cur:getReportDefinition", args ?? new GetReportDefinitionArgs(), options.WithDefaults());

        /// <summary>
        /// The AWS::CUR::ReportDefinition resource creates a Cost &amp; Usage Report with user-defined settings. You can use this resource to define settings like time granularity (hourly, daily, monthly), file format (Parquet, CSV), and S3 bucket for delivery of these reports.
        /// </summary>
        public static Output<GetReportDefinitionResult> Invoke(GetReportDefinitionInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetReportDefinitionResult>("aws-native:cur:getReportDefinition", args ?? new GetReportDefinitionInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetReportDefinitionArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the report that you want to create. The name must be unique, is case sensitive, and can't include spaces.
        /// </summary>
        [Input("reportName", required: true)]
        public string ReportName { get; set; } = null!;

        public GetReportDefinitionArgs()
        {
        }
    }

    public sealed class GetReportDefinitionInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the report that you want to create. The name must be unique, is case sensitive, and can't include spaces.
        /// </summary>
        [Input("reportName", required: true)]
        public Input<string> ReportName { get; set; } = null!;

        public GetReportDefinitionInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetReportDefinitionResult
    {
        /// <summary>
        /// A list of manifests that you want Amazon Web Services to create for this report.
        /// </summary>
        public readonly ImmutableArray<Pulumi.AwsNative.CUR.ReportDefinitionAdditionalArtifactsItem> AdditionalArtifacts;
        /// <summary>
        /// The compression format that AWS uses for the report.
        /// </summary>
        public readonly Pulumi.AwsNative.CUR.ReportDefinitionCompression? Compression;
        /// <summary>
        /// The format that AWS saves the report in.
        /// </summary>
        public readonly Pulumi.AwsNative.CUR.ReportDefinitionFormat? Format;
        /// <summary>
        /// Whether you want Amazon Web Services to update your reports after they have been finalized if Amazon Web Services detects charges related to previous months. These charges can include refunds, credits, or support fees.
        /// </summary>
        public readonly bool? RefreshClosedReports;
        /// <summary>
        /// The S3 bucket where AWS delivers the report.
        /// </summary>
        public readonly string? S3Bucket;
        /// <summary>
        /// The prefix that AWS adds to the report name when AWS delivers the report. Your prefix can't include spaces.
        /// </summary>
        public readonly string? S3Prefix;
        /// <summary>
        /// The region of the S3 bucket that AWS delivers the report into.
        /// </summary>
        public readonly string? S3Region;

        [OutputConstructor]
        private GetReportDefinitionResult(
            ImmutableArray<Pulumi.AwsNative.CUR.ReportDefinitionAdditionalArtifactsItem> additionalArtifacts,

            Pulumi.AwsNative.CUR.ReportDefinitionCompression? compression,

            Pulumi.AwsNative.CUR.ReportDefinitionFormat? format,

            bool? refreshClosedReports,

            string? s3Bucket,

            string? s3Prefix,

            string? s3Region)
        {
            AdditionalArtifacts = additionalArtifacts;
            Compression = compression;
            Format = format;
            RefreshClosedReports = refreshClosedReports;
            S3Bucket = s3Bucket;
            S3Prefix = s3Prefix;
            S3Region = s3Region;
        }
    }
}
