// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Schema for Package CloudFormation Resource
 */
export class Package extends pulumi.CustomResource {
    /**
     * Get an existing Package resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Package {
        return new Package(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:panorama:Package';

    /**
     * Returns true if the given object is an instance of Package.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Package {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Package.__pulumiType;
    }

    public /*out*/ readonly arn!: pulumi.Output<string>;
    public /*out*/ readonly createdTime!: pulumi.Output<number>;
    public /*out*/ readonly packageId!: pulumi.Output<string>;
    public readonly packageName!: pulumi.Output<string>;
    public /*out*/ readonly storageLocation!: pulumi.Output<outputs.panorama.PackageStorageLocation>;
    public readonly tags!: pulumi.Output<outputs.panorama.PackageTag[] | undefined>;

    /**
     * Create a Package resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: PackageArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["packageName"] = args ? args.packageName : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["createdTime"] = undefined /*out*/;
            resourceInputs["packageId"] = undefined /*out*/;
            resourceInputs["storageLocation"] = undefined /*out*/;
        } else {
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["createdTime"] = undefined /*out*/;
            resourceInputs["packageId"] = undefined /*out*/;
            resourceInputs["packageName"] = undefined /*out*/;
            resourceInputs["storageLocation"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Package.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Package resource.
 */
export interface PackageArgs {
    packageName?: pulumi.Input<string>;
    tags?: pulumi.Input<pulumi.Input<inputs.panorama.PackageTagArgs>[]>;
}
