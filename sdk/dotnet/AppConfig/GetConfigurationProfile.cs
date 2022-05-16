// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AppConfig
{
    public static class GetConfigurationProfile
    {
        /// <summary>
        /// Resource Type definition for AWS::AppConfig::ConfigurationProfile
        /// </summary>
        public static Task<GetConfigurationProfileResult> InvokeAsync(GetConfigurationProfileArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetConfigurationProfileResult>("aws-native:appconfig:getConfigurationProfile", args ?? new GetConfigurationProfileArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::AppConfig::ConfigurationProfile
        /// </summary>
        public static Output<GetConfigurationProfileResult> Invoke(GetConfigurationProfileInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetConfigurationProfileResult>("aws-native:appconfig:getConfigurationProfile", args ?? new GetConfigurationProfileInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetConfigurationProfileArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetConfigurationProfileArgs()
        {
        }
    }

    public sealed class GetConfigurationProfileInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetConfigurationProfileInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetConfigurationProfileResult
    {
        public readonly string? Description;
        public readonly string? Id;
        public readonly string? Name;
        public readonly string? RetrievalRoleArn;
        public readonly ImmutableArray<Outputs.ConfigurationProfileTags> Tags;
        public readonly ImmutableArray<Outputs.ConfigurationProfileValidators> Validators;

        [OutputConstructor]
        private GetConfigurationProfileResult(
            string? description,

            string? id,

            string? name,

            string? retrievalRoleArn,

            ImmutableArray<Outputs.ConfigurationProfileTags> tags,

            ImmutableArray<Outputs.ConfigurationProfileValidators> validators)
        {
            Description = description;
            Id = id;
            Name = name;
            RetrievalRoleArn = retrievalRoleArn;
            Tags = tags;
            Validators = validators;
        }
    }
}
