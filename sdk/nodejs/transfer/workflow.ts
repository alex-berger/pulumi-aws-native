// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::Transfer::Workflow
 */
export class Workflow extends pulumi.CustomResource {
    /**
     * Get an existing Workflow resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Workflow {
        return new Workflow(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:transfer:Workflow';

    /**
     * Returns true if the given object is an instance of Workflow.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Workflow {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Workflow.__pulumiType;
    }

    /**
     * Specifies the unique Amazon Resource Name (ARN) for the workflow.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * A textual description for the workflow.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Specifies the steps (actions) to take if any errors are encountered during execution of the workflow.
     */
    public readonly onExceptionSteps!: pulumi.Output<outputs.transfer.WorkflowStep[] | undefined>;
    /**
     * Specifies the details for the steps that are in the specified workflow.
     */
    public readonly steps!: pulumi.Output<outputs.transfer.WorkflowStep[]>;
    /**
     * Key-value pairs that can be used to group and search for workflows. Tags are metadata attached to workflows for any purpose.
     */
    public readonly tags!: pulumi.Output<outputs.transfer.WorkflowTag[] | undefined>;
    /**
     * A unique identifier for the workflow.
     */
    public /*out*/ readonly workflowId!: pulumi.Output<string>;

    /**
     * Create a Workflow resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: WorkflowArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.steps === undefined) && !opts.urn) {
                throw new Error("Missing required property 'steps'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["onExceptionSteps"] = args ? args.onExceptionSteps : undefined;
            resourceInputs["steps"] = args ? args.steps : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["workflowId"] = undefined /*out*/;
        } else {
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["onExceptionSteps"] = undefined /*out*/;
            resourceInputs["steps"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["workflowId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Workflow.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Workflow resource.
 */
export interface WorkflowArgs {
    /**
     * A textual description for the workflow.
     */
    description?: pulumi.Input<string>;
    /**
     * Specifies the steps (actions) to take if any errors are encountered during execution of the workflow.
     */
    onExceptionSteps?: pulumi.Input<pulumi.Input<inputs.transfer.WorkflowStepArgs>[]>;
    /**
     * Specifies the details for the steps that are in the specified workflow.
     */
    steps: pulumi.Input<pulumi.Input<inputs.transfer.WorkflowStepArgs>[]>;
    /**
     * Key-value pairs that can be used to group and search for workflows. Tags are metadata attached to workflows for any purpose.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.transfer.WorkflowTagArgs>[]>;
}