package transformer

import (
	"github.com/lucasfiduniv/HomeBroker-microservices-goLang/internal/market/dto"
	"github.com/lucasfiduniv/HomeBroker-microservices-goLang/internal/market/entity"
)

func TransformerInput(input dto.TradeInput) *entity.Order
