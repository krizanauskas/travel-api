package attractionservice

import (
	"fmt"
	"time"
	"travelapp/internal/config"
)

type AttractionService struct {
	ServerConfig config.ServerConfig
}

type Place struct {
	Lat                float64
	Lng                float64
	Id                 int
	Name               string
	Duration           time.Duration
	DistanceToPrevious float64
	CoverImage         string
}

func (p Place) GetCoverImgPath() string {
	return p.CoverImage
}

func New(config config.ServerConfig) AttractionService {
	return AttractionService{
		ServerConfig: config,
	}
}

func (s AttractionService) GetBreakfastPlaces() ([]Place, error) {
	breakfastPlaces := []Place{
		{
			Id:         15,
			Lat:        54.68404083458557,
			Lng:        25.28922250099148,
			Duration:   time.Duration(time.Minute * 20),
			Name:       "Karštos galvos",
			CoverImage: fmt.Sprintf("%s%s/uploads/breakfast/karstos-galvos.jpg", s.ServerConfig.Host, s.ServerConfig.Port),
		},
		{
			Id:         16,
			Lat:        54.67929956408939,
			Lng:        25.290305919414237,
			Duration:   time.Duration(time.Minute * 20),
			Name:       "Pirmas Blynas",
			CoverImage: fmt.Sprintf("%s%s/uploads/breakfast/pirmas-blynas.jpg", s.ServerConfig.Host, s.ServerConfig.Port),
		},
		{
			Id:         17,
			Lat:        54.68221482369828,
			Lng:        25.290617055640976,
			Duration:   time.Duration(time.Minute * 20),
			Name:       "Beigelistai",
			CoverImage: fmt.Sprintf("%s%s/uploads/breakfast/beigelistai.jpg", s.ServerConfig.Host, s.ServerConfig.Port),
		},
		{
			Id:         18,
			Lat:        54.69073023367748,
			Lng:        25.26324064204871,
			Duration:   time.Duration(time.Minute * 20),
			Name:       "Daily Poison",
			CoverImage: fmt.Sprintf("%s%s/uploads/breakfast/daily-poison.jpg", s.ServerConfig.Host, s.ServerConfig.Port),
		},
		{
			Id:         19,
			Lat:        54.68732380331815,
			Lng:        25.293748522165536,
			Duration:   time.Duration(time.Minute * 20),
			Name:       "StrangeLove",
			CoverImage: fmt.Sprintf("%s%s/uploads/breakfast/strangelove.jpg", s.ServerConfig.Host, s.ServerConfig.Port),
		},
		{
			Id:         20,
			Lat:        54.67235397656181,
			Lng:        25.27885439440505,
			Duration:   time.Duration(time.Minute * 20),
			Name:       "Špricas Brunch",
			CoverImage: fmt.Sprintf("%s%s/uploads/breakfast/spricas-brunch.jpg", s.ServerConfig.Host, s.ServerConfig.Port),
		},
		{
			Id:         21,
			Lat:        54.71224449729303,
			Lng:        25.286007222456234,
			Duration:   time.Duration(time.Minute * 20),
			Name:       "PAK FOODS",
			CoverImage: fmt.Sprintf("%s%s/uploads/breakfast/pak-foods.jpg", s.ServerConfig.Host, s.ServerConfig.Port),
		},
	}

	return breakfastPlaces, nil
}

func (s AttractionService) GetAttractionPlaces() ([]Place, error) {
	attractions := []Place{
		{
			Id:         1,
			Lat:        54.68537697300427,
			Lng:        25.287255310614462,
			Duration:   time.Duration(time.Minute * 40),
			Name:       "Vilniaus arkikatedra",
			CoverImage: fmt.Sprintf("%s%s/uploads/attractions/vilniaus-arkikatedra.jpg", s.ServerConfig.Host, s.ServerConfig.Port),
		},
		{
			Id:         2,
			Lat:        54.687372461209485,
			Lng:        25.301149205308647,
			Duration:   time.Duration(time.Minute * 20),
			Name:       "Stalo kalnas",
			CoverImage: fmt.Sprintf("%s%s/uploads/attractions/stalo-kalnas.jpg", s.ServerConfig.Host, s.ServerConfig.Port),
		},
		{
			Id:         3,
			Lat:        54.686978419941326,
			Lng:        25.29066424483757,
			Duration:   time.Duration(time.Minute * 20),
			Name:       "Gedimino kalnas",
			CoverImage: fmt.Sprintf("%s%s/uploads/attractions/gedimino-kalnas.jpg", s.ServerConfig.Host, s.ServerConfig.Port),
		},
		{
			Id:         4,
			Lat:        54.68387133039172,
			Lng:        54.68387133039172,
			Duration:   time.Duration(time.Minute * 20),
			Name:       "Bernardinų sodas",
			CoverImage: fmt.Sprintf("%s%s/uploads/attractions/bernardinu-sodas.jpg", s.ServerConfig.Host, s.ServerConfig.Port),
		},
		{
			Id:         5,
			Lat:        54.688050864252105,
			Lng:        25.280855180039417,
			Duration:   time.Duration(time.Minute * 20),
			Name:       "Vinco Kudirkos aikštė",
			CoverImage: fmt.Sprintf("%s%s/uploads/attractions/vinco-kudirkos-aikste.jpg", s.ServerConfig.Host, s.ServerConfig.Port),
		},
		{
			Id:         6,
			Lat:        54.70088139142579,
			Lng:        25.283196825139097,
			Duration:   time.Duration(time.Minute * 20),
			Name:       "Kalvarijų turgus",
			CoverImage: fmt.Sprintf("%s%s/uploads/attractions/kalvariju-turgus.jpg", s.ServerConfig.Host, s.ServerConfig.Port),
		},
		{
			Id:         7,
			Lat:        54.71093792384421,
			Lng:        25.261991788272386,
			Duration:   time.Duration(time.Minute * 20),
			Name:       "Akropolis",
			CoverImage: fmt.Sprintf("%s%s/uploads/attractions/akropolis.jpg", s.ServerConfig.Host, s.ServerConfig.Port),
		},
		{
			Id:         8,
			Lat:        54.6841645595128,
			Lng:        25.238656347351863,
			Duration:   time.Duration(time.Minute * 20),
			Name:       "Vingio parkas",
			CoverImage: fmt.Sprintf("%s%s/uploads/attractions/vingio-parkas.jpg", s.ServerConfig.Host, s.ServerConfig.Port),
		},
		{
			Id:         9,
			Lat:        54.748963427907555,
			Lng:        25.292848871531117,
			Duration:   time.Duration(time.Minute * 20),
			Name:       "Verkių Rūmai",
			CoverImage: fmt.Sprintf("%s%s/uploads/attractions/verkiu-rumai.jpg", s.ServerConfig.Host, s.ServerConfig.Port),
		},
		{
			Id:         10,
			Lat:        54.6793267864962,
			Lng:        25.285199001600606,
			Duration:   time.Duration(time.Minute * 20),
			Name:       "Iliuzijų muziejus",
			CoverImage: fmt.Sprintf("%s%s/uploads/attractions/iliuziju-muziejus.jpg", s.ServerConfig.Host, s.ServerConfig.Port),
		},
		{
			Id:         11,
			Lat:        54.67085554223728,
			Lng:        25.2843736073611,
			Duration:   time.Duration(time.Minute * 20),
			Name:       "Traukinių muziejus",
			CoverImage: fmt.Sprintf("%s%s/uploads/attractions/traukiniu-muziejus.jpg", s.ServerConfig.Host, s.ServerConfig.Port),
		},
		{
			Id:         12,
			Lat:        54.68817907309971,
			Lng:        25.214089597158058,
			Duration:   time.Duration(time.Minute * 20),
			Name:       "Vilniaus televizijos bokštas",
			CoverImage: fmt.Sprintf("%s%s/uploads/attractions/vilniaus-televizijos-bokstas.jpg", s.ServerConfig.Host, s.ServerConfig.Port),
		},
		{
			Id:         13,
			Lat:        54.69214910659866,
			Lng:        25.26724660324892,
			Duration:   time.Duration(time.Minute * 20),
			Name:       "Lukiškių kalėjimas 2.0",
			CoverImage: fmt.Sprintf("%s%s/uploads/attractions/lukiskiu-kalejimas-2.0.jpg", s.ServerConfig.Host, s.ServerConfig.Port),
		},
		{
			Id:         14,
			Lat:        54.692761881869835,
			Lng:        25.35247939674971,
			Duration:   time.Duration(time.Minute * 20),
			Name:       "Pūčkorių atodanga",
			CoverImage: fmt.Sprintf("%s%s/uploads/attractions/puckoriu-atodanga.jpg", s.ServerConfig.Host, s.ServerConfig.Port),
		},
	}

	return attractions, nil
}

func (s AttractionService) GetDinnerPlaces() ([]Place, error) {
	dinnerPlaces := []Place{
		{
			Id:       21,
			Lat:      54.68260243154007,
			Lng:      25.28079765963153,
			Duration: time.Duration(time.Minute * 20),
			Name:     "Trinity Restaurant and Cocktail House",
		},
		{
			Id:       22,
			Lat:      54.685055771994584,
			Lng:      25.28993539995256,
			Duration: time.Duration(time.Minute * 20),
			Name:     "Grey",
		},
		{
			Id:       23,
			Lat:      54.69051382880477,
			Lng:      25.28982756327076,
			Duration: time.Duration(time.Minute * 20),
			Name:     "River Town",
		},
		{
			Id:       24,
			Lat:      54.711698128492934,
			Lng:      25.296476806568975,
			Duration: time.Duration(time.Minute * 20),
			Name:     "Pomodoro",
		},
		{
			Id:       25,
			Lat:      54.69821667125098,
			Lng:      25.282399334532894,
			Duration: time.Duration(time.Minute * 20),
			Name:     "Rib Room",
		},
		{
			Id:       26,
			Lat:      54.69687895798521,
			Lng:      25.250506355067206,
			Duration: time.Duration(time.Minute * 20),
			Name:     "Veranda",
		},
	}

	return dinnerPlaces, nil
}
