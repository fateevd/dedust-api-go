package dedustApi

type Boost struct {
	Asset         string `json:"asset"`
	Budget        string `json:"budget"`
	EndAt         string `json:"endAt"`
	LiquidityPool string `json:"liquidityPool"`
	RewardPerDay  string `json:"rewardPerDay"`
	StartAt       string `json:"startAt"`
}

type Source struct {
	Chain   string `json:"chain"`
	Address string `json:"address"`
	Bridge  string `json:"bridge"`
	Symbol  string `json:"symbol"`
	Name    string `json:"name"`
}

type Asset struct {
	Type     string `json:"type"`
	Address  string `json:"address"`
	Name     string `json:"name"`
	Symbol   string `json:"symbol"`
	Image    string `json:"image"`
	Decimals int    `json:"decimals"`
	Source   Source `json:"source"`
}

type Pool struct {
	Address     string   `json:"address"`
	TotalSupply string   `json:"totalSupply"`
	Type        string   `json:"type"`
	TradeFee    string   `json:"tradeFee"`
	Assets      []string `json:"assets"`
	Reserves    []string `json:"reserves"`
	Fees        []string `json:"fees"`
	Volume      []string `json:"volume"`
}
type Price struct {
	Address string `json:"address"`
	Value   string `json:"value"`
}

type Promotions struct {
	Address string `json:"address"`
}
