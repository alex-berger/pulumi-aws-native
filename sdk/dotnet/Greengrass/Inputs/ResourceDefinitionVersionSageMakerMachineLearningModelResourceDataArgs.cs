// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Greengrass.Inputs
{

    public sealed class ResourceDefinitionVersionSageMakerMachineLearningModelResourceDataArgs : Pulumi.ResourceArgs
    {
        [Input("destinationPath", required: true)]
        public Input<string> DestinationPath { get; set; } = null!;

        [Input("ownerSetting")]
        public Input<Inputs.ResourceDefinitionVersionResourceDownloadOwnerSettingArgs>? OwnerSetting { get; set; }

        [Input("sageMakerJobArn", required: true)]
        public Input<string> SageMakerJobArn { get; set; } = null!;

        public ResourceDefinitionVersionSageMakerMachineLearningModelResourceDataArgs()
        {
        }
    }
}
