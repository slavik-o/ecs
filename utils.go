package ecs

// CreateComponentMask creates a component mask from a list of component IDs
func CreateComponentMask(components ...ComponentID) ComponentMask {
	mask := ComponentMask(0)

	for _, c := range components {
		mask = mask | (1 << c)
	}

	return mask
}
