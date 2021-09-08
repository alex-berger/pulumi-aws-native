// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AppSync
{
    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-graphqlapi.html
    /// </summary>
    [AwsNativeResourceType("aws-native:appsync:GraphQLApi")]
    public partial class GraphQLApi : Pulumi.CustomResource
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-graphqlapi.html#cfn-appsync-graphqlapi-additionalauthenticationproviders
        /// </summary>
        [Output("additionalAuthenticationProviders")]
        public Output<Outputs.GraphQLApiAdditionalAuthenticationProviders?> AdditionalAuthenticationProviders { get; private set; } = null!;

        [Output("apiId")]
        public Output<string> ApiId { get; private set; } = null!;

        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-graphqlapi.html#cfn-appsync-graphqlapi-authenticationtype
        /// </summary>
        [Output("authenticationType")]
        public Output<string> AuthenticationType { get; private set; } = null!;

        [Output("graphQLUrl")]
        public Output<string> GraphQLUrl { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-graphqlapi.html#cfn-appsync-graphqlapi-lambdaauthorizerconfig
        /// </summary>
        [Output("lambdaAuthorizerConfig")]
        public Output<Outputs.GraphQLApiLambdaAuthorizerConfig?> LambdaAuthorizerConfig { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-graphqlapi.html#cfn-appsync-graphqlapi-logconfig
        /// </summary>
        [Output("logConfig")]
        public Output<Outputs.GraphQLApiLogConfig?> LogConfig { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-graphqlapi.html#cfn-appsync-graphqlapi-name
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-graphqlapi.html#cfn-appsync-graphqlapi-openidconnectconfig
        /// </summary>
        [Output("openIDConnectConfig")]
        public Output<Outputs.GraphQLApiOpenIDConnectConfig?> OpenIDConnectConfig { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-graphqlapi.html#cfn-appsync-graphqlapi-tags
        /// </summary>
        [Output("tags")]
        public Output<Outputs.GraphQLApiTags?> Tags { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-graphqlapi.html#cfn-appsync-graphqlapi-userpoolconfig
        /// </summary>
        [Output("userPoolConfig")]
        public Output<Outputs.GraphQLApiUserPoolConfig?> UserPoolConfig { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-graphqlapi.html#cfn-appsync-graphqlapi-xrayenabled
        /// </summary>
        [Output("xrayEnabled")]
        public Output<bool?> XrayEnabled { get; private set; } = null!;


        /// <summary>
        /// Create a GraphQLApi resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public GraphQLApi(string name, GraphQLApiArgs args, CustomResourceOptions? options = null)
            : base("aws-native:appsync:GraphQLApi", name, args ?? new GraphQLApiArgs(), MakeResourceOptions(options, ""))
        {
        }

        private GraphQLApi(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:appsync:GraphQLApi", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing GraphQLApi resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static GraphQLApi Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new GraphQLApi(name, id, options);
        }
    }

    public sealed class GraphQLApiArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-graphqlapi.html#cfn-appsync-graphqlapi-additionalauthenticationproviders
        /// </summary>
        [Input("additionalAuthenticationProviders")]
        public Input<Inputs.GraphQLApiAdditionalAuthenticationProvidersArgs>? AdditionalAuthenticationProviders { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-graphqlapi.html#cfn-appsync-graphqlapi-authenticationtype
        /// </summary>
        [Input("authenticationType", required: true)]
        public Input<string> AuthenticationType { get; set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-graphqlapi.html#cfn-appsync-graphqlapi-lambdaauthorizerconfig
        /// </summary>
        [Input("lambdaAuthorizerConfig")]
        public Input<Inputs.GraphQLApiLambdaAuthorizerConfigArgs>? LambdaAuthorizerConfig { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-graphqlapi.html#cfn-appsync-graphqlapi-logconfig
        /// </summary>
        [Input("logConfig")]
        public Input<Inputs.GraphQLApiLogConfigArgs>? LogConfig { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-graphqlapi.html#cfn-appsync-graphqlapi-name
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-graphqlapi.html#cfn-appsync-graphqlapi-openidconnectconfig
        /// </summary>
        [Input("openIDConnectConfig")]
        public Input<Inputs.GraphQLApiOpenIDConnectConfigArgs>? OpenIDConnectConfig { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-graphqlapi.html#cfn-appsync-graphqlapi-tags
        /// </summary>
        [Input("tags")]
        public Input<Inputs.GraphQLApiTagsArgs>? Tags { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-graphqlapi.html#cfn-appsync-graphqlapi-userpoolconfig
        /// </summary>
        [Input("userPoolConfig")]
        public Input<Inputs.GraphQLApiUserPoolConfigArgs>? UserPoolConfig { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-graphqlapi.html#cfn-appsync-graphqlapi-xrayenabled
        /// </summary>
        [Input("xrayEnabled")]
        public Input<bool>? XrayEnabled { get; set; }

        public GraphQLApiArgs()
        {
        }
    }
}
