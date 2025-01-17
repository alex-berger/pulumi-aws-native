// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AppSync
{
    public static class GetResolver
    {
        /// <summary>
        /// Resource Type definition for AWS::AppSync::Resolver
        /// </summary>
        public static Task<GetResolverResult> InvokeAsync(GetResolverArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetResolverResult>("aws-native:appsync:getResolver", args ?? new GetResolverArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::AppSync::Resolver
        /// </summary>
        public static Output<GetResolverResult> Invoke(GetResolverInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetResolverResult>("aws-native:appsync:getResolver", args ?? new GetResolverInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetResolverArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetResolverArgs()
        {
        }
    }

    public sealed class GetResolverInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetResolverInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetResolverResult
    {
        public readonly Outputs.ResolverCachingConfig? CachingConfig;
        public readonly string? DataSourceName;
        public readonly string? Id;
        public readonly string? Kind;
        public readonly int? MaxBatchSize;
        public readonly Outputs.ResolverPipelineConfig? PipelineConfig;
        public readonly string? RequestMappingTemplate;
        public readonly string? RequestMappingTemplateS3Location;
        public readonly string? ResolverArn;
        public readonly string? ResponseMappingTemplate;
        public readonly string? ResponseMappingTemplateS3Location;
        public readonly Outputs.ResolverSyncConfig? SyncConfig;

        [OutputConstructor]
        private GetResolverResult(
            Outputs.ResolverCachingConfig? cachingConfig,

            string? dataSourceName,

            string? id,

            string? kind,

            int? maxBatchSize,

            Outputs.ResolverPipelineConfig? pipelineConfig,

            string? requestMappingTemplate,

            string? requestMappingTemplateS3Location,

            string? resolverArn,

            string? responseMappingTemplate,

            string? responseMappingTemplateS3Location,

            Outputs.ResolverSyncConfig? syncConfig)
        {
            CachingConfig = cachingConfig;
            DataSourceName = dataSourceName;
            Id = id;
            Kind = kind;
            MaxBatchSize = maxBatchSize;
            PipelineConfig = pipelineConfig;
            RequestMappingTemplate = requestMappingTemplate;
            RequestMappingTemplateS3Location = requestMappingTemplateS3Location;
            ResolverArn = resolverArn;
            ResponseMappingTemplate = responseMappingTemplate;
            ResponseMappingTemplateS3Location = responseMappingTemplateS3Location;
            SyncConfig = syncConfig;
        }
    }
}
