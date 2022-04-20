// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource type definition for AWS::SSMIncidents::ResponsePlan
 */
export class ResponsePlan extends pulumi.CustomResource {
    /**
     * Get an existing ResponsePlan resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ResponsePlan {
        return new ResponsePlan(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:ssmincidents:ResponsePlan';

    /**
     * Returns true if the given object is an instance of ResponsePlan.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ResponsePlan {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ResponsePlan.__pulumiType;
    }

    /**
     * The list of actions.
     */
    public readonly actions!: pulumi.Output<outputs.ssmincidents.ResponsePlanAction[] | undefined>;
    /**
     * The ARN of the response plan.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    public readonly chatChannel!: pulumi.Output<outputs.ssmincidents.ResponsePlanChatChannel | undefined>;
    /**
     * The display name of the response plan.
     */
    public readonly displayName!: pulumi.Output<string | undefined>;
    /**
     * The list of engagements to use.
     */
    public readonly engagements!: pulumi.Output<string[] | undefined>;
    public readonly incidentTemplate!: pulumi.Output<outputs.ssmincidents.ResponsePlanIncidentTemplate>;
    /**
     * The name of the response plan.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The tags to apply to the response plan.
     */
    public readonly tags!: pulumi.Output<outputs.ssmincidents.ResponsePlanTag[] | undefined>;

    /**
     * Create a ResponsePlan resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ResponsePlanArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.incidentTemplate === undefined) && !opts.urn) {
                throw new Error("Missing required property 'incidentTemplate'");
            }
            resourceInputs["actions"] = args ? args.actions : undefined;
            resourceInputs["chatChannel"] = args ? args.chatChannel : undefined;
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["engagements"] = args ? args.engagements : undefined;
            resourceInputs["incidentTemplate"] = args ? args.incidentTemplate : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["arn"] = undefined /*out*/;
        } else {
            resourceInputs["actions"] = undefined /*out*/;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["chatChannel"] = undefined /*out*/;
            resourceInputs["displayName"] = undefined /*out*/;
            resourceInputs["engagements"] = undefined /*out*/;
            resourceInputs["incidentTemplate"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ResponsePlan.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a ResponsePlan resource.
 */
export interface ResponsePlanArgs {
    /**
     * The list of actions.
     */
    actions?: pulumi.Input<pulumi.Input<inputs.ssmincidents.ResponsePlanActionArgs>[]>;
    chatChannel?: pulumi.Input<inputs.ssmincidents.ResponsePlanChatChannelArgs>;
    /**
     * The display name of the response plan.
     */
    displayName?: pulumi.Input<string>;
    /**
     * The list of engagements to use.
     */
    engagements?: pulumi.Input<pulumi.Input<string>[]>;
    incidentTemplate: pulumi.Input<inputs.ssmincidents.ResponsePlanIncidentTemplateArgs>;
    /**
     * The name of the response plan.
     */
    name?: pulumi.Input<string>;
    /**
     * The tags to apply to the response plan.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.ssmincidents.ResponsePlanTagArgs>[]>;
}
