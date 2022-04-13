// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ASK.Inputs
{

    public sealed class SkillAuthenticationConfigurationArgs : Pulumi.ResourceArgs
    {
        [Input("clientId", required: true)]
        public Input<string> ClientId { get; set; } = null!;

        [Input("clientSecret", required: true)]
        public Input<string> ClientSecret { get; set; } = null!;

        [Input("refreshToken", required: true)]
        public Input<string> RefreshToken { get; set; } = null!;

        public SkillAuthenticationConfigurationArgs()
        {
        }
    }
}
