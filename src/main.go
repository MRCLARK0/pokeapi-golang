package main

import (
	"fmt"
	"reflect"
	// resty "github.com/go-resty/resty/v2"
)

type Ability struct {
	EffectChanges []struct {
		EffectEntries []struct {
			Effect string `json:"effect"`
			Language struct {
				Name string `json:"name"`
				URL string `json:"url"`
			} `json:"language"`
		} `json:"effect_entries"`
		VersionGroup struct {
			Name string `json:"name"`
			URL string `json:"url"`
		} `json:"version_group"`
	} `json:"effect_changes"`
	EffectEntries []struct {
		Effect string `json:"effect"`
		Language struct {
			Name string `json:"name"`
			URL string `json:"url"`
		} `json:"language"`
		ShortEffect string `json:"short_effect"`
	} `json:"effect_entries"`
	FlavorTextEntries []struct {
		FlavorText string `json:"flavor_text"`
		Language struct {
			Name string `json:"name"`
			URL string `json:"url"`
		} `json:"language"`
		VersionGroup struct {
			Name string `json:"name"`
			URL string `json:"url"`
		} `json:"version_group"`
	} `json:"flavor_text_entries"`
	Generation struct {
		Name string `json:"name"`
		URL string `json:"url"`
	} `json:"generation"`
	ID int `json:"id"`
	IsMainSeries bool `json:"is_main_series"`
	Name string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	Pokemon []struct {
		IsHidden bool `json:"is_hidden"`
		Pokemon struct {
			Name string `json:"name"`
			URL string `json:"url"`
		} `json:"pokemon"`
		Slot int `json:"slot"`
	} `json:"pokemon"`
}

type Characteristic struct {
	Descriptions []struct {
		Description string `json:"description"`
		Language struct {
			Name string `json:"name"`
			URL string `json:"url"`
		} `json:"language"`
	} `json:"descriptions"`
	GeneModulo int `json:"gene_modulo"`
	HighestStat struct {
		Name string `json:"name"`
		URL string `json:"url"`
	} `json:"highest_stat"`
	ID int `json:"id"`
	PossibleValues []int `json:"possible_values"`
}

type EggGroup struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonSpecies []struct {
		Name string `json:"name"`
		URL string `json:"url"`
	} `json:"pokemon_species"`
}

type Gender struct {
	ID int `json:"id"`
	Name string `json:"name"`
	PokemonSpeciesDetails []struct {
		PokemonSpecies struct {
			Name string `json:"name"`
			URL string `json:"url"`
		} `json:"pokemon_species"`
		Rate int `json:"rate"`
	} `json:"pokemon_species_details"`
	RequiredForEvolution []struct {
		Name string `json:"name"`
		URL string `json:"url"`
	} `json:"required_for_evolution"`
}

type GrowthRate struct {
	Descriptions []struct {
		Description string `json:"description"`
		Language struct {
			Name string `json:"name"`
			URL string `json:"url"`
		} `json:"language"`
	} `json:"descriptions"`
	Formula string `json:"formula"`
	ID int `json:"id"`
	Levels []struct {
		Experience int `json:"experience"`
		Level int `json:"level"`
	} `json:"levels"`
	Name string `json:"name"`
	PokemonSpecies []struct {
		Name string `json:"name"`
		URL string `json:"url"`
	} `json:"pokemon_species"`
}

type Nature struct {
	DecreasedStat interface{} `json:"decreased_stat"`
	HatesFlavor interface{} `json:"hates_flavor"`
	ID int `json:"id"`
	LikesFlavor interface{} `json:"likes_flavor"`
	MoveBattleStylePreferences []struct {
		HighHpPreference int `json:"high_hp_preference"`
		LowHpPreference int `json:"low_hp_preference"`
		MoveBattleStyle struct {
			Name string `json:"name"`
			URL string `json:"url"`
		} `json:"move_battle_style"`
	} `json:"move_battle_style_preferences"`
	Name string `json:"name"`
	Names []struct {
		Langauge struct {
			Name string `json:"name"`
			URL string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokeathlonStatChanges []struct {
		MaxChange int `json:"max_change"`
		PokeathlonStat struct {
			Name string `json:"name"`
			URL string `json:"url"`
		} `json:"pokeathlon_stat"`
	} `json:"pokeathlon_stat_changes"`
}

type PokeathlonStat struct {
	AffectingNatures struct {
		Decrease []struct {
			MaxChange int `json:"max_change"`
			Nature struct {
				Name string `json:"name"`
				URL string `json:"url"`
			} `json:"nature"`
		} `json:"decrease"`
		Increase []struct {
			MaxChange int `json:"max_change"`
			Nature struct {
				Name string `json:"name"`
				URL string `json:"url"`
			} `json:"nature"`
		} `json:"increase"`
	} `json:"affecting_natures"`
	ID int `json:"id"`
	Name string `json:"name"`
	Names []struct {
		Langauge struct {
			Name string `json:"name"`
			URL string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
}

type Pokemon struct {
	Abilities []struct {
		Ability struct {
			Name string `json:"name"`
			URL string `json:"url"`
		} `json:"ability"`
		IsHidden bool `json:"is_hidden"`
		Slot int `json:"slot"`
	} `json:"abilities"`
	BaseExperience int `json:"base_experience"`
	Forms []struct {
		Name string `json:"name"`
		URL string `json:"url"`
	} `json:"forms"`
	GameIndices []struct {
		GameIndex int `json:"game_index"`
		Version struct {
			Name string `json:"name"`
			URL string `json:"url"`
		} `json:"version"`
	} `json:"game_indices"`
	Height int `json:"height"`
	HeldItems []interface{} `json:"held_items"`
	ID int `json:"id"`
	IsDefault bool `json:"is_default"`
	LocationAreaEncounters string `json:"location_area_encounters"`
	Moves []struct {
		Move struct {
			Name string `json:"name"`
			URL string `json:"url"`
		} `json:"move"`
		VersionGroupDetails []struct {
			LevelLearnedAt int `json:"level_learned_at"`
			MoveLearnMethod struct {
				Name string `json:"name"`
				URL string `json:"url"`
			} `json:"move_learn_method"`
			VersionGroup struct {
				Name string `json:"name"`
				URL string `json:"url"`
			} `json:"version_group"`
		} `json:"version_group_details"`
	} `json:"moves"`
	Name string `json:"name"`
	Order int `json:"order"`
	Species struct {
		Name string `json:"name"`
		URL string `json:"url"`
	} `json:"species"`
	Sprites struct {
		BackDefault string `json:"back_default"`
		BackFemale interface{} `json:"back_female"`
		BackShiny string `json:"back_shiny"`
		BackShinyFemale interface{} `json:"back_shiny_female"`
		FrontDefault string `json:"front_default"`
		FrontFemale interface{} `json:"front_female"`
		FrontShiny string `json:"front_shiny"`
		FrontShinyFemale interface{} `json:"front_shiny_female"`
	} `json:"sprites"`
	Stats []struct {
		BaseStat int `json:"base_stat"`
		Effort int `json:"effort"`
		Stat struct {
			Name string `json:"name"`
			URL string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			URL string `json:"url"`
		} `json:"type"`
	} `json:"types"`
	Weight int `json:"weight"`
}

type PokemonColor struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL string `json:"url"`
		} `json:"langauge"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonSpecies []struct {
		Name string `json:"name"`
		URL string `json:"url"`
	} `json:"pokemon_species"`
}

type PokemonForm struct {
	FormName string `json:"form_name"`
	FormNames []interface{} `json:"form_names"`
	FormOrder int `json:"form_order"`
	ID int `json:"id"`
	IsBattleOnly bool `json:"is_battle_only"`
	IsDefault bool `json:"is_default"`
	IsMega bool `json:"is_mega"`
	Name string `json:"name"`
	Names []interface{} `json:"names"`
	Order int `json:"order"`
	Pokemon struct {
		Name string `json:"name"`
		URL string `json:"url"`
	} `json:"pokemon"`
	Sprites struct {
		BackDefault string `json:"back_default"`
		BackShiny string `json:"back_shiny"`
		FrontDefault string `json:"front_default"`
		FrontShiny string `json:"front_shiny"`
	} `json:"sprites"`
	VersionGroup struct {
		Name string `json:"name"`
		URL string `json:"url"`
	} `json:"version_group"`
}

type PokemonHabitat struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonSpecies []struct {
		Name string `json:"name"`
		URL string `json:"url"`
	} `json:"pokemon_species"`
}

type PokemonShape struct {
	AwesomeNames []struct {
		AwesomeName string `json:"awesome_name"`
		Langauge struct {
			Name string `json:"name"`
			URL string `json:"url"`
		} `json:"langauge"`
	} `json:"awesome_names"`
	ID int `json:"id"`
	Name string `json:"name"`
	Names []struct {
		Langauge struct {
			Name string `json:"name"`
			URL string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonSpecies []struct {
		Name string `json:"name"`
		URL string `json:"url"`
	} `json:"pokemon_species"`
}

type PokemonSpecies struct {
	BaseHappiness int `json:"base_happiness"`
	CaptureRate int `json:"capture_rate"`
	Color struct {
		Name string `json:"name"`
		URL string `json:"url"`
	} `json:"color"`
	EggGroups []struct {
		Name string `json:"name"`
		URL string `json:"url"`
	} `json:"egg_groups"`
	EvolutionChain struct {
		URL string `json:"url"`
	} `json:"evolution_chain"`
	EvolvesFromSpecies interface{} `json:"evolves_from_species"`
	FlavorTextEntries []struct {
		FlavorText string `json:"flavor_text"`
		Langauge struct {
			Name string `json:"name"`
			URL string `json:"url"`
		} `json:"language"`
		Version struct {
			Name string `json:"name"`
			URL string `json:"url"`
		} `json:"version"`
	} `json:"flavor_text_entries"`
	FormDescriptions []interface{} `json:"form_descriptions"`
	FormsSwitchable bool `json:"forms_switchable"`
	GenderRate int `json:"gender_rate"`
	Genera []struct {
		Genus string `json:"genus"`
		Language struct {
			Name string `json:"name"`
			URL string `json:"url"`
		} `json:"language"`
	} `json:"genera"`
	Generation struct {
		Name string `json:"name"`
		URL string `json:"url"`
	} `json:"generation"`
	GrowthRate struct {
		Name string `json:"name"`
		URL string `json:"url"`
	} `json:"growth_rate"`
	Habitat struct {
		Name string `json:"name"`
		URL string `json:"url"`
	} `json:"habitat"`
	HasGenderDifferences bool `json:"has_gender_differences"`
	HatchCounter int `json:"hatch_counter"`
	ID int `json:"id"`
	IsBaby bool `json:"is_baby"`
	Name string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL string `json:"url"`
		} `json:"language"`
	} `json:"names"`
	Order int `json:"order"`
	PalParkEncounters []struct {
		Area struct {
			Name string `json:"name"`
			URL string `json:"url"`
		} `json:"area"`
		BaseScore int `json:"base_score"`
		Rate int `json:"rate"`
	} `json:"pal_park_encounters"`
	PokedexNumbers []struct {
		EntryNumber int `json:"entry_number"`
		Pokedex struct {
			Name string `json:"name"`
			URL string `json:"url"`
		} `json:"pokedex"`
	} `json:"pokedex_numbers"`
	Shape struct {
		Name string `json:"name"`
		URL string `json:"url"`
	} `json:"shape"`
	Varieties []struct {
		IsDefault bool `json:"is_default"`
		Pokemon struct {
			Name string `json:"name"`
			URL string `json:"url"`
		} `json:"pokemon"`
	} `json:"varieties"`
}

type Stat struct {
	AffectingMoves struct {
		Decrease []interface{} `json:"decrease"`
		Increase []struct {
			Change int `json:"change"`
			Move struct {
				Name string `json:"name"`
				URL string `json:"url"`
			} `json:"move"`
		} `json:"increase"`
	} `json:"affecting_moves"`
	AffectingNatures struct {
		Decrease []interface{} `json:"decrease"`
		Increase []interface{} `json:"increase"`
	} `json:"affecting_natures"`
	Characteristics []struct {
		URL string `json:"url"`
	} `json:"characteristics"`
	GameIndex int `json:"game_index"`
	ID int `json:"id"`
	IsBattleOnly bool `json:"is_battle_only"`
	MoveDamageClass interface{} `json:"move_damage_class"`
	Name string `json:"name"`
	Names []struct {
		Langauge struct {
			Name string `json:"name"`
			URL string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
}

type Type struct {
	DamageRelations struct {
		DoubleDamageFrom []struct {
			Name string `json:"name"`
			URL string `json:"url"`
		} `json:"double_damage_from"`
		DoubleDamageTo []interface{} `json:"double_damage_to"`
		HalfDamageFrom []interface{} `json:"half_damage_from"`
		HalfDamageTo []struct {
			Name string `json:"name"`
			URL string `json:"url"`
		} `json:"half_damage_to"`
		NoDamageFrom []struct {
			Name string `json:"name"`
			URL string `json:"url"`
		} `json:"no_damage_from"`
		NoDamageTo []struct {
			Name string `json:"name"`
			URL string `json:"url"`
		} `json:"no_damage_to"`
	} `json:"damage_relations"`
	GameIndices []struct {
		GameIndex int `json:"game_index"`
		Generation struct {
			Name string `json:"name"`
			URL string `json:"url"`
		} `json:"generation"`
		ID int `json:"id"`
		MoveDamageClass struct {
			Name string `json:"name"`
			URL string `json:"url"`
		} `json:"move_damage_class"`
		Moves []struct {
			Name string `json:"name"`
			URL string `json:"url"`
		} `json:"moves"`
		Name string `json:"name"`
		Names []struct {
			Language struct {
				Name string `json:"name"`
				URL string `json:"url"`
			} `json:"langauge"`
			Name string `json:"name"`
		} `json:"names"`
		Pokemon []struct {
			Pokemon struct {
				Name string `json:"name"`
				URL string `json:"url"`
			} `json:"pokemon"`
			Slot int `json:"slot"`
		} `json:"pokemon"`
	}
}

func main() {
	// client := resty.New()

	// resp, err := client.R().
	// 	EnableTrace().
	// 	Get("https://pokeapi.co/api/v2/pokemon/1")
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Println("Body: \n", resp)

	pokemon := Pokemon{}

	t := reflect.TypeOf(pokemon)

	fmt.Println("Struct Name:", t.Name())
	fmt.Println("Fields:")

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Printf("\t%s: %s\n", field.Name, field.Type)
	}
}