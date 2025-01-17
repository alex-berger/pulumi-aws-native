// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.DocDB
{
    public static class GetDBClusterParameterGroup
    {
        /// <summary>
        /// Resource Type definition for AWS::DocDB::DBClusterParameterGroup
        /// </summary>
        public static Task<GetDBClusterParameterGroupResult> InvokeAsync(GetDBClusterParameterGroupArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetDBClusterParameterGroupResult>("aws-native:docdb:getDBClusterParameterGroup", args ?? new GetDBClusterParameterGroupArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::DocDB::DBClusterParameterGroup
        /// </summary>
        public static Output<GetDBClusterParameterGroupResult> Invoke(GetDBClusterParameterGroupInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetDBClusterParameterGroupResult>("aws-native:docdb:getDBClusterParameterGroup", args ?? new GetDBClusterParameterGroupInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDBClusterParameterGroupArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetDBClusterParameterGroupArgs()
        {
        }
    }

    public sealed class GetDBClusterParameterGroupInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetDBClusterParameterGroupInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetDBClusterParameterGroupResult
    {
        public readonly string? Id;
        public readonly object? Parameters;
        public readonly ImmutableArray<Outputs.DBClusterParameterGroupTag> Tags;

        [OutputConstructor]
        private GetDBClusterParameterGroupResult(
            string? id,

            object? parameters,

            ImmutableArray<Outputs.DBClusterParameterGroupTag> tags)
        {
            Id = id;
            Parameters = parameters;
            Tags = tags;
        }
    }
}
