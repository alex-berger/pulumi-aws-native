// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.OpsWorks
{
    public static class GetApp
    {
        /// <summary>
        /// Resource Type definition for AWS::OpsWorks::App
        /// </summary>
        public static Task<GetAppResult> InvokeAsync(GetAppArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAppResult>("aws-native:opsworks:getApp", args ?? new GetAppArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::OpsWorks::App
        /// </summary>
        public static Output<GetAppResult> Invoke(GetAppInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetAppResult>("aws-native:opsworks:getApp", args ?? new GetAppInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAppArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetAppArgs()
        {
        }
    }

    public sealed class GetAppInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetAppInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetAppResult
    {
        public readonly Outputs.AppSource? AppSource;
        public readonly object? Attributes;
        public readonly ImmutableArray<Outputs.AppDataSource> DataSources;
        public readonly string? Description;
        public readonly ImmutableArray<string> Domains;
        public readonly bool? EnableSsl;
        public readonly ImmutableArray<Outputs.AppEnvironmentVariable> Environment;
        public readonly string? Id;
        public readonly string? Name;
        public readonly Outputs.AppSslConfiguration? SslConfiguration;
        public readonly string? Type;

        [OutputConstructor]
        private GetAppResult(
            Outputs.AppSource? appSource,

            object? attributes,

            ImmutableArray<Outputs.AppDataSource> dataSources,

            string? description,

            ImmutableArray<string> domains,

            bool? enableSsl,

            ImmutableArray<Outputs.AppEnvironmentVariable> environment,

            string? id,

            string? name,

            Outputs.AppSslConfiguration? sslConfiguration,

            string? type)
        {
            AppSource = appSource;
            Attributes = attributes;
            DataSources = dataSources;
            Description = description;
            Domains = domains;
            EnableSsl = enableSsl;
            Environment = environment;
            Id = id;
            Name = name;
            SslConfiguration = sslConfiguration;
            Type = type;
        }
    }
}