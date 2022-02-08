// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.MediaConvert
{
    public static class GetPreset
    {
        /// <summary>
        /// Resource Type definition for AWS::MediaConvert::Preset
        /// </summary>
        public static Task<GetPresetResult> InvokeAsync(GetPresetArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetPresetResult>("aws-native:mediaconvert:getPreset", args ?? new GetPresetArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::MediaConvert::Preset
        /// </summary>
        public static Output<GetPresetResult> Invoke(GetPresetInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetPresetResult>("aws-native:mediaconvert:getPreset", args ?? new GetPresetInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetPresetArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetPresetArgs()
        {
        }
    }

    public sealed class GetPresetInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetPresetInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetPresetResult
    {
        public readonly string? Arn;
        public readonly string? Category;
        public readonly string? Description;
        public readonly string? Id;
        public readonly object? SettingsJson;
        public readonly object? Tags;

        [OutputConstructor]
        private GetPresetResult(
            string? arn,

            string? category,

            string? description,

            string? id,

            object? settingsJson,

            object? tags)
        {
            Arn = arn;
            Category = category;
            Description = description;
            Id = id;
            SettingsJson = settingsJson;
            Tags = tags;
        }
    }
}