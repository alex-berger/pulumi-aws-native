// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::Connect::ContactFlow
 */
export class ContactFlow extends pulumi.CustomResource {
    /**
     * Get an existing ContactFlow resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ContactFlow {
        return new ContactFlow(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:connect:ContactFlow';

    /**
     * Returns true if the given object is an instance of ContactFlow.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ContactFlow {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ContactFlow.__pulumiType;
    }

    /**
     * The identifier of the contact flow (ARN).
     */
    public /*out*/ readonly contactFlowArn!: pulumi.Output<string>;
    /**
     * The content of the contact flow in JSON format.
     */
    public readonly content!: pulumi.Output<string>;
    /**
     * The description of the contact flow.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The identifier of the Amazon Connect instance (ARN).
     */
    public readonly instanceArn!: pulumi.Output<string>;
    /**
     * The name of the contact flow.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The state of the contact flow.
     */
    public readonly state!: pulumi.Output<enums.connect.ContactFlowState | undefined>;
    /**
     * One or more tags.
     */
    public readonly tags!: pulumi.Output<outputs.connect.ContactFlowTag[] | undefined>;
    /**
     * The type of the contact flow.
     */
    public readonly type!: pulumi.Output<enums.connect.ContactFlowType | undefined>;

    /**
     * Create a ContactFlow resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ContactFlowArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.content === undefined) && !opts.urn) {
                throw new Error("Missing required property 'content'");
            }
            if ((!args || args.instanceArn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceArn'");
            }
            resourceInputs["content"] = args ? args.content : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["instanceArn"] = args ? args.instanceArn : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["state"] = args ? args.state : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["contactFlowArn"] = undefined /*out*/;
        } else {
            resourceInputs["contactFlowArn"] = undefined /*out*/;
            resourceInputs["content"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["instanceArn"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ContactFlow.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a ContactFlow resource.
 */
export interface ContactFlowArgs {
    /**
     * The content of the contact flow in JSON format.
     */
    content: pulumi.Input<string>;
    /**
     * The description of the contact flow.
     */
    description?: pulumi.Input<string>;
    /**
     * The identifier of the Amazon Connect instance (ARN).
     */
    instanceArn: pulumi.Input<string>;
    /**
     * The name of the contact flow.
     */
    name?: pulumi.Input<string>;
    /**
     * The state of the contact flow.
     */
    state?: pulumi.Input<enums.connect.ContactFlowState>;
    /**
     * One or more tags.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.connect.ContactFlowTagArgs>[]>;
    /**
     * The type of the contact flow.
     */
    type?: pulumi.Input<enums.connect.ContactFlowType>;
}