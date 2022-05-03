// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Creates a fleet provisioning template.
 */
export class ProvisioningTemplate extends pulumi.CustomResource {
    /**
     * Get an existing ProvisioningTemplate resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ProvisioningTemplate {
        return new ProvisioningTemplate(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:iot:ProvisioningTemplate';

    /**
     * Returns true if the given object is an instance of ProvisioningTemplate.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ProvisioningTemplate {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ProvisioningTemplate.__pulumiType;
    }

    public readonly description!: pulumi.Output<string | undefined>;
    public readonly enabled!: pulumi.Output<boolean | undefined>;
    public readonly preProvisioningHook!: pulumi.Output<outputs.iot.ProvisioningTemplateProvisioningHook | undefined>;
    public readonly provisioningRoleArn!: pulumi.Output<string>;
    public readonly tags!: pulumi.Output<outputs.iot.ProvisioningTemplateTag[] | undefined>;
    public /*out*/ readonly templateArn!: pulumi.Output<string>;
    public readonly templateBody!: pulumi.Output<string>;
    public readonly templateName!: pulumi.Output<string | undefined>;

    /**
     * Create a ProvisioningTemplate resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ProvisioningTemplateArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.provisioningRoleArn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'provisioningRoleArn'");
            }
            if ((!args || args.templateBody === undefined) && !opts.urn) {
                throw new Error("Missing required property 'templateBody'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["enabled"] = args ? args.enabled : undefined;
            resourceInputs["preProvisioningHook"] = args ? args.preProvisioningHook : undefined;
            resourceInputs["provisioningRoleArn"] = args ? args.provisioningRoleArn : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["templateBody"] = args ? args.templateBody : undefined;
            resourceInputs["templateName"] = args ? args.templateName : undefined;
            resourceInputs["templateArn"] = undefined /*out*/;
        } else {
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["enabled"] = undefined /*out*/;
            resourceInputs["preProvisioningHook"] = undefined /*out*/;
            resourceInputs["provisioningRoleArn"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["templateArn"] = undefined /*out*/;
            resourceInputs["templateBody"] = undefined /*out*/;
            resourceInputs["templateName"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ProvisioningTemplate.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a ProvisioningTemplate resource.
 */
export interface ProvisioningTemplateArgs {
    description?: pulumi.Input<string>;
    enabled?: pulumi.Input<boolean>;
    preProvisioningHook?: pulumi.Input<inputs.iot.ProvisioningTemplateProvisioningHookArgs>;
    provisioningRoleArn: pulumi.Input<string>;
    tags?: pulumi.Input<pulumi.Input<inputs.iot.ProvisioningTemplateTagArgs>[]>;
    templateBody: pulumi.Input<string>;
    templateName?: pulumi.Input<string>;
}
