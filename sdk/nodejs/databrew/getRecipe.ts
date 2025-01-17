// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource schema for AWS::DataBrew::Recipe.
 */
export function getRecipe(args: GetRecipeArgs, opts?: pulumi.InvokeOptions): Promise<GetRecipeResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:databrew:getRecipe", {
        "name": args.name,
    }, opts);
}

export interface GetRecipeArgs {
    /**
     * Recipe name
     */
    name: string;
}

export interface GetRecipeResult {
    /**
     * Description of the recipe
     */
    readonly description?: string;
    readonly steps?: outputs.databrew.RecipeStep[];
}

export function getRecipeOutput(args: GetRecipeOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetRecipeResult> {
    return pulumi.output(args).apply(a => getRecipe(a, opts))
}

export interface GetRecipeOutputArgs {
    /**
     * Recipe name
     */
    name: pulumi.Input<string>;
}
