// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Create and manage wireless gateways, including LoRa gateways.
 */
export class WirelessGateway extends pulumi.CustomResource {
    /**
     * Get an existing WirelessGateway resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): WirelessGateway {
        return new WirelessGateway(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:iotwireless:WirelessGateway';

    /**
     * Returns true if the given object is an instance of WirelessGateway.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is WirelessGateway {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === WirelessGateway.__pulumiType;
    }

    /**
     * Arn for Wireless Gateway. Returned upon successful create.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * Description of Wireless Gateway.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The date and time when the most recent uplink was received.
     */
    public readonly lastUplinkReceivedAt!: pulumi.Output<string | undefined>;
    /**
     * The combination of Package, Station and Model which represents the version of the LoRaWAN Wireless Gateway.
     */
    public readonly loRaWAN!: pulumi.Output<outputs.iotwireless.WirelessGatewayLoRaWANGateway>;
    /**
     * Name of Wireless Gateway.
     */
    public readonly name!: pulumi.Output<string | undefined>;
    /**
     * A list of key-value pairs that contain metadata for the gateway.
     */
    public readonly tags!: pulumi.Output<outputs.iotwireless.WirelessGatewayTag[] | undefined>;
    /**
     * Thing Arn. Passed into Update to associate a Thing with the Wireless Gateway.
     */
    public readonly thingArn!: pulumi.Output<string | undefined>;
    /**
     * Thing Arn. If there is a Thing created, this can be returned with a Get call.
     */
    public /*out*/ readonly thingName!: pulumi.Output<string>;

    /**
     * Create a WirelessGateway resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: WirelessGatewayArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.loRaWAN === undefined) && !opts.urn) {
                throw new Error("Missing required property 'loRaWAN'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["lastUplinkReceivedAt"] = args ? args.lastUplinkReceivedAt : undefined;
            resourceInputs["loRaWAN"] = args ? args.loRaWAN : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["thingArn"] = args ? args.thingArn : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["thingName"] = undefined /*out*/;
        } else {
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["lastUplinkReceivedAt"] = undefined /*out*/;
            resourceInputs["loRaWAN"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["thingArn"] = undefined /*out*/;
            resourceInputs["thingName"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(WirelessGateway.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a WirelessGateway resource.
 */
export interface WirelessGatewayArgs {
    /**
     * Description of Wireless Gateway.
     */
    description?: pulumi.Input<string>;
    /**
     * The date and time when the most recent uplink was received.
     */
    lastUplinkReceivedAt?: pulumi.Input<string>;
    /**
     * The combination of Package, Station and Model which represents the version of the LoRaWAN Wireless Gateway.
     */
    loRaWAN: pulumi.Input<inputs.iotwireless.WirelessGatewayLoRaWANGatewayArgs>;
    /**
     * Name of Wireless Gateway.
     */
    name?: pulumi.Input<string>;
    /**
     * A list of key-value pairs that contain metadata for the gateway.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.iotwireless.WirelessGatewayTagArgs>[]>;
    /**
     * Thing Arn. Passed into Update to associate a Thing with the Wireless Gateway.
     */
    thingArn?: pulumi.Input<string>;
}
