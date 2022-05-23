// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./geofenceCollection";
export * from "./getGeofenceCollection";
export * from "./getMap";
export * from "./getPlaceIndex";
export * from "./getRouteCalculator";
export * from "./getTracker";
export * from "./map";
export * from "./placeIndex";
export * from "./routeCalculator";
export * from "./tracker";
export * from "./trackerConsumer";

// Export enums:
export * from "../types/enums/location";

// Import resources to register:
import { GeofenceCollection } from "./geofenceCollection";
import { Map } from "./map";
import { PlaceIndex } from "./placeIndex";
import { RouteCalculator } from "./routeCalculator";
import { Tracker } from "./tracker";
import { TrackerConsumer } from "./trackerConsumer";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "aws-native:location:GeofenceCollection":
                return new GeofenceCollection(name, <any>undefined, { urn })
            case "aws-native:location:Map":
                return new Map(name, <any>undefined, { urn })
            case "aws-native:location:PlaceIndex":
                return new PlaceIndex(name, <any>undefined, { urn })
            case "aws-native:location:RouteCalculator":
                return new RouteCalculator(name, <any>undefined, { urn })
            case "aws-native:location:Tracker":
                return new Tracker(name, <any>undefined, { urn })
            case "aws-native:location:TrackerConsumer":
                return new TrackerConsumer(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("aws-native", "location", _module)
