// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::AppStream::AppBlock
 */
export class AppBlock extends pulumi.CustomResource {
    /**
     * Get an existing AppBlock resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): AppBlock {
        return new AppBlock(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:appstream:AppBlock';

    /**
     * Returns true if the given object is an instance of AppBlock.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AppBlock {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AppBlock.__pulumiType;
    }

    public /*out*/ readonly arn!: pulumi.Output<string>;
    public /*out*/ readonly createdTime!: pulumi.Output<string>;
    public readonly description!: pulumi.Output<string | undefined>;
    public readonly displayName!: pulumi.Output<string | undefined>;
    public readonly name!: pulumi.Output<string>;
    public readonly setupScriptDetails!: pulumi.Output<outputs.appstream.AppBlockScriptDetails>;
    public readonly sourceS3Location!: pulumi.Output<outputs.appstream.AppBlockS3Location>;
    public readonly tags!: pulumi.Output<outputs.appstream.AppBlockTag[] | undefined>;

    /**
     * Create a AppBlock resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AppBlockArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.setupScriptDetails === undefined) && !opts.urn) {
                throw new Error("Missing required property 'setupScriptDetails'");
            }
            if ((!args || args.sourceS3Location === undefined) && !opts.urn) {
                throw new Error("Missing required property 'sourceS3Location'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["setupScriptDetails"] = args ? args.setupScriptDetails : undefined;
            resourceInputs["sourceS3Location"] = args ? args.sourceS3Location : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["createdTime"] = undefined /*out*/;
        } else {
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["createdTime"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["displayName"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["setupScriptDetails"] = undefined /*out*/;
            resourceInputs["sourceS3Location"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AppBlock.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a AppBlock resource.
 */
export interface AppBlockArgs {
    description?: pulumi.Input<string>;
    displayName?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    setupScriptDetails: pulumi.Input<inputs.appstream.AppBlockScriptDetailsArgs>;
    sourceS3Location: pulumi.Input<inputs.appstream.AppBlockS3LocationArgs>;
    tags?: pulumi.Input<pulumi.Input<inputs.appstream.AppBlockTagArgs>[]>;
}
