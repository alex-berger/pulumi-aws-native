// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.DAX
{
    public static class GetParameterGroup
    {
        /// <summary>
        /// Resource Type definition for AWS::DAX::ParameterGroup
        /// </summary>
        public static Task<GetParameterGroupResult> InvokeAsync(GetParameterGroupArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetParameterGroupResult>("aws-native:dax:getParameterGroup", args ?? new GetParameterGroupArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::DAX::ParameterGroup
        /// </summary>
        public static Output<GetParameterGroupResult> Invoke(GetParameterGroupInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetParameterGroupResult>("aws-native:dax:getParameterGroup", args ?? new GetParameterGroupInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetParameterGroupArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetParameterGroupArgs()
        {
        }
    }

    public sealed class GetParameterGroupInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetParameterGroupInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetParameterGroupResult
    {
        public readonly string? Description;
        public readonly string? Id;
        public readonly object? ParameterNameValues;

        [OutputConstructor]
        private GetParameterGroupResult(
            string? description,

            string? id,

            object? parameterNameValues)
        {
            Description = description;
            Id = id;
            ParameterNameValues = parameterNameValues;
        }
    }
}
