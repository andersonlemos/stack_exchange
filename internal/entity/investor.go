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

func NewInvestor(id string) *Investor {
	return &Investor{
		ID:            id,
		AssetPosition: []*InvestorAssetPosition{},
	}
}

func NewInvestorAssetPosition(assetId string, shares int) *InvestorAssetPosition {
	return &InvestorAssetPosition{
		AssetID: assetId,
		Shares:  shares,
	}
}

func (i *Investor) AddAssetPosition(assetPosition *InvestorAssetPosition) {
	i.AssetPosition = append(i.AssetPosition, assetPosition)
}

func (i *Investor) UpdatePosition(assetID string, shares int) {
	assetPosition := i.GetAssetPosition(assetID)

	if assetPosition == nil {
		i.AssetPosition = append(i.AssetPosition, NewInvestorAssetPosition(assetID, shares))
	} else {
		assetPosition.Shares = shares
	}

}

func (i *Investor) GetAssetPosition(assetId string) *InvestorAssetPosition {
	for _, assetPosition := range i.AssetPosition {
		if assetPosition.AssetID == assetId {
			return assetPosition
		}
	}
	return nil
}
