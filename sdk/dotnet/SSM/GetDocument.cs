// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SSM
{
    public static class GetDocument
    {
        /// <summary>
        /// The AWS::SSM::Document resource is an SSM document in AWS Systems Manager that defines the actions that Systems Manager performs, which can be used to set up and run commands on your instances.
        /// </summary>
        public static Task<GetDocumentResult> InvokeAsync(GetDocumentArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetDocumentResult>("aws-native:ssm:getDocument", args ?? new GetDocumentArgs(), options.WithDefaults());

        /// <summary>
        /// The AWS::SSM::Document resource is an SSM document in AWS Systems Manager that defines the actions that Systems Manager performs, which can be used to set up and run commands on your instances.
        /// </summary>
        public static Output<GetDocumentResult> Invoke(GetDocumentInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetDocumentResult>("aws-native:ssm:getDocument", args ?? new GetDocumentInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDocumentArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// A name for the Systems Manager document.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetDocumentArgs()
        {
        }
    }

    public sealed class GetDocumentInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// A name for the Systems Manager document.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public GetDocumentInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetDocumentResult
    {
        /// <summary>
        /// A list of key and value pairs that describe attachments to a version of a document.
        /// </summary>
        public readonly ImmutableArray<Outputs.DocumentAttachmentsSource> Attachments;
        /// <summary>
        /// The content for the Systems Manager document in JSON, YAML or String format.
        /// </summary>
        public readonly object? Content;
        /// <summary>
        /// Specify the document format for the request. The document format can be either JSON or YAML. JSON is the default format.
        /// </summary>
        public readonly Pulumi.AwsNative.SSM.DocumentFormat? DocumentFormat;
        /// <summary>
        /// A list of SSM documents required by a document. For example, an ApplicationConfiguration document requires an ApplicationConfigurationSchema document.
        /// </summary>
        public readonly ImmutableArray<Outputs.DocumentRequires> Requires;
        /// <summary>
        /// Optional metadata that you assign to a resource. Tags enable you to categorize a resource in different ways, such as by purpose, owner, or environment.
        /// </summary>
        public readonly ImmutableArray<Outputs.DocumentTag> Tags;
        /// <summary>
        /// Specify a target type to define the kinds of resources the document can run on.
        /// </summary>
        public readonly string? TargetType;
        /// <summary>
        /// Update method - when set to 'Replace', the update will replace the existing document; when set to 'NewVersion', the update will create a new version.
        /// </summary>
        public readonly Pulumi.AwsNative.SSM.DocumentUpdateMethod? UpdateMethod;
        /// <summary>
        /// An optional field specifying the version of the artifact you are creating with the document. This value is unique across all versions of a document, and cannot be changed.
        /// </summary>
        public readonly string? VersionName;

        [OutputConstructor]
        private GetDocumentResult(
            ImmutableArray<Outputs.DocumentAttachmentsSource> attachments,

            object? content,

            Pulumi.AwsNative.SSM.DocumentFormat? documentFormat,

            ImmutableArray<Outputs.DocumentRequires> requires,

            ImmutableArray<Outputs.DocumentTag> tags,

            string? targetType,

            Pulumi.AwsNative.SSM.DocumentUpdateMethod? updateMethod,

            string? versionName)
        {
            Attachments = attachments;
            Content = content;
            DocumentFormat = documentFormat;
            Requires = requires;
            Tags = tags;
            TargetType = targetType;
            UpdateMethod = updateMethod;
            VersionName = versionName;
        }
    }
}