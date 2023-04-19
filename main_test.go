package main

import (
	"reflect"
	"testing"
)

// Adding base tests as a starting point, More tests can be added to cover all the cases

func TestWorld_AddCity(t *testing.T) {
	type fields struct {
		Cities map[string]*City
		Aliens []*Alien
	}
	type args struct {
		name string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *City
	}{
		{
			name: "Test 1",
			fields: fields{
				Cities: map[string]*City{},
				Aliens: []*Alien{},
			},
			args: args{
				name: "City1",
			},
			want: &City{
				Name:   "City1",
				Aliens: make([]*Alien, 0),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &World{
				Cities: tt.fields.Cities,
				Aliens: tt.fields.Aliens,
			}
			if got := w.AddCity(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("World.AddCity() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWorld_AddAlien(t *testing.T) {
	type fields struct {
		Cities map[string]*City
		Aliens []*Alien
	}
	type args struct {
		a *Alien
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Test 1",
			fields: fields{
				Cities: map[string]*City{},
				Aliens: []*Alien{},
			},
			args: args{
				a: &Alien{
					ID: 1,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &World{
				Cities: tt.fields.Cities,
				Aliens: tt.fields.Aliens,
			}
			w.AddAlien(tt.args.a)
			if w.Aliens[0] != tt.args.a {
				t.Errorf("AddAlien() = %v, want %v", w.Aliens[0], tt.args.a)
			}
		})
	}
}

func TestWorld_AddRoad(t *testing.T) {
	type fields struct {
		Cities map[string]*City
		Aliens []*Alien
	}
	type args struct {
		from      string
		to        string
		direction string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Test 1",
			fields: fields{
				Cities: map[string]*City{
					"City1": {
						Name: "City1",
					},
					"City2": {
						Name: "City2",
					},
				},
				Aliens: []*Alien{},
			},
			args: args{
				from:      "City1",
				to:        "City2",
				direction: "north",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &World{
				Cities: tt.fields.Cities,
				Aliens: tt.fields.Aliens,
			}
			w.AddRoad(tt.args.from, tt.args.to, tt.args.direction)
			if w.Cities[tt.args.from].North != w.Cities[tt.args.to] {
				t.Errorf("AddRoad() = %v, want %v", w.Cities[tt.args.from].North, w.Cities[tt.args.to])
			}
		})
	}
}

func TestWorld_AddAlienToCity(t *testing.T) {
	type fields struct {
		Cities map[string]*City
		Aliens []*Alien
	}
	type args struct {
		a *Alien
		c *City
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Test 1",
			fields: fields{
				Cities: map[string]*City{
					"City1": {
						Name: "City1",
					},
				},
				Aliens: []*Alien{},
			},
			args: args{
				a: &Alien{
					ID: 1,
				},
				c: &City{
					Name: "City1",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &World{
				Cities: tt.fields.Cities,
				Aliens: tt.fields.Aliens,
			}
			w.AddCity(tt.args.c.Name)
			w.AddAlienToCity(tt.args.a, w.Cities[tt.args.c.Name])
			if w.Cities[tt.args.c.Name].Aliens[0] != tt.args.a {
				t.Errorf("AddAlienToCity() = %v, want %v", w.Cities[tt.args.c.Name].Aliens[0], tt.args.a)
			}
		})
	}
}

func TestWorld_RemoveAlienFromCity(t *testing.T) {
	type fields struct {
		Cities map[string]*City
		Aliens []*Alien
	}
	type args struct {
		a *Alien
		c *City
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Test 1",
			fields: fields{
				Cities: map[string]*City{
					"City1": {
						Name: "City1",
					},
				},
				Aliens: []*Alien{},
			},
			args: args{
				a: &Alien{
					ID: 1,
				},
				c: &City{
					Name: "City1",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &World{
				Cities: tt.fields.Cities,
				Aliens: tt.fields.Aliens,
			}
			w.AddAlienToCity(tt.args.a, tt.args.c)
			w.RemoveAlienFromCity(tt.args.a, tt.args.c)
			if len(w.Cities[tt.args.c.Name].Aliens) != 0 {
				t.Errorf("RemoveAlienFromCity() = %v, want %v", w.Cities[tt.args.c.Name].Aliens, 0)
			}
		})
	}
}

func TestWorld_RemoveCity(t *testing.T) {
	type fields struct {
		Cities map[string]*City
		Aliens []*Alien
	}
	type args struct {
		c *City
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Test 1",
			fields: fields{
				Cities: map[string]*City{
					"City1": {
						Name: "City1",
					},
				},
				Aliens: []*Alien{},
			},
			args: args{
				c: &City{
					Name: "City1",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &World{
				Cities: tt.fields.Cities,
				Aliens: tt.fields.Aliens,
			}
			w.RemoveCity(tt.args.c)
			if _, ok := w.Cities[tt.args.c.Name]; ok {
				t.Errorf("RemoveCity() = %v, want %v", w.Cities[tt.args.c.Name], nil)
			}
		})
	}
}

func TestWorld_RemoveAlien(t *testing.T) {
	type fields struct {
		Cities map[string]*City
		Aliens []*Alien
	}
	type args struct {
		a *Alien
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Test 1",
			fields: fields{
				Cities: map[string]*City{
					"City1": {
						Name: "City1",
					},
				},
				Aliens: []*Alien{},
			},
			args: args{
				a: &Alien{
					ID: 1,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &World{
				Cities: tt.fields.Cities,
				Aliens: tt.fields.Aliens,
			}
			w.AddAlien(tt.args.a)
			w.RemoveAlien(tt.args.a)
			if len(w.Aliens) != 0 {
				t.Errorf("RemoveAlien() = %v, want %v", w.Aliens, 0)
			}
		})
	}
}
