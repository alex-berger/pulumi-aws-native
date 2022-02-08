// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CloudFront
{
    public static class GetFunction
    {
        /// <summary>
        /// Resource Type definition for AWS::CloudFront::Function
        /// </summary>
        public static Task<GetFunctionResult> InvokeAsync(GetFunctionArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetFunctionResult>("aws-native:cloudfront:getFunction", args ?? new GetFunctionArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::CloudFront::Function
        /// </summary>
        public static Output<GetFunctionResult> Invoke(GetFunctionInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetFunctionResult>("aws-native:cloudfront:getFunction", args ?? new GetFunctionInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetFunctionArgs : Pulumi.InvokeArgs
    {
        [Input("functionARN", required: true)]
        public string FunctionARN { get; set; } = null!;

        public GetFunctionArgs()
        {
        }
    }

    public sealed class GetFunctionInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("functionARN", required: true)]
        public Input<string> FunctionARN { get; set; } = null!;

        public GetFunctionInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetFunctionResult
    {
        public readonly string? FunctionARN;
        public readonly Outputs.FunctionConfig? FunctionConfig;
        public readonly Outputs.FunctionMetadata? FunctionMetadata;
        public readonly string? Name;
        public readonly string? Stage;

        [OutputConstructor]
        private GetFunctionResult(
            string? functionARN,

            Outputs.FunctionConfig? functionConfig,

            Outputs.FunctionMetadata? functionMetadata,

            string? name,

            string? stage)
        {
            FunctionARN = functionARN;
            FunctionConfig = functionConfig;
            FunctionMetadata = functionMetadata;
            Name = name;
            Stage = stage;
        }
    }
}