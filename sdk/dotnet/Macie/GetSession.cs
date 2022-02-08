// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Macie
{
    public static class GetSession
    {
        /// <summary>
        /// The AWS::Macie::Session resource specifies a new Amazon Macie session. A session is an object that represents the Amazon Macie service. A session is required for Amazon Macie to become operational.
        /// </summary>
        public static Task<GetSessionResult> InvokeAsync(GetSessionArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetSessionResult>("aws-native:macie:getSession", args ?? new GetSessionArgs(), options.WithDefaults());

        /// <summary>
        /// The AWS::Macie::Session resource specifies a new Amazon Macie session. A session is an object that represents the Amazon Macie service. A session is required for Amazon Macie to become operational.
        /// </summary>
        public static Output<GetSessionResult> Invoke(GetSessionInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetSessionResult>("aws-native:macie:getSession", args ?? new GetSessionInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSessionArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// AWS account ID of customer
        /// </summary>
        [Input("awsAccountId", required: true)]
        public string AwsAccountId { get; set; } = null!;

        public GetSessionArgs()
        {
        }
    }

    public sealed class GetSessionInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// AWS account ID of customer
        /// </summary>
        [Input("awsAccountId", required: true)]
        public Input<string> AwsAccountId { get; set; } = null!;

        public GetSessionInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetSessionResult
    {
        /// <summary>
        /// AWS account ID of customer
        /// </summary>
        public readonly string? AwsAccountId;
        /// <summary>
        /// A enumeration value that specifies how frequently finding updates are published.
        /// </summary>
        public readonly Pulumi.AwsNative.Macie.SessionFindingPublishingFrequency? FindingPublishingFrequency;
        /// <summary>
        /// Service role used by Macie
        /// </summary>
        public readonly string? ServiceRole;
        /// <summary>
        /// A enumeration value that specifies the status of the Macie Session.
        /// </summary>
        public readonly Pulumi.AwsNative.Macie.SessionStatus? Status;

        [OutputConstructor]
        private GetSessionResult(
            string? awsAccountId,

            Pulumi.AwsNative.Macie.SessionFindingPublishingFrequency? findingPublishingFrequency,

            string? serviceRole,

            Pulumi.AwsNative.Macie.SessionStatus? status)
        {
            AwsAccountId = awsAccountId;
            FindingPublishingFrequency = findingPublishingFrequency;
            ServiceRole = serviceRole;
            Status = status;
        }
    }
}