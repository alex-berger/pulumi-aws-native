// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Inspector
{
    public static class GetAssessmentTarget
    {
        /// <summary>
        /// Resource Type definition for AWS::Inspector::AssessmentTarget
        /// </summary>
        public static Task<GetAssessmentTargetResult> InvokeAsync(GetAssessmentTargetArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAssessmentTargetResult>("aws-native:inspector:getAssessmentTarget", args ?? new GetAssessmentTargetArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::Inspector::AssessmentTarget
        /// </summary>
        public static Output<GetAssessmentTargetResult> Invoke(GetAssessmentTargetInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetAssessmentTargetResult>("aws-native:inspector:getAssessmentTarget", args ?? new GetAssessmentTargetInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAssessmentTargetArgs : Pulumi.InvokeArgs
    {
        [Input("arn", required: true)]
        public string Arn { get; set; } = null!;

        public GetAssessmentTargetArgs()
        {
        }
    }

    public sealed class GetAssessmentTargetInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("arn", required: true)]
        public Input<string> Arn { get; set; } = null!;

        public GetAssessmentTargetInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetAssessmentTargetResult
    {
        public readonly string? Arn;
        public readonly string? ResourceGroupArn;

        [OutputConstructor]
        private GetAssessmentTargetResult(
            string? arn,

            string? resourceGroupArn)
        {
            Arn = arn;
            ResourceGroupArn = resourceGroupArn;
        }
    }
}