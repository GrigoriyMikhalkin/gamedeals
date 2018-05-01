package games


import (
  "gamedeals/base"
  "gamedeals/database"
)


type Game struct {
  base.Model
  Name string
  Platforms []*Platform `gorm:"many2many:games_platforms"`
  GameEditions []GameEdition
}


type Platform struct {
  base.Model
  Name string
  LogoUrl string
  Games []*Game `gorm:"many2many:games_platforms"`
  Stores []*Store `gorm:"many2many:platforms_stores"`
  GameEditions []GameEdition
}


type GameEdition struct {
  base.Model
  Name string
  CoverUrl string
  Type uint `gorm:"default:0"`
  GameID uint
  PlatformID uint
  GamePages []StoreGamePage
}


type Store struct {
  base.Model
  Name string
  Country string
  Digital bool
  Physical bool
  Url string
  LogoUrl string
  GamePages []StoreGamePage
  Platforms []*Platform `gorm:"many2many:platforms_stores"`
}


type Currency struct {
  base.Model
  Name string
  GamePages []StoreGamePage
  Dividends []CurrencyExchangeRate `gorm:"association_foreignkey:DividendID"`
  Divisors  []CurrencyExchangeRate `gorm:"association_foreignkey:DivisorID"`
}


type CurrencyExchangeRate struct {
  base.Model
  Rate float32
  DividendID uint
  DivisorID  uint
}


type StoreGamePage struct {
  base.Model
  GameEditionID uint
  StoreID uint
  Price float32
  Url string
  CurrencyID uint
}


func InitDB() {
  database.DB.AutoMigrate(&Game{})
  database.DB.AutoMigrate(&GameEdition{})
  database.DB.AutoMigrate(&Store{})
  database.DB.AutoMigrate(&StoreGamePage{})
  database.DB.AutoMigrate(&Platform{})
  database.DB.AutoMigrate(&Currency{})
  database.DB.AutoMigrate(&CurrencyExchangeRate{})
}
