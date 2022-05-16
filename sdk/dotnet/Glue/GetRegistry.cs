// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Glue
{
    public static class GetRegistry
    {
        /// <summary>
        /// This resource creates a Registry for authoring schemas as part of Glue Schema Registry.
        /// </summary>
        public static Task<GetRegistryResult> InvokeAsync(GetRegistryArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetRegistryResult>("aws-native:glue:getRegistry", args ?? new GetRegistryArgs(), options.WithDefaults());

        /// <summary>
        /// This resource creates a Registry for authoring schemas as part of Glue Schema Registry.
        /// </summary>
        public static Output<GetRegistryResult> Invoke(GetRegistryInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetRegistryResult>("aws-native:glue:getRegistry", args ?? new GetRegistryInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetRegistryArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Amazon Resource Name for the created Registry.
        /// </summary>
        [Input("arn", required: true)]
        public string Arn { get; set; } = null!;

        public GetRegistryArgs()
        {
        }
    }

    public sealed class GetRegistryInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Amazon Resource Name for the created Registry.
        /// </summary>
        [Input("arn", required: true)]
        public Input<string> Arn { get; set; } = null!;

        public GetRegistryInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetRegistryResult
    {
        /// <summary>
        /// Amazon Resource Name for the created Registry.
        /// </summary>
        public readonly string? Arn;
        /// <summary>
        /// A description of the registry. If description is not provided, there will not be any default value for this.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// List of tags to tag the Registry
        /// </summary>
        public readonly ImmutableArray<Outputs.RegistryTag> Tags;

        [OutputConstructor]
        private GetRegistryResult(
            string? arn,

            string? description,

            ImmutableArray<Outputs.RegistryTag> tags)
        {
            Arn = arn;
            Description = description;
            Tags = tags;
        }
    }
}
