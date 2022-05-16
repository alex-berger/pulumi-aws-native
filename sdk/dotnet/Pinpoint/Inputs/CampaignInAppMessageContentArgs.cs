// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Pinpoint.Inputs
{

    public sealed class CampaignInAppMessageContentArgs : Pulumi.ResourceArgs
    {
        [Input("backgroundColor")]
        public Input<string>? BackgroundColor { get; set; }

        [Input("bodyConfig")]
        public Input<Inputs.CampaignInAppMessageBodyConfigArgs>? BodyConfig { get; set; }

        [Input("headerConfig")]
        public Input<Inputs.CampaignInAppMessageHeaderConfigArgs>? HeaderConfig { get; set; }

        [Input("imageUrl")]
        public Input<string>? ImageUrl { get; set; }

        [Input("primaryBtn")]
        public Input<Inputs.CampaignInAppMessageButtonArgs>? PrimaryBtn { get; set; }

        [Input("secondaryBtn")]
        public Input<Inputs.CampaignInAppMessageButtonArgs>? SecondaryBtn { get; set; }

        public CampaignInAppMessageContentArgs()
        {
        }
    }
}