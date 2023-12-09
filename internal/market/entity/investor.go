package entity

type Investor struct {
	ID            string
	Name          string
	AssetPosition []*InvestorAssetPosition
}
type InvestorAssetPosition struct {
	AssetID string
	Shares  int
}
