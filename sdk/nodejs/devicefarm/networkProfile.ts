// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * AWS::DeviceFarm::NetworkProfile creates a new DF Network Profile
 */
export class NetworkProfile extends pulumi.CustomResource {
    /**
     * Get an existing NetworkProfile resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): NetworkProfile {
        return new NetworkProfile(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:devicefarm:NetworkProfile';

    /**
     * Returns true if the given object is an instance of NetworkProfile.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NetworkProfile {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NetworkProfile.__pulumiType;
    }

    public /*out*/ readonly arn!: pulumi.Output<string>;
    public readonly description!: pulumi.Output<string | undefined>;
    public readonly downlinkBandwidthBits!: pulumi.Output<number | undefined>;
    public readonly downlinkDelayMs!: pulumi.Output<number | undefined>;
    public readonly downlinkJitterMs!: pulumi.Output<number | undefined>;
    public readonly downlinkLossPercent!: pulumi.Output<number | undefined>;
    public readonly name!: pulumi.Output<string>;
    public readonly projectArn!: pulumi.Output<string>;
    public readonly tags!: pulumi.Output<outputs.devicefarm.NetworkProfileTag[] | undefined>;
    public readonly uplinkBandwidthBits!: pulumi.Output<number | undefined>;
    public readonly uplinkDelayMs!: pulumi.Output<number | undefined>;
    public readonly uplinkJitterMs!: pulumi.Output<number | undefined>;
    public readonly uplinkLossPercent!: pulumi.Output<number | undefined>;

    /**
     * Create a NetworkProfile resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NetworkProfileArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.projectArn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'projectArn'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["downlinkBandwidthBits"] = args ? args.downlinkBandwidthBits : undefined;
            resourceInputs["downlinkDelayMs"] = args ? args.downlinkDelayMs : undefined;
            resourceInputs["downlinkJitterMs"] = args ? args.downlinkJitterMs : undefined;
            resourceInputs["downlinkLossPercent"] = args ? args.downlinkLossPercent : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["projectArn"] = args ? args.projectArn : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["uplinkBandwidthBits"] = args ? args.uplinkBandwidthBits : undefined;
            resourceInputs["uplinkDelayMs"] = args ? args.uplinkDelayMs : undefined;
            resourceInputs["uplinkJitterMs"] = args ? args.uplinkJitterMs : undefined;
            resourceInputs["uplinkLossPercent"] = args ? args.uplinkLossPercent : undefined;
            resourceInputs["arn"] = undefined /*out*/;
        } else {
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["downlinkBandwidthBits"] = undefined /*out*/;
            resourceInputs["downlinkDelayMs"] = undefined /*out*/;
            resourceInputs["downlinkJitterMs"] = undefined /*out*/;
            resourceInputs["downlinkLossPercent"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["projectArn"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["uplinkBandwidthBits"] = undefined /*out*/;
            resourceInputs["uplinkDelayMs"] = undefined /*out*/;
            resourceInputs["uplinkJitterMs"] = undefined /*out*/;
            resourceInputs["uplinkLossPercent"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(NetworkProfile.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a NetworkProfile resource.
 */
export interface NetworkProfileArgs {
    description?: pulumi.Input<string>;
    downlinkBandwidthBits?: pulumi.Input<number>;
    downlinkDelayMs?: pulumi.Input<number>;
    downlinkJitterMs?: pulumi.Input<number>;
    downlinkLossPercent?: pulumi.Input<number>;
    name?: pulumi.Input<string>;
    projectArn: pulumi.Input<string>;
    tags?: pulumi.Input<pulumi.Input<inputs.devicefarm.NetworkProfileTagArgs>[]>;
    uplinkBandwidthBits?: pulumi.Input<number>;
    uplinkDelayMs?: pulumi.Input<number>;
    uplinkJitterMs?: pulumi.Input<number>;
    uplinkLossPercent?: pulumi.Input<number>;
}