package dedustApi

import (
	"context"
	"fmt"
	"github.com/machinebox/graphql"
)

const baseUrl = "https://api.dedust.io/v3/graphql"

var defaultVariables = map[string]interface{}{}

func request[T any](schema string, variables map[string]interface{}) (T, error) {
	client := graphql.NewClient(baseUrl)

	req := graphql.NewRequest(schema)

	if len(variables) > 0 {
		fmt.Println("Variables:", variables)
		for key, value := range variables {
			req.Var(key, value)
		}
	}

	var data T

	if err := client.Run(context.Background(), req, &data); err != nil {
		fmt.Println("Error making GraphQL request:", err)
		return data, err
	}

	return data, nil
}

type ResponseBoost struct {
	Boosts []Boost `json:"boosts"`
}

func GetBoost() ([]Boost, error) {
	resp, err := request[ResponseBoost](queryGetBoosts, defaultVariables)
	return resp.Boosts, err
}

type ResponseAssets struct {
	Asset []Asset `json:"assets"`
}

func GetAssets() ([]Asset, error) {
	resp, err := request[ResponseAssets](queryGetAssets, defaultVariables)
	return resp.Asset, err
}

type ResponseAsset struct {
	Asset Asset `json:"asset"`
}

func GetAsset(assetAddress string) (Asset, error) {
	query := map[string]interface{}{
		"id": assetAddress,
	}

	resp, err := request[ResponseAsset](queryGetAsset, query)
	return resp.Asset, err
}

type ResponsePrice struct {
	Price Price `json:"price"`
}

func GetPrice(assetAddress string, decimals int) (Price, error) {
	variables := map[string]interface{}{
		"id":       assetAddress,
		"decimals": decimals,
	}
	resp, err := request[ResponsePrice](queryGetPrices, variables)
	return resp.Price, err
}

type ResponsePrices struct {
	Price []Price `json:"price"`
}

func GetPrices(assetAddress string, decimals int) ([]Price, error) {
	variables := map[string]interface{}{
		"id":       assetAddress,
		"decimals": decimals,
	}
	resp, err := request[ResponsePrices](queryGetPrices, variables)
	return resp.Price, err
}

type ResponsePools struct {
	Pools []Pool `json:"pools"`
}

func GetPools() ([]Pool, error) {
	resp, err := request[ResponsePools](queryGetPools, defaultVariables)
	return resp.Pools, err
}

type ResponsePool struct {
	Pool Pool `json:"pool"`
}

func GetPool(poolAddress string) (Pool, error) {
	variables := map[string]interface{}{
		"address": poolAddress,
	}
	resp, err := request[ResponsePool](queryGetPool, variables)
	return resp.Pool, err
}

type ResponsePromotions struct {
	Promotions []Promotions `json:"promotions"`
}

func GetPromotions() ([]Promotions, error) {
	resp, err := request[ResponsePromotions](queryGetPromotions, defaultVariables)
	return resp.Promotions, err
}
