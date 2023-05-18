package utils

import (
	"encoding/json"
	"errors"
	"reflect"
)

func MatchObjects(queryData interface{}, items []interface{}, multiple bool) ([]interface{}, error) {
	queryDataJSON, err := json.Marshal(queryData)
	if err != nil {
		return nil, errors.New("Error marshaling queryData")
	}
	queryDataMap := make(map[string]interface{})
	if err := json.Unmarshal(queryDataJSON, &queryDataMap); err != nil {
		return nil, errors.New("Error unmarshaling queryData")
	}

	// Find the matching items
	matchingItems := make([]interface{}, 0)

	for _, item := range items {
		itemJSON, err := json.Marshal(item)
		if err != nil {
			return nil, errors.New("Error marshaling item")
		}

		var itemMap map[string]interface{}
		if err := json.Unmarshal(itemJSON, &itemMap); err != nil {
			return nil, errors.New("Error unmarshaling item")
		}

		// Check if all keys and values in queryDataMap are present in itemMap
		match := true
		for key, value := range queryDataMap {
			if itemMapValue, ok := itemMap[key]; !ok || !reflect.DeepEqual(value, itemMapValue) {
				match = false
				break
			}
		}

		if match {
			matchingItems = append(matchingItems, item)
			if !multiple {
				break
			}
		}
	}

	if len(matchingItems) > 0 {
		return matchingItems, nil
	} else {
		return nil, nil
	}
}
