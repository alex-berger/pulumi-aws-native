// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AppFlow.Inputs
{

    public sealed class ConnectorProfileTrendmicroConnectorProfileCredentialsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Secret Access Key portion of the credentials.
        /// </summary>
        [Input("apiSecretKey", required: true)]
        public Input<string> ApiSecretKey { get; set; } = null!;

        public ConnectorProfileTrendmicroConnectorProfileCredentialsArgs()
        {
        }
    }
}
