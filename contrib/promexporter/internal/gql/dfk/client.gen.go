// Code generated by github.com/Yamashou/gqlgenc, DO NOT EDIT.

package dfk

import (
	"context"
	"net/http"

	"github.com/Yamashou/gqlgenc/client"
)

type Client struct {
	Client *client.Client
}

func NewClient(cli *http.Client, baseURL string, options ...client.HTTPRequestOption) *Client {
	return &Client{Client: client.NewClient(cli, baseURL, options...)}
}

type Query struct {
	Profile            *Profile             "json:\"profile\" graphql:\"profile\""
	Profiles           []*Profile           "json:\"profiles\" graphql:\"profiles\""
	AssistingAuction   *Auction             "json:\"assistingAuction\" graphql:\"assistingAuction\""
	AssistingAuctions  []*Auction           "json:\"assistingAuctions\" graphql:\"assistingAuctions\""
	SaleAuction        *Auction             "json:\"saleAuction\" graphql:\"saleAuction\""
	SaleAuctions       []*Auction           "json:\"saleAuctions\" graphql:\"saleAuctions\""
	Hero               *Hero                "json:\"hero\" graphql:\"hero\""
	Heroes             []*Hero              "json:\"heroes\" graphql:\"heroes\""
	Pet                *Pet                 "json:\"pet\" graphql:\"pet\""
	Pets               []*Pet               "json:\"pets\" graphql:\"pets\""
	PetAuction         *PetAuction          "json:\"petAuction\" graphql:\"petAuction\""
	PetAuctions        []*PetAuction        "json:\"petAuctions\" graphql:\"petAuctions\""
	BazaarOrder        *BazaarOrder         "json:\"bazaarOrder\" graphql:\"bazaarOrder\""
	BazaarOrders       []*BazaarOrder       "json:\"bazaarOrders\" graphql:\"bazaarOrders\""
	BazaarTransactions []*BazaarTransaction "json:\"bazaarTransactions\" graphql:\"bazaarTransactions\""
	Accessory          *Accessory           "json:\"accessory\" graphql:\"accessory\""
	Accessories        []*Accessory         "json:\"accessories\" graphql:\"accessories\""
	AccessoryAuction   *AccessoryAuction    "json:\"accessoryAuction\" graphql:\"accessoryAuction\""
	AccessoryAuctions  []*AccessoryAuction  "json:\"accessoryAuctions\" graphql:\"accessoryAuctions\""
	Armor              *Armor               "json:\"armor\" graphql:\"armor\""
	Armors             []*Armor             "json:\"armors\" graphql:\"armors\""
	ArmorAuction       *ArmorAuction        "json:\"armorAuction\" graphql:\"armorAuction\""
	ArmorAuctions      []*ArmorAuction      "json:\"armorAuctions\" graphql:\"armorAuctions\""
	Weapon             *Weapon              "json:\"weapon\" graphql:\"weapon\""
	Weapons            []*Weapon            "json:\"weapons\" graphql:\"weapons\""
	WeaponAuction      *WeaponAuction       "json:\"weaponAuction\" graphql:\"weaponAuction\""
	WeaponAuctions     []*WeaponAuction     "json:\"weaponAuctions\" graphql:\"weaponAuctions\""
}
type StuckHeroes struct {
	Heroes []*struct {
		ID *string "json:\"id\" graphql:\"id\""
	} "json:\"heroes\" graphql:\"heroes\""
}

const StuckHeroesDocument = `query StuckHeroes ($skip: Int, $owner: String) {
	heroes(skip: $skip, where: {owner:$owner}) {
		id
	}
}
`

func (c *Client) StuckHeroes(ctx context.Context, skip *int64, owner *string, httpRequestOptions ...client.HTTPRequestOption) (*StuckHeroes, error) {
	vars := map[string]interface{}{
		"skip":  skip,
		"owner": owner,
	}

	var res StuckHeroes
	if err := c.Client.Post(ctx, "StuckHeroes", StuckHeroesDocument, &res, vars, httpRequestOptions...); err != nil {
		return nil, err
	}

	return &res, nil
}