// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Definition of AWS::AmplifyUIBuilder::Component Resource Type
 */
export function getComponent(args: GetComponentArgs, opts?: pulumi.InvokeOptions): Promise<GetComponentResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:amplifyuibuilder:getComponent", {
        "appId": args.appId,
        "environmentName": args.environmentName,
        "id": args.id,
    }, opts);
}

export interface GetComponentArgs {
    appId: string;
    environmentName: string;
    id: string;
}

export interface GetComponentResult {
    readonly appId?: string;
    readonly bindingProperties?: outputs.amplifyuibuilder.ComponentBindingProperties;
    readonly children?: outputs.amplifyuibuilder.ComponentChild[];
    readonly collectionProperties?: outputs.amplifyuibuilder.ComponentCollectionProperties;
    readonly componentType?: string;
    readonly environmentName?: string;
    readonly events?: outputs.amplifyuibuilder.ComponentEvents;
    readonly id?: string;
    readonly name?: string;
    readonly overrides?: outputs.amplifyuibuilder.ComponentOverrides;
    readonly properties?: outputs.amplifyuibuilder.ComponentProperties;
    readonly schemaVersion?: string;
    readonly sourceId?: string;
    readonly variants?: outputs.amplifyuibuilder.ComponentVariant[];
}

export function getComponentOutput(args: GetComponentOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetComponentResult> {
    return pulumi.output(args).apply(a => getComponent(a, opts))
}

export interface GetComponentOutputArgs {
    appId: pulumi.Input<string>;
    environmentName: pulumi.Input<string>;
    id: pulumi.Input<string>;
}
