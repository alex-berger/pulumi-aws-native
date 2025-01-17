// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Cognito
{
    public static class GetIdentityPool
    {
        /// <summary>
        /// Resource Type definition for AWS::Cognito::IdentityPool
        /// </summary>
        public static Task<GetIdentityPoolResult> InvokeAsync(GetIdentityPoolArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetIdentityPoolResult>("aws-native:cognito:getIdentityPool", args ?? new GetIdentityPoolArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::Cognito::IdentityPool
        /// </summary>
        public static Output<GetIdentityPoolResult> Invoke(GetIdentityPoolInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetIdentityPoolResult>("aws-native:cognito:getIdentityPool", args ?? new GetIdentityPoolInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetIdentityPoolArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetIdentityPoolArgs()
        {
        }
    }

    public sealed class GetIdentityPoolInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetIdentityPoolInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetIdentityPoolResult
    {
        public readonly bool? AllowClassicFlow;
        public readonly bool? AllowUnauthenticatedIdentities;
        public readonly object? CognitoEvents;
        public readonly ImmutableArray<Outputs.IdentityPoolCognitoIdentityProvider> CognitoIdentityProviders;
        public readonly Outputs.IdentityPoolCognitoStreams? CognitoStreams;
        public readonly string? DeveloperProviderName;
        public readonly string? Id;
        public readonly string? IdentityPoolName;
        public readonly string? Name;
        public readonly ImmutableArray<string> OpenIdConnectProviderARNs;
        public readonly Outputs.IdentityPoolPushSync? PushSync;
        public readonly ImmutableArray<string> SamlProviderARNs;
        public readonly object? SupportedLoginProviders;

        [OutputConstructor]
        private GetIdentityPoolResult(
            bool? allowClassicFlow,

            bool? allowUnauthenticatedIdentities,

            object? cognitoEvents,

            ImmutableArray<Outputs.IdentityPoolCognitoIdentityProvider> cognitoIdentityProviders,

            Outputs.IdentityPoolCognitoStreams? cognitoStreams,

            string? developerProviderName,

            string? id,

            string? identityPoolName,

            string? name,

            ImmutableArray<string> openIdConnectProviderARNs,

            Outputs.IdentityPoolPushSync? pushSync,

            ImmutableArray<string> samlProviderARNs,

            object? supportedLoginProviders)
        {
            AllowClassicFlow = allowClassicFlow;
            AllowUnauthenticatedIdentities = allowUnauthenticatedIdentities;
            CognitoEvents = cognitoEvents;
            CognitoIdentityProviders = cognitoIdentityProviders;
            CognitoStreams = cognitoStreams;
            DeveloperProviderName = developerProviderName;
            Id = id;
            IdentityPoolName = identityPoolName;
            Name = name;
            OpenIdConnectProviderARNs = openIdConnectProviderARNs;
            PushSync = pushSync;
            SamlProviderARNs = samlProviderARNs;
            SupportedLoginProviders = supportedLoginProviders;
        }
    }
}
