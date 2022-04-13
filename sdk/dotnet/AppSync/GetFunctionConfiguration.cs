// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AppSync
{
    public static class GetFunctionConfiguration
    {
        /// <summary>
        /// Resource Type definition for AWS::AppSync::FunctionConfiguration
        /// </summary>
        public static Task<GetFunctionConfigurationResult> InvokeAsync(GetFunctionConfigurationArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetFunctionConfigurationResult>("aws-native:appsync:getFunctionConfiguration", args ?? new GetFunctionConfigurationArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::AppSync::FunctionConfiguration
        /// </summary>
        public static Output<GetFunctionConfigurationResult> Invoke(GetFunctionConfigurationInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetFunctionConfigurationResult>("aws-native:appsync:getFunctionConfiguration", args ?? new GetFunctionConfigurationInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetFunctionConfigurationArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetFunctionConfigurationArgs()
        {
        }
    }

    public sealed class GetFunctionConfigurationInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetFunctionConfigurationInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetFunctionConfigurationResult
    {
        public readonly string? DataSourceName;
        public readonly string? Description;
        public readonly string? FunctionArn;
        public readonly string? FunctionId;
        public readonly string? FunctionVersion;
        public readonly string? Id;
        public readonly int? MaxBatchSize;
        public readonly string? Name;
        public readonly string? RequestMappingTemplate;
        public readonly string? RequestMappingTemplateS3Location;
        public readonly string? ResponseMappingTemplate;
        public readonly string? ResponseMappingTemplateS3Location;
        public readonly Outputs.FunctionConfigurationSyncConfig? SyncConfig;

        [OutputConstructor]
        private GetFunctionConfigurationResult(
            string? dataSourceName,

            string? description,

            string? functionArn,

            string? functionId,

            string? functionVersion,

            string? id,

            int? maxBatchSize,

            string? name,

            string? requestMappingTemplate,

            string? requestMappingTemplateS3Location,

            string? responseMappingTemplate,

            string? responseMappingTemplateS3Location,

            Outputs.FunctionConfigurationSyncConfig? syncConfig)
        {
            DataSourceName = dataSourceName;
            Description = description;
            FunctionArn = functionArn;
            FunctionId = functionId;
            FunctionVersion = functionVersion;
            Id = id;
            MaxBatchSize = maxBatchSize;
            Name = name;
            RequestMappingTemplate = requestMappingTemplate;
            RequestMappingTemplateS3Location = requestMappingTemplateS3Location;
            ResponseMappingTemplate = responseMappingTemplate;
            ResponseMappingTemplateS3Location = responseMappingTemplateS3Location;
            SyncConfig = syncConfig;
        }
    }
}
