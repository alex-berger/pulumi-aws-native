// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Cloud9
{
    public static class GetEnvironmentEC2
    {
        /// <summary>
        /// Resource Type definition for AWS::Cloud9::EnvironmentEC2
        /// </summary>
        public static Task<GetEnvironmentEC2Result> InvokeAsync(GetEnvironmentEC2Args args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetEnvironmentEC2Result>("aws-native:cloud9:getEnvironmentEC2", args ?? new GetEnvironmentEC2Args(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::Cloud9::EnvironmentEC2
        /// </summary>
        public static Output<GetEnvironmentEC2Result> Invoke(GetEnvironmentEC2InvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetEnvironmentEC2Result>("aws-native:cloud9:getEnvironmentEC2", args ?? new GetEnvironmentEC2InvokeArgs(), options.WithDefaults());
    }


    public sealed class GetEnvironmentEC2Args : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetEnvironmentEC2Args()
        {
        }
    }

    public sealed class GetEnvironmentEC2InvokeArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetEnvironmentEC2InvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetEnvironmentEC2Result
    {
        public readonly string? Arn;
        public readonly string? Description;
        public readonly string? Id;
        public readonly string? Name;
        public readonly ImmutableArray<Outputs.EnvironmentEC2Tag> Tags;

        [OutputConstructor]
        private GetEnvironmentEC2Result(
            string? arn,

            string? description,

            string? id,

            string? name,

            ImmutableArray<Outputs.EnvironmentEC2Tag> tags)
        {
            Arn = arn;
            Description = description;
            Id = id;
            Name = name;
            Tags = tags;
        }
    }
}
