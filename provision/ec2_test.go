package provision

import "testing"

func Test_parsePorts_empty(t *testing.T) {

	ports, err := parsePorts("")
	if err != nil {
		t.Fatal(err)
	}

	if len(ports) != 0 {
		t.Fatalf("Expected empty slice, got %d", len(ports))
	}
}

func Test_parsePorts_single(t *testing.T) {

	wantPort := 80
	str := "80"
	ports, err := parsePorts(str)
	if err != nil {
		t.Fatal(err)
	}

	if len(ports) != 1 {
		t.Fatalf("Want single port, got %d", len(ports))
	}

	if ports[0] != wantPort {
		t.Fatalf("Want port %d, got %d", wantPort, ports[0])
	}
}

func Test_parsePorts_multiple(t *testing.T) {

	wantPorts := []int{27017, 22}

	str := "27017,22"

	ports, err := parsePorts(str)
	if err != nil {
		t.Fatal(err)
	}

	if len(ports) != len(wantPorts) {
		t.Fatalf("Want %d ports, got %d", len(wantPorts), len(ports))
	}

	found := 0

	for _, port := range ports {
		for _, wantPort := range wantPorts {
			if port == wantPort {
				found++
			}
		}
	}

	if found != len(wantPorts) {
		t.Fatalf("Want %v ports, got %v", wantPorts, ports)
	}
}
