package compute

import (
	"context"

	googleCompute "google.golang.org/api/compute/v1"
)

type ComputeVMs struct {
  VMs *googleCompute.InstanceList
}

// GetVMs: Simplified Function to Get VMs
// TODO(mccallry) resuse Context and Services
func GetVMs(project string, zone string) (ComputeVMs, error) {

	ctx := context.Background()
	computeService, _ := googleCompute.NewService(ctx)
	listClient := computeService.Instances.List(project, zone)
	vms, err := listClient.Do()

  result := ComputeVMs{
    VMs: vms,
  }
  return result, err
}
