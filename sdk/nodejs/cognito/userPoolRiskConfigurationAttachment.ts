// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::Cognito::UserPoolRiskConfigurationAttachment
 *
 * @deprecated UserPoolRiskConfigurationAttachment is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
 */
export class UserPoolRiskConfigurationAttachment extends pulumi.CustomResource {
    /**
     * Get an existing UserPoolRiskConfigurationAttachment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): UserPoolRiskConfigurationAttachment {
        pulumi.log.warn("UserPoolRiskConfigurationAttachment is deprecated: UserPoolRiskConfigurationAttachment is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        return new UserPoolRiskConfigurationAttachment(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:cognito:UserPoolRiskConfigurationAttachment';

    /**
     * Returns true if the given object is an instance of UserPoolRiskConfigurationAttachment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is UserPoolRiskConfigurationAttachment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === UserPoolRiskConfigurationAttachment.__pulumiType;
    }

    public readonly accountTakeoverRiskConfiguration!: pulumi.Output<outputs.cognito.UserPoolRiskConfigurationAttachmentAccountTakeoverRiskConfigurationType | undefined>;
    public readonly clientId!: pulumi.Output<string>;
    public readonly compromisedCredentialsRiskConfiguration!: pulumi.Output<outputs.cognito.UserPoolRiskConfigurationAttachmentCompromisedCredentialsRiskConfigurationType | undefined>;
    public readonly riskExceptionConfiguration!: pulumi.Output<outputs.cognito.UserPoolRiskConfigurationAttachmentRiskExceptionConfigurationType | undefined>;
    public readonly userPoolId!: pulumi.Output<string>;

    /**
     * Create a UserPoolRiskConfigurationAttachment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated UserPoolRiskConfigurationAttachment is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible. */
    constructor(name: string, args: UserPoolRiskConfigurationAttachmentArgs, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("UserPoolRiskConfigurationAttachment is deprecated: UserPoolRiskConfigurationAttachment is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.clientId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clientId'");
            }
            if ((!args || args.userPoolId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'userPoolId'");
            }
            resourceInputs["accountTakeoverRiskConfiguration"] = args ? args.accountTakeoverRiskConfiguration : undefined;
            resourceInputs["clientId"] = args ? args.clientId : undefined;
            resourceInputs["compromisedCredentialsRiskConfiguration"] = args ? args.compromisedCredentialsRiskConfiguration : undefined;
            resourceInputs["riskExceptionConfiguration"] = args ? args.riskExceptionConfiguration : undefined;
            resourceInputs["userPoolId"] = args ? args.userPoolId : undefined;
        } else {
            resourceInputs["accountTakeoverRiskConfiguration"] = undefined /*out*/;
            resourceInputs["clientId"] = undefined /*out*/;
            resourceInputs["compromisedCredentialsRiskConfiguration"] = undefined /*out*/;
            resourceInputs["riskExceptionConfiguration"] = undefined /*out*/;
            resourceInputs["userPoolId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(UserPoolRiskConfigurationAttachment.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a UserPoolRiskConfigurationAttachment resource.
 */
export interface UserPoolRiskConfigurationAttachmentArgs {
    accountTakeoverRiskConfiguration?: pulumi.Input<inputs.cognito.UserPoolRiskConfigurationAttachmentAccountTakeoverRiskConfigurationTypeArgs>;
    clientId: pulumi.Input<string>;
    compromisedCredentialsRiskConfiguration?: pulumi.Input<inputs.cognito.UserPoolRiskConfigurationAttachmentCompromisedCredentialsRiskConfigurationTypeArgs>;
    riskExceptionConfiguration?: pulumi.Input<inputs.cognito.UserPoolRiskConfigurationAttachmentRiskExceptionConfigurationTypeArgs>;
    userPoolId: pulumi.Input<string>;
}