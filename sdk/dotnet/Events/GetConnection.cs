// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Events
{
    public static class GetConnection
    {
        /// <summary>
        /// Resource Type definition for AWS::Events::Connection.
        /// </summary>
        public static Task<GetConnectionResult> InvokeAsync(GetConnectionArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetConnectionResult>("aws-native:events:getConnection", args ?? new GetConnectionArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::Events::Connection.
        /// </summary>
        public static Output<GetConnectionResult> Invoke(GetConnectionInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetConnectionResult>("aws-native:events:getConnection", args ?? new GetConnectionInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetConnectionArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the connection.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetConnectionArgs()
        {
        }
    }

    public sealed class GetConnectionInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the connection.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public GetConnectionInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetConnectionResult
    {
        /// <summary>
        /// The arn of the connection resource.
        /// </summary>
        public readonly string? Arn;
        public readonly Outputs.AuthParametersProperties? AuthParameters;
        public readonly Pulumi.AwsNative.Events.ConnectionAuthorizationType? AuthorizationType;
        /// <summary>
        /// Description of the connection.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// The arn of the secrets manager secret created in the customer account.
        /// </summary>
        public readonly string? SecretArn;

        [OutputConstructor]
        private GetConnectionResult(
            string? arn,

            Outputs.AuthParametersProperties? authParameters,

            Pulumi.AwsNative.Events.ConnectionAuthorizationType? authorizationType,

            string? description,

            string? secretArn)
        {
            Arn = arn;
            AuthParameters = authParameters;
            AuthorizationType = authorizationType;
            Description = description;
            SecretArn = secretArn;
        }
    }
}