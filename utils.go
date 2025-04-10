package ecs

// CreateComponentMask creates a component mask from a list of component IDs
func CreateComponentMask(components ...ComponentID) ComponentMask {
	// Initialize mask
	mask := ComponentMask(0)

	// Iterate over all components
	for _, c := range components {
		// Set bit for component
		mask = mask | (1 << c)
	}

	return mask
}
