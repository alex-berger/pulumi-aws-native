// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SSM
{
    public static class GetAssociation
    {
        /// <summary>
        /// The AWS::SSM::Association resource associates an SSM document in AWS Systems Manager with EC2 instances that contain a configuration agent to process the document.
        /// </summary>
        public static Task<GetAssociationResult> InvokeAsync(GetAssociationArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAssociationResult>("aws-native:ssm:getAssociation", args ?? new GetAssociationArgs(), options.WithDefaults());

        /// <summary>
        /// The AWS::SSM::Association resource associates an SSM document in AWS Systems Manager with EC2 instances that contain a configuration agent to process the document.
        /// </summary>
        public static Output<GetAssociationResult> Invoke(GetAssociationInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetAssociationResult>("aws-native:ssm:getAssociation", args ?? new GetAssociationInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAssociationArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Unique identifier of the association.
        /// </summary>
        [Input("associationId", required: true)]
        public string AssociationId { get; set; } = null!;

        public GetAssociationArgs()
        {
        }
    }

    public sealed class GetAssociationInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Unique identifier of the association.
        /// </summary>
        [Input("associationId", required: true)]
        public Input<string> AssociationId { get; set; } = null!;

        public GetAssociationInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetAssociationResult
    {
        public readonly bool? ApplyOnlyAtCronInterval;
        /// <summary>
        /// Unique identifier of the association.
        /// </summary>
        public readonly string? AssociationId;
        /// <summary>
        /// The name of the association.
        /// </summary>
        public readonly string? AssociationName;
        public readonly string? AutomationTargetParameterName;
        public readonly ImmutableArray<string> CalendarNames;
        public readonly Pulumi.AwsNative.SSM.AssociationComplianceSeverity? ComplianceSeverity;
        /// <summary>
        /// The version of the SSM document to associate with the target.
        /// </summary>
        public readonly string? DocumentVersion;
        /// <summary>
        /// The ID of the instance that the SSM document is associated with.
        /// </summary>
        public readonly string? InstanceId;
        public readonly string? MaxConcurrency;
        public readonly string? MaxErrors;
        /// <summary>
        /// The name of the SSM document.
        /// </summary>
        public readonly string? Name;
        public readonly Outputs.AssociationInstanceAssociationOutputLocation? OutputLocation;
        /// <summary>
        /// Parameter values that the SSM document uses at runtime.
        /// </summary>
        public readonly object? Parameters;
        /// <summary>
        /// A Cron or Rate expression that specifies when the association is applied to the target.
        /// </summary>
        public readonly string? ScheduleExpression;
        public readonly int? ScheduleOffset;
        public readonly Pulumi.AwsNative.SSM.AssociationSyncCompliance? SyncCompliance;
        /// <summary>
        /// The targets that the SSM document sends commands to.
        /// </summary>
        public readonly ImmutableArray<Outputs.AssociationTarget> Targets;
        public readonly int? WaitForSuccessTimeoutSeconds;

        [OutputConstructor]
        private GetAssociationResult(
            bool? applyOnlyAtCronInterval,

            string? associationId,

            string? associationName,

            string? automationTargetParameterName,

            ImmutableArray<string> calendarNames,

            Pulumi.AwsNative.SSM.AssociationComplianceSeverity? complianceSeverity,

            string? documentVersion,

            string? instanceId,

            string? maxConcurrency,

            string? maxErrors,

            string? name,

            Outputs.AssociationInstanceAssociationOutputLocation? outputLocation,

            object? parameters,

            string? scheduleExpression,

            int? scheduleOffset,

            Pulumi.AwsNative.SSM.AssociationSyncCompliance? syncCompliance,

            ImmutableArray<Outputs.AssociationTarget> targets,

            int? waitForSuccessTimeoutSeconds)
        {
            ApplyOnlyAtCronInterval = applyOnlyAtCronInterval;
            AssociationId = associationId;
            AssociationName = associationName;
            AutomationTargetParameterName = automationTargetParameterName;
            CalendarNames = calendarNames;
            ComplianceSeverity = complianceSeverity;
            DocumentVersion = documentVersion;
            InstanceId = instanceId;
            MaxConcurrency = maxConcurrency;
            MaxErrors = maxErrors;
            Name = name;
            OutputLocation = outputLocation;
            Parameters = parameters;
            ScheduleExpression = scheduleExpression;
            ScheduleOffset = scheduleOffset;
            SyncCompliance = syncCompliance;
            Targets = targets;
            WaitForSuccessTimeoutSeconds = waitForSuccessTimeoutSeconds;
        }
    }
}
