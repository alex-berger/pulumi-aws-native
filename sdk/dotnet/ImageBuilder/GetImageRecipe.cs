// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ImageBuilder
{
    public static class GetImageRecipe
    {
        /// <summary>
        /// Resource schema for AWS::ImageBuilder::ImageRecipe
        /// </summary>
        public static Task<GetImageRecipeResult> InvokeAsync(GetImageRecipeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetImageRecipeResult>("aws-native:imagebuilder:getImageRecipe", args ?? new GetImageRecipeArgs(), options.WithDefaults());

        /// <summary>
        /// Resource schema for AWS::ImageBuilder::ImageRecipe
        /// </summary>
        public static Output<GetImageRecipeResult> Invoke(GetImageRecipeInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetImageRecipeResult>("aws-native:imagebuilder:getImageRecipe", args ?? new GetImageRecipeInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetImageRecipeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the image recipe.
        /// </summary>
        [Input("arn", required: true)]
        public string Arn { get; set; } = null!;

        public GetImageRecipeArgs()
        {
        }
    }

    public sealed class GetImageRecipeInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the image recipe.
        /// </summary>
        [Input("arn", required: true)]
        public Input<string> Arn { get; set; } = null!;

        public GetImageRecipeInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetImageRecipeResult
    {
        /// <summary>
        /// Specify additional settings and launch scripts for your build instances.
        /// </summary>
        public readonly Outputs.ImageRecipeAdditionalInstanceConfiguration? AdditionalInstanceConfiguration;
        /// <summary>
        /// The Amazon Resource Name (ARN) of the image recipe.
        /// </summary>
        public readonly string? Arn;

        [OutputConstructor]
        private GetImageRecipeResult(
            Outputs.ImageRecipeAdditionalInstanceConfiguration? additionalInstanceConfiguration,

            string? arn)
        {
            AdditionalInstanceConfiguration = additionalInstanceConfiguration;
            Arn = arn;
        }
    }
}
