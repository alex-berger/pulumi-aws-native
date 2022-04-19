// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./getJobTemplate";
export * from "./getPreset";
export * from "./getQueue";
export * from "./jobTemplate";
export * from "./preset";
export * from "./queue";

// Import resources to register:
import { JobTemplate } from "./jobTemplate";
import { Preset } from "./preset";
import { Queue } from "./queue";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "aws-native:mediaconvert:JobTemplate":
                return new JobTemplate(name, <any>undefined, { urn })
            case "aws-native:mediaconvert:Preset":
                return new Preset(name, <any>undefined, { urn })
            case "aws-native:mediaconvert:Queue":
                return new Queue(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("aws-native", "mediaconvert", _module)
