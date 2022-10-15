//--Summary:
//  Copy your rcv-func solution to this directory and write unit tests.
//
//--Requirements:
//* Write unit tests that ensure:
//  - Health & energy can not go above their maximums
//  - Health & energy can not go below 0
//* If any of your  tests fail, make the necessary corrections
//  in the copy of your rcv-func solution file.
//
//--Notes:
//* Use `go test -v ./exercise/testing` to run these specific tests

package main

import "testing"

func newPlayer() Player {
	return Player{name: "test", health: 100, maxHealth: 100, energy: 500, maxEnergy: 500}
}

func TestHealth(t *testing.T) {
	player := newPlayer()
	player.addHealth(999)

	if player.health > player.maxHealth {
		/* abort the testing because next part of the code will be invalid,if this particular test fails(there's no point in continuing the tests if this one fails) */
		t.Fatalf("Health went over limit: %v, want: %v", player.health, player.maxHealth)
	}

	player.applyDamage(player.maxHealth + 1)
	if player.health < 0 {
		t.Fatalf("Health: %v. Minimum: 0", player.health)
	}

	// test the overflow behavior of uint(4 billion is the overflowed value)
	if player.health > player.maxHealth {
		t.Fatalf("Health: %v. Maximum: %v", player.health, player.maxHealth)
	}
}

func TestEnergy(t *testing.T) {
	player := newPlayer()
	player.addEnergy(999)

	if player.energy > player.maxEnergy {
		t.Fatalf("Energy went over limit: %v, want: %v", player.energy, player.maxEnergy)
	}

	player.consumeEnergy(player.maxEnergy + 1)
	if player.energy < 0 {
		t.Fatalf("Energy: %v. Minimum: 0", player.energy)
	}

	if player.energy > player.maxEnergy {
		t.Fatalf("Energy: %v. Maximum: %v", player.energy, player.maxEnergy)
	}
}
