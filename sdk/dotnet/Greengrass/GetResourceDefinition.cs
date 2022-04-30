// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Greengrass
{
    public static class GetResourceDefinition
    {
        /// <summary>
        /// Resource Type definition for AWS::Greengrass::ResourceDefinition
        /// </summary>
        public static Task<GetResourceDefinitionResult> InvokeAsync(GetResourceDefinitionArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetResourceDefinitionResult>("aws-native:greengrass:getResourceDefinition", args ?? new GetResourceDefinitionArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::Greengrass::ResourceDefinition
        /// </summary>
        public static Output<GetResourceDefinitionResult> Invoke(GetResourceDefinitionInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetResourceDefinitionResult>("aws-native:greengrass:getResourceDefinition", args ?? new GetResourceDefinitionInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetResourceDefinitionArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetResourceDefinitionArgs()
        {
        }
    }

    public sealed class GetResourceDefinitionInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetResourceDefinitionInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetResourceDefinitionResult
    {
        public readonly string? Arn;
        public readonly string? Id;
        public readonly string? LatestVersionArn;
        public readonly string? Name;
        public readonly object? Tags;

        [OutputConstructor]
        private GetResourceDefinitionResult(
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
