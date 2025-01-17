// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight
{
    public static class GetDashboard
    {
        /// <summary>
        /// Definition of the AWS::QuickSight::Dashboard Resource Type.
        /// </summary>
        public static Task<GetDashboardResult> InvokeAsync(GetDashboardArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetDashboardResult>("aws-native:quicksight:getDashboard", args ?? new GetDashboardArgs(), options.WithDefaults());

        /// <summary>
        /// Definition of the AWS::QuickSight::Dashboard Resource Type.
        /// </summary>
        public static Output<GetDashboardResult> Invoke(GetDashboardInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetDashboardResult>("aws-native:quicksight:getDashboard", args ?? new GetDashboardInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDashboardArgs : Pulumi.InvokeArgs
    {
        [Input("awsAccountId", required: true)]
        public string AwsAccountId { get; set; } = null!;

        [Input("dashboardId", required: true)]
        public string DashboardId { get; set; } = null!;

        public GetDashboardArgs()
        {
        }
    }

    public sealed class GetDashboardInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("awsAccountId", required: true)]
        public Input<string> AwsAccountId { get; set; } = null!;

        [Input("dashboardId", required: true)]
        public Input<string> DashboardId { get; set; } = null!;

        public GetDashboardInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetDashboardResult
    {
        /// <summary>
        /// &lt;p&gt;The Amazon Resource Name (ARN) of the resource.&lt;/p&gt;
        /// </summary>
        public readonly string? Arn;
        /// <summary>
        /// &lt;p&gt;The last time that this dataset was published.&lt;/p&gt;
        /// </summary>
        public readonly string? LastPublishedTime;
        /// <summary>
        /// &lt;p&gt;The display name of the dashboard.&lt;/p&gt;
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// &lt;p&gt;A structure that contains the permissions of the dashboard. You can use this structure
        ///             for granting permissions by providing a list of IAM action information for each
        ///             principal ARN. &lt;/p&gt;
        /// 
        ///         &lt;p&gt;To specify no permissions, omit the permissions list.&lt;/p&gt;
        /// </summary>
        public readonly ImmutableArray<Outputs.DashboardResourcePermission> Permissions;
        /// <summary>
        /// &lt;p&gt;Contains a map of the key-value pairs for the resource tag or tags assigned to the
        ///             dashboard.&lt;/p&gt;
        /// </summary>
        public readonly ImmutableArray<Outputs.DashboardTag> Tags;

        [OutputConstructor]
        private GetDashboardResult(
            string? arn,

            string? lastPublishedTime,

            string? name,

            ImmutableArray<Outputs.DashboardResourcePermission> permissions,

            ImmutableArray<Outputs.DashboardTag> tags)
        {
            Arn = arn;
            LastPublishedTime = lastPublishedTime;
            Name = name;
            Permissions = permissions;
            Tags = tags;
        }
    }
}
