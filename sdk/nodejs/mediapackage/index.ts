// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./asset";
export * from "./channel";
export * from "./getAsset";
export * from "./getChannel";
export * from "./getOriginEndpoint";
export * from "./getPackagingConfiguration";
export * from "./getPackagingGroup";
export * from "./originEndpoint";
export * from "./packagingConfiguration";
export * from "./packagingGroup";

// Export enums:
export * from "../types/enums/mediapackage";

// Import resources to register:
import { Asset } from "./asset";
import { Channel } from "./channel";
import { OriginEndpoint } from "./originEndpoint";
import { PackagingConfiguration } from "./packagingConfiguration";
import { PackagingGroup } from "./packagingGroup";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "aws-native:mediapackage:Asset":
                return new Asset(name, <any>undefined, { urn })
            case "aws-native:mediapackage:Channel":
                return new Channel(name, <any>undefined, { urn })
            case "aws-native:mediapackage:OriginEndpoint":
                return new OriginEndpoint(name, <any>undefined, { urn })
            case "aws-native:mediapackage:PackagingConfiguration":
                return new PackagingConfiguration(name, <any>undefined, { urn })
            case "aws-native:mediapackage:PackagingGroup":
                return new PackagingGroup(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("aws-native", "mediapackage", _module)