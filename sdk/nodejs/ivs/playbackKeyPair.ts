// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::IVS::PlaybackKeyPair
 */
export class PlaybackKeyPair extends pulumi.CustomResource {
    /**
     * Get an existing PlaybackKeyPair resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): PlaybackKeyPair {
        return new PlaybackKeyPair(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:ivs:PlaybackKeyPair';

    /**
     * Returns true if the given object is an instance of PlaybackKeyPair.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PlaybackKeyPair {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PlaybackKeyPair.__pulumiType;
    }

    /**
     * Key-pair identifier.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * Key-pair identifier.
     */
    public /*out*/ readonly fingerprint!: pulumi.Output<string>;
    /**
     * An arbitrary string (a nickname) assigned to a playback key pair that helps the customer identify that resource. The value does not need to be unique.
     */
    public readonly name!: pulumi.Output<string | undefined>;
    /**
     * The public portion of a customer-generated key pair.
     */
    public readonly publicKeyMaterial!: pulumi.Output<string>;
    /**
     * A list of key-value pairs that contain metadata for the asset model.
     */
    public readonly tags!: pulumi.Output<outputs.ivs.PlaybackKeyPairTag[] | undefined>;

    /**
     * Create a PlaybackKeyPair resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PlaybackKeyPairArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.publicKeyMaterial === undefined) && !opts.urn) {
                throw new Error("Missing required property 'publicKeyMaterial'");
            }
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["publicKeyMaterial"] = args ? args.publicKeyMaterial : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["fingerprint"] = undefined /*out*/;
        } else {
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["fingerprint"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["publicKeyMaterial"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(PlaybackKeyPair.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a PlaybackKeyPair resource.
 */
export interface PlaybackKeyPairArgs {
    /**
     * An arbitrary string (a nickname) assigned to a playback key pair that helps the customer identify that resource. The value does not need to be unique.
     */
    name?: pulumi.Input<string>;
    /**
     * The public portion of a customer-generated key pair.
     */
    publicKeyMaterial: pulumi.Input<string>;
    /**
     * A list of key-value pairs that contain metadata for the asset model.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.ivs.PlaybackKeyPairTagArgs>[]>;
}
