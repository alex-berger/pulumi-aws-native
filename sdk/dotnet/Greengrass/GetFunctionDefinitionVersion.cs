// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Greengrass
{
    public static class GetFunctionDefinitionVersion
    {
        /// <summary>
        /// Resource Type definition for AWS::Greengrass::FunctionDefinitionVersion
        /// </summary>
        public static Task<GetFunctionDefinitionVersionResult> InvokeAsync(GetFunctionDefinitionVersionArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetFunctionDefinitionVersionResult>("aws-native:greengrass:getFunctionDefinitionVersion", args ?? new GetFunctionDefinitionVersionArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::Greengrass::FunctionDefinitionVersion
        /// </summary>
        public static Output<GetFunctionDefinitionVersionResult> Invoke(GetFunctionDefinitionVersionInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetFunctionDefinitionVersionResult>("aws-native:greengrass:getFunctionDefinitionVersion", args ?? new GetFunctionDefinitionVersionInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetFunctionDefinitionVersionArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetFunctionDefinitionVersionArgs()
        {
        }
    }

    public sealed class GetFunctionDefinitionVersionInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetFunctionDefinitionVersionInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetFunctionDefinitionVersionResult
    {
        public readonly string? Id;

        [OutputConstructor]
        private GetFunctionDefinitionVersionResult(string? id)
        {
            Id = id;
        }
    }
}
