// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::AppStream::Entitlement
 */
export class Entitlement extends pulumi.CustomResource {
    /**
     * Get an existing Entitlement resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Entitlement {
        return new Entitlement(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:appstream:Entitlement';

    /**
     * Returns true if the given object is an instance of Entitlement.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Entitlement {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Entitlement.__pulumiType;
    }

    public readonly appVisibility!: pulumi.Output<string>;
    public readonly attributes!: pulumi.Output<outputs.appstream.EntitlementAttribute[]>;
    public /*out*/ readonly createdTime!: pulumi.Output<string>;
    public readonly description!: pulumi.Output<string | undefined>;
    public /*out*/ readonly lastModifiedTime!: pulumi.Output<string>;
    public readonly name!: pulumi.Output<string>;
    public readonly stackName!: pulumi.Output<string>;

    /**
     * Create a Entitlement resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: EntitlementArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.appVisibility === undefined) && !opts.urn) {
                throw new Error("Missing required property 'appVisibility'");
            }
            if ((!args || args.attributes === undefined) && !opts.urn) {
                throw new Error("Missing required property 'attributes'");
            }
            if ((!args || args.stackName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'stackName'");
            }
            resourceInputs["appVisibility"] = args ? args.appVisibility : undefined;
            resourceInputs["attributes"] = args ? args.attributes : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["stackName"] = args ? args.stackName : undefined;
            resourceInputs["createdTime"] = undefined /*out*/;
            resourceInputs["lastModifiedTime"] = undefined /*out*/;
        } else {
            resourceInputs["appVisibility"] = undefined /*out*/;
            resourceInputs["attributes"] = undefined /*out*/;
            resourceInputs["createdTime"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["lastModifiedTime"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["stackName"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Entitlement.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Entitlement resource.
 */
export interface EntitlementArgs {
    appVisibility: pulumi.Input<string>;
    attributes: pulumi.Input<pulumi.Input<inputs.appstream.EntitlementAttributeArgs>[]>;
    description?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    stackName: pulumi.Input<string>;
}
