// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AppFlow.Inputs
{

    public sealed class ConnectorProfileZendeskConnectorProfileCredentialsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The credentials used to access protected resources.
        /// </summary>
        [Input("accessToken")]
        public Input<string>? AccessToken { get; set; }

        /// <summary>
        /// The identiﬁer for the desired client.
        /// </summary>
        [Input("clientId", required: true)]
        public Input<string> ClientId { get; set; } = null!;

        /// <summary>
        /// The client secret used by the oauth client to authenticate to the authorization server.
        /// </summary>
        [Input("clientSecret", required: true)]
        public Input<string> ClientSecret { get; set; } = null!;

        /// <summary>
        /// The oauth needed to request security tokens from the connector endpoint.
        /// </summary>
        [Input("connectorOAuthRequest")]
        public Input<Inputs.ConnectorProfileConnectorOAuthRequestArgs>? ConnectorOAuthRequest { get; set; }

        public ConnectorProfileZendeskConnectorProfileCredentialsArgs()
        {
        }
    }
}
