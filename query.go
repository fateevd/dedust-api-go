package dedustApi

const queryGetAssets = `query GetAssets {	
	assets {
  		type
   		address
   		name
   		symbol
   		image
   		decimals
   		aliased
   		price
   		source {
   			chain
   			address
   			bridge
   			symbol
   			name
    	}
  	}
}`

const queryGetBoosts = `query GetBoosts {
	boosts {
		asset
   		budget
   		endAt
   		liquidityPool
   		rewardPerDay
   		startAt
  	}
}`

const queryGetAsset = `query GetAsset($id: ID!) {
	asset(address: $id) {
		type
		address
		name
		symbol
		image
		decimals
		aliased
		price
		source {
			chain
			address
			bridge
			symbol
     		name
    	}
  	}
}`

const queryGetPrices = `query GetPrice($id: ID!, $decimals: Int!) {
  price(address: $id, decimals:$decimals) {
  	address,
    value
  }
}`

const queryGetPools = `query GetPools($filter: PoolsFiltersInput) {
  pools(filter: $filter) {
    address
    totalSupply
    type
    tradeFee
    assets
    reserves
    fees
    volume
  }
}`

const queryGetPool = `query GetPools($address: ID!) {
  pool(address: $address) {
    address
    totalSupply
    type
    tradeFee
    assets
    reserves
    fees
    volume
  }
}`

const queryGetPromotions = `query GetPromotions {
  promotions {
	address
  }
}`
