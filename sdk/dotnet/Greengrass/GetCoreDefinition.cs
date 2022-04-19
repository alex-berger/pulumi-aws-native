// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Greengrass
{
    public static class GetCoreDefinition
    {
        /// <summary>
        /// Resource Type definition for AWS::Greengrass::CoreDefinition
        /// </summary>
        public static Task<GetCoreDefinitionResult> InvokeAsync(GetCoreDefinitionArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetCoreDefinitionResult>("aws-native:greengrass:getCoreDefinition", args ?? new GetCoreDefinitionArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::Greengrass::CoreDefinition
        /// </summary>
        public static Output<GetCoreDefinitionResult> Invoke(GetCoreDefinitionInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetCoreDefinitionResult>("aws-native:greengrass:getCoreDefinition", args ?? new GetCoreDefinitionInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetCoreDefinitionArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetCoreDefinitionArgs()
        {
        }
    }

    public sealed class GetCoreDefinitionInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetCoreDefinitionInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetCoreDefinitionResult
    {
        public readonly string? Arn;
        public readonly string? Id;
        public readonly string? LatestVersionArn;
        public readonly string? Name;
        public readonly object? Tags;

        [OutputConstructor]
        private GetCoreDefinitionResult(
            string? arn,

            string? id,

            string? latestVersionArn,

            string? name,

            object? tags)
        {
            Arn = arn;
            Id = id;
            LatestVersionArn = latestVersionArn;
            Name = name;
            Tags = tags;
        }
    }
}
