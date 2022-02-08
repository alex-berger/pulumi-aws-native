// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.NimbleStudio
{
    public static class GetStudioComponent
    {
        /// <summary>
        /// Represents a studio component which connects a non-Nimble Studio resource in your account to your studio
        /// </summary>
        public static Task<GetStudioComponentResult> InvokeAsync(GetStudioComponentArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetStudioComponentResult>("aws-native:nimblestudio:getStudioComponent", args ?? new GetStudioComponentArgs(), options.WithDefaults());

        /// <summary>
        /// Represents a studio component which connects a non-Nimble Studio resource in your account to your studio
        /// </summary>
        public static Output<GetStudioComponentResult> Invoke(GetStudioComponentInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetStudioComponentResult>("aws-native:nimblestudio:getStudioComponent", args ?? new GetStudioComponentInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetStudioComponentArgs : Pulumi.InvokeArgs
    {
        [Input("studioComponentId", required: true)]
        public string StudioComponentId { get; set; } = null!;

        /// <summary>
        /// &lt;p&gt;The studioId. &lt;/p&gt;
        /// </summary>
        [Input("studioId", required: true)]
        public string StudioId { get; set; } = null!;

        public GetStudioComponentArgs()
        {
        }
    }

    public sealed class GetStudioComponentInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("studioComponentId", required: true)]
        public Input<string> StudioComponentId { get; set; } = null!;

        /// <summary>
        /// &lt;p&gt;The studioId. &lt;/p&gt;
        /// </summary>
        [Input("studioId", required: true)]
        public Input<string> StudioId { get; set; } = null!;

        public GetStudioComponentInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetStudioComponentResult
    {
        public readonly Outputs.StudioComponentConfiguration? Configuration;
        /// <summary>
        /// &lt;p&gt;The description.&lt;/p&gt;
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// &lt;p&gt;The EC2 security groups that control access to the studio component.&lt;/p&gt;
        /// </summary>
        public readonly ImmutableArray<string> Ec2SecurityGroupIds;
        /// <summary>
        /// &lt;p&gt;Initialization scripts for studio components.&lt;/p&gt;
        /// </summary>
        public readonly ImmutableArray<Outputs.StudioComponentInitializationScript> InitializationScripts;
        /// <summary>
        /// &lt;p&gt;The name for the studio component.&lt;/p&gt;
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// &lt;p&gt;Parameters for the studio component scripts.&lt;/p&gt;
        /// </summary>
        public readonly ImmutableArray<Outputs.StudioComponentScriptParameterKeyValue> ScriptParameters;
        public readonly string? StudioComponentId;
        public readonly Pulumi.AwsNative.NimbleStudio.StudioComponentType? Type;

        [OutputConstructor]
        private GetStudioComponentResult(
            Outputs.StudioComponentConfiguration? configuration,

            string? description,

            ImmutableArray<string> ec2SecurityGroupIds,

            ImmutableArray<Outputs.StudioComponentInitializationScript> initializationScripts,

            string? name,

            ImmutableArray<Outputs.StudioComponentScriptParameterKeyValue> scriptParameters,

            string? studioComponentId,

            Pulumi.AwsNative.NimbleStudio.StudioComponentType? type)
        {
            Configuration = configuration;
            Description = description;
            Ec2SecurityGroupIds = ec2SecurityGroupIds;
            InitializationScripts = initializationScripts;
            Name = name;
            ScriptParameters = scriptParameters;
            StudioComponentId = studioComponentId;
            Type = type;
        }
    }
}