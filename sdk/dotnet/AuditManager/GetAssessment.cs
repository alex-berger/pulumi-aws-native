// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AuditManager
{
    public static class GetAssessment
    {
        /// <summary>
        /// An entity that defines the scope of audit evidence collected by AWS Audit Manager.
        /// </summary>
        public static Task<GetAssessmentResult> InvokeAsync(GetAssessmentArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAssessmentResult>("aws-native:auditmanager:getAssessment", args ?? new GetAssessmentArgs(), options.WithDefaults());

        /// <summary>
        /// An entity that defines the scope of audit evidence collected by AWS Audit Manager.
        /// </summary>
        public static Output<GetAssessmentResult> Invoke(GetAssessmentInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetAssessmentResult>("aws-native:auditmanager:getAssessment", args ?? new GetAssessmentInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAssessmentArgs : Pulumi.InvokeArgs
    {
        [Input("assessmentId", required: true)]
        public string AssessmentId { get; set; } = null!;

        public GetAssessmentArgs()
        {
        }
    }

    public sealed class GetAssessmentInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("assessmentId", required: true)]
        public Input<string> AssessmentId { get; set; } = null!;

        public GetAssessmentInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetAssessmentResult
    {
        public readonly string? Arn;
        public readonly string? AssessmentId;
        public readonly Outputs.AssessmentReportsDestination? AssessmentReportsDestination;
        public readonly double? CreationTime;
        /// <summary>
        /// The list of delegations.
        /// </summary>
        public readonly ImmutableArray<Outputs.AssessmentDelegation> Delegations;
        /// <summary>
        /// The list of roles for the specified assessment.
        /// </summary>
        public readonly ImmutableArray<Outputs.AssessmentRole> Roles;
        public readonly Outputs.AssessmentScope? Scope;
        public readonly Pulumi.AwsNative.AuditManager.AssessmentStatus? Status;
        /// <summary>
        /// The tags associated with the assessment.
        /// </summary>
        public readonly ImmutableArray<Outputs.AssessmentTag> Tags;

        [OutputConstructor]
        private GetAssessmentResult(
            string? arn,

            string? assessmentId,

            Outputs.AssessmentReportsDestination? assessmentReportsDestination,

            double? creationTime,

            ImmutableArray<Outputs.AssessmentDelegation> delegations,

            ImmutableArray<Outputs.AssessmentRole> roles,

            Outputs.AssessmentScope? scope,

            Pulumi.AwsNative.AuditManager.AssessmentStatus? status,

            ImmutableArray<Outputs.AssessmentTag> tags)
        {
            Arn = arn;
            AssessmentId = assessmentId;
            AssessmentReportsDestination = assessmentReportsDestination;
            CreationTime = creationTime;
            Delegations = delegations;
            Roles = roles;
            Scope = scope;
            Status = status;
            Tags = tags;
        }
    }
}
