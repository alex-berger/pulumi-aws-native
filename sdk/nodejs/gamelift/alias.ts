// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * The AWS::GameLift::Alias resource creates an alias for an Amazon GameLift (GameLift) fleet destination.
 */
export class Alias extends pulumi.CustomResource {
    /**
     * Get an existing Alias resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Alias {
        return new Alias(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:gamelift:Alias';

    /**
     * Returns true if the given object is an instance of Alias.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Alias {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Alias.__pulumiType;
    }

    /**
     * Unique alias ID
     */
    public /*out*/ readonly aliasId!: pulumi.Output<string>;
    /**
     * A human-readable description of the alias.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * A descriptive label that is associated with an alias. Alias names do not need to be unique.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * A routing configuration that specifies where traffic is directed for this alias, such as to a fleet or to a message.
     */
    public readonly routingStrategy!: pulumi.Output<outputs.gamelift.AliasRoutingStrategy>;

    /**
     * Create a Alias resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AliasArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.routingStrategy === undefined) && !opts.urn) {
                throw new Error("Missing required property 'routingStrategy'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["routingStrategy"] = args ? args.routingStrategy : undefined;
            resourceInputs["aliasId"] = undefined /*out*/;
        } else {
            resourceInputs["aliasId"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["routingStrategy"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Alias.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Alias resource.
 */
export interface AliasArgs {
    /**
     * A human-readable description of the alias.
     */
    description?: pulumi.Input<string>;
    /**
     * A descriptive label that is associated with an alias. Alias names do not need to be unique.
     */
    name?: pulumi.Input<string>;
    /**
     * A routing configuration that specifies where traffic is directed for this alias, such as to a fleet or to a message.
     */
    routingStrategy: pulumi.Input<inputs.gamelift.AliasRoutingStrategyArgs>;
}
