// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.DMS.Inputs
{

    public sealed class EndpointDocDbSettingsArgs : Pulumi.ResourceArgs
    {
        [Input("docsToInvestigate")]
        public Input<int>? DocsToInvestigate { get; set; }

        [Input("extractDocId")]
        public Input<bool>? ExtractDocId { get; set; }

        [Input("nestingLevel")]
        public Input<string>? NestingLevel { get; set; }

        [Input("secretsManagerAccessRoleArn")]
        public Input<string>? SecretsManagerAccessRoleArn { get; set; }

        [Input("secretsManagerSecretId")]
        public Input<string>? SecretsManagerSecretId { get; set; }

        public EndpointDocDbSettingsArgs()
        {
        }
    }
}