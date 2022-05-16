// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./getLaunchProfile";
export * from "./getStreamingImage";
export * from "./getStudio";
export * from "./getStudioComponent";
export * from "./launchProfile";
export * from "./streamingImage";
export * from "./studio";
export * from "./studioComponent";

// Export enums:
export * from "../types/enums/nimblestudio";

// Import resources to register:
import { LaunchProfile } from "./launchProfile";
import { StreamingImage } from "./streamingImage";
import { Studio } from "./studio";
import { StudioComponent } from "./studioComponent";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "aws-native:nimblestudio:LaunchProfile":
                return new LaunchProfile(name, <any>undefined, { urn })
            case "aws-native:nimblestudio:StreamingImage":
                return new StreamingImage(name, <any>undefined, { urn })
            case "aws-native:nimblestudio:Studio":
                return new Studio(name, <any>undefined, { urn })
            case "aws-native:nimblestudio:StudioComponent":
                return new StudioComponent(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("aws-native", "nimblestudio", _module)