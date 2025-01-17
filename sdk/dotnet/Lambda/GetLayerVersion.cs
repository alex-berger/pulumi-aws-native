// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Lambda
{
    public static class GetLayerVersion
    {
        /// <summary>
        /// Resource Type definition for AWS::Lambda::LayerVersion
        /// </summary>
        public static Task<GetLayerVersionResult> InvokeAsync(GetLayerVersionArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetLayerVersionResult>("aws-native:lambda:getLayerVersion", args ?? new GetLayerVersionArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::Lambda::LayerVersion
        /// </summary>
        public static Output<GetLayerVersionResult> Invoke(GetLayerVersionInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetLayerVersionResult>("aws-native:lambda:getLayerVersion", args ?? new GetLayerVersionInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetLayerVersionArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetLayerVersionArgs()
        {
        }
    }

    public sealed class GetLayerVersionInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetLayerVersionInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetLayerVersionResult
    {
        public readonly string? Id;

        [OutputConstructor]
        private GetLayerVersionResult(string? id)
        {
            Id = id;
        }
    }
}
