package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

// City is a city in the world
type City struct {
	Name   string
	North  *City
	South  *City
	East   *City
	West   *City
	Aliens []*Alien
	InMap  bool
}

// Alien is an alien in the world
type Alien struct {
	ID       int
	Location *City
}

// World is the world
type World struct {
	Cities map[string]*City
	Aliens []*Alien
}

// NewWorld creates a new world
func NewWorld() *World {
	return &World{
		Cities: make(map[string]*City),
	}
}

// NewCity creates a new city
func NewCity(name string) *City {
	return &City{
		Name:   name,
		Aliens: make([]*Alien, 0),
	}
}

// NewAlien creates a new alien
func NewAlien(id int) *Alien {
	return &Alien{
		ID: id,
	}
}

// AddCity adds a city to the world
func (w *World) AddCity(name string) *City {
	c := NewCity(name)
	w.Cities[name] = c
	return c
}

// AddAlien adds an alien to the world
func (w *World) AddAlien(a *Alien) {
	w.Aliens = append(w.Aliens, a)
}

// AddRoad adds a road between two cities
func (w *World) AddRoad(from, to string, direction string) {
	fromCity := w.Cities[from]
	toCity := w.Cities[to]
	if toCity == nil {
		toCity = w.AddCity(to)
	}
	switch direction {
	case "north":
		fromCity.North = toCity
		toCity.South = fromCity
	case "south":
		fromCity.South = toCity
		toCity.North = fromCity
	case "east":
		fromCity.East = toCity
		toCity.West = fromCity
	case "west":
		fromCity.West = toCity
		toCity.East = fromCity
	}
}

// AddAlienToCity adds an alien to a city
func (w *World) AddAlienToCity(a *Alien, c *City) {
	c.Aliens = append(c.Aliens, a)
	a.Location = c
}

// RemoveAlienFromCity removes an alien from a city
func (w *World) RemoveAlienFromCity(a *Alien, c *City) {
	for i, alien := range c.Aliens {
		if alien == a {
			c.Aliens = append(c.Aliens[:i], c.Aliens[i+1:]...)
			break
		}
	}
}

// RemoveCity removes a city from the world
func (w *World) RemoveCity(c *City) {
	cityName := c.Name
	delete(w.Cities, c.Name)
	for _, city := range w.Cities {
		if city.North != nil && city.North.Name == cityName {
			city.North = nil
		}
		if city.South != nil && city.South.Name == cityName {
			city.South = nil
		}
		if city.East != nil && city.East.Name == cityName {
			city.East = nil
		}
		if city.West != nil && city.West.Name == cityName {
			city.West = nil
		}
	}
}

// GetRandomCity gets a random city from the world
func (w *World) GetRandomCity() *City {
	totalCities := len(w.Cities)
	randomCity := rand.Intn(totalCities)
	i := 0
	for _, c := range w.Cities {
		if i == randomCity {
			return c
		}
		i++
	}
	return nil
}

// RemoveAlien removes an alien from the world
func (w *World) RemoveAlien(a *Alien) {
	for i, alien := range w.Aliens {
		if alien == a {
			w.Aliens = append(w.Aliens[:i], w.Aliens[i+1:]...)
			break
		}
	}
}

// MoveAlien moves an alien to a new city
func (w *World) MoveAlien(a *Alien, c *City) {
	w.RemoveAlienFromCity(a, a.Location)
	w.AddAlienToCity(a, c)
}

// FightAliens fights aliens in a city
func (w *World) FightAliens(c *City) {
	if len(c.Aliens) > 1 {
		message := fmt.Sprintf("%s has been destroyed by", c.Name)
		// There could be more than 2 aliens in a city, So we are going to kill them all as of now, Can change this logic to kill only 2 aliens
		for _, a := range c.Aliens {
			w.RemoveAlien(a)
			message += fmt.Sprintf(" alien %d", a.ID)
		}
		message += "!"
		log.Println(message)
		w.RemoveCity(c)
	}
}

// MoveAliens moves aliens in the world
func (w *World) MoveAliens() {
	for _, a := range w.Aliens {

		// Check for alien trapped in a city
		if a.Location == nil || a.Location.North == nil && a.Location.South == nil && a.Location.East == nil && a.Location.West == nil {
			continue
		}

		var c *City
		for c == nil {
			switch rand.Intn(4) {
			case 0:
				c = a.Location.North
			case 1:
				c = a.Location.South
			case 2:
				c = a.Location.East
			case 3:
				c = a.Location.West
			}
		}
		w.MoveAlien(a, c)
	}
}

// FightAliensInCities fights aliens in cities
func (w *World) FightAliensInCities() {
	for _, c := range w.Cities {
		w.FightAliens(c)
	}
}

// Run runs the world
func (w *World) Run() {
	// Run the simulation 10000 times
	for i := 0; i < 10000; i++ {
		// Need atleast 2 aliens to fight
		if len(w.Aliens) < 2 {
			break
		}
		w.MoveAliens()
		w.FightAliensInCities()
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	w := NewWorld()
	file, err := os.Open("city_map.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		c := w.AddCity(parts[0])
		c.InMap = true
		for _, part := range parts[1:] {
			parts2 := strings.Split(part, "=")
			w.AddRoad(parts[0], parts2[1], parts2[0])
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	args := os.Args
	if len(args) < 2 {
		log.Fatal("Please provide the number of aliens")
	}
	totalAliens, err := strconv.Atoi(args[1])
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < totalAliens; i++ {
		a := NewAlien(i)
		w.AddAlien(a)
		w.AddAlienToCity(a, w.GetRandomCity())
	}

	w.Run()

	for _, c := range w.Cities {
		if !c.InMap {
			continue
		}
		fmt.Printf("%s", c.Name)
		if c.East != nil {
			fmt.Printf(" east=%s", c.East.Name)
		}
		if c.West != nil {
			fmt.Printf(" west=%s", c.West.Name)
		}
		if c.North != nil {
			fmt.Printf(" north=%s", c.North.Name)
		}
		if c.South != nil {
			fmt.Printf(" south=%s", c.South.Name)
		}
		fmt.Println()
	}
}
