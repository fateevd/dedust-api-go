### Documentation for `dedustApi.go`

#### Getting Started
To install the library, run:

```shell
go get github.com/fateevd/dedust-api-go
```

#### Package `dedustApi`



This package provides functions to interact with the DeDust API using GraphQL. It includes functions to fetch various data such as boosts, assets, prices, pools, and promotions.

#### Constants

- `baseUrl`: The base URL for the DeDust GraphQL API.

#### Variables

- `defaultVariables`: A map containing default variables for GraphQL requests.

#### Functions

Fetches the list of boosts.

- **Returns:**
    - `[]Boost`: A slice of `Boost` objects.
    - `error`: An error if the request fails.

##### `GetAssets() ([]Asset, error)`

Fetches the list of assets.

- **Returns:**
    - `[]Asset`: A slice of `Asset` objects.
    - `error`: An error if the request fails.

##### `GetAsset(assetAddress string) (Asset, error)`

Fetches a single asset by its address.

- **Parameters:**
    - `assetAddress`: The address of the asset.
- **Returns:**
    - `Asset`: The `Asset` object.
    - `error`: An error if the request fails.

##### `GetPrice(assetAddress string, decimals int) (Price, error)`

Fetches the price of an asset.

- **Parameters:**
    - `assetAddress`: The address of the asset.
    - `decimals`: The number of decimals for the price.
- **Returns:**
    - `Price`: The `Price` object.
    - `error`: An error if the request fails.

##### `GetPrices(assetAddress string, decimals int) ([]Price, error)`

Fetches the prices of an asset.

- **Parameters:**
    - `assetAddress`: The address of the asset.
    - `decimals`: The number of decimals for the prices.
- **Returns:**
    - `[]Price`: A slice of `Price` objects.
    - `error`: An error if the request fails.

##### `GetPools() ([]Pool, error)`

Fetches the list of pools.

- **Returns:**
    - `[]Pool`: A slice of `Pool` objects.
    - `error`: An error if the request fails.

##### `GetPool(poolAddress string) (Pool, error)`

Fetches a single pool by its address.

- **Parameters:**
    - `poolAddress`: The address of the pool.
- **Returns:**
    - `Pool`: The `Pool` object.
    - `error`: An error if the request fails.

##### `GetPromotions() ([]Promotions, error)`

Fetches the list of promotions.

- **Returns:**
    - `[]Promotions`: A slice of `Promotions` objects.
    - `error`: An error if the request fails.
