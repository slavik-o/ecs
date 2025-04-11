package shared

import "github.com/slavik-o/ecs"

// DamageSystem handles applying damage when collisions occur
type DamageSystem struct {
	world *ecs.World
}

func NewDamageSystem(world *ecs.World) *DamageSystem {
	system := &DamageSystem{world: world}

	// Subscribe to collision events
	world.EventManager.Subscribe(EVENT_COLLISION, system.handleCollision)

	return system
}

func (s *DamageSystem) ComponentMask() ecs.ComponentMask {
	return 0 // This system doesn't process entities directly
}

func (s *DamageSystem) Update(dt float32, world *ecs.World) {
	// No direct entity processing - this system works via event handling
}

func (s *DamageSystem) handleCollision(event ecs.Event) {
	collisionEvent := event.(*CollisionEvent)

	// Check if entity1 has health
	if s.world.HasComponent(collisionEvent.Entity1, COMPONENT_HEALTH) {
		health := s.world.GetComponent(collisionEvent.Entity1, COMPONENT_HEALTH).(*Health)

		previousHealth := health.Current

		// Apply some damage
		health.Current -= 10
		if health.Current < 0 {
			health.Current = 0
		}

		// Publish health changed event
		s.world.EventManager.Publish(&HealthChangedEvent{
			Entity:         collisionEvent.Entity1,
			PreviousHealth: previousHealth,
			NewHealth:      health.Current,
		})

		// If health reached zero, publish entity died event
		if health.Current == 0 {
			s.world.EventManager.Publish(&EntityDiedEvent{
				Entity: collisionEvent.Entity1,
			})
		}
	}

	// Similarly for entity2
	if s.world.HasComponent(collisionEvent.Entity2, COMPONENT_HEALTH) {
		health := s.world.GetComponent(collisionEvent.Entity2, COMPONENT_HEALTH).(*Health)
		previousHealth := health.Current

		health.Current -= 10
		if health.Current < 0 {
			health.Current = 0
		}

		s.world.EventManager.Publish(&HealthChangedEvent{
			Entity:         collisionEvent.Entity2,
			PreviousHealth: previousHealth,
			NewHealth:      health.Current,
		})

		if health.Current == 0 {
			s.world.EventManager.Publish(&EntityDiedEvent{
				Entity: collisionEvent.Entity2,
			})
		}
	}
}
