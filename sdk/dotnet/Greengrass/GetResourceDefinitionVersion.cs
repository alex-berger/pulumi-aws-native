// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Greengrass
{
    public static class GetResourceDefinitionVersion
    {
        /// <summary>
        /// Resource Type definition for AWS::Greengrass::ResourceDefinitionVersion
        /// </summary>
        public static Task<GetResourceDefinitionVersionResult> InvokeAsync(GetResourceDefinitionVersionArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetResourceDefinitionVersionResult>("aws-native:greengrass:getResourceDefinitionVersion", args ?? new GetResourceDefinitionVersionArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::Greengrass::ResourceDefinitionVersion
        /// </summary>
        public static Output<GetResourceDefinitionVersionResult> Invoke(GetResourceDefinitionVersionInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetResourceDefinitionVersionResult>("aws-native:greengrass:getResourceDefinitionVersion", args ?? new GetResourceDefinitionVersionInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetResourceDefinitionVersionArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetResourceDefinitionVersionArgs()
        {
        }
    }

    public sealed class GetResourceDefinitionVersionInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetResourceDefinitionVersionInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetResourceDefinitionVersionResult
    {
        public readonly string? Id;

        [OutputConstructor]
        private GetResourceDefinitionVersionResult(string? id)
        {
            Id = id;
        }
    }
}
