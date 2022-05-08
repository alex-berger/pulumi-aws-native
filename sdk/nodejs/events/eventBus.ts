// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::Events::EventBus
 *
 * @deprecated EventBus is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
 */
export class EventBus extends pulumi.CustomResource {
    /**
     * Get an existing EventBus resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): EventBus {
        pulumi.log.warn("EventBus is deprecated: EventBus is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        return new EventBus(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:events:EventBus';

    /**
     * Returns true if the given object is an instance of EventBus.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is EventBus {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === EventBus.__pulumiType;
    }

    public /*out*/ readonly arn!: pulumi.Output<string>;
    public readonly eventSourceName!: pulumi.Output<string | undefined>;
    public readonly name!: pulumi.Output<string>;
    public /*out*/ readonly policy!: pulumi.Output<string>;
    public readonly tags!: pulumi.Output<outputs.events.EventBusTagEntry[] | undefined>;

    /**
     * Create a EventBus resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated EventBus is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible. */
    constructor(name: string, args?: EventBusArgs, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("EventBus is deprecated: EventBus is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["eventSourceName"] = args ? args.eventSourceName : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["policy"] = undefined /*out*/;
        } else {
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["eventSourceName"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["policy"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(EventBus.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a EventBus resource.
 */
export interface EventBusArgs {
    eventSourceName?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    tags?: pulumi.Input<pulumi.Input<inputs.events.EventBusTagEntryArgs>[]>;
}
