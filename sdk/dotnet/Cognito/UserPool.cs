// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Cognito
{
    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpool.html
    /// </summary>
    [AwsNativeResourceType("aws-native:cognito:UserPool")]
    public partial class UserPool : Pulumi.CustomResource
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpool.html#cfn-cognito-userpool-accountrecoverysetting
        /// </summary>
        [Output("accountRecoverySetting")]
        public Output<Outputs.UserPoolAccountRecoverySetting?> AccountRecoverySetting { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpool.html#cfn-cognito-userpool-admincreateuserconfig
        /// </summary>
        [Output("adminCreateUserConfig")]
        public Output<Outputs.UserPoolAdminCreateUserConfig?> AdminCreateUserConfig { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpool.html#cfn-cognito-userpool-aliasattributes
        /// </summary>
        [Output("aliasAttributes")]
        public Output<ImmutableArray<string>> AliasAttributes { get; private set; } = null!;

        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpool.html#cfn-cognito-userpool-autoverifiedattributes
        /// </summary>
        [Output("autoVerifiedAttributes")]
        public Output<ImmutableArray<string>> AutoVerifiedAttributes { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpool.html#cfn-cognito-userpool-deviceconfiguration
        /// </summary>
        [Output("deviceConfiguration")]
        public Output<Outputs.UserPoolDeviceConfiguration?> DeviceConfiguration { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpool.html#cfn-cognito-userpool-emailconfiguration
        /// </summary>
        [Output("emailConfiguration")]
        public Output<Outputs.UserPoolEmailConfiguration?> EmailConfiguration { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpool.html#cfn-cognito-userpool-emailverificationmessage
        /// </summary>
        [Output("emailVerificationMessage")]
        public Output<string?> EmailVerificationMessage { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpool.html#cfn-cognito-userpool-emailverificationsubject
        /// </summary>
        [Output("emailVerificationSubject")]
        public Output<string?> EmailVerificationSubject { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpool.html#cfn-cognito-userpool-enabledmfas
        /// </summary>
        [Output("enabledMfas")]
        public Output<ImmutableArray<string>> EnabledMfas { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpool.html#cfn-cognito-userpool-lambdaconfig
        /// </summary>
        [Output("lambdaConfig")]
        public Output<Outputs.UserPoolLambdaConfig?> LambdaConfig { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpool.html#cfn-cognito-userpool-mfaconfiguration
        /// </summary>
        [Output("mfaConfiguration")]
        public Output<string?> MfaConfiguration { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpool.html#cfn-cognito-userpool-policies
        /// </summary>
        [Output("policies")]
        public Output<Outputs.UserPoolPolicies?> Policies { get; private set; } = null!;

        [Output("providerName")]
        public Output<string> ProviderName { get; private set; } = null!;

        [Output("providerURL")]
        public Output<string> ProviderURL { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpool.html#cfn-cognito-userpool-schema
        /// </summary>
        [Output("schema")]
        public Output<ImmutableArray<Outputs.UserPoolSchemaAttribute>> Schema { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpool.html#cfn-cognito-userpool-smsauthenticationmessage
        /// </summary>
        [Output("smsAuthenticationMessage")]
        public Output<string?> SmsAuthenticationMessage { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpool.html#cfn-cognito-userpool-smsconfiguration
        /// </summary>
        [Output("smsConfiguration")]
        public Output<Outputs.UserPoolSmsConfiguration?> SmsConfiguration { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpool.html#cfn-cognito-userpool-smsverificationmessage
        /// </summary>
        [Output("smsVerificationMessage")]
        public Output<string?> SmsVerificationMessage { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpool.html#cfn-cognito-userpool-userpooladdons
        /// </summary>
        [Output("userPoolAddOns")]
        public Output<Outputs.UserPoolUserPoolAddOns?> UserPoolAddOns { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpool.html#cfn-cognito-userpool-userpoolname
        /// </summary>
        [Output("userPoolName")]
        public Output<string?> UserPoolName { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpool.html#cfn-cognito-userpool-userpooltags
        /// </summary>
        [Output("userPoolTags")]
        public Output<Union<System.Text.Json.JsonElement, string>?> UserPoolTags { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpool.html#cfn-cognito-userpool-usernameattributes
        /// </summary>
        [Output("usernameAttributes")]
        public Output<ImmutableArray<string>> UsernameAttributes { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpool.html#cfn-cognito-userpool-usernameconfiguration
        /// </summary>
        [Output("usernameConfiguration")]
        public Output<Outputs.UserPoolUsernameConfiguration?> UsernameConfiguration { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpool.html#cfn-cognito-userpool-verificationmessagetemplate
        /// </summary>
        [Output("verificationMessageTemplate")]
        public Output<Outputs.UserPoolVerificationMessageTemplate?> VerificationMessageTemplate { get; private set; } = null!;


        /// <summary>
        /// Create a UserPool resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public UserPool(string name, UserPoolArgs? args = null, CustomResourceOptions? options = null)
            : base("aws-native:cognito:UserPool", name, args ?? new UserPoolArgs(), MakeResourceOptions(options, ""))
        {
        }

        private UserPool(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:cognito:UserPool", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing UserPool resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static UserPool Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new UserPool(name, id, options);
        }
    }

    public sealed class UserPoolArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpool.html#cfn-cognito-userpool-accountrecoverysetting
        /// </summary>
        [Input("accountRecoverySetting")]
        public Input<Inputs.UserPoolAccountRecoverySettingArgs>? AccountRecoverySetting { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpool.html#cfn-cognito-userpool-admincreateuserconfig
        /// </summary>
        [Input("adminCreateUserConfig")]
        public Input<Inputs.UserPoolAdminCreateUserConfigArgs>? AdminCreateUserConfig { get; set; }

        [Input("aliasAttributes")]
        private InputList<string>? _aliasAttributes;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpool.html#cfn-cognito-userpool-aliasattributes
        /// </summary>
        public InputList<string> AliasAttributes
        {
            get => _aliasAttributes ?? (_aliasAttributes = new InputList<string>());
            set => _aliasAttributes = value;
        }

        [Input("autoVerifiedAttributes")]
        private InputList<string>? _autoVerifiedAttributes;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpool.html#cfn-cognito-userpool-autoverifiedattributes
        /// </summary>
        public InputList<string> AutoVerifiedAttributes
        {
            get => _autoVerifiedAttributes ?? (_autoVerifiedAttributes = new InputList<string>());
            set => _autoVerifiedAttributes = value;
        }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpool.html#cfn-cognito-userpool-deviceconfiguration
        /// </summary>
        [Input("deviceConfiguration")]
        public Input<Inputs.UserPoolDeviceConfigurationArgs>? DeviceConfiguration { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpool.html#cfn-cognito-userpool-emailconfiguration
        /// </summary>
        [Input("emailConfiguration")]
        public Input<Inputs.UserPoolEmailConfigurationArgs>? EmailConfiguration { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpool.html#cfn-cognito-userpool-emailverificationmessage
        /// </summary>
        [Input("emailVerificationMessage")]
        public Input<string>? EmailVerificationMessage { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpool.html#cfn-cognito-userpool-emailverificationsubject
        /// </summary>
        [Input("emailVerificationSubject")]
        public Input<string>? EmailVerificationSubject { get; set; }

        [Input("enabledMfas")]
        private InputList<string>? _enabledMfas;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpool.html#cfn-cognito-userpool-enabledmfas
        /// </summary>
        public InputList<string> EnabledMfas
        {
            get => _enabledMfas ?? (_enabledMfas = new InputList<string>());
            set => _enabledMfas = value;
        }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpool.html#cfn-cognito-userpool-lambdaconfig
        /// </summary>
        [Input("lambdaConfig")]
        public Input<Inputs.UserPoolLambdaConfigArgs>? LambdaConfig { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpool.html#cfn-cognito-userpool-mfaconfiguration
        /// </summary>
        [Input("mfaConfiguration")]
        public Input<string>? MfaConfiguration { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpool.html#cfn-cognito-userpool-policies
        /// </summary>
        [Input("policies")]
        public Input<Inputs.UserPoolPoliciesArgs>? Policies { get; set; }

        [Input("schema")]
        private InputList<Inputs.UserPoolSchemaAttributeArgs>? _schema;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpool.html#cfn-cognito-userpool-schema
        /// </summary>
        public InputList<Inputs.UserPoolSchemaAttributeArgs> Schema
        {
            get => _schema ?? (_schema = new InputList<Inputs.UserPoolSchemaAttributeArgs>());
            set => _schema = value;
        }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpool.html#cfn-cognito-userpool-smsauthenticationmessage
        /// </summary>
        [Input("smsAuthenticationMessage")]
        public Input<string>? SmsAuthenticationMessage { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpool.html#cfn-cognito-userpool-smsconfiguration
        /// </summary>
        [Input("smsConfiguration")]
        public Input<Inputs.UserPoolSmsConfigurationArgs>? SmsConfiguration { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpool.html#cfn-cognito-userpool-smsverificationmessage
        /// </summary>
        [Input("smsVerificationMessage")]
        public Input<string>? SmsVerificationMessage { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpool.html#cfn-cognito-userpool-userpooladdons
        /// </summary>
        [Input("userPoolAddOns")]
        public Input<Inputs.UserPoolUserPoolAddOnsArgs>? UserPoolAddOns { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpool.html#cfn-cognito-userpool-userpoolname
        /// </summary>
        [Input("userPoolName")]
        public Input<string>? UserPoolName { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpool.html#cfn-cognito-userpool-userpooltags
        /// </summary>
        [Input("userPoolTags")]
        public InputUnion<System.Text.Json.JsonElement, string>? UserPoolTags { get; set; }

        [Input("usernameAttributes")]
        private InputList<string>? _usernameAttributes;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpool.html#cfn-cognito-userpool-usernameattributes
        /// </summary>
        public InputList<string> UsernameAttributes
        {
            get => _usernameAttributes ?? (_usernameAttributes = new InputList<string>());
            set => _usernameAttributes = value;
        }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpool.html#cfn-cognito-userpool-usernameconfiguration
        /// </summary>
        [Input("usernameConfiguration")]
        public Input<Inputs.UserPoolUsernameConfigurationArgs>? UsernameConfiguration { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpool.html#cfn-cognito-userpool-verificationmessagetemplate
        /// </summary>
        [Input("verificationMessageTemplate")]
        public Input<Inputs.UserPoolVerificationMessageTemplateArgs>? VerificationMessageTemplate { get; set; }

        public UserPoolArgs()
        {
        }
    }
}
