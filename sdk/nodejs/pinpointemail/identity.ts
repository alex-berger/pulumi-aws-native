// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::PinpointEmail::Identity
 *
 * @deprecated Identity is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
 */
export class Identity extends pulumi.CustomResource {
    /**
     * Get an existing Identity resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Identity {
        pulumi.log.warn("Identity is deprecated: Identity is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        return new Identity(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:pinpointemail:Identity';

    /**
     * Returns true if the given object is an instance of Identity.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Identity {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Identity.__pulumiType;
    }

    public readonly dkimSigningEnabled!: pulumi.Output<boolean | undefined>;
    public readonly feedbackForwardingEnabled!: pulumi.Output<boolean | undefined>;
    public /*out*/ readonly identityDNSRecordName1!: pulumi.Output<string>;
    public /*out*/ readonly identityDNSRecordName2!: pulumi.Output<string>;
    public /*out*/ readonly identityDNSRecordName3!: pulumi.Output<string>;
    public /*out*/ readonly identityDNSRecordValue1!: pulumi.Output<string>;
    public /*out*/ readonly identityDNSRecordValue2!: pulumi.Output<string>;
    public /*out*/ readonly identityDNSRecordValue3!: pulumi.Output<string>;
    public readonly mailFromAttributes!: pulumi.Output<outputs.pinpointemail.IdentityMailFromAttributes | undefined>;
    public readonly name!: pulumi.Output<string>;
    public readonly tags!: pulumi.Output<outputs.pinpointemail.IdentityTags[] | undefined>;

    /**
     * Create a Identity resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated Identity is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible. */
    constructor(name: string, args?: IdentityArgs, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("Identity is deprecated: Identity is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["dkimSigningEnabled"] = args ? args.dkimSigningEnabled : undefined;
            resourceInputs["feedbackForwardingEnabled"] = args ? args.feedbackForwardingEnabled : undefined;
            resourceInputs["mailFromAttributes"] = args ? args.mailFromAttributes : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["identityDNSRecordName1"] = undefined /*out*/;
            resourceInputs["identityDNSRecordName2"] = undefined /*out*/;
            resourceInputs["identityDNSRecordName3"] = undefined /*out*/;
            resourceInputs["identityDNSRecordValue1"] = undefined /*out*/;
            resourceInputs["identityDNSRecordValue2"] = undefined /*out*/;
            resourceInputs["identityDNSRecordValue3"] = undefined /*out*/;
        } else {
            resourceInputs["dkimSigningEnabled"] = undefined /*out*/;
            resourceInputs["feedbackForwardingEnabled"] = undefined /*out*/;
            resourceInputs["identityDNSRecordName1"] = undefined /*out*/;
            resourceInputs["identityDNSRecordName2"] = undefined /*out*/;
            resourceInputs["identityDNSRecordName3"] = undefined /*out*/;
            resourceInputs["identityDNSRecordValue1"] = undefined /*out*/;
            resourceInputs["identityDNSRecordValue2"] = undefined /*out*/;
            resourceInputs["identityDNSRecordValue3"] = undefined /*out*/;
            resourceInputs["mailFromAttributes"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Identity.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Identity resource.
 */
export interface IdentityArgs {
    dkimSigningEnabled?: pulumi.Input<boolean>;
    feedbackForwardingEnabled?: pulumi.Input<boolean>;
    mailFromAttributes?: pulumi.Input<inputs.pinpointemail.IdentityMailFromAttributesArgs>;
    name?: pulumi.Input<string>;
    tags?: pulumi.Input<pulumi.Input<inputs.pinpointemail.IdentityTagsArgs>[]>;
}
