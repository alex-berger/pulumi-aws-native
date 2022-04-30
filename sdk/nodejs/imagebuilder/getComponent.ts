// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource schema for AWS::ImageBuilder::Component
 */
export function getComponent(args: GetComponentArgs, opts?: pulumi.InvokeOptions): Promise<GetComponentResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:imagebuilder:getComponent", {
        "arn": args.arn,
    }, opts);
}

export interface GetComponentArgs {
    /**
     * The Amazon Resource Name (ARN) of the component.
     */
    arn: string;
}

export interface GetComponentResult {
    /**
     * The Amazon Resource Name (ARN) of the component.
     */
    readonly arn?: string;
    /**
     * The encryption status of the component.
     */
    readonly encrypted?: boolean;
    /**
     * The type of the component denotes whether the component is used to build the image or only to test it. 
     */
    readonly type?: enums.imagebuilder.ComponentType;
}

export function getComponentOutput(args: GetComponentOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetComponentResult> {
    return pulumi.output(args).apply(a => getComponent(a, opts))
}

export interface GetComponentOutputArgs {
    /**
     * The Amazon Resource Name (ARN) of the component.
     */
    arn: pulumi.Input<string>;
}
