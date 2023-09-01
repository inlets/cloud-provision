package provision

import "testing"

func TestCustomGCEIDConstAndDest(t *testing.T) {
	inputInstanceName := "inlets"
	inputZone := "us-central1-a"
	inputProjectID := "playground"
	inputRegion := "us-central"

	customID := toGCEID(inputInstanceName, inputZone, inputProjectID, inputRegion)

	outputInstanceName, outputZone, outputProjectID, outputRegion, err := getGCEFieldsFromID(customID)
	if err != nil {
		t.Error(err)
	}
	if inputInstanceName != outputInstanceName ||
		inputZone != outputZone ||
		inputProjectID != outputProjectID ||
		inputRegion != outputRegion {
		t.Errorf("Input fields: %s, %s, %s are not equal to the ouput fields: %s, %s, %s",
			inputInstanceName, inputZone, inputProjectID, outputInstanceName, outputZone, outputProjectID)
	}

}
