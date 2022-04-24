// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource for Greengrass component version.
 */
export class ComponentVersion extends pulumi.CustomResource {
    /**
     * Get an existing ComponentVersion resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ComponentVersion {
        return new ComponentVersion(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:greengrassv2:ComponentVersion';

    /**
     * Returns true if the given object is an instance of ComponentVersion.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ComponentVersion {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ComponentVersion.__pulumiType;
    }

    public /*out*/ readonly arn!: pulumi.Output<string>;
    public /*out*/ readonly componentName!: pulumi.Output<string>;
    public /*out*/ readonly componentVersion!: pulumi.Output<string>;
    public readonly inlineRecipe!: pulumi.Output<string | undefined>;
    public readonly lambdaFunction!: pulumi.Output<outputs.greengrassv2.ComponentVersionLambdaFunctionRecipeSource | undefined>;
    public readonly tags!: pulumi.Output<any | undefined>;

    /**
     * Create a ComponentVersion resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ComponentVersionArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["inlineRecipe"] = args ? args.inlineRecipe : undefined;
            resourceInputs["lambdaFunction"] = args ? args.lambdaFunction : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["componentName"] = undefined /*out*/;
            resourceInputs["componentVersion"] = undefined /*out*/;
        } else {
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["componentName"] = undefined /*out*/;
            resourceInputs["componentVersion"] = undefined /*out*/;
            resourceInputs["inlineRecipe"] = undefined /*out*/;
            resourceInputs["lambdaFunction"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ComponentVersion.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a ComponentVersion resource.
 */
export interface ComponentVersionArgs {
    inlineRecipe?: pulumi.Input<string>;
    lambdaFunction?: pulumi.Input<inputs.greengrassv2.ComponentVersionLambdaFunctionRecipeSourceArgs>;
    tags?: any;
}
