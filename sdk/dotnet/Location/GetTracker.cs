// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Location
{
    public static class GetTracker
    {
        /// <summary>
        /// Definition of AWS::Location::Tracker Resource Type
        /// </summary>
        public static Task<GetTrackerResult> InvokeAsync(GetTrackerArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetTrackerResult>("aws-native:location:getTracker", args ?? new GetTrackerArgs(), options.WithDefaults());

        /// <summary>
        /// Definition of AWS::Location::Tracker Resource Type
        /// </summary>
        public static Output<GetTrackerResult> Invoke(GetTrackerInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetTrackerResult>("aws-native:location:getTracker", args ?? new GetTrackerInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetTrackerArgs : Pulumi.InvokeArgs
    {
        [Input("trackerName", required: true)]
        public string TrackerName { get; set; } = null!;

        public GetTrackerArgs()
        {
        }
    }

    public sealed class GetTrackerInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("trackerName", required: true)]
        public Input<string> TrackerName { get; set; } = null!;

        public GetTrackerInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetTrackerResult
    {
        public readonly string? Arn;
        public readonly string? CreateTime;
        public readonly string? TrackerArn;
        public readonly string? UpdateTime;

        [OutputConstructor]
        private GetTrackerResult(
            string? arn,

            string? createTime,

            string? trackerArn,

            string? updateTime)
        {
            Arn = arn;
            CreateTime = createTime;
            TrackerArn = trackerArn;
            UpdateTime = updateTime;
        }
    }
}
