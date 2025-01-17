// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource schema for StateMachine
 */
export class StateMachine extends pulumi.CustomResource {
    /**
     * Get an existing StateMachine resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): StateMachine {
        return new StateMachine(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:stepfunctions:StateMachine';

    /**
     * Returns true if the given object is an instance of StateMachine.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is StateMachine {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === StateMachine.__pulumiType;
    }

    public /*out*/ readonly arn!: pulumi.Output<string>;
    public readonly definition!: pulumi.Output<outputs.stepfunctions.StateMachineDefinition | undefined>;
    public readonly definitionS3Location!: pulumi.Output<outputs.stepfunctions.StateMachineS3Location | undefined>;
    public readonly definitionString!: pulumi.Output<string | undefined>;
    public readonly definitionSubstitutions!: pulumi.Output<outputs.stepfunctions.StateMachineDefinitionSubstitutions | undefined>;
    public readonly loggingConfiguration!: pulumi.Output<outputs.stepfunctions.StateMachineLoggingConfiguration | undefined>;
    public /*out*/ readonly name!: pulumi.Output<string>;
    public readonly roleArn!: pulumi.Output<string>;
    public readonly stateMachineName!: pulumi.Output<string | undefined>;
    public readonly stateMachineType!: pulumi.Output<enums.stepfunctions.StateMachineType | undefined>;
    public readonly tags!: pulumi.Output<outputs.stepfunctions.StateMachineTagsEntry[] | undefined>;
    public readonly tracingConfiguration!: pulumi.Output<outputs.stepfunctions.StateMachineTracingConfiguration | undefined>;

    /**
     * Create a StateMachine resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: StateMachineArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.roleArn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'roleArn'");
            }
            resourceInputs["definition"] = args ? args.definition : undefined;
            resourceInputs["definitionS3Location"] = args ? args.definitionS3Location : undefined;
            resourceInputs["definitionString"] = args ? args.definitionString : undefined;
            resourceInputs["definitionSubstitutions"] = args ? args.definitionSubstitutions : undefined;
            resourceInputs["loggingConfiguration"] = args ? args.loggingConfiguration : undefined;
            resourceInputs["roleArn"] = args ? args.roleArn : undefined;
            resourceInputs["stateMachineName"] = args ? args.stateMachineName : undefined;
            resourceInputs["stateMachineType"] = args ? args.stateMachineType : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["tracingConfiguration"] = args ? args.tracingConfiguration : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
        } else {
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["definition"] = undefined /*out*/;
            resourceInputs["definitionS3Location"] = undefined /*out*/;
            resourceInputs["definitionString"] = undefined /*out*/;
            resourceInputs["definitionSubstitutions"] = undefined /*out*/;
            resourceInputs["loggingConfiguration"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["roleArn"] = undefined /*out*/;
            resourceInputs["stateMachineName"] = undefined /*out*/;
            resourceInputs["stateMachineType"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["tracingConfiguration"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(StateMachine.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a StateMachine resource.
 */
export interface StateMachineArgs {
    definition?: pulumi.Input<inputs.stepfunctions.StateMachineDefinitionArgs>;
    definitionS3Location?: pulumi.Input<inputs.stepfunctions.StateMachineS3LocationArgs>;
    definitionString?: pulumi.Input<string>;
    definitionSubstitutions?: pulumi.Input<inputs.stepfunctions.StateMachineDefinitionSubstitutionsArgs>;
    loggingConfiguration?: pulumi.Input<inputs.stepfunctions.StateMachineLoggingConfigurationArgs>;
    roleArn: pulumi.Input<string>;
    stateMachineName?: pulumi.Input<string>;
    stateMachineType?: pulumi.Input<enums.stepfunctions.StateMachineType>;
    tags?: pulumi.Input<pulumi.Input<inputs.stepfunctions.StateMachineTagsEntryArgs>[]>;
    tracingConfiguration?: pulumi.Input<inputs.stepfunctions.StateMachineTracingConfigurationArgs>;
}
