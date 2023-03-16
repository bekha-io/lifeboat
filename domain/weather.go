package domain

import "fmt"

type Weather struct {
	WeatherType WeatherType
	Description string
	Callback    func()
}

func (w Weather) CardType() CardType {
	return WeatherCardType
}

func (w Weather) Title() string {
	// Тут нужно возвращать название погодной карты и ее описание
	return fmt.Sprintf("%v (%v)", w.WeatherType, w.Description)
}

type WeatherType string

var (
	SundayWeatherType    WeatherType = "Солчнечный день"
	StormyWeatherType    WeatherType = "Шторм"
	MuggyWeatherType     WeatherType = "Знойно"
	FoggyWeatherType     WeatherType = "Туман"
	RoughSeasWeatherType WeatherType = "Качка"
	RainWeatherType      WeatherType = "Дождь"
	CalmWeatherType      WeatherType = "Штиль"
	WindyWeatherType     WeatherType = "Ветер"
	HeatWeatherType      WeatherType = "Жара"
	ClearWeatherType     WeatherType = "Чистое небо"
)

var weatherAll = []Weather{
	SundayWeather, StormyWeather, MuggyWeather, FoggyWeather, RoughSeasWeather, RainWeather, CalmWeather,
	WindyWeather, HeatWeather, ClearWeather,
}

func GetWeatherAll() []Weather {
	return weatherAll
}

var SundayWeather = Weather{
	WeatherType: SundayWeatherType,
	Description: "В этом раунде будет еще одна дележка припасов сразу после первой. Аминь",
	Callback: func() {

	},
}

var StormyWeather = Weather{
	WeatherType: StormyWeatherType,
	Description: "Те, кто должны получить жажду за греблю в этом ходу, вместо этого падают за борт",
	Callback: func() {

	},
}

var MuggyWeather = Weather{
	WeatherType: MuggyWeatherType,
	Description: "Чтобы избавиться от жажды, нужно по две фляжки воды",
	Callback: func() {

	},
}

var FoggyWeather = Weather{
	WeatherType: FoggyWeatherType,
	Description: "Проигнорируйте всех птиц с карты навигации в этом ходу",
	Callback: func() {

	},
}

var RoughSeasWeather = Weather{
	WeatherType: RoughSeasWeatherType,
	Description: "Те, кто должны получить жажду за драку в этом ходу, вместо этого падают за борт",
	Callback: func() {

	},
}

var RainWeather = Weather{
	WeatherType: RainWeatherType,
	Description: "Никто не испытывает жажду в этом ходу",
	Callback: func() {

	},
}

var CalmWeather = Weather{
	WeatherType: CalmWeatherType,
	Description: "В этом ходу нет фазы навигации",
	Callback: func() {

	},
}

var WindyWeather = Weather{
	WeatherType: WindyWeatherType,
	Description: "Сыграйте верхнюю карту из колоды навигации прямо перед фазой навигации " +
		"(итого в этом ходу будет сыграно 2 карты навигации)",
	Callback: func() {

	},
}

var HeatWeather = Weather{
	WeatherType: HeatWeatherType,
	Description: "Все испытывают дополнительную жажду в этом ходу",
	Callback: func() {

	},
}

var ClearWeather = Weather{
	WeatherType: ClearWeatherType,
	Description: "Замешайте сброшенные карты погоды в колоду погоды на носу лодки",
	Callback: func() {

	},
}
