package pages

import (
	"math/rand"
	"net/http"

	"cadavre-exquis/ces"

	"github.com/labstack/echo/v4"
)

func Index(c echo.Context) error {
	var ids []string
	ids = append(ids, "17RnDMF2H0EYjQHbAZrW", "ENBFuDYUyQ3bYT2BD4AY", "TAowjTz9Tn6YauwN5OGx", "WaSHZczbvI2KGyfiqNPc", "cHPGV4qzMRtaI52UEQlz", "jUKaofFA8GIKjI74vLXK", "mlubeO3CYUgNJBqFmf2F", "ugcMXY9CkuuO2GsnMBTN", "ycnZHCRETpGmMSHa2fzr")
	id := ids[rand.Intn(8)]
	ce, err := ces.GetCEById(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Not a valid id.")
	}
	return c.Render(http.StatusOK, "index.html", map[string]interface{}{
		"id":     ce.ID,
		"reveal": ce.Reveal,
	})
}

func Home(c echo.Context) error {
	var ids []string
	ids = append(ids, "17RnDMF2H0EYjQHbAZrW", "ENBFuDYUyQ3bYT2BD4AY", "TAowjTz9Tn6YauwN5OGx", "WaSHZczbvI2KGyfiqNPc", "cHPGV4qzMRtaI52UEQlz", "jUKaofFA8GIKjI74vLXK", "mlubeO3CYUgNJBqFmf2F", "ugcMXY9CkuuO2GsnMBTN", "ycnZHCRETpGmMSHa2fzr")
	id := ids[rand.Intn(8)]

	ce, err := ces.GetCEById(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Not a valid id.")
	}

	return c.Render(http.StatusOK, "home.html", map[string]interface{}{
		"id":     ce.ID,
		"reveal": ce.Reveal,
	})
}

func User(c echo.Context) error {
	return c.Render(http.StatusOK, "signin.html", map[string]interface{}{})
}
