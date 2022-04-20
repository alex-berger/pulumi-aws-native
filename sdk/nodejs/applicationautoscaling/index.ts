// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./getScalableTarget";
export * from "./getScalingPolicy";
export * from "./scalableTarget";
export * from "./scalingPolicy";

// Import resources to register:
import { ScalableTarget } from "./scalableTarget";
import { ScalingPolicy } from "./scalingPolicy";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "aws-native:applicationautoscaling:ScalableTarget":
                return new ScalableTarget(name, <any>undefined, { urn })
            case "aws-native:applicationautoscaling:ScalingPolicy":
                return new ScalingPolicy(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("aws-native", "applicationautoscaling", _module)
