// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Pinpoint.Outputs
{

    [OutputType]
    public sealed class InAppTemplateOverrideButtonConfiguration
    {
        public readonly Pulumi.AwsNative.Pinpoint.InAppTemplateButtonAction? ButtonAction;
        public readonly string? Link;

        [OutputConstructor]
        private InAppTemplateOverrideButtonConfiguration(
            Pulumi.AwsNative.Pinpoint.InAppTemplateButtonAction? buttonAction,

            string? link)
        {
            ButtonAction = buttonAction;
            Link = link;
        }
    }
}
