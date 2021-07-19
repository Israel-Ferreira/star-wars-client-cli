package models



type Planet struct {
	Name string `json:"name"`
	RotationalPeriod rune `json:"rotational_period"`
	OrbitalPeriod rune `json:"orbital_period"`
	Climate string `json:"climate"`
	Gravity string `json:"gravity"`
	Terrain string `json:"terrain"`
}




type PlanetResult struct {
	Result
	Planets []Planet `json:"results"`
}