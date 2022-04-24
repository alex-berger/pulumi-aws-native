// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ASK
{
    public static class GetSkill
    {
        /// <summary>
        /// Resource Type definition for Alexa::ASK::Skill
        /// </summary>
        public static Task<GetSkillResult> InvokeAsync(GetSkillArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetSkillResult>("aws-native:ask:getSkill", args ?? new GetSkillArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for Alexa::ASK::Skill
        /// </summary>
        public static Output<GetSkillResult> Invoke(GetSkillInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetSkillResult>("aws-native:ask:getSkill", args ?? new GetSkillInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSkillArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetSkillArgs()
        {
        }
    }

    public sealed class GetSkillInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetSkillInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetSkillResult
    {
        public readonly Outputs.SkillAuthenticationConfiguration? AuthenticationConfiguration;
        public readonly string? Id;
        public readonly Outputs.SkillPackage? SkillPackage;

        [OutputConstructor]
        private GetSkillResult(
            Outputs.SkillAuthenticationConfiguration? authenticationConfiguration,

            string? id,

            Outputs.SkillPackage? skillPackage)
        {
            AuthenticationConfiguration = authenticationConfiguration;
            Id = id;
            SkillPackage = skillPackage;
        }
    }
}
