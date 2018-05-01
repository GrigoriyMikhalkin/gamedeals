package games


import (
  "net/http"
  "github.com/jinzhu/gorm"
  "github.com/labstack/echo"

  "gamedeals/database"
)


type GameDetail struct {
  Id                uint
  Name              string
  CoverUrl          string
  MinPhysicalPrice  float32
  MaxPhysicalPrice  float32
  MinDigitalPrice   float32
  MaxDigitalPrice   float32
  EditionsCount     uint
  Platforms         string
}


type GameEditionDetail struct {
  Id        uint
  Name      string
  CoverUrl  string
  MinPrice  float32
  MaxPrice  float32
  Type      uint
  Platform  uint
}


type StoreDetail struct {
  Id        uint
  Name      string
  LogoUrl   string
  Url       string
  Physical  bool
  Digital   bool
  Platforms string
}


type StoreGamePageDetail struct {
  Id        uint
  name      string
  Price     float32
  Url       string
  Currency  uint
}


func createGameDetailQuery(query *gorm.DB) *gorm.DB {
  return query.
      Joins("JOIN game_editions ON games.id = game_editions.game_id").
      Joins("JOIN games_platforms ON games.id = games_platforms.game_id").
      Joins("JOIN platforms ON games_platforms.platform_id = platforms.id").
      Joins("JOIN store_game_pages ON game_editions.id = store_game_pages.game_edition_id").
      Joins("JOIN stores ON store_game_pages.store_id = stores.id").
      Group("games.id").
      Select([]string{
          "games.id AS id",
          "games.name AS name",
          "count(DISTINCT game_editions.id) AS editions_count",
          "string_agg(DISTINCT platforms.name, ',') AS platforms",
          "(array_agg(game_editions.cover_url ORDER BY game_editions.id ASC))[1] AS cover_url",
          "min(store_game_pages.price) FILTER (WHERE stores.digital is True) as min_digital_price",
          "max(store_game_pages.price) FILTER (WHERE stores.digital is True) as max_digital_price",
          "min(store_game_pages.price) FILTER (WHERE stores.physical is True) as min_physical_price",
          "max(store_game_pages.price) FILTER (WHERE stores.physical is True) as max_physical_price",
      })
}


func getGamesList(c echo.Context) error {
  gameDetails := createGameDetailQuery(database.DB.
    Table("games"),
  ).Find(&[]GameDetail{})

  return c.JSON(http.StatusOK, gameDetails)
}


func getGameDetail(c echo.Context) error {
  id := c.Param("id")
  gameDetail := createGameDetailQuery(database.DB.
    Table("games").
    Where("games.id = ?", id),
  ).First(&GameDetail{})

  return c.JSON(http.StatusOK, gameDetail)
}


func createGameEditionDetailQuery(query *gorm.DB) *gorm.DB {
  return query.
      Joins("JOIN store_game_pages ON game_editions.id = store_game_pages.game_edition_id").
      Group("game_editions.id").
      Select([]string{
          "game_editions.id AS id",
          "game_editions.name AS name",
          "game_editions.type AS type",
          "game_editions.cover_url AS cover_url",
          "game_editions.platform_id AS platform",
          "min(store_game_pages.price) as min_price",
          "max(store_game_pages.price) as max_price",
      })
}


func getGameEditionsList(c echo.Context) error {
  gameEditionDetails := createGameEditionDetailQuery(database.DB.
    Table("game_editions"),
  ).Find(&[]GameEditionDetail{})

  return c.JSON(http.StatusOK, gameEditionDetails)
}


func getGameEditionDetail(c echo.Context) error {
  id := c.Param("id")
  gameEditionDetail := createGameEditionDetailQuery(database.DB.
    Table("game_editions").
    Where("game_editions.id = ?", id),
  ).First(&GameEditionDetail{})

  return c.JSON(http.StatusOK, gameEditionDetail)
}


func createStoreDetailQuery(query *gorm.DB) *gorm.DB {
  return query.
      Joins("JOIN platforms_stores ON stores.id = platforms_stores.store_id").
      Joins("JOIN platforms ON platforms_stores.platform_id = platforms.id").
      Group("stores.id").
      Select([]string{
          "stores.id AS id",
          "stores.name AS name",
          "stores.logo_url AS logo_url",
          "stores.url AS url",
          "stores.physical AS physical",
          "stores.digital AS digital",
          "string_agg(DISTINCT platforms.name, ',') AS platforms",
      })
}


func getStoresList(c echo.Context) error {
  storeDetails := createStoreDetailQuery(database.DB.
    Table("stores"),
  ).Find(&[]StoreDetail{})

  return c.JSON(http.StatusOK, storeDetails)
}


func getStoreDetail(c echo.Context) error {
  id := c.Param("id")
  storeDetail := createStoreDetailQuery(database.DB.
    Table("stores").
    Where("stores.id = ?", id),
  ).First(&StoreDetail{})

  return c.JSON(http.StatusOK, storeDetail)
}


func getStoreGamePage(c echo.Context) error {
  id := c.Param("gameid")
  storeGameDetails := database.DB.
    Table("game_editions").
    Joins("JOIN store_game_pages ON game_editions.id = store_game_pages.game_edition_id AND game_editions.id = ?", id).
    Joins("JOIN stores ON store_game_pages.store_id = stores.id").
    Group("store_game_pages.id").
    Select([]string{
      "store_game_pages.id AS id",
      "(array_agg(stores.name))[1] AS name",
      "store_game_pages.price AS price",
      "store_game_pages.url AS url",
      "store_game_pages.currency_id AS currency",
    }).
    Find(&[]StoreGamePageDetail{})

  return c.JSON(http.StatusOK, storeGameDetails)
}


func CreateGameEndpoints(e *echo.Echo) {
  e.GET("/games/", getGamesList)
  e.GET("/games/:id", getGameDetail)
}


func CreateGameEditionsEndpoints(e *echo.Echo) {
  e.GET("/gameeditions/", getGameEditionsList)
  e.GET("/gameeditions/:id", getGameEditionDetail)
}


func CreateStoresEndpoints(e *echo.Echo) {
  e.GET("/stores/", getStoresList)
  e.GET("/stores/:id", getStoreDetail)
}


func CreateStoreGamePagesEndpoints(e *echo.Echo) {
  e.GET("/storegamepage/:gameid", getStoreGamePage)
}


func GetSearchResult(c echo.Context) error {
  /*
   * Determines, what to return to user
   *  1) Available game editions
   *  2) Stores
   */
   searchString := c.QueryParam("q")

   return c.String(http.StatusOK, searchString)
}
