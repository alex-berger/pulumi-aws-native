package schema

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
)

func getItemTypeReplaces(schema CloudFormationSchema, itemType, resourceType, path string, diff resource.ValueDiff) ([]string, error) {
	propertyTypeName := resourceType + "." + itemType
	propertyTypeSpec, ok := schema.PropertyTypes[propertyTypeName]
	if !ok {
		propertyTypeSpec, ok = schema.PropertyTypes[itemType]
		if !ok {
			return nil, errors.Errorf("could not find property type %v in schema", propertyTypeName)
		}
	}
	return getPropertiesReplaces(schema, propertyTypeSpec, resourceType, path, diff.Object)
}

func getPropertyReplaces(schema CloudFormationSchema, spec PropertySpec, resourceType, path string, diff resource.ValueDiff) ([]string, error) {
	// If this property is immutable, it will force a replacement. We can stop here.
	if spec.UpdateType == "Immutable" {
		return []string{path}, nil
	}

	// Check the type of the property.
	if spec.PrimitiveType != "" {
		// Nothing to do.
		return nil, nil
	}
	switch spec.Type {
	case "List":
		if diff.Array == nil || spec.PrimitiveItemType != "" {
			return nil, nil
		}
		var paths []string
		for i, itemDiff := range diff.Array.Updates {
			itemPath := fmt.Sprintf("%s[%d]", path, i)
			ps, err := getItemTypeReplaces(schema, spec.ItemType, resourceType, itemPath, itemDiff)
			if err != nil {
				return nil, err
			}
			paths = append(paths, ps...)
		}
		return paths, nil
	case "Map":
		if diff.Object == nil || spec.PrimitiveItemType != "" {
			return nil, nil
		}
		var paths []string
		for k, itemDiff := range diff.Object.Updates {
			itemPath := fmt.Sprintf("%s.%s", path, k)
			ps, err := getItemTypeReplaces(schema, spec.ItemType, resourceType, itemPath, itemDiff)
			if err != nil {
				return nil, err
			}
			paths = append(paths, ps...)
		}
		return paths, nil
	case "":
		return nil, errors.Errorf("schema is missing type information for %v", path)
	default:
		return getItemTypeReplaces(schema, spec.Type, resourceType, path, diff)
	}
}

func getPropertiesReplaces(schema CloudFormationSchema, spec PropertyTypeSpec, resourceType, path string, diff *resource.ObjectDiff) ([]string, error) {
	if diff == nil {
		return nil, nil
	}

	var paths []string
	for n, propertySpec := range spec.Properties {
		sdkName := ToPropertyName(n)
		if itemDiff, ok := diff.Updates[resource.PropertyKey(sdkName)]; ok {
			var propertyPath string
			if path == "" {
				propertyPath = sdkName
			} else {
				propertyPath = fmt.Sprintf("%s.%s", path, sdkName)
			}
			ps, err := getPropertyReplaces(schema, propertySpec, resourceType, propertyPath, itemDiff)
			if err != nil {
				return nil, err
			}
			paths = append(paths, ps...)
		}
	}
	return paths, nil
}

func GetResourceReplaces(schema CloudFormationSchema, resourceType string, diff *resource.ObjectDiff) ([]string, error) {
	// Fetch the schema for the resource type.
	resourceSpec, ok := schema.ResourceTypes[resourceType]
	if !ok {
		return nil, errors.Errorf("unknown resource type %v", resourceType)
	}
	return getPropertiesReplaces(schema, resourceSpec.PropertyTypeSpec, resourceType, "", diff)
}
