// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Glue
{
    public static class GetDatabase
    {
        /// <summary>
        /// Resource Type definition for AWS::Glue::Database
        /// </summary>
        public static Task<GetDatabaseResult> InvokeAsync(GetDatabaseArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetDatabaseResult>("aws-native:glue:getDatabase", args ?? new GetDatabaseArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::Glue::Database
        /// </summary>
        public static Output<GetDatabaseResult> Invoke(GetDatabaseInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetDatabaseResult>("aws-native:glue:getDatabase", args ?? new GetDatabaseInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDatabaseArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetDatabaseArgs()
        {
        }
    }

    public sealed class GetDatabaseInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetDatabaseInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetDatabaseResult
    {
        public readonly Outputs.DatabaseInput? DatabaseInput;
        public readonly string? Id;

        [OutputConstructor]
        private GetDatabaseResult(
            Outputs.DatabaseInput? databaseInput,

            string? id)
        {
            DatabaseInput = databaseInput;
            Id = id;
        }
    }
}